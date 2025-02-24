<script setup lang="ts" generic="T">
import type { TableTool } from '#/utils/table/use-table-data';

import { computed } from 'vue';

import TableDisplay from './display.vue';

defineOptions({ name: 'TableTool' });

const props = defineProps<{
  position: 'bottom' | 'top';
  tableTool: TableTool<T>;
}>();

const isTop = computed(() => props.position === 'top');
</script>

<template>
  <div
    :class="{
      'flex-col-reverse': !isTop,
    }"
    class="flex w-full flex-col"
  >
    <div
      :class="{
        'mb-2': isTop,
        'mt-2': !isTop,
      }"
      class="flex w-full flex-wrap justify-between gap-2"
    >
      <div class="flex">
        <TableDisplay :table-tool="tableTool" />
      </div>
      <slot></slot>
    </div>

    <slot name="table"></slot>
  </div>
</template>
