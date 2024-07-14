import req, { useAPI } from '@/axios'
import { createQueryPayload } from '@/hooks/web/usePagi'
import { CreateNode, GetManyNode, GetNode, Node, UpdateNode } from './types'

export class ResourceAPI<T = Node<{}>> {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}`
  }

  get({ id, opt }: GetNode) {
    return useAPI({
      fn: () => req.get<Node<T>>({ url: `${this.base}/${id}` }),
      opt
    })
  }

  getMany({ query, opt }: GetManyNode) {
    return useAPI({
      fn: () => req.get<Node<T>[]>({ url: this.base, params: createQueryPayload(query) }),
      opt
    })
  }

  create({ body, opt }: CreateNode) {
    return useAPI({
      fn: () => req.post<Node<T>>({ url: this.base, data: body }),
      opt
    })
  }

  update({ id, body, opt }: UpdateNode) {
    return useAPI({
      fn: () => req.put<Node<T>>({ url: `${this.base}/${id}`, data: body }),
      opt
    })
  }

  delete({ id, opt }: GetNode) {
    return useAPI({
      fn: () => req.delete<Node<T>>({ url: `${this.base}/${id}` }),
      opt
    })
  }
}
