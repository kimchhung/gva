export type SuccessRes<DataT, MetaT> = {
  code: number
  message: string

  // succes data
  data: DataT
  meta?: MetaT
  success: true
}

export type FailedRes<T> = T & {
  code: number
  message: string
  success: false
}

type UnionRes<DataT, MetaT> = {
  data?: DataT
  meta?: MetaT
  code: number
  message: string
  success: false
}

// Union type for API response
export type APIRes<DataT, MetaT = any, FailedT extends Recordable = {}> =
  | UnionRes<DataT, MetaT>
  | (SuccessRes<DataT, MetaT> | FailedRes<FailedT>)
