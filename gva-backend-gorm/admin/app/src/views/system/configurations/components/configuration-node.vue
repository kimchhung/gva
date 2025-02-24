<script setup lang="ts">
import type { Configuration } from '#/api/configuration/types';

import { useModel } from 'vue';

import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import {
  Collapse,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Tag,
} from 'ant-design-vue';

import { CONFIGURATION_PERMISSION } from '#/constants';
import { getConfigTitle } from '#/utils/helper/configuration';
import { copyToClipboard } from '#/utils/helper/copy';
import { withSuper } from '#/utils/helper/permissions';
import GroupButton from '#/views/_core/button/group-button.vue';

const props = defineProps<{
  activeKey: number[];
  list: Configuration[];
}>();

defineEmits(['edit', 'delete']);

const collapseKey = useModel(props, 'activeKey');
</script>

<template>
  <Collapse v-model:active-key="collapseKey" collapsible="icon">
    <Collapse.Panel v-for="config in list" :key="config.id">
      <template #header>
        <div class="flex items-center justify-between">
          {{ getConfigTitle(config) }}
          <Popconfirm
            v-if="!config.children || config.children.length === 0"
            :ok-text="$t('common.confirm')"
            :title="$t('common.confirmation')"
            placement="topRight"
            @confirm="$emit('delete', config.id)"
          >
            <button
              class="text-red-500"
              danger
              v-permissions="withSuper(CONFIGURATION_PERMISSION.DELETE)"
            >
              <VbenIcon
                class="text-destructive text-lg"
                icon="lucide:trash-2"
              />
            </button>
          </Popconfirm>
        </div>
      </template>
      <Form horizontal>
        <FormItem
          v-for="child in config.children.filter(
            (child) => child.type !== 'group',
          )"
          :key="child.key"
          :html-for="child.description"
        >
          <template #label>
            <Tag class="leading-7" color="blue">
              {{ child.key }}
              <button @click="() => copyToClipboard(child.key)">
                <VbenIcon icon="carbon:copy" />
              </button>
            </Tag>
          </template>

          <div class="flex gap-4">
            <div class="flex-1">
              <Input :value="JSON.stringify(child.value)" disabled />
            </div>
            <div>
              <GroupButton
                :buttons="[
                  {
                    vPermissions: withSuper(CONFIGURATION_PERMISSION.EDIT),
                    actionType: 'edit',
                    onClick: () => $emit('edit', child),
                  },
                  {
                    vPermissions: withSuper(CONFIGURATION_PERMISSION.DELETE),
                    actionType: 'delete',
                    onClick: () => $emit('delete', child.id),
                  },
                ]"
              />
            </div>
          </div>
        </FormItem>
      </Form>
      <configuration-node
        v-if="config.children.some((child) => child.type === 'group')"
        v-model:active-key="collapseKey"
        :list="config.children.filter((child) => child.type === 'group')"
        @delete="$emit('delete', $event)"
        @edit="$emit('edit', $event)"
      />
    </Collapse.Panel>
  </Collapse>
</template>
