<script lang="ts" setup>
import { computed } from 'vue';

import { useAntdDesignTokens } from '@vben/hooks';
import { preferences, usePreferences } from '@vben/preferences';
import { useAccessStore, useLockStore } from '@vben/stores';

import { App, ConfigProvider, theme } from 'ant-design-vue';

import { antdLocale } from '#/locales';

import useAFK from './hooks/use-afk';
import { useAuthStore } from './store';

import '@wangeditor/editor/dist/css/style.css';

defineOptions({ name: 'App' });

const { isDark } = usePreferences();
const { tokens } = useAntdDesignTokens();

const { logout } = useAuthStore();
const accessStore = useAccessStore();
const lockStore = useLockStore();

const tokenTheme = computed(() => {
  const algorithm = isDark.value
    ? [theme.darkAlgorithm]
    : [theme.defaultAlgorithm];

  // antd Compact mode algorithm
  if (preferences.app.compact) {
    algorithm.push(theme.compactAlgorithm);
  }

  return {
    algorithm,
    token: tokens,
  };
});

const durationToBeAFK = 1000 * 60 * 30; // 30min

useAFK(durationToBeAFK, () => {
  if (accessStore.accessToken && !lockStore.isLockScreen) {
    logout();
  }
});
</script>

<template>
  <ConfigProvider :locale="antdLocale" :theme="tokenTheme">
    <App>
      <RouterView />
    </App>
  </ConfigProvider>
</template>
