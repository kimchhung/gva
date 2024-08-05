import { MenuTypeEnum } from '#/constants/menuType'
import { WithEdgeTree } from '#/types/edges'

export type MenuRoute = WithEdgeTree<{
  createdAt: string
  isEnable: boolean
  path: string
  component: string
  redirect: string
  name: string
  meta: Recordable
  type: MenuTypeEnum
}>
