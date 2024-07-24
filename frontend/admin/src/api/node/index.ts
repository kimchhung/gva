import { req } from '@/axios'
import { parseQuery } from '@/hooks/web/usePagi'
import { APIRes } from '../types'
import { CreateNode, GetManyMeta, GetManyNode, GetNode, Node, UpdateNode } from './types'

export class ResourceAPI<T extends object = Node<{}>> {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}`
  }

  get({ id }: GetNode) {
    return req.get<APIRes<Node<T>>>({ url: `${this.base}/${id}` })
  }

  getMany({ query }: GetManyNode<T>) {
    const params = parseQuery(query)
    return req.get<APIRes<Node<T>[], GetManyMeta>>({ url: this.base, params })
  }

  create({ body }: CreateNode<T>) {
    return req.post<APIRes<Node<T>>>({ url: this.base, data: body })
  }

  update({ id, body }: UpdateNode<T>) {
    return req.put<APIRes<Node<T>>>({ url: `${this.base}/${id}`, data: body })
  }

  delete({ id }: GetNode) {
    return req.delete<APIRes<Node<T>>>({ url: `${this.base}/${id}` })
  }
}
