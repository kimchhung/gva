import type { SingleFilter, SingleSort } from '#/utils/pagi/query';

export interface SuccessRes<DataT, MetaT> {
  code: number;
  // succes data
  data: DataT;

  message: string;
  meta?: MetaT;
  success: true;
}

export type FailedRes<T> = {
  code: number;
  message: string;
  success: false;
} & T;

interface UnionRes<DataT, MetaT> {
  code: number;
  data?: DataT;
  message: string;
  meta?: MetaT;
  success: false;
}

// Union type for API response
export type APIRes<
  DataT,
  MetaT = any,
  FailedT extends Record<string, any> = any,
> = (FailedRes<FailedT> | SuccessRes<DataT, MetaT>) | UnionRes<DataT, MetaT>;

export interface AdminIndexConfig {
  now: string;
  publicIp: string;
  country: string;
  countryCode: string;
}

export interface CursorParam<
  FilterableKey extends string,
  SortableKey extends string = FilterableKey,
> {
  isCount?: boolean;
  first?: number;
  last?: number;
  after?: string;
  before?: string;
  sorts: SingleSort<SortableKey>[];
  filters: SingleFilter<FilterableKey>[];
}

export interface PageInfo {
  hasNextPage?: boolean;
  hasPreviousPage?: boolean;
  startCursor?: string;
  endCursor?: string;
}

export interface Edge<T> {
  cursor: string;
  node: T;
}

export interface CursorData<T> {
  totalCount?: number;
  edges: Edge<T>[];
  pageInfo?: PageInfo;
  isLoading?: boolean;
}

export interface UploadedImage {
  filename: string;
  url: string;
}
