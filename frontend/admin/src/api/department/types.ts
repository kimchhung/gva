import { WithEdgeTree } from '@/types/edges'

export type Department = WithEdgeTree<{
  id: string
  name: string
  nameId: string
  isEnable: boolean
}>
