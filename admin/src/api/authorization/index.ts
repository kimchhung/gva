import request from '@/axios'
import { QueryPagi, createQueryPayload } from '@/hooks/web/usePagi'
import { MenuRoute } from './types'

export const getRouters = (query: QueryPagi = { limit: 500, page: 1 }) => {
  return request.get<MenuRoute[]>({
    url: '/authorization/routes',
    params: createQueryPayload(query)
  })
}

export const createRouter = (body: MenuRoute) => {
  return request.post<MenuRoute>({ url: '/authorization/route', data: body })
}

export const updateRouter = (body: { id?: number } & MenuRoute) => {
  return request.put<MenuRoute>({ url: `/authorization/route/${body?.id}`, data: body })
}

export const deleteRouter = (id: number) => {
  return request.delete<any>({ url: `/authorization/route/${id}` })
}
