import type { RouteRecordRaw } from 'vue-router';

import { filterTree, mapTree } from '@vben-core/shared';

async function generateRoutesByFrontend(
  routes: RouteRecordRaw[],
  roles: string[],
  permissions?: string[],
  isSuperAdmin?: boolean,
  forbiddenComponent?: RouteRecordRaw['component'],
): Promise<RouteRecordRaw[]> {
  // 根据角色标识过滤路由表,判断当前用户是否拥有指定权限
  const finalRoutes = filterTree(routes, (route) => {
    return hasPermission(
      route,
      roles,
      permissions ?? [],
      isSuperAdmin ?? false,
    );
  });

  if (!forbiddenComponent) {
    return finalRoutes;
  }

  // 如果有禁止访问的页面，将禁止访问的页面替换为403页面
  return mapTree(finalRoutes, (route) => {
    if (menuHasVisibleWithForbidden(route)) {
      route.component = forbiddenComponent;
    }
    return route;
  });
}

/**
 * 判断路由是否有权限访问
 * @param route
 * @param roles
 * @param permissions
 * @param isSuperAdmin
 */
function hasPermission(
  route: RouteRecordRaw,
  roles: string[],
  permissions: string[],
  isSuperAdmin?: boolean,
) {
  const permissions_: string[] = route.meta?.permissions ?? [];
  const roles_: string[] = route.meta?.roles ?? [];
  if (permissions_.length === 0 && roles_.length === 0) {
    return true;
  }

  if (isSuperAdmin) {
    return true;
  }

  const hasPermissionAccess =
    permissions_?.length === 0
      ? true
      : permissions.some((value) => permissions_?.includes(value));

  const hasRoleAccess = roles.some((value) => roles_?.includes(value));

  const canAccess =
    roles_ && roles_?.length > 0 ? hasRoleAccess : hasPermissionAccess;

  return canAccess || (!canAccess && menuHasVisibleWithForbidden(route));
}

/**
 * 判断路由是否在菜单中显示，但是访问会被重定向到403
 * @param route
 */
function menuHasVisibleWithForbidden(route: RouteRecordRaw) {
  return (
    !!route.meta?.permissions &&
    Reflect.has(route.meta || {}, 'menuVisibleWithForbidden') &&
    !!route.meta?.menuVisibleWithForbidden
  );
}

export { generateRoutesByFrontend, hasPermission };
