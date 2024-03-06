import { NO_REDIRECT_WHITE_LIST } from '@/constants'
import { useNProgress } from '@/hooks/web/useNProgress'
import { usePageLoading } from '@/hooks/web/usePageLoading'
import { useTitle } from '@/hooks/web/useTitle'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { useAppStoreWithOut } from '@/store/modules/app'
import { usePermissionStoreWithOut } from '@/store/modules/permission'
import type { RouteRecordRaw } from 'vue-router'
import router from './router'

const { start, done } = useNProgress()

const { loadStart, loadDone } = usePageLoading()

router.beforeEach(async (to, from, next) => {
  start()
  loadStart()

  const permissionStore = usePermissionStoreWithOut()
  const appStore = useAppStoreWithOut()
  const adminStore = useAdminStoreWithOut()

  if (adminStore.token) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      if (permissionStore.getIsAddRouters) {
        next()
        return
      }

      if (adminStore.isNeedReFetchUserInfo) {
        await adminStore.fetchUserInfo()
        await adminStore.fetchAdminRouters()
      }

      // Developers can modify it according to the actual situation
      const roleRouters = adminStore.getRoleRouters || []

      // Whether to use dynamic routing
      if (appStore.getDynamicRouter && appStore.serverDynamicRouter) {
        await permissionStore.generateRoutes('server', roleRouters as AppCustomRouteRecordRaw[])
      } else {
        await permissionStore.generateRoutes('static', [])
      }

      permissionStore.getAddRouters.forEach((route) => {
        router.addRoute(route as unknown as RouteRecordRaw) // Dynamic adding accessable routing table
      })

      const redirectPath = from.query.redirect || to.path
      const redirect = decodeURIComponent(redirectPath as string)
      const nextData = to.path === redirect ? { ...to, replace: true } : { path: redirect }
      permissionStore.setIsAddRouters(true)

      next(nextData)
    }
  } else {
    if (NO_REDIRECT_WHITE_LIST.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.path}`) // Otherwise, all redirect to the login page
    }
  }
})

router.afterEach((to) => {
  useTitle(to?.meta?.title as string)
  done() // 结束Progress
  loadDone()
})
