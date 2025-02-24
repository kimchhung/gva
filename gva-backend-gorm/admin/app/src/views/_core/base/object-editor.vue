<script setup lang="ts">
import { reactive, watch } from 'vue';

import { VbenIcon } from '@vben-core/shadcn-ui';

import { Button, Col, FormItem, Input, Row } from 'ant-design-vue';

const props = defineProps<{
  modelValue: Record<string, any>;
}>();

const emit = defineEmits(['update:modelValue']);

const state = reactive({
  localObject: { ...props.modelValue },
  editedKeys: Object.keys(props.modelValue),
});

watch(
  () => props.modelValue,
  (newVal) => {
    state.localObject = { ...newVal };
    state.editedKeys = Object.keys(newVal);
  },
  { immediate: true },
);

const handleAddObject = () => {
  const key = `key-${Object.keys(state.localObject).length + 1}`;
  const value = `value-${Object.keys(state.localObject).length + 1}`;

  state.localObject[key] = value;
  state.editedKeys.push(key);

  emit('update:modelValue', state.localObject);
};

const updateKey = (oldKey: string, newKey: string) => {
  if (!newKey || oldKey === newKey) return;

  if (state.localObject[oldKey] !== undefined) {
    state.localObject[newKey] = state.localObject[oldKey];
  }
  delete state.localObject[oldKey];

  state.editedKeys = state.editedKeys.map((key) =>
    key === oldKey ? newKey : key,
  );

  emit('update:modelValue', state.localObject);
};

const handleDeleteObject = (key: string) => {
  if (Object.keys(state.localObject).length === 1) return;

  delete state.localObject[key];

  const index = state.editedKeys.indexOf(key);
  if (index !== -1) state.editedKeys.splice(index, 1);

  emit('update:modelValue', state.localObject);
};
</script>

<template>
  <div
    v-for="(_, key, index) in state.localObject"
    :key="`object-${index}`"
    class="mb-2"
  >
    <Row
      align="middle"
      class="my-1 first:mb-1 first:mt-0"
      justify="space-between"
    >
      <Col :span="10">
        <FormItem :name="`object.${key}`" class="mb-0 pb-0">
          <Input
            v-model:value="state.editedKeys[index]"
            :placeholder="`Key ${index + 1}`"
            @blur="
              () =>
                state.editedKeys[index] !== undefined
                  ? updateKey(key as string, state.editedKeys[index])
                  : {}
            "
          />
        </FormItem>
      </Col>
      :
      <Col :span="10">
        <FormItem :name="`object.${key}.value`" class="mb-0 pb-0">
          <Input
            v-model:value="state.localObject[key]"
            :placeholder="`Value ${index + 1}`"
            @change="() => emit('update:modelValue', state.localObject)"
          />
        </FormItem>
      </Col>

      <Col>
        <Button
          :disabled="Object.keys(state.localObject).length === 1"
          danger
          type="dashed"
          @click="handleDeleteObject(key as string)"
        >
          <VbenIcon icon="lucide:trash" />
        </Button>
      </Col>
    </Row>
    <div
      v-if="index === Object.keys(state.localObject).length - 1"
      class="mt-4"
    >
      <Button block type="dashed" @click="handleAddObject">
        Add Key-Value
      </Button>
    </div>
  </div>
</template>
