<script setup lang="ts">
import { Popconfirm, type PopconfirmProps } from 'ant-design-vue';

const props = defineProps<PopconfirmProps>();
</script>

<template>
  <Popconfirm
    v-bind="props"
    :ok-text="$t('common.confirm')"
    :title="$t('common.confirmation')"
  >
    <slot></slot>
  </Popconfirm>
</template>
