import { UseAPIOption } from '@/axios'
import { QueryPagi } from '@/hooks/web/usePagi'

export type Node = {
  id: string
  name: string
  createdAt: string
}

export type CreateNode = UseAPIOption & {
  body: Node
}

export type UpdateNode = UseAPIOption & {
  id: string
  body: Partial<Node>
}

export type GetNode = UseAPIOption & {
  id: string
}

export type GetManyNode = UseAPIOption & {
  query: QueryPagi
}
