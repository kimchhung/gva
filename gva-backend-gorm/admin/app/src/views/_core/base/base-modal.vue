<script setup lang="ts">
import { defineProps } from 'vue';

import { Modal, type ModalProps } from 'ant-design-vue';

const props = defineProps<
  {
    entity?: string;
    mode?: 'create' | 'edit';
    okLoading?: boolean;
    okText?: string;
  } & ModalProps
>();
</script>

<template>
  <Modal
    v-bind="props"
    v-if="!$slots.title"
    :body-style="{ paddingTop: '10px' }"
    :ok-text="props.okText ?? $t('common.save')"
    :title="
      title ??
      $t(props.mode === 'create' ? 'common.create' : 'common.edit', {
        name: entity,
      })
    "
    closable
    mask
  >
    <template v-if="$slots.title" #title>
      <slot name="title"></slot>
    </template>
    <slot></slot>
    <template v-if="$slots.footer" #footer>
      <slot name="footer"></slot>
    </template>
  </Modal>
</template>
