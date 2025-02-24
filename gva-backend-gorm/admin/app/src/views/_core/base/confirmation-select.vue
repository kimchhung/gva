<script setup lang="ts">
import type { DefaultOptionType, SelectValue } from 'ant-design-vue/es/select';

import { ref, watch } from 'vue';

import { onClickOutside } from '@vueuse/core';
import { Popconfirm, Tag } from 'ant-design-vue';

import StatusSelect from './status-select.vue';

const props = defineProps<{
  defaultValue: number | string;
  disabled?: boolean;
  onConfirm?: (value: number | string) => Promise<any> | void;
  options: ({ color?: string } & DefaultOptionType)[];
}>();

const visible = ref(false);
const selectValue = ref(props.defaultValue);

const popConfirmRef = ref<HTMLElement>();

const handleChange = (value: SelectValue) => {
  selectValue.value = value as string;
  visible.value = true;
};

const handleCancel = () => {
  visible.value = false;
  selectValue.value = props.defaultValue;
};

const handleConfirm = async () => {
  if (props.onConfirm) {
    return props.onConfirm(selectValue.value)?.then((data) => {
      if (Array.isArray(data) && data[1]) {
        return;
      }
      visible.value = false;
    });
  }
};

watch(
  () => props.defaultValue,
  (value) => {
    selectValue.value = value;
  },
);

onClickOutside(popConfirmRef, () => {
  if (!visible.value) return;
  visible.value = false;
  selectValue.value = props.defaultValue;
});

const getPopupContainer = () => {
  if (popConfirmRef.value) {
    return popConfirmRef.value;
  }

  return document.body;
};
</script>

<template>
  <template v-if="disabled">
    <Tag :color="options.find((item) => item.value === selectValue)?.color">
      {{ options.find((item) => item.value === selectValue)?.label }}
    </Tag>
  </template>
  <template v-else>
    <Popconfirm
      :get-popup-container="getPopupContainer"
      :ok-text="$t('common.confirm')"
      :on-cancel="handleCancel"
      :open="visible"
      :title="$t('common.confirmation')"
      @confirm="handleConfirm"
    >
      <div
        ref="popConfirmRef"
        class="absolute left-1/2 top-0 z-10 w-[300px] -translate-x-1/2"
      ></div>
      <StatusSelect
        :disabled="disabled"
        :model-value="selectValue"
        :option-style="
          (item) => ({
            display: item.value === selectValue ? 'none' : 'block',
          })
        "
        :options="options"
        @change="handleChange"
      />
    </Popconfirm>
  </template>
</template>
