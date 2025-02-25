<script setup lang="ts">
import type { Configuration } from '#/api/configuration/types';

import { computed, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';

import { Button, notification } from 'ant-design-vue';

import { api } from '#/api';
import { CONFIGURATION_PERMISSION } from '#/constants';
import { $t } from '#/locales';
import { mapConfigTree } from '#/utils/helper/configuration';
import { withSuper } from '#/utils/helper/permissions';
import { defaultQuery } from '#/utils/pagi/query';
import ActionButton from '#/views/_core/button/action-button.vue';

import ConfigurationForm from './components/configuration-form.vue';
import ConfigurationNode from './components/configuration-node.vue';

const openSheet = ref<boolean>(false);

const editData = ref<Configuration | null>(null);

const { data: configurationList } = api().configuration.getMany.useQuery({
  query: () => defaultQuery({ limit: 500, selects: ['list'] }),
});

const tree = computed(() => {
  return mapConfigTree(configurationList.value?.data ?? []);
});

const groups = computed(() => {
  return configurationList.value?.data.filter((item) => item.type === 'group');
});

const collapseKey = ref<number[]>([]);

const handleEdit = (configuration: Configuration) => {
  editData.value = configuration;
  openSheet.value = true;
};

const handleAdd = () => {
  editData.value = null;
  openSheet.value = true;
};

const toggleCollapseAll = () => {
  collapseKey.value =
    collapseKey.value.length <= 0
      ? (groups.value?.map((item) => item.id) ?? [])
      : [];
};

const handleDelete = async (id: number) => {
  await api().configuration.delete({
    id,
    opt: {
      onSuccess: () => {
        api().configuration.getMany.invalidate();
        notification.success({
          message: $t('message.deleteSuccess'),
        });
      },
    },
  });
};

watch(
  tree,
  () => {
    if (tree.value) {
      collapseKey.value =
        tree.value.length > 0 ? tree.value.map((item) => item.id) : [];
    }
  },
  {
    once: true,
  },
);
</script>

<template>
  <Page :full-height="false" :title="$t('page.system.configurations.title')">
    <template #appendHeader>
      <ActionButton
        :value="
          $t('common.add', { name: $t('page.system.configurations.singular') })
        "
        action-type="create"
        v-permissions="withSuper(CONFIGURATION_PERMISSION.ADD)"
        @click="handleAdd"
      />
    </template>
    <ConfigurationForm
      v-model:open="openSheet"
      :groups="groups || []"
      :initial-data="editData"
    />
    <div class="mb-4">
      <Button type="primary" @click="toggleCollapseAll">
        {{
          collapseKey.length === 0
            ? $t('common.expandAll')
            : $t('common.collapseAll')
        }}
      </Button>
    </div>
    <ConfigurationNode
      v-model:active-key="collapseKey"
      :list="tree"
      @delete="handleDelete"
      @edit="handleEdit"
    />
  </Page>
</template>
