<script lang="ts" setup>
import type { BreadcrumbStyleType } from '@vben/types';
import type { IBreadcrumb } from '@vben-core/shadcn-ui';

import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { $locale, $t } from '@vben/locales';
import { translateMetaTitle } from '@vben/utils';
import { VbenBackgroundBreadcrumb, VbenBreadcrumb } from '@vben-core/shadcn-ui';

interface Props {
  hideWhenOnlyOne?: boolean;
  showHome?: boolean;
  showIcon?: boolean;
  type?: BreadcrumbStyleType;
}

const props = withDefaults(defineProps<Props>(), {
  showHome: false,
  showIcon: false,
  type: 'normal',
});

const route = useRoute();
const router = useRouter();

const breadcrumbs = computed((): IBreadcrumb[] => {
  const matched = route.matched;

  const resultBreadcrumb: IBreadcrumb[] = [];

  for (const match of matched) {
    const { meta, path } = match;
    const { hideChildrenInMenu, hideInBreadcrumb, icon, title } = meta || {};
    if (hideInBreadcrumb || hideChildrenInMenu || !path) {
      continue;
    }

    const translatedMeta = translateMetaTitle(meta, { $t, $locale });

    resultBreadcrumb.push({
      icon,
      path: path || route.path,
      title: title ? translatedMeta.title : '',
    });
  }
  if (props.showHome) {
    resultBreadcrumb.unshift({
      icon: 'mdi:home-outline',
      isHome: true,
      path: '/',
    });
  }
  if (props.hideWhenOnlyOne && resultBreadcrumb.length === 1) {
    return [];
  }

  return resultBreadcrumb;
});

function handleSelect(path: string) {
  router.push(path);
}
</script>
<template>
  <VbenBreadcrumb
    v-if="type === 'normal'"
    :breadcrumbs="breadcrumbs"
    :show-icon="showIcon"
    class="ml-2"
    @select="handleSelect"
  />
  <VbenBackgroundBreadcrumb
    v-if="type === 'background'"
    :breadcrumbs="breadcrumbs"
    :show-icon="showIcon"
    class="ml-2"
    @select="handleSelect"
  />
</template>
