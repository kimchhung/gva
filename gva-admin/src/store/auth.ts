import type { LoginAndRegisterParams } from '@gva/common-ui';
import type { UserInfo } from '@gva/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { DEFAULT_HOME_PATH, LOGIN_PATH } from '@gva/constants';
import { resetAllStores, useAccessStore, useUserStore } from '@gva/stores';

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
  async function authLogin(params: LoginAndRegisterParams, onSuccess?: () => Promise<void> | void) {
    // Asynchronous processing user login operation and acquisition accessToken
    let userInfo: null | UserInfo = null;
    try {
      loginLoading.value = true;
      const [res] = await api.auth.login({
        body: params,
      });
      console.log({ params, res });
      if (res?.data) {
        userInfo = res.data.admin as any;
        accessStore.setAccessToken(res.data.token);
        accessStore.setRefreshToken(res.data.token);
        userStore.setUserInfo(res.data.admin as any);
        // accessStore.setAccessCodes(accessCodes);

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
    resetAllStores();
    accessStore.setLoginExpired(false);

    //Back to the login page belt on the current routing address
    await router.replace({
      path: LOGIN_PATH,
      query: {
        redirect: encodeURIComponent(router.currentRoute.value.fullPath),
      },
    });
  }

  async function fetchUserInfo() {
    let userInfo: null | UserInfo = null;
    const [res, err] = await api.auth.me();
    if (!err) {
      userInfo = {
        id: res.data.id,
        avatar: '',
        homePath: '/',
        name: res.data.username,
        username: res.data.username,
      };
    }
    userStore.setUserInfo(userInfo);
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
