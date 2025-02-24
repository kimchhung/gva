import type { RequestOption } from '../../utils/axios';
import type { CursorParam } from '../types';

import type { QueryPagi } from '#/utils/pagi/query';

export type Node<T = Record<string, any>> = Record<string, any> & T;

export interface CreateNode<T = any> {
  body: Node<T>;
}

export interface UpdateNode<T = any> {
  body: Node<T>;
  id: number | string;
}

export interface UpdateNodePartial<T = any> {
  body: Partial<Node<T>>;
  id: number | string;
}

export interface GetNode {
  id: number | string;
}

export interface GetManyNode<
  T extends Record<string, any>,
  FilterableKey extends string = AllKeys<T>,
  SortableKey extends string = FilterableKey,
> {
  query?: Partial<QueryPagi<FilterableKey, SortableKey>>;
  cursor?: Partial<CursorParam<FilterableKey, SortableKey>>;
}

export interface GetManyMeta {
  page: number;
  limit: number;
  totalCount: number;
  hasNext: boolean;
}

export interface Option<T> {
  opt?: RequestOption<T>;
}
