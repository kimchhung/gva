import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
  type RouteRecordRaw,
} from 'vue-router';

import {
  mergeRouteModules,
  resetStaticRoutes,
  traverseTreeValues,
} from '@vben/utils';

import { coreRoutes, fallbackNotFoundRoute } from './core';
import { createRouterGuard } from './guard';

const dynamicRouteFiles = import.meta.glob('./modules/**/*.ts', {
  eager: true,
});

// If necessary, you can open the comment by yourself and create a folder
// const staticRouteFiles = import.meta.glob('./static/**/*.ts', { eager: true });

/** Dynamic route */
const dynamicRoutes: RouteRecordRaw[] = mergeRouteModules(dynamicRouteFiles);

/** Static routing list, access to these pages can not require permissions */
// const staticRoutes: RouteRecordRaw[] = mergeRouteModules(staticRouteFiles);
const staticRoutes: RouteRecordRaw[] = [];

/** Route list, consisting of basic routing+static routing */
const routes: RouteRecordRaw[] = [
  ...coreRoutes,
  ...staticRoutes,
  fallbackNotFoundRoute,
];

/** Basic routing list, these routes do not need to enter authority interception */
const coreRouteNames = traverseTreeValues(coreRoutes, (route) => route.name);

export { coreRouteNames, dynamicRoutes, routes };

/**
 *  Create a Vue-Router instance
 */
const router = createRouter({
  history:
    import.meta.env.VITE_ROUTER_HISTORY === 'hash'
      ? createWebHashHistory(import.meta.env.VITE_BASE)
      : createWebHistory(import.meta.env.VITE_BASE),
  // The initial route list of the route should be added.
  routes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
  // Should the tail slash be prohibited.
  // strict: true,
});

const resetRoutes = () => resetStaticRoutes(router, routes);

// Create route guard
createRouterGuard(router);

export { resetRoutes, router };
