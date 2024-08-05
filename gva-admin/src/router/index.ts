import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router';

import { resetStaticRoutes } from '@gva/utils';

import { createRouterGuard } from './guard';
import { routes } from './routes';

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
