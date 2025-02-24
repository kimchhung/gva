<script setup lang="ts">
import { useI18n } from '@vben/locales';

import cronstrue from 'cronstrue/i18n';

defineProps<{
  cron: string;
}>();

const { locale } = useI18n();
</script>

<template>
  <p>
    {{
      cron &&
      cronstrue.toString(cron, {
        locale: locale === 'zh-CN' ? 'zh_CN' : 'en',
      })
    }}
  </p>
</template>
