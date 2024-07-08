import req, { useAPI } from '@/axios'
import { createQueryPayload } from '@/hooks/web/usePagi'
import { CreateRoute, GetManyRoute, GetRoute, MenuRoute, UpdateRoute } from './types'

export type Resource = ReturnType<typeof resource>

export const resource = (base = '/routes') => {
  const get = (id: string) => {
    return req.get<MenuRoute>({ url: `${base}/${id}` })
  }

  get.getMany = ({ query, opt }: GetManyRoute) => {
    return useAPI({
      fn: () => req.get<MenuRoute[]>({ url: base, params: createQueryPayload(query) }),
      opt
    })
  }

  get.create = ({ body, opt }: CreateRoute) => {
    return useAPI({
      fn: () => req.post<MenuRoute>({ url: base, data: body }),
      opt
    })
  }

  get.update = ({ id, body, opt }: UpdateRoute) => {
    return useAPI({
      fn: () => req.put<MenuRoute>({ url: `${base}/${id}`, data: body }),
      opt
    })
  }

  get.delete = ({ id, opt }: GetRoute) => {
    return useAPI({
      fn: () => req.delete<MenuRoute>({ url: `${base}/${id}` }),
      opt
    })
  }

  return get
}

const module: APIModule = {
  name: 'route',
  resource
}

export default module
