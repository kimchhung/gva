<script lang="ts" setup>
import type { LoginCodeParams } from '@vben/common-ui';

import { ref } from 'vue';

import { AuthenticationCodeLogin } from '@vben/common-ui';
import { LOGIN_PATH } from '@vben/constants';

defineOptions({ name: 'CodeLogin' });

const loading = ref(false);

/**
 * Asynchronous processing login operation
 * Asynchronously handle the login process
 * @param values Login form data
 */
async function handleLogin(values: LoginCodeParams) {
  // eslint-disable-next-line no-console
  console.log(values);
}
</script>

<template>
  <AuthenticationCodeLogin
    :loading="loading"
    :login-path="LOGIN_PATH"
    @submit="handleLogin"
  />
</template>
