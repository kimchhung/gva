<script setup lang="ts">
import type { Configuration } from '#/api/configuration/types';

import { computed, reactive, ref, watch } from 'vue';

import { VbenIcon } from '@vben-core/shadcn-ui';
import { preferences } from '@vben/preferences';

import {
  Button,
  Col,
  DatePicker,
  Drawer,
  Form,
  type FormInstance,
  FormItem,
  Input,
  InputNumber,
  notification,
  RangePicker,
  Row,
  Space,
  Switch,
  Textarea,
  TreeSelect,
} from 'ant-design-vue';
import dayjs from 'dayjs';

import { useValidator } from '#/hooks/use-validator';
import { $t } from '#/locales';
import { getFormInfos } from '#/utils/form/label';
import ImageUploader from '#/views/_core/base/image-uploader.vue';
import ObjectEditor from '#/views/_core/base/object-editor.vue';

import { api } from '#/api';
import Block from './block.vue';

const props = defineProps<{
  groups: Configuration[];
  initialData?: Configuration | null;
  open: boolean;
}>();

const emit = defineEmits(['update:open']);

const groupOptions = computed(() =>
  props.groups.map((item) => ({
    id: item.id,
    label: item.key,
    value: item.id,
    metadata: item.metadata,
    pId: item.parentId,
    rootId: item.rootId,
  })),
);

const dataTypeOptions = computed(() => [
  {
    label: $t('common.dataTypes.group'),
    value: 'group',
    icon: 'lucide:folder',
  },
  {
    label: $t('common.dataTypes.string'),
    value: 'string',
    icon: 'lucide:text',
  },
  {
    label: $t('common.dataTypes.textarea'),
    value: 'textarea',
    icon: 'lucide:file-text',
  },
  { label: $t('common.dataTypes.number'), value: 'int', icon: 'lucide:hash' },
  {
    label: $t('common.dataTypes.boolean'),
    value: 'bool',
    icon: 'lucide:square-check',
  },
  {
    label: $t('common.dataTypes.array'),
    value: 'array',
    icon: 'lucide:layers',
  },
  { label: $t('common.dataTypes.json'), value: 'json', icon: 'lucide:code' },
  {
    label: $t('common.dataTypes.date'),
    value: 'date',
    icon: 'lucide:calendar',
  },
  {
    label: $t('common.dataTypes.dateRange'),
    value: 'dateRange',
    icon: 'lucide:calendar-range',
  },
  { label: $t('common.dataTypes.image'), value: 'image', icon: 'lucide:image' },
  {
    label: $t('common.dataTypes.multiImage'),
    value: 'multiImage',
    icon: 'lucide:images',
  },
  {
    label: $t('common.dataTypes.object'),
    value: 'object',
    icon: 'lucide:octagon',
  },
]);

const editMode = computed(() => !!props.initialData);
const loading = ref<boolean>(false);

const defaultFormState = () => ({
  key: '',
  value: '' as any,
  date: '' as unknown as dayjs.Dayjs,
  dateRange: ['' as unknown, '' as unknown] as [dayjs.Dayjs, dayjs.Dayjs],
  description: '',
  metadata: {} as Record<string, any>,
  parentId: undefined as number | undefined,
  rootId: undefined as number | undefined,
  type: 'json',
  array: [''] as string[],
  json: [{ key: '', value: '' }] as { key: string; value: string }[],
  image: '',
  multiImage: [{ imageUrl: '', redirectUrl: '' }] as {
    imageUrl: string;
    redirectUrl: string;
  }[],
  object: {
    key: 'value',
  } as Record<string, any>,
});

const formRef = ref<FormInstance>();
const formState = reactive(defaultFormState());

const formInfos = computed(() =>
  getFormInfos({
    key: $t('common.key'),
    value: $t('common.value'),
    metadata: $t('common.metadata'),
    description: $t('common.description'),
    parentId: $t('common.group'),
    group: $t('common.group'),
    type: $t('common.type'),
    json: $t('common.json'),
    array: $t('common.array'),
    date: $t('common.date'),
    dateRange: $t('common.dateRange'),
  }),
);

const { required } = useValidator();
const rules = computed(() => ({
  key: [required(formInfos.value.key.label)],
  ...(formState.type === 'group'
    ? {}
    : {
        parentId: [required(formInfos.value.value.label)],
      }),
  metadata: {
    required: true,
  },
  type: [required(formInfos.value.type.label)],
  value: [required(formInfos.value.value.label)],
  json: {
    required: true,
    each: {
      key: [required(formInfos.value.json.label)],
      value: [required(formInfos.value.json.label)],
    },
  },
  object: {
    required: true,
  },
  array: [required(formInfos.value.array.label)],
  date: [required(formInfos.value.date.label)],
  dateRange: [required(formInfos.value.dateRange.label)],
}));

const resetFormState = () => {
  formState.key = '';
  formState.value = '';
  formState.metadata = {
    key: 'value',
  };
  formState.date = '' as unknown as dayjs.Dayjs;
  formState.dateRange = [
    '' as unknown as dayjs.Dayjs,
    '' as unknown as dayjs.Dayjs,
  ];
  formState.description = '';
  formState.parentId = undefined;
  formState.type = 'json';
  formState.array = [''];
  formState.object = {
    key: 'value',
  } as Record<string, any>;
  formState.json = [{ key: '', value: '' }];
  formState.image = '';
  formState.multiImage = [{ imageUrl: '', redirectUrl: '' }];
};

const populateFormState = (val: Configuration | null) => {
  if (!val) {
    resetFormState();
    return;
  }

  formState.key = val.key;
  formState.description = val.description;
  formState.type = val.type;
  formState.metadata = val.metadata ?? {
    key: 'value',
  };
  formState.parentId = val.parentId;
  formState.rootId = val.rootId ?? val.parentId;

  switch (val.type) {
    case 'array': {
      formState.array = JSON.parse(val.value);
      break;
    }
    case 'date': {
      formState.value = dayjs(val.value, 'HH:mm:ss');
      break;
    }
    case 'dateRange': {
      const [start, end] = val.value.split(' - ');
      formState.dateRange = [dayjs(start), dayjs(end)];
      break;
    }
    case 'image': {
      formState.image = val.value;
      break;
    }
    case 'int': {
      formState.value = Number.parseInt(val.value, 10);
      break;
    }
    case 'json': {
      formState.json = JSON.parse(val.value);
      break;
    }
    case 'multiImage': {
      formState.multiImage = JSON.parse(val.value);
      break;
    }
    case 'number': {
      formState.value = Number(val.value);
      break;
    }
    case 'object': {
      formState.object = val.value ?? {
        key: 'value',
      };
      break;
    }
    default: {
      formState.value = val.value;
    }
  }
};

const handleAddArray = () => formState.array.push('');
const handleDeleteArray = (index: number) =>
  formState.array.length > 1 && formState.array.splice(index, 1);

const handleAddJSON = () => formState.json.push({ key: '', value: '' });
const handleDeleteJSON = (index: number) =>
  formState.json.length > 1 && formState.json.splice(index, 1);

const onSubmit = async () => {
  if (!formRef.value) return;
  await formRef.value.validate();

  const payload: Partial<Configuration> = {
    key: formState.key,
    description: formState.description,
    type: formState.type,
    metadata: formState.metadata,
    parentId: formState.parentId,
    rootId: formState.rootId,
  };

  if (formState.type !== 'group') {
    switch (formState.type) {
      case 'array': {
        payload.value = JSON.stringify(formState.array);
        break;
      }
      case 'date': {
        payload.value = dayjs(formState.date).format('HH:mm:ss');
        break;
      }
      case 'dateRange': {
        payload.value = `${dayjs(formState.dateRange[0]).format('YYYY-MM-DD')} - ${dayjs(formState.dateRange[1]).format('YYYY-MM-DD')}`;
        break;
      }
      case 'image': {
        payload.value = formState.image;
        break;
      }
      case 'json': {
        payload.value = JSON.stringify(formState.json);
        break;
      }
      case 'multiImage': {
        payload.value = JSON.stringify(formState.multiImage);
        break;
      }
      case 'object': {
        payload.value = formState.object;
        break;
      }

      default: {
        payload.value = formState.value;
      }
    }
  }

  await (editMode.value
    ? api().configuration.update({
        id: props.initialData?.id as number,
        body: payload,
        opt: {
          loading,
          onSuccess: () => {
            loading.value = false;
            api().configuration.getMany.invalidate();
            notification.success({
              message: $t('message.updateSuccess'),
            });
            emit('update:open', false);
            resetFormState();
            formRef.value?.resetFields();
          },
        },
      })
    : api().configuration.create({
        body: payload,
        opt: {
          loading,
          onSuccess: () => {
            api().configuration.getMany.invalidate();
            notification.success({
              message: $t('message.createSuccess'),
            });
            emit('update:open', false);
            resetFormState();
            formRef.value?.resetFields();
          },
        },
      }));
};

watch(
  () => props.initialData,
  () => populateFormState(props.initialData ?? null),
  { immediate: true },
);

const handleSelect = (dataType: { label: string; value: string }) => {
  formState.type = dataType.value;
};

const handleCancel = () => {
  emit('update:open', false);
};

const getLabel = (config: (typeof groupOptions.value)[0]): string => {
  if (config.metadata?.labelEn || config.metadata?.labelZh) {
    return preferences.app.locale === 'en-US'
      ? config.metadata?.labelEn
      : config.metadata?.labelZh;
  }

  return config.label;
};
</script>

<template>
  <Drawer
    :closable="false"
    :destroy-on-close="true"
    :mask-closable="false"
    :open="open"
    :width="400"
    close-icon
    @close="emit('update:open', false)"
  >
    <template #title>
      <div class="flex items-center justify-between">
        <div>
          {{ editMode ? $t('common.edit') : $t('common.create') }}
          {{ $t('page.system.configurations.singular') }}
        </div>
        <VbenIcon
          class="cursor-pointer"
          icon="lucide:x"
          @click="emit('update:open', false)"
        />
      </div>
    </template>
    <Form ref="formRef" :model="formState" :rules="rules" layout="vertical">
      <Block :title="$t('common.key')">
        <FormItem class="mb-2" name="key">
          <Input v-model:value="formState.key" />
        </FormItem>
      </Block>
      <Block :title="$t('common.description')">
        <FormItem class="mb-2" name="description">
          <Input v-model:value="formState.description" />
        </FormItem>
      </Block>
      <Block :title="$t('common.metadata')">
        <ObjectEditor v-model="formState.metadata" />
      </Block>
      <Block :title="$t('common.dataTypes.title')">
        <div class="flex w-full flex-wrap justify-between">
          <template v-for="dataType in dataTypeOptions" :key="dataType.value">
            <div
              class="flex h-20 w-20 cursor-pointer flex-col"
              @click="handleSelect(dataType)"
            >
              <div
                :class="{
                  'outline-box-active': dataType.value === formState.type,
                }"
                class="outline-box flex-center group cursor-pointer"
              >
                <div class="py-1.5">
                  <div
                    class="flex-center relative size-5 items-center justify-center rounded-sm align-middle"
                  >
                    <VbenIcon
                      :class="{
                        'text-primary opacity-100':
                          dataType.value === formState.type,
                        'opacity-60': dataType.value !== formState.type,
                      }"
                      :icon="dataType.icon"
                      class="absolute z-10 size-5 group-hover:opacity-100"
                    />
                  </div>
                </div>
              </div>
              <div class="text-muted-foreground my-2 text-center text-xs">
                {{ dataType.label }}
              </div>
            </div>
          </template>
        </div>
      </Block>
      <Block :title="$t('common.group')">
        <FormItem class="mb-2" name="parentId">
          <TreeSelect
            v-model:value="formState.parentId"
            :placeholder="$t('common.group')"
            :tree-data="groupOptions"
            tree-data-simple-mode
            tree-default-expand-all
            @change="
              (value) => {
                const _item = groupOptions.find((item) => item.value === value);
                if (_item) {
                  formState.rootId = _item.rootId || _item.pId || _item.value;
                }
              }
            "
          >
            <template #title="item">
              {{ getLabel(item) }}
            </template>
          </TreeSelect>
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'int'" :title="$t('common.value')">
        <FormItem class="mb-2" name="value">
          <InputNumber v-model:value="formState.value" class="w-full" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'string'" :title="$t('common.value')">
        <FormItem class="mb-2" name="value">
          <Input v-model:value="formState.value" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'textarea'" :title="$t('common.value')">
        <FormItem class="mb-2" name="value">
          <Textarea v-model:value="formState.value" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'array'" :title="$t('common.value')">
        <div
          v-for="(_, idx) in formState.array"
          :key="`arr-value-${idx}`"
          class="py-2"
        >
          <Row align="middle" justify="space-between">
            <Col span="20">
              <FormItem :name="`array.${idx}`" class="mb-2">
                <Input
                  v-model:value="formState.array[idx]"
                  :placeholder="`${$t('common.value')} ${idx + 1}`"
                />
              </FormItem>
            </Col>
            <Button
              :disabled="formState.array.length === 1"
              danger
              type="dashed"
              @click="() => handleDeleteArray(idx)"
            >
              <VbenIcon icon="lucide:trash" />
            </Button>
          </Row>
        </div>
        <Button block class="mt-2" type="dashed" @click="handleAddArray">
          Add Key-Value
        </Button>
      </Block>
      <Block v-if="formState.type === 'bool'" :title="$t('common.value')">
        <FormItem class="mb-2">
          <Switch v-model:checked="formState.value" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'date'" :title="$t('common.value')">
        <FormItem class="mb-2" name="date">
          <DatePicker v-model:value="formState.date" class="w-full" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'dateRange'" :title="$t('common.value')">
        <FormItem class="mb-2" name="dateRange">
          <RangePicker v-model:value="formState.dateRange" class="w-full" />
        </FormItem>
      </Block>
      <Block v-if="formState.type === 'json'" :title="$t('common.value')">
        <div
          v-for="(item, idx) in formState.json"
          :key="`json-${idx}`"
          class="mb-2 last:mb-4"
        >
          <Row align="top" justify="space-between">
            <Col :span="10">
              <FormItem :name="`json.${idx}.key`" class="mb-2">
                <Input
                  v-model:value="item.key"
                  :placeholder="`${$t('common.key')} ${idx + 1}`"
                />
              </FormItem>
            </Col>
            <Col :span="10">
              <FormItem :name="`json.${idx}.value`" class="mb-2">
                <Input
                  v-model:value="item.value"
                  :placeholder="`${$t('common.value')} ${idx + 1}`"
                />
              </FormItem>
            </Col>
            <Col>
              <Button
                :disabled="formState.json.length === 1"
                danger
                type="dashed"
                @click="() => handleDeleteJSON(idx)"
              >
                <VbenIcon icon="lucide:trash" />
              </Button>
            </Col>
          </Row>
        </div>

        <Button block type="dashed" @click="handleAddJSON">
          Add Key-Value
        </Button>
      </Block>

      <Block v-if="formState.type === 'image'" :title="$t('common.value')">
        <FormItem class="mb-2" name="value">
          <ImageUploader v-model="formState.image" />
        </FormItem>
      </Block>

      <Block v-if="formState.type === 'multiImage'" :title="$t('common.value')">
        <div
          v-for="(item, idx) in formState.multiImage"
          :key="`multi-image-${idx}`"
          class="mb-2 last:mb-4"
        >
          <FormItem class="mb-2" name="imageUrl">
            <ImageUploader
              v-model="item.imageUrl"
              :placeholder="`${$t('common.image')} ${idx + 1}`"
            />
          </FormItem>

          <Row align="top" justify="space-between">
            <Col :span="20">
              <FormItem class="mb-2" name="redirectUrl">
                <Input
                  v-model:value="item.redirectUrl"
                  :placeholder="`${$t('common.redirectUrl')} ${idx + 1}`"
                />
              </FormItem>
            </Col>
            <Col>
              <Button
                :disabled="formState.multiImage.length === 1"
                danger
                type="dashed"
                @click="() => formState.multiImage.splice(idx, 1)"
              >
                <VbenIcon icon="lucide:trash" />
              </Button>
            </Col>
          </Row>
        </div>
        <Button
          block
          type="dashed"
          @click="
            () => formState.multiImage.push({ imageUrl: '', redirectUrl: '' })
          "
        >
          {{ $t('common.addImage') }}
        </Button>
      </Block>
      <Block v-if="formState.type === 'object'" :title="$t('common.value')">
        <ObjectEditor v-model="formState.object" />
      </Block>
    </Form>

    <template #footer>
      <Space>
        <Button class="w-full" @click="handleCancel">
          {{ $t('common.cancel') }}
        </Button>
        <Button class="w-full" type="primary" @click="onSubmit">
          {{ editMode ? $t('common.save') : $t('common.create') }}
        </Button>
      </Space>
    </template>
  </Drawer>
</template>

<style>
.ant-drawer-body,
.ant-drawer-header {
  padding: 16px !important;
}
</style>
