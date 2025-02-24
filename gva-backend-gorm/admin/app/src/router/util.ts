import { type RouteRecordRaw } from 'vue-router';

// sync children permissions to parent //
const syncPermissions = (routes: RouteRecordRaw[]) => {
  for (const route of routes) {
    if (route.children) {
      route.children = syncPermissions(route.children);
    }

    if (route.meta && route.children && route.children.length > 0) {
      route.meta.permissions = [
        ...new Set(
          route.children?.map((child) => child.meta?.permissions ?? []).flat(),
        ),
      ];
      route.meta.roles = [
        ...new Set(
          route.children?.map((child) => child.meta?.roles ?? []).flat(),
        ),
      ];
    }
  }
  return routes;
};

// type heper //
export const defineRoute = (routes: RouteRecordRaw[]) => {
  return syncPermissions(routes);
};
