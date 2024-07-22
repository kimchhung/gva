/* eslint-disable @typescript-eslint/consistent-type-definitions */
import { defineComponent } from 'vue'
import type { RouteRecordRaw } from 'vue-router'

/**
* redirect: noredirect         When set to noredirect, this route cannot be clicked in the breadcrumb navigation.
* name:'router-name'          Set the name of the route. It must be filled in, otherwise, problems will occur when using <keep-alive>.
* meta : {
    hidden: true              When set to true, this route will not appear in the sidebar, such as 404, login pages (default false).

    alwaysShow: true          When you have more than one child route declared under a route, it will automatically become a nested mode.
                              If you want to always show your root route regardless of the number of child routes declared,
                              you can set alwaysShow: true. This way, it will ignore the previously defined rules and always show the root route (default false).

    title: 'title'             Set the name displayed for this route in the sidebar and breadcrumb.

    icon: 'svg-name'          Set the icon for this route.

    noCache: true             If set to true, it will not be cached by <keep-alive> (default false).

    breadcrumb: false         If set to false, it will not be displayed in the breadcrumb (default true).

    affix: true               If set to true, it will always be fixed in the tag (default false).

    noTagsView: true          If set to true, it will not appear in the tag (default false).

    activeMenu: '/dashboard' The highlighted route path.

    canTo: true               Set to true to allow route jump even if hidden is true (default false).

    scopes: ['user:edit','user:add', 'user:delete'] Set the permissions for this route.
 }
**/
interface RouteMetaCustom extends Record<string | number | symbol, unknown> {
  hidden?: boolean
  alwaysShow?: boolean
  title?: string
  icon?: string
  noCache?: boolean
  breadcrumb?: boolean
  affix?: boolean
  activeMenu?: string
  noTagsView?: boolean
  canTo?: boolean
  permission?: string[]
}

declare module 'vue-router' {
  interface RouteMeta extends RouteMetaCustom {}
}

type Component<T = any> =
  | ReturnType<typeof defineComponent>
  | (() => Promise<typeof import('*.vue')>)
  | (() => Promise<T>)

declare global {
  declare interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta' | 'children'> {
    name: string
    meta: RouteMetaCustom
    component?: Component | string
    children?: AppRouteRecordRaw[]
    props?: Record
    fullPath?: string
  }

  declare interface AppCustomRouteRecordRaw
    extends Omit<RouteRecordRaw, 'meta' | 'component' | 'children'> {
    name: string
    meta: RouteMetaCustom
    component: string
    path: string
    redirect: string
    children?: AppCustomRouteRecordRaw[]
  }
}
