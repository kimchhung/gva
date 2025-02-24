<script setup lang="ts">
import type { Querier } from '#/utils/pagi/form';

import { computed } from 'vue';

import { VbenIcon } from '@vben-core/shadcn-ui';

import { Button, Pagination, Select } from 'ant-design-vue';

const props = defineProps<{
  hasData: boolean;
  pagination: {
    current: number;
    hasNext: boolean;
    pageSize: number;
    total: number;
  };
  querier: Querier;
  refetch: () => void;
}>();

const handleChangePage = (opts: { limit?: number; page?: number }) => {
  props.querier.setQuery((q: any) => {
    if (opts.page) {
      q.page = opts.page;
    }
    if (opts.limit) {
      q.limit = opts.limit;
    }
    return q;
  });

  props.refetch();
};

const handlePrev = () =>
  handleChangePage({ page: props.pagination.current - 1 });
const handleNext = () =>
  handleChangePage({ page: props.pagination.current + 1 });

const isSelectListOnly = computed(() => {
  return !props.querier.defaultQuery.selects.includes('totalCount');
});
</script>

<template>
  <div v-if="hasData" class="mt-3">
    <div v-if="isSelectListOnly" class="flex items-center justify-end gap-2">
      <Button
        :disabled="pagination.current === 1"
        class="flex items-center"
        size="small"
        type="text"
        @click="handlePrev"
      >
        <VbenIcon icon="ooui:next-rtl" />
      </Button>
      <Button
        :disabled="!pagination.hasNext"
        class="flex items-center"
        size="small"
        type="text"
        @click="handleNext"
      >
        <VbenIcon icon="ooui:next-ltr" />
      </Button>
      <Select
        :dropdown-match-select-width="false"
        :options="
          [10, 20, 50, 100, 200].map((item) => ({
            label: $t('common.paginatePage', {
              num: item,
            }),
            value: item,
          }))
        "
        :value="pagination.pageSize"
        size="small"
        @change="
          (value) => {
            handleChangePage({ limit: value as number });
          }
        "
      />
    </div>
    <div v-else class="flex justify-end">
      <Pagination
        :current="pagination.current"
        :page-size="pagination.pageSize"
        :total="pagination.total"
        show-size-changer
        size="small"
        @show-size-change="
          (current, pageSize) => {
            handleChangePage({ limit: pageSize, page: current });
          }
        "
        @update:current="(current) => handleChangePage({ page: current })"
      />
    </div>
  </div>
</template>

<style scoped lang="scss">
.cursor-paginate {
  :deep(.ant-pagination-item) {
    display: none !important;
  }
}
</style>
