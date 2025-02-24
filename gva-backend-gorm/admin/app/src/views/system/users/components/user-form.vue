<script setup lang="ts">
import type { Rule } from 'ant-design-vue/es/form';

import type { Admin } from '#/api/admin/types';

import { computed, reactive } from 'vue';

import { $t } from '@vben/locales';

import { Form, FormItem, Input, InputPassword, Select } from 'ant-design-vue';

import { api } from '#/api';
import { useBaseModalForm } from '#/hooks/use-base-modal-form';
import { useValidator } from '#/hooks/use-validator';
import { getFormInfos } from '#/utils/form/label';
import { defaultQuery, Operation } from '#/utils/pagi/query';
import BaseModal from '#/views/_core/base/base-modal.vue';

const props = defineProps<{
  open: boolean;
  record: Admin | null;
}>();

const defaultFormState = () => ({
  username: '',
  name: '',
  password: '',
  roles: [] as number[],
});

const formState = reactive(defaultFormState());

const {
  formRef,
  openRef,
  getFormDirtyState,
  handleSubmit,
  loadingRef,
  closeModal,
} = useBaseModalForm(props, {
  onOpen: ({ record }) => {
    if (record) {
      formState.username = record.username;
      formState.name = record.name;
      formState.roles = record.roles.map((role) => role.id);
    } else {
      Object.assign(formState, defaultFormState());
    }
    return formState;
  },
});

const { data: adminRoles, isLoading } = api().adminRole.getMany.useQuery({
  query: () =>
    defaultQuery({
      filters: [
        {
          column: 'status' as any,
          operation: Operation.Equal,
          value: '1',
        },
      ],
    }),
});

const formInfos = computed(() =>
  getFormInfos({
    username: $t('common.username'),
    password: $t('common.password'),
    name: $t('common.name'),
    roles: $t('common.roles'),
  }),
);

const { required, notSpace, notSpecialCharacters, lengthRange, text } =
  useValidator();

const rules: Record<string, Rule[]> = {
  username: [
    required(formInfos.value.username.label),
    notSpace(),
    notSpecialCharacters(),
    lengthRange(5, 30),
  ],
  name: [required(formInfos.value.name.label), text({ allowNumber: true })],
  password: [required(formInfos.value.password.label), lengthRange(6, 20)],
  roles: [required(formInfos.value.roles.label)],
};

const onSubmit = handleSubmit(async () => {
  const payload = {
    username: formState.username,
    name: formState.name,
    roles: formState.roles.map((id) => ({ id })),
  };

  const options = {
    onSuccess: () => {
      api().admin.getMany.invalidate();
      closeModal({
        showSuccessMessage: true,
      });
    },
    loading: loadingRef,
  };

  await (props.record
    ? api().admin.update({
        id: props.record!.id,
        body: payload,
        opt: options,
      })
    : api().admin.create({
        body: {
          ...payload,
          password: formState.password,
        },
        opt: options,
      }));
});
</script>

<template>
  <BaseModal
    v-model:open="openRef"
    :entity="$t('page.admin.users.title')"
    :mode="!props.record ? 'create' : 'edit'"
    :ok-button-props="{
      loading: loadingRef,
      disabled: !getFormDirtyState({ ...formState }),
    }"
    @ok="onSubmit"
  >
    <Form ref="formRef" :model="formState" :rules="rules" layout="vertical">
      <FormItem :label="formInfos.username.label" name="username">
        <Input
          v-model:value="formState.username"
          :placeholder="formInfos.username.placeholder"
        />
      </FormItem>
      <template v-if="!record">
        <FormItem :label="formInfos.password.label" name="password">
          <InputPassword
            v-model:value="formState.password"
            :placeholder="formInfos.password.placeholder"
            autocomplete="new-password"
          />
        </FormItem>
      </template>
      <FormItem :label="formInfos.name.label" name="name">
        <Input
          v-model:value="formState.name"
          :placeholder="formInfos.name.placeholder"
        />
      </FormItem>
      <FormItem :label="formInfos.roles.label" name="roles">
        <Select
          v-model:value="formState.roles"
          :field-names="{ label: 'name', value: 'id' }"
          :loading="isLoading"
          :options="adminRoles?.data ?? []"
          :placeholder="formInfos.roles.placeholder"
          allow-clear
          mode="multiple"
          option-filter-prop="name"
          show-arrow
          show-search
        />
      </FormItem>
    </Form>
  </BaseModal>
</template>
