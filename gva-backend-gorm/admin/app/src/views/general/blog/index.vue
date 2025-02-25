<script setup lang="tsx">
import type { Blog } from '#/api/blog/types';

import { computed } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  notification,
  Pagination,
  type TableColumnsType,
} from 'ant-design-vue';

import { api } from '#/api';
import { BLOG_PERMISSION } from '#/constants';
import { withSuper } from '#/utils/helper/permissions';
import { TypeDate, TypeString, useQueryForm } from '#/utils/pagi/form';
import { tableColumns } from '#/utils/table/column';
import { useGroupActionWidth } from '#/utils/table/use-group-action';
import { useTableData } from '#/utils/table/use-table-data';
import ActionButton from '#/views/_core/button/action-button.vue';
import GroupButton from '#/views/_core/button/group-button.vue';
import Filter from '#/views/_core/filter/index.vue';
import Table from '#/views/_core/table/table.vue';
import TableTool from '#/views/_core/table/tool.vue';

const { actionGroupRef, actionWidth } = useGroupActionWidth();

const columns = computed<TableColumnsType<Blog>>(() => [
  tableColumns.id(),
  tableColumns.createdAt(),
  tableColumns.action({
    width: actionWidth.value,
  }),
]);

const querier = useQueryForm({
  model: {
    createdAt: TypeDate(),
    id: TypeString(),
  },
  modelConfig: computed(() => ({
    createdAt: {
      label: computed(() => $t('common.created_at')),
    },
    id: {
      label: computed(() => $t('common.id')),
    },
  })),
});

const handleDeleteBlog = async (record: Blog) => {
  return api().blog.delete({
    id: record.id,
    opt: {
      onSuccess: () => {
        api().blog.getMany.invalidate();
        notification.success({
          message: $t('message.deleteSuccess'),
        });
      },
    },
  });
};

const { tableState, tableFunction, tableTool } = useTableData<Blog>({
  querier: querier as any,
  columns,
  fetcher: (api) => api.blog.getMany,
  table: {
    getHasNext: (res) => !!res?.meta?.hasNext,
    getList: (res) => res?.data,
    getTotal: (res) => res?.meta?.totalCount,
  },
});
</script>

<template>
  <Page
    :badge-text="tableState.total().toString()"
    :title="$t('page.blog.title')"
    show-footer
  >
    <template #appendHeader>
      <Filter
        :querier="querier as any"
        @search="() => tableFunction.forceRefetch()"
      >
        <ActionButton
          action-type="create"
          v-permissions="withSuper(BLOG_PERMISSION.ADD)"
        />
      </Filter>
    </template>
    <Table
      :page-size="{
        limit: querier.query.limit,
        page: querier.query.page,
      }"
      :table-props="{
        onChange: (_, __, sort) => tableFunction.handleSort(sort),
        loading: tableState.isLoading(),
        dataSource: tableState.list(),
        columns: tableState.columns(),
      }"
    >
      <template #action="{ record }">
        <GroupButton
          ref="actionGroupRef"
          :buttons="[
            {
              value: $t('common.edit'),
              actionType: 'edit',
              vPermissions: withSuper(BLOG_PERMISSION.VIEW),
            },
            {
              value: $t('common.edit'),
              actionType: 'edit',
              vPermissions: withSuper(BLOG_PERMISSION.EDIT),
              onClick: () => handleDeleteBlog(record),
            },
            {
              value: $t('common.delete'),
              actionType: 'delete',
              vPermissions: withSuper(BLOG_PERMISSION.DELETE),
              onClick: () => handleDeleteBlog(record),
            },
          ]"
          :max-display="3"
          @no-permissions="() => tableFunction.hideColumn('action')"
        />
      </template>
    </Table>
    <template #footer>
      <TableTool :table-tool="tableTool" position="bottom">
        <Pagination
          :current="querier.query.page"
          :page-size="querier.query.limit"
          :total="tableState.total()"
          show-size-changer
          @show-size-change="(_, pageSize) => tableFunction.setLimit(pageSize)"
          @update:current="(current) => tableFunction.goToPage(current)"
        />
      </TableTool>
    </template>
  </Page>
</template>
