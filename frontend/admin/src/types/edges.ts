export type WithEdge<T, Edges extends Object = {}> = T & {
  edges: Edges
}

type EdgeChildren<T, Member extends Object = {}> = Member & {
  parent?: T
  children?: T[]
}

export type WithEdgeTree<
  T,
  EdgeMember extends Object = {},
  Edges = EdgeChildren<T, EdgeMember>
> = T & {
  id: string
  pid?: string
  edges: Edges
}

export type WithTree<T> = T & {
  id: string
  pid?: string
  children?: WithTree<T>[]
}
