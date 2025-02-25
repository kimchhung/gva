<script setup lang="ts">
import type { Admin } from '#/api/admin/types';

import { reactive } from 'vue';

import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import { Button, Form, FormItem, Input } from 'ant-design-vue';

import { api } from '#/api';
import { useBaseModalForm } from '#/hooks/use-base-modal-form';
import { useValidator } from '#/hooks/use-validator';
import BaseModal from '#/views/_core/base/base-modal.vue';

const props = defineProps<{
  open: boolean;
  record: Admin | null;
}>();

const defaultFormState = () => ({
  ipWhiteList: [''] as string[],
});

const formState = reactive(defaultFormState());

const {
  formRef,
  openRef,
  handleSubmit,
  loadingRef,
  closeModal,
  getFormDirtyState,
} = useBaseModalForm(props, {
  onOpen: ({ record }) => {
    if (record) {
      formState.ipWhiteList = record.ipWhiteList || [];
      if (formState.ipWhiteList.length === 0) {
        formState.ipWhiteList = [''];
      }
    }
    return formState;
  },
});

const { ip } = useValidator();

const onSubmit = handleSubmit(async () => {
  if (!props.record) return;
  const payload = {
    ipWhiteList: formState.ipWhiteList.filter(Boolean),
  };

  await api().admin.updatePartial({
    id: props.record.id,
    body: payload,
    opt: {
      onSuccess: () => {
        api().admin.getMany.invalidate();
        closeModal({
          showSuccessMessage: true,
        });
      },
      loading: loadingRef,
    },
  });
});

const onAddIp = () => {
  formState.ipWhiteList = [...formState.ipWhiteList, ''];
};

const onRemoveIp = (index: number) => {
  formState.ipWhiteList =
    formState.ipWhiteList.length === 1
      ? ['']
      : formState.ipWhiteList.filter((_, i) => i !== index);
};
</script>

<template>
  <BaseModal
    v-model:open="openRef"
    :entity="$t('page.admin.users.form.ip-whitelist')"
    :ok-button-props="{
      loading: loadingRef,
      disabled: !getFormDirtyState({ ...formState }),
    }"
    mode="edit"
    @ok="onSubmit"
  >
    <Form ref="formRef" :model="formState" class="pt-2">
      <template v-for="(_, index) in formState.ipWhiteList" :key="index">
        <FormItem
          :label="`IP ${index + 1}`"
          :name="['ipWhiteList', index]"
          :rules="[ip()]"
        >
          <div class="mb-2 space-y-2">
            <div class="flex space-x-2">
              <Input
                v-model:value="formState.ipWhiteList[index]"
                class="w-full"
                placeholder="0.0.0.0"
              />
              <Button
                class="flex items-center justify-center"
                type="ghost"
                @click="onRemoveIp(index)"
              >
                <template #icon>
                  <VbenIcon
                    class="text-destructive text-lg"
                    icon="lucide:trash-2"
                  />
                </template>
              </Button>
            </div>
          </div>
        </FormItem>
      </template>
      <Button :disabled="formState.ipWhiteList.length === 5" @click="onAddIp">
        {{
          $t('common.add', {
            name: $t('page.admin.users.form.ip-whitelist'),
          })
        }}
      </Button>
    </Form>
  </BaseModal>
</template>
