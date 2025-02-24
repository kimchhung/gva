<script setup lang="ts">
import type { Permission } from '#/api/permission/types';

import { computed, onBeforeMount, reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Button,
  Form,
  type FormInstance,
  FormItem,
  Input,
  notification,
  Space,
  Textarea,
  Tree,
} from 'ant-design-vue';

import { api } from '#/api';
import { getPermissionTree } from '#/constants/permission-tree';
import { useValidator } from '#/hooks/use-validator';
import { getFormInfos } from '#/utils/form/label';

const defaultFormState = () => ({
  name: '',
  description: '',
  permissions: [] as string[],
});

const formRef = ref<FormInstance>();
const formState = reactive(defaultFormState());

const formInfos = computed(() =>
  getFormInfos({
    name: $t('common.name'),
    description: $t('common.description'),
    permissions: $t('page.admin.roles.form.permissions'),
  }),
);

const treeData = computed(getPermissionTree);

const route = useRoute();
const router = useRouter();

const onSave = async () => {
  formRef.value?.validate().then(async () => {
    const [id, body] = [
      typeof route.params.id === 'string' && route.params.id,
      {
        ...formState,
        permissions: formState.permissions.filter((item) => item.includes(':')),
      },
    ];

    const onSuccess = () => {
      api().adminRole.getMany.invalidate();
      const currentPath = router.currentRoute.value.path;
      const parentPath = currentPath.slice(0, currentPath.lastIndexOf('/'));
      notification.success({
        message: $t('message.createSuccess'),
      });
      router.push({
        path: parentPath,
      });
    };

    if (id) {
      api().adminRole.update({
        id,
        body,
        opt: {
          onSuccess,
        },
      });
    } else {
      api().adminRole.create({ body, opt: { onSuccess } });
    }
  });
};

const { required, notSpecialCharacters } = useValidator();
const rules = {
  name: [required(formInfos.value.name.label), notSpecialCharacters()],
  permissions: [required(formInfos.value.permissions.label)],
};

onBeforeMount(() => {
  if (route.params.id) {
    api().adminRole.get({
      id: route.params.id as string,
      opt: {
        onSuccess: ({ data }) => {
          Object.assign(formState, {
            name: data.name,
            description: data.description,
            permissions: data.permissions.map((item: Permission) => item.scope),
          });
        },
      },
    });
  }
});
</script>

<template>
  <Page
    :title="
      $t(route.params.id ? 'common.edit' : 'common.create', {
        name: $t('page.admin.roles.title'),
      })
    "
  >
    <template #appendHeader>
      <Space>
        <Button type="primary" @click="onSave">{{ $t('common.save') }}</Button>
        <Button @click="() => formRef?.resetFields()">
          {{ $t('common.reset') }}
        </Button>
      </Space>
    </template>
    <Form
      ref="formRef"
      :model="formState"
      :rules="rules"
      class="pt-2"
      layout="vertical"
    >
      <FormItem :label="formInfos.name.label" name="name">
        <Input
          v-model:value="formState.name"
          :placeholder="formInfos.name.placeholder"
        />
      </FormItem>
      <FormItem :label="formInfos.description.label" name="description">
        <Textarea
          v-model:value="formState.description"
          :placeholder="formInfos.description.placeholder"
        />
      </FormItem>
      <FormItem :label="formInfos.permissions.label" name="permissions">
        <Tree
          v-model:checked-keys="formState.permissions"
          :tree-data="treeData"
          auto-expand-parent
          checkable
        />
      </FormItem>
    </Form>
  </Page>
</template>
