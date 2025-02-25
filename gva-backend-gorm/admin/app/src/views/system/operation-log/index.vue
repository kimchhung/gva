<script setup lang="tsx">
import type { OperationLog } from '#/api/operation-log/types';

import { computed, ref } from 'vue';

import { Page } from '@vben/common-ui';

import { Pagination, type TableColumnsType, Tag } from 'ant-design-vue';

import { api } from '#/api';
import { OPERATION_LOG_PERMISSION } from '#/constants';
import { METHODS_COLOR } from '#/constants/operation-log';
import { $t } from '#/locales';
import { todayOnly } from '#/utils/helper/date-util';
import { withSuper } from '#/utils/helper/permissions';
import {
  TypeDate,
  TypeNumber,
  TypeSelect,
  TypeString,
  useQueryForm,
} from '#/utils/pagi/form';
import { defaultQueryPagination } from '#/utils/pagi/query';
import { tableColumns } from '#/utils/table/column';
import { useGroupActionWidth } from '#/utils/table/use-group-action';
import { useTableData } from '#/utils/table/use-table-data';
import GroupButton from '#/views/_core/button/group-button.vue';
import Filter from '#/views/_core/filter/index.vue';
import Table from '#/views/_core/table/table.vue';
import TableTool from '#/views/_core/table/tool.vue';

import OperationLogDetail from './components/operation-log-detail.vue';

const { actionGroupRef, actionWidth } = useGroupActionWidth();

const columns = computed<TableColumnsType<OperationLog>>(() => [
  tableColumns.id(),
  tableColumns.dateColumn(
    'createdAt',
    $t('page.system.operationLog.createdAt'),
    {
      width: 120,
      key: 'operatingTime',
    },
  ),
  {
    dataIndex: ['admin', 'username'],
    title: $t('page.system.operationLog.operator'),
    width: 100,
    key: 'operator',
  },
  {
    dataIndex: 'ip',
    title: $t('page.system.operationLog.ip'),
    width: 100,
  },
  {
    width: 160,
    dataIndex: 'scope',
    title: $t('page.system.operationLog.scope'),
    customRender: ({ value }) => {
      if (!value) return tableColumns.nullText;
      const [scope, action] = value.split(':');
      return (
        <Tag color="orange">
          {`${$t(`permission.${action}`, {
            name: $t(`permission.${scope}`),
          })}`}{' '}
        </Tag>
      );
    },
  },
  {
    dataIndex: 'path',
    title: $t('page.system.operationLog.path'),
    width: 220,
    customRender: ({ record }) => (
      <p>
        <Tag color={METHODS_COLOR?.[record.method]}>{record.method}</Tag>
        <span class="text-sm">{record.path}</span>
      </p>
    ),
  },
  tableColumns.tagColumn('code', $t('page.system.operationLog.statusCode'), {
    getProps: (value) => ({
      color: value >= 0 ? 'green' : 'red',
    }),
    width: 60,
  }),
  tableColumns.tagColumn('latency', $t('page.system.operationLog.latency'), {
    getItem: (value) => `${value}ms`,
    width: 80,
  }),
  tableColumns.action({
    width: actionWidth.value,
  }),
]);

const { data: permissionRes } = api().permission.getMany.useQuery();
const permissionOptions = computed(() => {
  return (
    permissionRes.value?.data?.reduce(
      (acc, item) => {
        const [scope, action] = item.scope.split(':');
        if (action === 'super') return acc;
        acc.push({
          label: `${$t(`permission.${action}`, {
            name: $t(`permission.${scope}`),
          })}`,
          value: item.scope,
        });

        return acc;
      },
      [] as { label: string; value: string }[],
    ) || []
  );
});

const querier = useQueryForm({
  model: {
    operatingTime: TypeDate(todayOnly()),
    method: TypeSelect(
      () => [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' },
        { label: 'PATCH', value: 'PATCH' },
      ],
      [
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' },
        { label: 'PATCH', value: 'PATCH' },
      ],
    ),
    action: TypeSelect(() => permissionOptions.value),
    latency: TypeNumber(),
    code: TypeNumber(),
    ip: TypeString(),
  },
  modelConfig: computed(() => ({
    operatingTime: {
      label: $t('page.system.operationLog.createdAt'),
      isFixed: true,
    },
    method: {
      label: $t('page.system.operationLog.method'),
    },
    latency: {
      label: $t('page.system.operationLog.latency'),
    },
    code: {
      label: $t('page.system.operationLog.statusCode'),
    },
    action: {
      label: $t('common.action'),
    },
    ip: {
      label: $t('page.system.operationLog.ip'),
    },
  })),
  query: defaultQueryPagination({
    sorts: [
      {
        column: 'operatingTime',
        direction: 'desc',
      },
    ],
    selects: ['list'],
  }),
});

const openDetail = ref(false);
const selectedRecord = ref<null | OperationLog>(null);

const { tableState, tableFunction, tableTool } = useTableData<OperationLog>({
  querier: querier as any,
  columns,
  fetcher: (api) => api.operationLog.getMany,
  table: {
    getHasNext: (res) => !!res?.meta?.hasNext,
    getList: (res) => res?.data,
    getTotal: (res) => res?.meta?.totalCount,
  },
});
</script>

<template>
  <Page :title="$t('page.system.operationLog.title')" show-footer>
    <OperationLogDetail v-model:open="openDetail" :record="selectedRecord" />
    <template #appendHeader>
      <Filter
        :querier="querier as any"
        @search="() => tableFunction.forceRefetch()"
      />
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
              vPermissions: withSuper(OPERATION_LOG_PERMISSION.VIEW),
              actionType: 'detail',
              onClick: () => {
                selectedRecord = record;
                openDetail = true;
              },
            },
          ]"
          @no-permissions="() => tableFunction.hideColumn('action')"
        />
      </template>
    </Table>
    <template #footer>
      <TableTool :table-tool="tableTool" position="bottom">
        <Pagination
          :current="querier.query.page"
          :page-size="querier.query.limit"
          :total="
            tableState.hasNext()
              ? (querier.query.page + 1) * querier.query.limit
              : querier.query.page * querier.query.limit
          "
          show-size-changer
          @show-size-change="(_, pageSize) => tableFunction.setLimit(pageSize)"
          @update:current="(current) => tableFunction.goToPage(current)"
        />
      </TableTool>
    </template>
  </Page>
</template>
