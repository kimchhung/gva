<script setup lang="tsx">
import type { AdminRole } from '#/api/admin-role/types';

import { computed } from 'vue';
import { useRouter } from 'vue-router';

import { useAccess } from '@vben/access';
import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  notification,
  Pagination,
  type TableColumnsType,
} from 'ant-design-vue';

import { api } from '#/api';
import { ADMIN_ROLE_PERMISSION, ROLE } from '#/constants';
import { withSuper } from '#/utils/helper/permissions';
import { TypeDate, TypeString, useQueryForm } from '#/utils/pagi/form';
import { tableColumns } from '#/utils/table/column';
import { useGroupActionWidth } from '#/utils/table/use-group-action';
import { useTableData } from '#/utils/table/use-table-data';
import ConfirmationSelect from '#/views/_core/base/confirmation-select.vue';
import ActionButton from '#/views/_core/button/action-button.vue';
import GroupButton from '#/views/_core/button/group-button.vue';
import Filter from '#/views/_core/filter/index.vue';
import Table from '#/views/_core/table/table.vue';
import TableTool from '#/views/_core/table/tool.vue';

const { hasAccessByPermissions } = useAccess();

const { actionGroupRef, actionWidth } = useGroupActionWidth();
const columns = computed<TableColumnsType<AdminRole>>(() => [
  tableColumns.id(),
  tableColumns.name({
    width: 200,
  }),
  tableColumns.textColumn('description', $t('common.description'), {
    width: 380,
  }),
  tableColumns.status({
    width: 120,
    customRender: ({ value, record }) => {
      const options = [
        { label: $t('common.enable'), value: 1, color: 'green' },
        { label: $t('common.disable'), value: 0, color: 'red' },
      ];

      return (
        <ConfirmationSelect
          defaultValue={value}
          disabled={
            record.nameId === ROLE.SUPER_ADMIN ||
            !hasAccessByPermissions(withSuper(ADMIN_ROLE_PERMISSION.EDIT))
          }
          onConfirm={(newValue) => {
            return api().adminRole.updatePartial({
              body: {
                status: newValue as number,
              },
              id: record.id,
              opt: {
                onSuccess: () => {
                  api().adminRole.getMany.invalidate();
                  notification.success({
                    message: $t('message.updateSuccess'),
                  });
                },
              },
            });
          }}
          options={options}
        />
      );
    },
  }),
  tableColumns.createdAt(),
  tableColumns.updatedAt(),
  tableColumns.action({
    width: actionWidth.value,
  }),
]);

const querier = useQueryForm({
  model: {
    createdAt: TypeDate(),
    name: TypeString(),
  },
  modelConfig: computed(() => ({
    createdAt: {
      label: computed(() => $t('common.created_at')),
    },
    name: {
      label: computed(() => $t('common.name')),
    },
  })),
});

const handleDeleteUser = async (record: AdminRole) => {
  return api().adminRole.delete({
    id: record.id,
    opt: {
      onSuccess: () => {
        api().adminRole.getMany.invalidate();
        notification.success({
          message: $t('message.deleteSuccess'),
        });
      },
    },
  });
};

const router = useRouter();
const handleCreate = () => {
  router.push({
    path: `${router.currentRoute.value.path}/create`,
  });
};

const handleEdit = (record: AdminRole) => {
  router.push({
    path: `${router.currentRoute.value.path}/${record.id}`,
  });
};

const { tableState, tableFunction, tableTool } = useTableData<AdminRole>({
  querier: querier as any,
  columns,
  fetcher: (api) => api.adminRole.getMany,
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
    :title="$t('page.admin.roles.title')"
    show-footer
  >
    <template #appendHeader>
      <Filter
        :querier="querier as any"
        @search="() => tableFunction.forceRefetch()"
      >
        <ActionButton
          action-type="create"
          v-permissions="withSuper(ADMIN_ROLE_PERMISSION.ADD)"
          @click="handleCreate"
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
        <div :class="{ hidden: record.nameId === ROLE.SUPER_ADMIN }">
          <GroupButton
            ref="actionGroupRef"
            :buttons="[
              {
                vPermissions: withSuper(ADMIN_ROLE_PERMISSION.EDIT),
                actionType: 'edit',
                onClick: () => handleEdit(record),
              },
              {
                vPermissions: withSuper(ADMIN_ROLE_PERMISSION.DELETE),
                actionType: 'delete',
                onClick: () => handleDeleteUser(record),
              },
            ]"
            @no-permissions="() => tableFunction.hideColumn('action')"
          />
        </div>
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
