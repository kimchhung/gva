import type { RouteRecordRaw } from 'vue-router';

import { mergeRouteModules, traverseTreeValues } from '@gva/utils';

import { coreRoutes, fallbackNotFoundRoute } from './core';

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
const routes: RouteRecordRaw[] = [...coreRoutes, ...staticRoutes, fallbackNotFoundRoute];

/** Basic routing list, these routes do not need to enter authority interception */
const coreRouteNames = traverseTreeValues(coreRoutes, (route) => route.name);

export { coreRouteNames, dynamicRoutes, routes };
