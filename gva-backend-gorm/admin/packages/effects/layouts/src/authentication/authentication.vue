<script setup lang="ts">
import { computed } from 'vue';

import { preferences, usePreferences } from '@vben/preferences';

import AuthenticationFormView from './form.vue';
import SloganIcon from './icons/slogan.vue';

defineOptions({ name: 'Authentication' });

const { authPanelCenter, authPanelLeft, authPanelRight } = usePreferences();
const appName = computed(() => preferences.app.name);
const logoSource = computed(() => preferences.logo.source);
</script>

<template>
  <div class="flex min-h-full flex-1 select-none overflow-x-hidden">
    <!-- 左侧认证面板 -->
    <AuthenticationFormView
      v-if="authPanelLeft"
      class="min-h-full w-2/5"
      transition-name="slide-left"
    />

    <!-- 头部 Logo 和应用名称 -->
    <div class="absolute left-0 top-0 z-10 flex flex-1">
      <div
        class="text-foreground ml-4 mt-4 flex flex-1 items-center sm:left-6 sm:top-6"
      >
        <img :alt="appName" :src="logoSource" class="mr-2" width="42" />
        <p class="text-xl font-medium">
          {{ appName }}
        </p>
      </div>
    </div>

    <!-- 中间内容 -->
    <div v-if="!authPanelCenter" class="relative hidden w-0 flex-1 lg:block">
      <div class="bg-authentication absolute inset-0 h-full w-full">
        <div class="flex-col-center -enter-x mr-20 h-full">
          <SloganIcon :alt="appName" class="animate-float h-80 w-1/2" />
        </div>
      </div>
    </div>

    <!-- 中心认证面板 -->
    <div v-if="authPanelCenter" class="flex-center bg-authentication w-full">
      <AuthenticationFormView
        class="md:bg-background w-full rounded-3xl pb-20 shadow-2xl md:w-2/3 lg:w-1/2 xl:w-2/5"
      />
    </div>

    <!-- 右侧认证面板 -->
    <AuthenticationFormView
      v-if="authPanelRight"
      class="min-h-full w-2/5 flex-1"
    />
  </div>
</template>
