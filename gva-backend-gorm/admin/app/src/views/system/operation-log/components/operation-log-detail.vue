<script setup lang="ts">
import type { OperationLog } from '#/api/operation-log/types';

import VueJsonPretty from 'vue-json-pretty';

import { $t } from '@vben/locales';

import { Modal, Tag, Textarea } from 'ant-design-vue';

import { METHODS_COLOR } from '#/constants/operation-log';
import { tableColumns } from '#/utils/table/column';

import 'vue-json-pretty/lib/styles.css';

defineProps<{
  open: boolean;
  record: null | OperationLog;
}>();

const emit = defineEmits(['update:open']);
</script>

<template>
  <Modal
    :footer="null"
    :open="open"
    :title="$t('common.detail')"
    :width="700"
    @cancel="() => emit('update:open', false)"
  >
    <div
      v-if="record"
      class="info grid"
      style="grid-template-columns: 100px 3fr"
    >
      <div class="p-2">
        <p class="font-bold">{{ $t('page.system.operationLog.method') }}:</p>
      </div>
      <div class="p-2">
        <Tag :color="METHODS_COLOR[record.method]">{{ record.method }}</Tag>
      </div>
      <div class="p-2">
        <p class="font-bold">{{ $t('page.system.operationLog.path') }}:</p>
      </div>
      <div class="p-2">
        <code>{{ record.path }}</code>
      </div>

      <template v-if="Object.keys(record.data).length > 0">
        <div class="p-2">
          <p class="font-bold">{{ $t('page.system.operationLog.params') }}:</p>
        </div>
        <div class="p-2">
          <VueJsonPretty :data="record.data ?? {}" />
        </div>
      </template>
      <div class="p-2">
        <p class="font-bold">
          {{ $t('page.system.operationLog.statusCode') }}:
        </p>
      </div>
      <div class="p-2">
        <Tag :color="record.code >= 0 ? 'green' : 'red'">{{ record.code }}</Tag>
      </div>
      <template v-if="record.msg">
        <div class="p-2">
          <p class="font-bold">{{ $t('page.system.operationLog.message') }}:</p>
        </div>
        <div class="p-2">
          <p>
            {{ record?.msg ?? tableColumns.nullText }}
          </p>
        </div>
      </template>
      <template v-if="record.error">
        <div class="p-2">
          <p class="font-bold">{{ $t('page.system.operationLog.error') }}:</p>
        </div>
        <div class="p-2">
          <Textarea
            :default-value="record.error"
            :rows="8"
            class="resize-none"
            readonly
          />
        </div>
      </template>
    </div>
  </Modal>
</template>

<style lang="scss">
.vjs-tree-node.is-highlight,
.vjs-tree-node:hover {
  @apply bg-slate-500/10 dark:bg-slate-100/10;
}
</style>
