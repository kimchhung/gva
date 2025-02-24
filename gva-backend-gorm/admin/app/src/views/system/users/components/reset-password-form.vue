<script setup lang="ts">
import type { Admin } from '#/api/admin/types';

import { computed, reactive, ref } from 'vue';

import { $t } from '@vben/locales';

import {
  Form,
  type FormInstance,
  FormItem,
  InputPassword,
  notification,
} from 'ant-design-vue';

import { api } from '#/api';
import { useValidator } from '#/hooks/use-validator';
import { getFormInfos } from '#/utils/form/label';
import BaseModal from '#/views/_core/base/base-modal.vue';

const props = defineProps<{
  open: boolean;
  record: Admin | null;
}>();

const emit = defineEmits(['update:open']);

const defaultFormState = () => ({
  password: '',
});

const formRef = ref<FormInstance>();
const formState = reactive(defaultFormState());

const loading = ref(false);

const onSuccess = () => {
  notification.success({
    message: $t(
      props.record ? 'message.updateSuccess' : 'message.createSuccess',
    ),
  });
  api().admin.getMany.invalidate();
  emit('update:open', false);
  formRef.value?.resetFields();
};

const { required, lengthRange } = useValidator();

const onSubmit = async () => {
  if (!formRef.value || !props.record) return;
  const isValid = await formRef.value.validate();
  if (!isValid) return;

  loading.value = true;
  const payload = {
    password: formState.password,
  };

  await api().admin.updatePartial({
    id: props.record.id,
    body: payload,
    opt: {
      onSuccess,
      loading,
    },
  });
};

const formInfos = computed(() =>
  getFormInfos({
    password: $t('common.password'),
  }),
);
</script>

<template>
  <BaseModal
    :ok-button-props="{
      loading,
      disabled: !formState.password,
    }"
    :open="open"
    :title="$t('page.admin.users.form.reset-password')"
    mode="edit"
    @cancel="() => emit('update:open', false)"
    @ok="onSubmit"
  >
    <Form ref="formRef" :model="formState" class="pt-2" layout="vertical">
      <FormItem
        :label="formInfos.password.label"
        :rules="[required(formInfos.password.label), lengthRange(6, 20)]"
        name="password"
      >
        <InputPassword
          v-model:value="formState.password"
          :placeholder="formInfos.password.placeholder"
          autocomplete="off"
        />
      </FormItem>
    </Form>
  </BaseModal>
</template>
