import type {
  AccessModeType,
  GenerateMenuAndRoutesOptions,
  RouteRecordRaw,
} from '@vben/types';

import {
  cloneDeep,
  generateMenus,
  generateRoutesByBackend,
  generateRoutesByFrontend,
  generateRoutesDocs,
  mapTree,
} from '@vben/utils';

async function generateAccessible(
  mode: AccessModeType,
  options: GenerateMenuAndRoutesOptions,
) {
  const { router } = options;

  options.routes = cloneDeep(options.routes);
  // Generate
  const accessibleRoutes = await generateRoutes(mode, options);

  // Dynamically add to the router instance
  accessibleRoutes.forEach((route) => {
    router.addRoute(route);
  });

  // Generate menu
  const accessibleMenus = await generateMenus(accessibleRoutes, options.router);
  return { accessibleMenus, accessibleRoutes };
}

/**
 * Generate routes
 * @param mode
 */
async function generateRoutes(
  mode: AccessModeType,
  options: GenerateMenuAndRoutesOptions,
) {
  const { forbiddenComponent, isSuperAdmin, permissions, roles, routes } =
    options;

  let resultRoutes: RouteRecordRaw[] = routes;

  const docRoutes = await generateRoutesDocs(options);

  switch (mode) {
    case 'backend': {
      resultRoutes = await generateRoutesByBackend(options);
      break;
    }
    case 'frontend': {
      resultRoutes = await generateRoutesByFrontend(
        routes,
        roles,
        permissions,
        isSuperAdmin,
        forbiddenComponent,
      );
      break;
    }
  }

  resultRoutes.push(...docRoutes);
  /**
   * 调整路由树，做以下处理：
   * 1. 对未添加redirect的路由添加redirect
   */
  resultRoutes = mapTree(resultRoutes, (route) => {
    // 如果有redirect或者没有子路由，则直接返回
    if (route.redirect || !route.children || route.children.length === 0) {
      return route;
    }
    const firstChild = route.children[0];

    // 如果子路由不是以/开头，则直接返回,这种情况需要计算全部父级的path才能得出正确的path，这里不做处理
    if (!firstChild?.path || !firstChild.path.startsWith('/')) {
      return route;
    }

    route.redirect = firstChild.path;
    return route;
  });

  return resultRoutes;
}

export { generateAccessible };
