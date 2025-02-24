<script setup lang="ts">
import type { Admin, AdminSetTOTPResponse } from '#/api/admin/types';

import { computed, reactive, ref } from 'vue';

import { Copy } from '@vben/icons';
import { $t } from '@vben/locales';

import {
  Alert,
  Button,
  Form,
  type FormInstance,
  FormItem,
  Input,
  notification,
  QRCode,
} from 'ant-design-vue';

import { useValidator } from '#/hooks/use-validator';
import { getFormInfos } from '#/utils/form/label';
import { copyToClipboard } from '#/utils/helper/copy';
import BaseModal from '#/views/_core/base/base-modal.vue';

const props = defineProps<{
  open: boolean;
  record: Admin | null;
}>();

const emit = defineEmits(['update:open']);

const defaultFormState = () => ({
  totp: '',
});

enum STEP {
  // eslint-disable-next-line no-unused-vars
  TOTP_UPDATED,
  // eslint-disable-next-line no-unused-vars
  TOTP_VALIDATION,
}

const formRef = ref<FormInstance>();
const formState = reactive(defaultFormState());
const step = ref(STEP.TOTP_VALIDATION);

const setTOTPResponse = ref<AdminSetTOTPResponse | null>();

const loading = ref(false);

const { required } = useValidator();

const onSubmit = async () => {
  if (!formRef.value || !props.record) return;
  const isValid = await formRef.value.validate();
  if (!isValid) return;

  loading.value = true;
  const payload = {
    totp: formState.totp,
  };

  await api.admin.setTOTP({
    id: props.record.id,
    body: payload,
    opt: {
      onSuccess: (res) => {
        notification.success({
          message: $t('message.updateSuccess'),
        });
        formRef.value?.resetFields();
        setTOTPResponse.value = res.data;
        step.value = STEP.TOTP_UPDATED;
      },
      loading,
    },
  });
};

const formInfos = computed(() =>
  getFormInfos({
    totp: $t('page.admin.users.form.totp'),
  }),
);

const qrcodeCanvasRef = ref();

const downloadQR = async () => {
  const url = await qrcodeCanvasRef.value.toDataURL();
  const a = document.createElement('a');
  a.download = 'QRCode.png';
  a.href = url;
  document.body.append(a);
  a.click();
  a.remove();
};

const handleCopy = () => {
  if (setTOTPResponse.value?.totpKey) {
    copyToClipboard(setTOTPResponse.value?.totpKey);
  }
};
</script>

<template>
  <BaseModal
    :ok-button-props="{
      loading,
      disabled: !formState.totp,
    }"
    :ok-text="$t('common.edit')"
    :open="open"
    :title="$t('page.admin.users.form.totp')"
    @cancel="
      () => {
        formRef?.resetFields();
        emit('update:open', false);
        step = STEP.TOTP_VALIDATION;
      }
    "
    @ok="onSubmit"
    v-bind="{
      ...(step === STEP.TOTP_UPDATED && {
        footer: null,
      }),
    }"
  >
    <Form
      v-if="step === STEP.TOTP_VALIDATION"
      ref="formRef"
      :model="formState"
      layout="vertical"
    >
      <Alert
        :message="$t('page.admin.users.form.totpAlert')"
        class="mb-4"
        type="info"
      />
      <FormItem :rules="[required(formInfos.totp.label)]" name="totp">
        <Input
          v-model:value="formState.totp"
          :placeholder="$t('page.admin.users.form.totp')"
          autocomplete="off"
          @press-enter="onSubmit"
        />
      </FormItem>
    </Form>
    <div v-else class="flex justify-center">
      <div class="flex flex-col items-center">
        <Alert
          :message="$t('page.admin.users.form.totpUpdated')"
          type="success"
        />
        <div class="my-4">
          <QRCode
            ref="qrcodeCanvasRef"
            :value="setTOTPResponse?.totpURL"
            class="mt-4"
            color="black"
          />
        </div>
        <div class="mb-4 flex items-center justify-center">
          <p class="text-base">{{ setTOTPResponse?.totpKey }}</p>
          <Button class="ml-2" @click="handleCopy">
            <Copy class="size-4" />
          </Button>
        </div>
        <Button @click="downloadQR">
          {{ $t('page.admin.users.form.downloadQR') }}
        </Button>
      </div>
    </div>
  </BaseModal>
</template>
