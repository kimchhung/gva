import { req } from '#/utils/axios'
import { parseQuery } from '#/hooks/web/use-pagi'

import { CreateNode, GetManyMeta, GetManyNode, GetNode, Node, UpdateNode } from './types'

export class ResourceAPI<T extends object = Node<{}>> {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}`
  }

  get({ id }: GetNode) {
    return req.get<Node<T>>({ url: `${this.base}/${id}` })
  }

  getMany({ query }: GetManyNode<T>) {
    const params = parseQuery(query)
    return req.get<Node<T>[], GetManyMeta>({ url: this.base, params })
  }

  create({ body }: CreateNode<T>) {
    return req.post<Node<T>>({ url: this.base, data: body })
  }

  update({ id, body }: UpdateNode<T>) {
    return req.put<Node<T>>({ url: `${this.base}/${id}`, data: body })
  }

  delete({ id }: GetNode) {
    return req.delete<Node<T>>({ url: `${this.base}/${id}` })
  }
}
