import type { AccessModeType, GenerateMenuAndRoutesOptions, RouteRecordRaw } from '@gva/types';

import {
  cloneDepp,
  generateMenus,
  generateRoutesByBackend,
  generateRoutesByFrontend,
  mapTree,
} from '@gva/utils';

async function generateAccessible(mode: AccessModeType, options: GenerateMenuAndRoutesOptions) {
  const { router } = options;

  options.routes = cloneDepp(options.routes);
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
async function generateRoutes(mode: AccessModeType, options: GenerateMenuAndRoutesOptions) {
  const { forbiddenComponent, roles, routes } = options;

  let resultRoutes: RouteRecordRaw[] = routes;
  switch (mode) {
    case 'frontend': {
      resultRoutes = await generateRoutesByFrontend(routes, roles || [], forbiddenComponent);
      break;
    }
    case 'backend': {
      resultRoutes = await generateRoutesByBackend(options);
      break;
    }
  }

  /**
   * Adjust the routing tree and do the following processing:
   * 1.Add Redirect to the routing of unlaxed Redirect
   */
  resultRoutes = mapTree(resultRoutes, (route) => {
    // If there is a redirect or no sub -route, return directly
    if (route.redirect || !route.children || route.children.length === 0) {
      return route;
    }
    const firstChild = route.children[0];

    // If the subway is not/beginning, it returns directly. In this case, you need to calculate all the parent Path to get the correct PATH.
    if (!firstChild.path || !firstChild.path.startsWith('/')) {
      return route;
    }

    route.redirect = firstChild.path;
    return route;
  });

  return resultRoutes;
}

export { generateAccessible };
