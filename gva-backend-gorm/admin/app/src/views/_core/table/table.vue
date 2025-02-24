<script setup lang="ts" generic="T extends object = any">
import { Table, type TableColumnType, type TableProps } from 'ant-design-vue';

const props = defineProps<{
  pageSize: {
    limit: number;
    page: number;
  };
  tableProps?: { style?: Record<string, string> } & TableProps<T>;
}>();

defineSlots<{
  action?: (props: {
    column: TableColumnType<T>;
    index: number;
    record: T;
  }) => void;
  bodyCell?: (props: {
    column: TableColumnType<T>;
    index: number;
    record: T;
    text: any;
    value: any;
  }) => void;
}>();
</script>

<template>
  <div class="h-full flex-grow overflow-hidden">
    <Table
      :pagination="false"
      :scroll="{ x: 1000, y: 'auto' }"
      class="custom-table h-full"
      size="small"
      v-bind="props.tableProps"
    >
      <template #bodyCell="{ column, index, record, text, value }">
        <slot
          name="bodyCell"
          v-bind="{ column, index, record: record as T, text, value }"
        ></slot>
        <template v-if="column.key === 'index'">
          {{ ((pageSize?.page ?? 1) - 1) * (pageSize.limit ?? 20) + index + 1 }}
        </template>
        <template v-if="column.key === 'action'">
          <slot
            name="action"
            v-bind="{ column, index, record: record as T }"
          ></slot>
        </template>
      </template>
    </Table>
  </div>
</template>

<style scoped lang="scss">
.custom-table {
  :deep(.ant-table-cell-fix-right) {
    @apply max-[820px]:right-[unset] !important;
  }

  :deep(.ant-spin-nested-loading),
  :deep(.ant-spin-container),
  :deep(.ant-table),
  :deep(.ant-table-container) {
    height: 100%;
  }

  :deep(.ant-table-body) {
    max-height: calc(100% - 38px);
  }
}
</style>
