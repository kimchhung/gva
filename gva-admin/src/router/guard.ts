import type { Router } from 'vue-router';

import { LOGIN_PATH } from '@gva/constants';
import { preferences } from '@gva/preferences';
import { useAccessStore, useUserStore } from '@gva/stores';
import { startProgress, stopProgress } from '@gva/utils';

import { useTitle } from '@vueuse/core';

import { $t } from '#/locales';
import { coreRouteNames, dynamicRoutes } from '#/router/routes';
import { useAuthStore } from '#/store';

import { generateAccess } from './access';

/**
 * General Guard Configuration
 * @param router
 */
function setupCommonGuard(router: Router) {
  // The page that has been loaded
  const loadedPaths = new Set<string>();

  router.beforeEach(async (to) => {
    to.meta.loaded = loadedPaths.has(to.path);

    // Page loading progress bar
    if (!to.meta.loaded && preferences.transition.progress) {
      startProgress();
    }
    return true;
  });

  router.afterEach((to) => {
    // Whether the record page is loaded, if it has been loaded, the subsequent page switching animation equivalent effects are not repeatedly executed

    if (preferences.tabbar.enable) {
      loadedPaths.add(to.path);
    }

    // Close the page loading progress bar
    if (preferences.transition.progress) {
      stopProgress();
    }

    // Dynamic modification title
    if (preferences.app.dynamicTitle) {
      const { title } = to.meta;
      // useTitle(`${$t(title)} - ${preferences.app.name}`);
      useTitle(`${$t(title)} - ${preferences.app.name}`);
    }
  });
}

/**
 * Permanent access guard configuration
 * @param router
 */
function setupAccessGuard(router: Router) {
  router.beforeEach(async (to, from) => {
    const accessStore = useAccessStore();
    const userStore = useUserStore();
    const authStore = useAuthStore();

    // accessToken examine
    if (!accessStore.accessToken) {
      if (
        // Basic routing, these routes do not need to enter authority interception
        coreRouteNames.includes(to.name as string) ||
        // Determine the conflict of permission access permissions, you can access
        to.meta.ignoreAccess
      ) {
        return true;
      }

      // No access permissions, jump login page
      if (to.fullPath !== LOGIN_PATH) {
        return {
          path: LOGIN_PATH,
          // If you don't need it, delete it directly query
          query: { redirect: encodeURIComponent(to.fullPath) },
          // Carry the current jump page, and then jump to the page after logging in
          replace: true,
        };
      }
      return to;
    }

    const accessRoutes = accessStore.accessRoutes;

    // Have you ever generated dynamic routing
    if (accessRoutes && accessRoutes.length > 0) {
      return true;
    }

    // Generate route table
    // Currently logging in the role identity list of character logo
    const userInfo = userStore.userInfo || (await authStore.fetchUserInfo());
    const userRoles = userInfo?.roles ?? [];

    // Generate menu and routing
    const { accessibleMenus, accessibleRoutes } = await generateAccess({
      roles: userRoles,
      router,
      // It will be displayed in the menu, but the access will be redirected to 403
      routes: dynamicRoutes,
    });

    // Save menu information and routing information
    accessStore.setAccessMenus(accessibleMenus);
    accessStore.setAccessRoutes(accessibleRoutes);
    const redirectPath = (from.query.redirect ?? to.path) as string;

    return {
      path: decodeURIComponent(redirectPath),
      replace: true,
    };
  });
}

/**
 * Project guard configuration
 * @param router
 */
function createRouterGuard(router: Router) {
  /** Universal */
  setupCommonGuard(router);
  /** Permission access */
  setupAccessGuard(router);
}

export { createRouterGuard };
