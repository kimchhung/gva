<script setup lang="ts" generic="T">
import type { TableTool } from '#/utils/table/use-table-data';

import { VbenIcon } from '@vben-core/shadcn-ui';

import { Button, Flex, Popover } from 'ant-design-vue';

defineOptions({ name: 'TableDisplay' });

const props = defineProps<{ tableTool: TableTool<T> }>();

const isSelected = (k: string) => {
  return props.tableTool.selectedViewColumns.includes(k);
};

const toggleSelect = (k: string) => {
  if (isSelected(k)) {
    props.tableTool.setSelectedViewColumns(
      props.tableTool.selectedViewColumns.filter((key) => key !== k),
    );
  } else {
    const lists = [...props.tableTool.selectedViewColumns, k];
    props.tableTool.setSelectedViewColumns(lists);
  }
};
</script>

<template>
  <div class="flex">
    <Popover
      class="flex items-center gap-2"
      placement="bottomLeft"
      trigger="click"
    >
      <Button class="px-2">
        <template #icon>
          <VbenIcon icon="lucide:sliders-horizontal" />
        </template>
        {{ $t('common.display') }}
      </Button>
      <template #content>
        <Flex gap="small" vertical>
          <template
            v-for="opt in props.tableTool.viewColumnOptions"
            :key="opt.value"
          >
            <Button
              class="flex items-center justify-between"
              @click="() => toggleSelect(String(opt.value))"
            >
              <span>
                {{ opt.label }}
              </span>
              <span class="ml-1 w-2">
                <VbenIcon
                  v-if="isSelected(String(opt.value))"
                  class="text-primary"
                  icon="lucide:check"
                />
              </span>
            </Button>
          </template>
        </Flex>
      </template>
    </Popover>
  </div>
</template>
