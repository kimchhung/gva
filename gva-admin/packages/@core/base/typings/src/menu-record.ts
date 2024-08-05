import type { RouteRecordRaw } from 'vue-router';

/**
 * 扩展路由原始对象
 */
type ExRouteRecordRaw = {
  parent?: string;
  parents?: string[];
  path?: any;
} & RouteRecordRaw;

interface MenuRecordBadgeRaw {
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
  badgeVariants?: 'destructive' | 'primary' | string;
}

/**
 * Menu original object
 */
interface MenuRecordRaw extends MenuRecordBadgeRaw {
  /**
   * Icon name when activation
   */
  activeIcon?: string;
  /**
   * Submenu
   */
  children?: MenuRecordRaw[];
  /**
   * Whether to disable the menu
   * @default false
   */
  disabled?: boolean;
  /**
   * Icon name
   */
  icon?: string;
  /**
   * Menu name
   */
  name: string;
  /**
   * queue number
   */
  order?: number;
  /**
   * Parent -level path
   */
  parent?: string;
  /**
   * All parent path
   */
  parents?: string[];
  /**
   * Menu path, the only one, can be used as key
   */
  path: string;
  /**
   * Whether to display the menu
   * @default true
   */
  show?: boolean;
}

export type { ExRouteRecordRaw, MenuRecordBadgeRaw, MenuRecordRaw };
