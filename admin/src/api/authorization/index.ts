import request from '@/axios'
import { MenuRoute } from './types'

export const getRouters = (params?: any) => {
  return request.get<MenuRoute[]>({ url: '/authorization/routes', params })
}

export const createRouter = (body?: any) => {
  return request.post<MenuRoute>({ url: '/authorization/route', data: body })
}
