<script setup lang="ts">
import type { DefaultOptionType } from 'ant-design-vue/es/select';

import { computed } from 'vue';

import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import { Select, SelectOption, Tag } from 'ant-design-vue';

const props = defineProps<{
  disabled?: boolean;
  modelValue: number | string;
  options?: ({ color?: string } & DefaultOptionType)[];
  optionStyle?: (item: DefaultOptionType) => Record<string, string>;
  value?: number | string;
}>();

const emit = defineEmits(['update:modelValue']);

const computedValue = computed<number | string>({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit('update:modelValue', value);
  },
});

const defaultOptions = props?.options ?? [
  { label: $t('common.enable'), value: 1, color: 'green' },
  { label: $t('common.disable'), value: 0, color: 'red' },
];
</script>

<template>
  <Select
    v-model:value="computedValue"
    :disabled="disabled"
    :dropdown-match-select-width="false"
    class="custom-ant-select"
    style="border: 0"
  >
    <template #suffixIcon>
      <VbenIcon icon="teenyicons:down-solid" />
    </template>

    <SelectOption
      v-for="item in options ?? defaultOptions"
      :key="item.value"
      :style="{
        padding: '5px 0px 5px 4px',
        display: item.value === computedValue ? 'none' : 'block',
        ...(optionStyle ? optionStyle(item) : {}),
      }"
      :value="item.value"
      class="custom-ant-select-option"
    >
      <Tag :color="item?.color">
        {{ item.label }}
      </Tag>
    </SelectOption>
  </Select>
</template>

<style scoped lang="scss">
.custom-ant-select.ant-select {
  :deep(.ant-select-selector) {
    padding: 0 0 0 4px;
  }
}
</style>
