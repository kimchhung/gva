<script lang="ts" setup>
import { AuthenticationLogin } from '@vben/common-ui';

import { useAuthStore } from '#/store';

defineOptions({ name: 'Login' });

const authStore = useAuthStore();
</script>

<template>
  <div>
    <AuthenticationLogin
      :loading="authStore.loginLoading"
      password-placeholder="123456"
      username-placeholder="admin"
      @submit="authStore.authLogin"
    />
  </div>
</template>
