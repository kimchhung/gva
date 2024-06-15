import { MenuRoute } from '../route/types'

export type AdminLoginType = {
  username: string
  password: string
}

export type LoginResponse = {
  token: string
  admin: AdminInfo
}

export type AdminInfo = {
  id: number
  username: string
  isSuperAdmin?: boolean
  edges: {
    roles: any[]
  }

  [key: string]: any
}

export const convertEdgeChildren = (
  list: MenuRoute[],
  appRoutes: AppCustomRouteRecordRaw[] = []
): AppCustomRouteRecordRaw[] => {
  list.forEach((r) => {
    const { edges, ...more } = r
    const appRoute: AppCustomRouteRecordRaw = { ...more }

    if (edges?.children?.length > 0) {
      appRoute.children = convertEdgeChildren(edges.children) // Pass the list and result to the recursive call
    }

    appRoutes.push(appRoute)
  })

  return appRoutes
}
