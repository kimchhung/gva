import { QueryUrl } from '@/hooks/web/usePagi'

export type Node<T = Record<string, any>> = T & {
  id: string
  createdAt: string
}

export type CreateNode<T = any> = {
  body: Node<T>
}

export type UpdateNode<T = any> = {
  id: string
  body: Partial<Node<T>>
}

export type GetNode = {
  id: string
}

export type GetManyNode<T extends Object = {}> = {
  query: QueryUrl<T>
}

export type GetManyMeta = {
  total?: number
  offset: number
  limit: number
}
