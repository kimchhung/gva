import type { PERMISSION, ROLE } from '#/constants';

import 'vue';
import 'vue-router';

declare module 'vue-router' {
  interface _RouteRecordBase {
    basePath?: string;
    fullPath?: string;
  }

  interface RouteMeta {
    permissions?: PERMISSION[];
    roles?: ROLE[];
  }

  /**
   * Defines a route record with an optional `order` field.
   * This extends Vue Router's `RouteRecordRaw` by removing the `children` property
   * and adding an optional `order` field for custom sorting.
   */
  declare type DefineRouteRecord = { order?: number } & Omit<
    RouteRecordRaw,
    'children'
  >;

  /**
   * Defines a route that can have nested routes via the `children` property.
   * The generic parameter `T` is intended to constrain the keys of the `children` object.
   * This allows for more precise typing of route keys, improving type safety.
   *
   * @template T The type of keys allowed in the `children` object.
   */
  type DefineRoute<T extends string = string> = {
    children?: Record<T, DefineRoute<T>>;
  } & DefineRouteRecord;
}

declare module 'vue' {
  export interface API {
    name: string;
  }

  export interface ComponentCustomProperties {
    vSafeHtml: string;
    vPermissions: PERMISSION[];
    vRoles: string[];
    vAutoAnimate: {
      // they don’t want them via prefers-reduced-motion.
      disrespectUserMotionPreference: false;
      // Animation duration in milliseconds (default: 250)
      duration: 250;
      // When true, this will enable animations even if the user has indicated
      // Easing for motion (default: 'ease-in-out')
      easing: 'ease-in-out';
    };
  }
}

declare module 'vue/types/vue' {
  interface Vue {
    api: API;
  }
}

export {};
