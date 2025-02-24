<script setup lang="ts" generic="T extends object = any">
import type { Querier } from '#/utils/pagi/form';

import { computed } from 'vue';

import {
  Table,
  type TableColumnsType,
  type TableColumnType,
  type TableProps,
} from 'ant-design-vue';

import { useTableSort } from '#/utils/table/use-table-sort';
import Filter from '#/views/_core/filter/index.vue';

import DataTablePaginate from './data-table-paginate.vue';

const props = withDefaults(
  defineProps<{
    columns: TableColumnsType<T>;
    getApi: (getMany: typeof api) => any;
    hideFilter?: boolean;
    minTableWidth?: number;
    querier: any;
    tableProps?: { style?: Record<string, string> } & TableProps<T>;
  }>(),
  {
    hideFilter: false,
    tableProps: undefined,
    minTableWidth: 1000,
  },
);

defineSlots<{
  action?: (props: {
    column: TableColumnType<T>;
    index: number;
    record: Record<string, any> & T;
  }) => void;
  bodyCell?: (props: {
    column: TableColumnType<T>;
    index: number;
    record: Record<string, any> & T;
    text: any;
    value: any;
  }) => void;
}>();

const querier: Querier = props.querier;

const getMany = props.getApi(api);

if (!('useQuery' in getMany)) {
  throw new Error('getMany must have a use method');
}

const query = computed(() => querier.query);

const {
  data: resp,
  isFetching,
  refetch,
} = getMany.useQuery({
  query,
});

const pagination = computed(() => {
  const { limit, page } = query.value;
  return {
    current: page || resp.value?.meta?.page,
    pageSize: limit || resp.value?.meta?.limit,
    total: resp.value?.meta?.totalCount,
    hasNext: resp.value?.meta?.hasNext,
  };
});

const forceRefetch = () => {
  const queryNow = querier.getQuery();
  querier.setQuery(() => queryNow);
  refetch();
};

const { handleSort } = useTableSort(querier);
</script>

<template>
  <div class="flex h-full flex-col">
    <Filter
      v-if="!hideFilter"
      :querier="querier"
      @search="() => forceRefetch()"
    />
    <div class="flex-grow overflow-hidden">
      <Table
        :columns="columns"
        :data-source="resp?.data"
        :loading="isFetching"
        :pagination="false"
        :scroll="{ x: minTableWidth, y: 'auto' }"
        class="custom-table h-full"
        size="small"
        v-bind="tableProps"
        @change="
          (_, __, sort) => {
            handleSort(sort);
          }
        "
      >
        <template #bodyCell="{ column, index, record, text, value }">
          <slot
            name="bodyCell"
            v-bind="{ column, index, record: record as T, text, value }"
          ></slot>
          <template v-if="column.key === 'index'">
            {{
              ((pagination?.current ?? 1) - 1) * (pagination?.pageSize ?? 10) +
              index +
              1
            }}
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
    <DataTablePaginate
      :has-data="resp?.data.length > 0"
      :pagination="pagination"
      :querier="querier"
      :refetch="() => {}"
    />
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
    max-height: calc(100% - 39px);
  }
}
</style>
