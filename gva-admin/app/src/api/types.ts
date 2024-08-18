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
