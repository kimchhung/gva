import type { Router, RouteRecordRaw } from 'vue-router';

import type { Component } from 'vue';

interface RouteMeta {
  /**
   * Activation icon (menu (menu/tab）
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
   * iframe 地址
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
   * permissions
   */
  permissions?: string[];
  /**
   * roles
   */
  roles?: string[];
  /**
   * Title name
   */
  title:
    | ((props: { t: (key: string, options?: any) => string }) => string)
    | string;

  /**
   * dynamic title
   */

  titleEn?: string;
  titleZh?: string;
}

// Define the recursive type to change the component property of RoutereCordraw to String
type RouteRecordStringComponent<T = string> = {
  children?: RouteRecordStringComponent<T>[];
  component: T;
} & Omit<RouteRecordRaw, 'children' | 'component'>;

type ComponentRecordType = Record<string, () => Promise<Component>>;

interface GenerateMenuAndRoutesOptions {
  fetchMenuListAsync?: () => Promise<RouteRecordStringComponent[]>;
  forbiddenComponent?: RouteRecordRaw['component'];
  isSuperAdmin: boolean;
  layoutMap?: ComponentRecordType;
  pageMap?: ComponentRecordType;
  permissions: string[];
  roles: string[];
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
