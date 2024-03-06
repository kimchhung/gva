import { asyncRouterMap, constantRouterMap } from '@/router'
import { flatMultiLevelRoutes, generateRoutesByServer } from '@/utils/routerHelper'
import { cloneDeep } from 'lodash-es'
import { defineStore } from 'pinia'
import { store } from '../index'

export type PermissionState = {
  routers: AppRouteRecordRaw[]
  addRouters: AppRouteRecordRaw[]
  isAddRouters: boolean
  menuTabRouters: AppRouteRecordRaw[]
}

export const usePermissionStore = defineStore('permission', {
  state: (): PermissionState => ({
    routers: [],
    addRouters: [],
    isAddRouters: false,
    menuTabRouters: []
  }),
  getters: {
    getRouters(): AppRouteRecordRaw[] {
      return this.routers
    },
    getAddRouters(): AppRouteRecordRaw[] {
      return flatMultiLevelRoutes(cloneDeep(this.addRouters))
    },
    getIsAddRouters(): boolean {
      return this.isAddRouters
    },
    getMenuTabRouters(): AppRouteRecordRaw[] {
      return this.menuTabRouters
    }
  },
  actions: {
    setIsAddRouters(state: boolean): void {
      this.isAddRouters = state
    },
    setMenuTabRouters(routers: AppRouteRecordRaw[]): void {
      this.menuTabRouters = routers
    },
    async generateRoutes(
      type: 'server' | 'frontEnd' | 'static',
      routers: AppCustomRouteRecordRaw[]
    ) {
      try {
        const routerMap =
          type === 'server' ? await generateRoutesByServer(routers) : cloneDeep(asyncRouterMap)

        // Dynamic routing, 404 must be put to the end
        this.addRouters = routerMap.concat([
          {
            path: '/:path(.*)*',
            redirect: '/404',
            name: '404Page',
            meta: {
              hidden: true,
              breadcrumb: false
            }
          }
        ])

        // All routes of the rendering menu
        this.routers = cloneDeep(constantRouterMap).concat(routerMap)
      } catch (error) {
        // Handle any errors that occur during the generation of routes
        console.error('Error generating routes:', error)
      }
    }
  },
  persist: false
  // persist: {
  //   paths: ['routers', 'addRouters', 'menuTabRouters']
  // }
})

export const usePermissionStoreWithOut = () => {
  return usePermissionStore(store)
}
