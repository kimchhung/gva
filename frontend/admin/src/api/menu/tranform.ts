import { edgelistToTree, edgeTreeToList } from '@/utils/edgeTree'
import { convertEdgeChildren } from '@/utils/tree'
import { Node } from '../node/types'
import { MenuRoute } from './types'

export const menuToRoute = (data: Node<MenuRoute>[]): AppCustomRouteRecordRaw[] => {
  const nested = edgelistToTree(data)
  const list = convertEdgeChildren(nested)
  return list as AppCustomRouteRecordRaw[]
}

export const menuToFlat = (data: Node<MenuRoute>[]) => {
  const flated = edgeTreeToList(data)
  return flated
}

export const menuToNested = (data: Node<MenuRoute>[]) => {
  const nested = edgelistToTree(data)
  return nested
}
