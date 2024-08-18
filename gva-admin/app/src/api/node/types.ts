export type Node<T = Record<string, any>> = {
  createdAt: string;
  id: string;
} & T;

export interface CreateNode<T = any> {
  body: Node<T>;
}

export interface UpdateNode<T = any> {
  body: Partial<Node<T>>;
  id: string;
}

export interface GetNode {
  id: string;
}

export interface GetManyNode<T extends object = any> {
  query: T;
}

export interface GetManyMeta {
  limit: number;
  offset: number;
  total?: number;
}
