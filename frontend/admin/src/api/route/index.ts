import request from '@/axios'
import { QueryPagi, createQueryPayload } from '@/hooks/web/usePagi'
import { MenuRoute } from './types'

type getTypeQuery = QueryPagi & { isGroupNested?: boolean }

export const getRouters = (query: getTypeQuery = { limit: 100, page: 1 }) => {
  return request.get<MenuRoute[]>({
    url: '/routes',
    params: createQueryPayload(query)
  })
}

export const createRouter = (body: MenuRoute) => {
  return request.post<MenuRoute>({ url: '/routes', data: body })
}

export const updateRouter = (body: { id?: number } & MenuRoute) => {
  return request.put<MenuRoute>({ url: `/routes/${body?.id}`, data: body })
}

export const deleteRouter = (id: number) => {
  return request.delete<any>({ url: `/routes/${id}` })
}
