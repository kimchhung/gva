import { UseAPIOption } from '@/axios'
import { MenuTypeEnum } from '@/constants/menuType'
import { QueryPagi } from '@/hooks/web/usePagi'

export type GetManyRoute = UseAPIOption & {
  query: QueryPagi & { isGroupNested?: boolean }
}

export type MenuRoute = {
  id: string
  createdAt: string
  parentId?: number
  isEnable: boolean
  path: string
  component: string
  redirect: string
  name: string
  meta: Recordable
  edges: Edges
  type: MenuTypeEnum
}

export type NestedRoute = {
  createdAt: string
  updatedAt: string
  isEnable: boolean
  path: string
  component: string
  redirect: string
  name: string
  meta: Recordable
  children: NestedRoute[]
}

export type Edges = {
  children: MenuRoute[]
}
