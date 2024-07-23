import { UseAPIOption } from '@/axios'
import { QueryUrl } from '@/hooks/web/usePagi'

export type Node<T = Record<string, any>> = T & {
  id: string
  createdAt: string
}

export type CreateNode<T = any> = UseAPIOption & {
  body: Node<T>
}

export type UpdateNode<T = any> = UseAPIOption & {
  id: string
  body: Partial<Node<T>>
}

export type GetNode = UseAPIOption & {
  id: string
}

export type GetManyNode<T extends Object = {}> = UseAPIOption & {
  query: QueryUrl<T>
}

export type APIResponseMeta = {
  total?: number
  offset: number
  limit: number
}

export type ManyNodesResponse<T = any, Meta = APIResponseMeta> = APIResponse<Node<T>[], Meta>
