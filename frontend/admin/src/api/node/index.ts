import { req, useAPI } from '@/axios'
import { parseQuery } from '@/hooks/web/usePagi'
import { CreateNode, GetManyNode, GetNode, ManyNodesResponse, Node, UpdateNode } from './types'

export class ResourceAPI<T extends object = Node<{}>> {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}`
  }

  get({ id, opt }: GetNode) {
    return useAPI({
      fn: () => req.get<APIResponse<Node<T>>>({ url: `${this.base}/${id}` }),
      opt
    })
  }

  getMany({ query, opt }: GetManyNode<T>) {
    return useAPI({
      fn: () =>
        req.get<ManyNodesResponse<T>>({
          url: this.base,
          params: parseQuery(query)
        }),
      opt
    })
  }

  create({ body, opt }: CreateNode) {
    return useAPI({
      fn: () => req.post<APIResponse<Node<T>>>({ url: this.base, data: body }),
      opt
    })
  }

  update({ id, body, opt }: UpdateNode) {
    return useAPI({
      fn: () => req.put<APIResponse<Node<T>>>({ url: `${this.base}/${id}`, data: body }),
      opt
    })
  }

  delete({ id, opt }: GetNode) {
    return useAPI({
      fn: () => req.delete<APIResponse<Node<T>>>({ url: `${this.base}/${id}` }),
      opt
    })
  }
}
