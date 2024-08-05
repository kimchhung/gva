<script lang="ts" setup>
import { AuthenticationQrCodeLogin } from '@gva/common-ui';
import { LOGIN_PATH } from '@gva/constants';

defineOptions({ name: 'QrCodeLogin' });
</script>

<template>
  <AuthenticationQrCodeLogin :login-path="LOGIN_PATH" />
</template>
