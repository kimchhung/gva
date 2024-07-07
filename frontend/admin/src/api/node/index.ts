import req, { useAPI } from '@/axios'
import { createQueryPayload } from '@/hooks/web/usePagi'
import { CreateNode, GetManyNode, GetNode, Node, UpdateNode } from './types'

export const nodeResource = (base = '/nodes') => {
  const get = (id: string) => {
    return req.get<Node>({ url: `${base}/${id}` })
  }

  get.getMany = ({ query, opt }: GetManyNode) => {
    return useAPI({
      fn: () => req.get<Node[]>({ url: base, params: createQueryPayload(query) }),
      opt
    })
  }

  get.create = ({ body, opt }: CreateNode) => {
    return useAPI({
      fn: () => req.post<Node>({ url: base, data: body }),
      opt
    })
  }

  get.update = ({ id, body, opt }: UpdateNode) => {
    return useAPI({
      fn: () => req.put<Node>({ url: `${base}/${id}`, data: body }),
      opt
    })
  }

  get.delete = ({ id, opt }: GetNode) => {
    return useAPI({
      fn: () => req.delete<Node>({ url: `${base}/${id}` }),
      opt
    })
  }

  return get
}
