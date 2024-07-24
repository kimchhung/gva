import { Admin } from '@/api/admin/types'
import { useI18n } from '@/hooks/web/useI18n'
import router from '@/router'

import { menuToRoute } from '@/api/menu/tranform'
import { QueryUrl } from '@/hooks/web/usePagi'
import { ElMessageBox } from 'element-plus'
import { defineStore } from 'pinia'
import { store } from '../index'
import { useTagsViewStore } from './tagsView'

type AdminState = {
  adminInfo?: Admin
  loading: boolean
  tokenKey: string
  token: string
  routers?: AppCustomRouteRecordRaw[]
}

export const useAdminStore = defineStore('admin', {
  state: (): AdminState => {
    return {
      adminInfo: undefined,
      loading: false,
      tokenKey: 'Authorization',
      token: '',
      routers: undefined
    }
  },
  getters: {
    getTokenKey: (state) => {
      return state.tokenKey
    },
    getToken: (state) => {
      return state.token
    },
    getUserInfo: (state) => {
      return state.adminInfo
    },
    isNeedReFetchUserInfo: (state) => {
      return !state.adminInfo?.id
    },
    isSuperAdmin: (state) => {
      return !!state.adminInfo?.edges?.roles?.find?.((r: any) => r.name === 'SUPER_ADMIN')
    },
    getRoleRouters: (state) => {
      return state.routers
    }
  },
  actions: {
    setTokenKey(tokenKey: string) {
      this.tokenKey = tokenKey
    },
    setToken(token: string) {
      this.token = token
    },
    setAdminInfo(userInfo?: Admin) {
      this.adminInfo = userInfo
    },
    setRoleRouters(roleRouters: AppCustomRouteRecordRaw[]) {
      this.routers = roleRouters
    },
    async fetchUserInfo() {
      this.loading = true
      const [res] = await api.auth.me()
      if (!res?.success) return

      this.setAdminInfo(res.data)
      this.loading = false
    },

    async fetchAdminRouters() {
      const [res] = await api.menu.getMany({
        query: new QueryUrl(100)
      })

      if (!res?.success) return this.routers

      const routes = menuToRoute(res.data)
      this.setRoleRouters(routes)

      return this.routers
    },
    logoutConfirm() {
      const { t } = useI18n()
      ElMessageBox.confirm(t('common.loginOutMessage'), t('common.reminder'), {
        confirmButtonText: t('common.ok'),
        cancelButtonText: t('common.cancel'),
        type: 'warning'
      })
        .then(async () => {
          this.reset()
        })
        .catch(() => {})
    },
    reset() {
      const tagsViewStore = useTagsViewStore()
      tagsViewStore.delAllViews()
      this.setToken('')
      this.setAdminInfo(undefined)
      this.setRoleRouters([])
      router.replace('/login')
    },
    logout() {
      this.reset()
    }
  },
  persist: {
    paths: ['token']
  }
})

export const useAdminStoreWithOut = () => {
  return useAdminStore(store)
}
