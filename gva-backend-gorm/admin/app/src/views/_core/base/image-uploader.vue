<script setup lang="ts">
import type { UploadListType } from 'ant-design-vue/es/upload/interface';
import type { UploadRequestOption } from 'ant-design-vue/es/vc-upload/interface';

import { onMounted, ref, watch } from 'vue';

import { VbenIcon } from '@vben-core/shadcn-ui';
import { $t } from '@vben/locales';

import { api } from '#/api';
import { Form, message, Upload, type UploadFile } from 'ant-design-vue';

interface Props {
  listType?: UploadListType;
  modelValue?: string;
  sizeLimit?: number; // size limit in MB
}

const props = withDefaults(defineProps<Props>(), {
  listType: 'picture-card',
  modelValue: '',
  sizeLimit: 0.5,
});

const emit = defineEmits(['update:modelValue']);

const formItemContext = Form.useInjectFormItemContext();

const triggerChange = (changedValue: string) => {
  emit('update:modelValue', changedValue);
  formItemContext.onFieldChange();
};

const uploadFiles = ref<UploadFile[]>([]);

const initialImage = () => {
  if (props.modelValue) {
    const newImage = new Date().toDateString();
    uploadFiles.value = [
      {
        status: 'done',
        uid: newImage,
        name: `${newImage}.png`,
        url: props.modelValue,
      },
    ];
  } else {
    uploadFiles.value = [];
  }
};

onMounted(() => {
  initialImage();
});

watch(
  () => props.modelValue,
  () => {
    initialImage();
  },
);

const loading = ref(false);

const handleUploadImage = async (opts: UploadRequestOption) => {
  if (opts.onSuccess) {
    const formData = new FormData();
    formData.append('file', opts.file);
    const [res] = await api().uploadImage(formData, {
      loading,
    });
    triggerChange(res?.data.url ?? '');
  }
};

const handleDeleteImage = () => {
  triggerChange('');
};

const beforeUpload = (file: File) => {
  const isLtSizeLimit = file.size / 1024 / 1024 < props.sizeLimit!;
  if (!isLtSizeLimit) {
    message.error(
      $t('validator.error.imageSizeExeeded', { size: props.sizeLimit }),
    );
  }
  return isLtSizeLimit;
};
</script>

<template>
  <Upload
    :before-upload="beforeUpload"
    :custom-request="handleUploadImage"
    :file-list="uploadFiles"
    :list-type="props.listType"
    :max-count="1"
    :multiple="false"
    accept=".png,.jpeg,.jpg,.gif,.webp"
    name="icon"
    @remove="handleDeleteImage"
  >
    <div class="flex flex-col items-center justify-center">
      <template v-if="loading">
        <VbenIcon icon="svg-spinners:90-ring" />
      </template>
      <template v-else>
        <VbenIcon icon="ri:upload-2-line" />
      </template>
      <div style="margin-top: 8px">
        {{ `${$t('common.upload')}${loading ? '...' : ''}` }}
      </div>
    </div>
  </Upload>
</template>
