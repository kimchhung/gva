<script setup lang="tsx">
import type { Admin } from '#/api/admin/types';

import { computed, ref } from 'vue';

import { useAccess } from '@vben/access';
import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { useUserStore } from '@vben/stores';

import {
  notification,
  Pagination,
  type TableColumnsType,
} from 'ant-design-vue';

import { ADMIN_PERMISSION } from '#/constants';
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

import { api } from '#/api';
import IpWhitelistForm from './components/ip-whitelist-form.vue';
import ResetPasswordForm from './components/reset-password-form.vue';
import TotpForm from './components/totp-form.vue';
import UserForm from './components/user-form.vue';

const { hasAccessByPermissions } = useAccess();
const { userInfo } = useUserStore();

const { actionGroupRef, actionWidth } = useGroupActionWidth();

const columns = computed<TableColumnsType<Admin>>(() => [
  tableColumns.id(),
  tableColumns.name({
    width: 140,
  }),
  {
    dataIndex: 'username',
    title: $t('common.username'),
    width: 140,
  },
  tableColumns.tagColumn('roles', $t('common.roles'), {
    getItem: (role) => role.name,
    width: 260,
  }),
  tableColumns.status({
    title: $t('page.admin.users.form.totp'),
    dataIndex: 'enableTOTP',
    width: 140,
  }),
  tableColumns.status({
    width: 140,
    disabled: (record) =>
      !hasAccessByPermissions(withSuper(ADMIN_PERMISSION.EDIT)) ||
      (!!userInfo?.userId && record.id === userInfo.userId),
    onConfirm: async (newValue: any, record: Admin) => {
      return api().admin.updatePartial({
        body: {
          status: newValue,
        },
        id: record.id,
        opt: {
          onSuccess: () => {
            notification.success({
              message: $t('message.updateSuccess'),
            });
          },
        },
      });
    },
  }),
  tableColumns.tagColumn(
    'currentLoginIp',
    $t('page.admin.users.form.lastLoginIP'),
    {
      width: 140,
    },
  ),
  tableColumns.tagColumn(
    'currentRegion',
    $t('page.admin.users.form.lastLoginRegion'),
    {
      width: 140,
    },
  ),
  tableColumns.dateColumn(
    'currentLoginAt',
    $t('page.admin.users.form.lastLoginAt'),
    {
      width: 200,
    },
  ),
  tableColumns.createdAt(),
  tableColumns.action({
    width: actionWidth.value,
  }),
]);

const querier = useQueryForm({
  model: {
    createdAt: TypeDate(),
    name: TypeString(),
    username: TypeString(),
  },
  modelConfig: computed(() => ({
    createdAt: {
      label: computed(() => $t('common.created_at')),
    },
    name: {
      label: computed(() => $t('common.name')),
    },
    username: {
      label: computed(() => $t('common.username')),
    },
  })),
});

const selectedRecord = ref<Admin | null>(null);

const userForm = ref(false);
const openUserForm = (record: Admin | null) => {
  selectedRecord.value = record;
  userForm.value = true;
};

const ipWhitelistForm = ref(false);
const openWhiteListForm = (record: Admin | null) => {
  selectedRecord.value = record;
  ipWhitelistForm.value = true;
};

const totpForm = ref(false);
const openTotpForm = (record: Admin | null) => {
  selectedRecord.value = record;
  totpForm.value = true;
};

const handleDeleteUser = async (record: Admin) => {
  return api().admin.delete({
    id: record.id,
    opt: {
      onSuccess: () => {
        api().admin.getMany.invalidate();
        notification.success({
          message: $t('message.deleteSuccess'),
        });
      },
    },
  });
};

const resetPasswordForm = ref(false);
const openResetPasswordForm = (record: Admin | null) => {
  selectedRecord.value = record;
  resetPasswordForm.value = true;
};
const { tableState, tableFunction, tableTool } = useTableData<Admin>({
  querier: querier as any,
  columns,
  fetcher: (api) => api.admin.getMany,
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
    :title="$t('page.admin.users.title')"
    show-footer
  >
    <UserForm v-model:open="userForm" :record="selectedRecord" />
    <IpWhitelistForm v-model:open="ipWhitelistForm" :record="selectedRecord" />
    <ResetPasswordForm
      v-model:open="resetPasswordForm"
      :record="selectedRecord"
    />
    <TotpForm v-model:open="totpForm" :record="selectedRecord" />
    <template #appendHeader>
      <Filter
        :querier="querier as any"
        @search="() => tableFunction.forceRefetch()"
      >
        <ActionButton
          action-type="create"
          v-permissions="withSuper(ADMIN_PERMISSION.ADD)"
          @click="openUserForm(null)"
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
              vPermissions: withSuper(ADMIN_PERMISSION.EDIT),
              onClick: () => openUserForm(record),
            },
            {
              value: $t('common.delete'),
              actionType: 'delete',
              disabled: record.id === userInfo?.userId,
              vPermissions: withSuper(ADMIN_PERMISSION.DELETE),
              onClick: () => handleDeleteUser(record),
            },
            {
              value: $t('page.admin.users.form.ip-whitelist'),
              icon: 'eos-icons:ip',
              vPermissions: withSuper(ADMIN_PERMISSION.EDIT),
              onClick: () => openWhiteListForm(record),
            },
            {
              value: $t('page.admin.users.form.reset-password'),
              icon: 'fluent:password-reset-48-regular',
              vPermissions: withSuper(ADMIN_PERMISSION.EDIT),
              onClick: () => openResetPasswordForm(record),
            },
            {
              value: $t('page.admin.users.form.totp'),
              icon: 'uim:google',
              vPermissions: withSuper(ADMIN_PERMISSION.EDIT),
              onClick: () => openTotpForm(record),
            },
          ]"
          :max-display="2"
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
