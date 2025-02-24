import type { MenuRecordRaw } from '@vben-core/typings';
import type { RouteRecordRaw } from 'vue-router';

import { acceptHMRUpdate, defineStore } from 'pinia';

type AccessToken = null | string;

interface AccessState {
  /**
   * Permission code
   */
  accessCodes: string[];
  /**
   * Access menu list
   */
  accessMenus: MenuRecordRaw[];
  /**
   * Can access lists
   */
  accessRoutes: RouteRecordRaw[];
  /**
   * Log in accessToken
   */
  accessToken: AccessToken;
  /**
   * Login whether to expire
   */
  loginExpired: boolean;
  /**
   * Log in accessToken
   */
  refreshToken: AccessToken;
}

/**
 * View permissions related
 */
export const useAccessStore = defineStore('core-access', {
  actions: {
    setAccessCodes(codes: string[]) {
      this.accessCodes = codes;
    },
    setAccessMenus(menus: MenuRecordRaw[]) {
      this.accessMenus = menus;
    },
    setAccessRoutes(routes: RouteRecordRaw[]) {
      this.accessRoutes = routes;
    },
    setAccessToken(token: AccessToken) {
      this.accessToken = token;
    },
    setLoginExpired(loginExpired: boolean) {
      this.loginExpired = loginExpired;
    },
    setRefreshToken(token: AccessToken) {
      this.refreshToken = token;
    },
  },
  persist: {
    // Endurance
    paths: ['accessToken', 'refreshToken', 'accessCodes'],
  },
  state: (): AccessState => ({
    accessCodes: [],
    accessMenus: [],
    accessRoutes: [],
    accessToken: null,
    loginExpired: false,
    refreshToken: null,
  }),
});

// Solve the problem of hot updates
const hot = import.meta.hot;
if (hot) {
  hot.accept(acceptHMRUpdate(useAccessStore, hot));
}
