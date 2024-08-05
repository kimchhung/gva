
import { convertEdgeChildren, edgelistToTree, edgeTreeToList } from '#/utils/edge-tree'
import type { RouteRecordStringComponent } from 'packages/types/src'
import { Node } from '../node/types'
import { MenuRoute } from './types'




export const menuToRoute = (data: Node<MenuRoute>[]): RouteRecordStringComponent[] => {
  const nested = edgelistToTree(data)
  const list = convertEdgeChildren(nested)
  return list as RouteRecordStringComponent[]
}

export const menuToFlat = (data: Node<MenuRoute>[]) => {
  const flated = edgeTreeToList(data)
  return flated
}

export const menuToNested = (data: Node<MenuRoute>[]) => {
  const nested = edgelistToTree(data)
  return nested
}
