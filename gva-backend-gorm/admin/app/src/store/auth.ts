import type { LoginAndRegisterParams } from '@vben/common-ui';
import type { BasicUserInfo, UserInfo } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { DEFAULT_HOME_PATH, LOGIN_PATH } from '@vben/constants';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';

import { notification } from 'ant-design-vue';
import { defineStore } from 'pinia';

import { $t } from '#/locales';

export const useAuthStore = defineStore('auth', () => {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const router = useRouter();

  const loginLoading = ref(false);

  /**
   * Asynchronous processing login operation
   * Asynchronously handle the login process
   * @param params Login form data
   */
  async function authLogin(
    params: LoginAndRegisterParams,
    onSuccess?: () => Promise<void> | void,
  ) {
    // Asynchronous processing user login operation and acquisition accessToken
    let userInfo: null | UserInfo = null;

    // Clear AccessRoutes here so when login new account it will regenerate route access
    accessStore.setAccessRoutes([]);

    try {
      loginLoading.value = true;
      const [res] = await api.auth.login({
        body: {
          password: params.password,
          totp: params.totp ?? '',
          username: params.username,
        },
      });
      if (res?.data) {
        userInfo = res.data.admin as any;
        accessStore.setAccessToken(res.data.token);
        accessStore.setRefreshToken(res.data.token);

        userStore.setUserInfo({
          permissions: res.data.admin.permissionScope,
          roles: res.data.admin.roleNameId,
          realName: res.data.admin.name,
          userId: res.data.admin.id,
          avatar: '',
          ...userInfo,
        } as any);

        if (accessStore.loginExpired) {
          accessStore.setLoginExpired(false);
        } else {
          onSuccess
            ? await onSuccess?.()
            : await router.push(userInfo?.homePath || DEFAULT_HOME_PATH);
        }

        if (userInfo?.realName) {
          notification.success({
            description: `${$t('authentication.loginSuccessDesc')}:${userInfo?.realName}`,
            duration: 3,
            message: $t('authentication.loginSuccess'),
          });
        }
      }
    } finally {
      loginLoading.value = false;
    }

    return {
      userInfo,
    };
  }

  async function logout() {
    if (router.currentRoute.value.path === LOGIN_PATH) return;

    resetAllStores();
    accessStore.setLoginExpired(false);

    // Back to the login page belt on the current routing address
    const currentRoutePath = encodeURIComponent(
      router.currentRoute.value.fullPath,
    );
    await router.replace({
      path: LOGIN_PATH,
      query: { redirect: currentRoutePath },
    });
  }

  async function fetchUserInfo() {
    let userInfo: BasicUserInfo | null = null;
    const [res, err] = await api.auth.me();
    if (err) return null;

    userInfo = {
      ...res.data,
      permissions: res.data.permissionScope,
      realName: res.data.name,
      roles: res.data.roleNameId,
      userId: Number(res.data.id),
      avatar: '',
    };

    if (userInfo) {
      userStore.setUserInfo(userInfo);
    }

    return userInfo;
  }

  function $reset() {
    loginLoading.value = false;
  }

  return {
    $reset,
    authLogin,
    fetchUserInfo,
    loginLoading,
    logout,
  };
});
