<script lang="ts" setup>
import { useRouter } from 'vue-router';

import { Fallback } from '@gva/common-ui';

import { Button } from 'ant-design-vue';

const router = useRouter();

function details() {
  router.push({ name: 'BreadcrumbLateralDetailDemo' });
}
</script>

<template>
  <Fallback
    description="Click to view details and observe the change of bread crumbs."
    status="coming-soon"
    title="Bread Dandruff navigation-level mode"
  >
    <template #action>
      <Button type="primary" @click="details">点击查看详情</Button>
    </template>
  </Fallback>
</template>
