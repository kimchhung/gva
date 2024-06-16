import { getAdminInfoReq } from '@/api/admin'
import { AdminInfo, convertEdgeChildren } from '@/api/admin/types'

import { getRouters } from '@/api/route'
import { useApi } from '@/axios'
import { useI18n } from '@/hooks/web/useI18n'
import router from '@/router'
import { ElMessageBox } from 'element-plus'
import { defineStore } from 'pinia'
import { store } from '../index'
import { useTagsViewStore } from './tagsView'

type AdminState = {
  adminInfo?: AdminInfo
  tokenKey: string
  token: string
  routers?: AppCustomRouteRecordRaw[]
}

export const useAdminStore = defineStore('admin', {
  state: (): AdminState => {
    return {
      adminInfo: undefined,
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
    setAdminInfo(userInfo?: AdminInfo) {
      this.adminInfo = userInfo
    },
    setRoleRouters(roleRouters: AppCustomRouteRecordRaw[]) {
      this.routers = roleRouters
    },
    async fetchUserInfo() {
      const [data] = await useApi(() => getAdminInfoReq())
      if (data) this.setAdminInfo(data)
    },
    async fetchAdminRouters() {
      const [data] = await useApi(() => getRouters({ limit: 100, page: 1, isGroupNested: true }))
      if (data) this.setRoleRouters(convertEdgeChildren(data as any))
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
