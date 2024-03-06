import request from '@/axios'

export const getRoleListApi = () => {
  return request.get<RouteReponse>({ url: '/mock/role/table' })
}

export type RouteReponse = {
  list: AdminRoute[]
}

export type AdminRoute = {
  createdAt: string
  updatedAt: string
  isEnable: boolean
  path: string
  component: string
  redirect: string
  name: string
  meta: Recordable
  edges: Edges
}

export type NestedRoute = {
  createdAt: string
  updatedAt: string
  isEnable: boolean
  path: string
  component: string
  redirect: string
  name: string
  meta: Recordable
  children: NestedRoute[]
}

export type Edges = {
  children: AdminRoute[]
}
