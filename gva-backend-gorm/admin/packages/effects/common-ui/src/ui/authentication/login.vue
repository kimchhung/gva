<script setup lang="ts">
import type { AuthenticationProps, LoginEmits } from './types';

import { computed, reactive } from 'vue';

import { $t } from '@vben/locales';
import { VbenButton, VbenInput, VbenInputPassword } from '@vben-core/shadcn-ui';

import Title from './auth-title.vue';

interface Props extends AuthenticationProps {}

defineOptions({
  name: 'AuthenticationLogin',
});

withDefaults(defineProps<Props>(), {
  codeLoginPath: '/auth/code-login',
  forgetPasswordPath: '/auth/forget-password',
  loading: false,
  passwordPlaceholder: '',
  qrCodeLoginPath: '/auth/qrcode-login',
  registerPath: '/auth/register',
  showCodeLogin: true,
  showForgetPassword: true,
  showQrcodeLogin: true,
  showRegister: true,
  showRememberMe: true,
  showThirdPartyLogin: true,
  subTitle: '',
  title: '',
  usernamePlaceholder: '',
});

const emit = defineEmits<{
  submit: LoginEmits['submit'];
}>();

const REMEMBER_ME_KEY = `REMEMBER_ME_USERNAME_${location.hostname}`;

const localUsername = localStorage.getItem(REMEMBER_ME_KEY) || '';

const formState = reactive({
  password: '',
  rememberMe: !!localUsername,
  submitted: false,
  username: localUsername,
  totp: '',
});

const usernameStatus = computed(() => {
  return formState.submitted && !formState.username ? 'error' : 'default';
});

const passwordStatus = computed(() => {
  return formState.submitted && !formState.password ? 'error' : 'default';
});

const totpStatus = computed(() => {
  return formState.submitted && !formState.totp ? 'error' : 'default';
});

function handleSubmit() {
  formState.submitted = true;

  if (
    usernameStatus.value !== 'default' ||
    passwordStatus.value !== 'default'
  ) {
    return;
  }

  localStorage.setItem(
    REMEMBER_ME_KEY,
    formState.rememberMe ? formState.username : '',
  );

  emit('submit', {
    password: formState.password,
    username: formState.username,
    totp: formState.totp,
  });
}
</script>

<template>
  <div @keypress.enter.prevent="handleSubmit">
    <Title>
      <div class="flex space-x-2">
        <span>
          {{ title || `${$t('authentication.welcomeBack')}` }}
        </span>
      </div>
    </Title>

    <VbenInput
      v-model="formState.username"
      :autofocus="false"
      :error-tip="$t('authentication.usernameTip')"
      :label="$t('authentication.username')"
      :placeholder="$t('authentication.username')"
      :status="usernameStatus"
      name="username"
      required
      type="text"
    />
    <VbenInputPassword
      v-model="formState.password"
      :error-tip="$t('authentication.passwordTip')"
      :label="$t('authentication.password')"
      :placeholder="$t('authentication.password')"
      :status="passwordStatus"
      name="password"
      required
      type="password"
    />
    <VbenInput
      v-model="formState.totp"
      :autofocus="false"
      :error-tip="$t('authentication.totpTip')"
      :label="$t('authentication.totp')"
      :placeholder="$t('authentication.totp')"
      :status="totpStatus"
      name="totp"
      required
    />
    <VbenButton :loading="loading" class="w-full" @click="handleSubmit">
      {{ $t('common.login') }}
    </VbenButton>

    <div class="mt-4">
      <slot name="bottom"></slot>
    </div>
  </div>
</template>
