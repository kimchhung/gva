import request from '@/axios'
import { MenuRoute } from './types'

export const getRouters = (params?: any) => {
  return request.get<MenuRoute[]>({ url: '/authorization/routes', params })
}

export const createRouter = (body: MenuRoute) => {
  return request.post<MenuRoute>({ url: '/authorization/route', data: body })
}

export const updateRouter = (body: { id?: number } & MenuRoute) => {
  return request.put<MenuRoute>({ url: `/authorization/route/:${body?.id}`, data: body })
}
