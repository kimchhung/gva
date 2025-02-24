<script lang="ts" setup>
import type { NotificationItem } from '@vben/layouts';

import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';

import { AuthenticationLoginExpiredModal } from '@vben/common-ui';
import { LOGIN_PATH } from '@vben/constants';
import { useAppConfig } from '@vben/hooks';
import { BookOpenText } from '@vben/icons';
import {
  BasicLayout,
  LockScreen,
  Notification,
  UserDropdown,
} from '@vben/layouts';
import { preferences } from '@vben/preferences';
import {
  resetAllStores,
  storeToRefs,
  useAccessStore,
  useUserStore,
} from '@vben/stores';
import { openWindow } from '@vben/utils';

import { $t } from '#/locales';
import { resetRoutes } from '#/router';
import { useAuthStore } from '#/store';

const notifications = ref<NotificationItem[]>([]);

const userStore = useUserStore();
const authStore = useAuthStore();
const accessStore = useAccessStore();
const showDot = computed(() =>
  notifications.value.some((item) => !item.isRead),
);

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);
const menus = computed(() => [
  {
    handler: () => {
      openWindow(`${apiURL}/docs`, {
        target: '_blank',
      });
    },
    icon: BookOpenText,
    text: $t('widgets.document'),
  },
]);

const { loginLoading } = storeToRefs(authStore);

const avatar = computed(() => {
  return userStore.userInfo?.avatar ?? preferences.app.defaultAvatar;
});

const router = useRouter();

async function handleLogout() {
  resetAllStores();
  resetRoutes();
  await router.replace(LOGIN_PATH);
}

function handleNoticeClear() {
  notifications.value = [];
}

function handleMakeAll() {
  notifications.value.forEach((item) => (item.isRead = true));
}
</script>

<template>
  <BasicLayout @clear-preferences-and-logout="handleLogout">
    <template #user-dropdown>
      <UserDropdown
        :avatar
        :menus
        :text="userStore.userInfo?.realName"
        @logout="handleLogout"
      />
    </template>
    <template #notification>
      <Notification
        :dot="showDot"
        :notifications="notifications"
        @clear="handleNoticeClear"
        @make-all="handleMakeAll"
      />
    </template>
    <template #extra>
      <AuthenticationLoginExpiredModal
        v-model:open="accessStore.loginExpired"
        :avatar
        :loading="loginLoading"
        password-placeholder="123456"
        username-placeholder="admin"
        @submit="authStore.authLogin"
      />
    </template>
    <template #lock-screen>
      <LockScreen :avatar @to-login="handleLogout" />
    </template>
  </BasicLayout>
</template>
