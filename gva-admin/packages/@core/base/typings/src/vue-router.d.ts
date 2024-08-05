import type { Router, RouteRecordRaw } from 'vue-router';

import type { Component } from 'vue';

interface RouteMeta {
  /**
   * Activate icon (menu/tab)
   */
  activeIcon?: string;
  /**
   * The currently activated menu, sometimes do not want to activate the existing menu, use it when the parent menu needs to be activated
   * @default false
   */
  activePath?: string;
  /**
   * Whether to fix the tab page
   * @default false
   */
  affixTab?: boolean;
  /**
   * The order of fixed tab page
   * @default 0
   */
  affixTabOrder?: number;
  /**
   * You need a specific character identification to access it
   * @default []
   */
  authority?: string[];
  /**
   * Label
   */
  badge?: string;
  /**
   * Logo
   */
  badgeType?: 'dot' | 'normal';
  /**
   * Logo color
   */
  badgeVariants?:
    | 'default'
    | 'destructive'
    | 'primary'
    | 'success'
    | 'warning'
    | string;
  /**
   * The sub -level of the current routing is not shown in the menu
   * @default false
   */
  hideChildrenInMenu?: boolean;
  /**
   * The current route is not displayed in bread debris
   * @default false
   */
  hideInBreadcrumb?: boolean;
  /**
   * The current route is not displayed in the menu
   * @default false
   */
  hideInMenu?: boolean;
  /**
   * The current routing is not displayed on the tab page
   * @default false
   */
  hideInTab?: boolean;
  /**
   *Icon (menu/tab)
   */
  icon?: string;
  /**
   * iframe address
   */
  iframeSrc?: string;
  /**
   * Ignore permissions, you can directly access
   * @default false
   */
  ignoreAccess?: boolean;
  /**
   * Open Keepalive cache
   */
  keepAlive?: boolean;
  /**
   * External link-jump path
   */
  link?: string;
  /**
   * Whether the route has been loaded
   */
  loaded?: boolean;
  /**
   * Maximum opening number of tab pages
   * @default -1
   */
  maxNumOfOpenTab?: number;
  /**
   * The menu can be seen, but the access will be redirected to 403
   */
  menuVisibleWithForbidden?: boolean;
  /**
   * Used for routing-> menu sorting
   */
  order?: number;
  /**
   * Title name
   */
  title: string;
}

// 定义递归类型以将 RouteRecordRaw 的 component 属性更改为 string
type RouteRecordStringComponent<T = string> = {
  children?: RouteRecordStringComponent<T>[];
  component: T;
} & Omit<RouteRecordRaw, 'children' | 'component'>;

type ComponentRecordType = Record<string, () => Promise<Component>>;

interface GenerateMenuAndRoutesOptions {
  fetchMenuListAsync?: () => Promise<RouteRecordStringComponent[]>;
  forbiddenComponent?: RouteRecordRaw['component'];
  layoutMap?: ComponentRecordType;
  pageMap?: ComponentRecordType;
  roles?: string[];
  router: Router;
  routes: RouteRecordRaw[];
}

export type {
  ComponentRecordType,
  GenerateMenuAndRoutesOptions,
  RouteMeta,
  RouteRecordRaw,
  RouteRecordStringComponent,
};
