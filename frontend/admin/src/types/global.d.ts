import { RawAxiosRequestHeaders } from 'axios'
declare global {
  declare type Fn<T = any> = {
    (...arg: T[]): T
  }

  declare type Nullable<T> = T | null

  declare type ElRef<T extends HTMLElement = HTMLDivElement> = Nullable<T>

  declare type Recordable<T = any, K = string> = Record<K extends null | undefined ? string : K, T>
  declare type RecordWithID = Recordable & {
    id: string
  }

  declare type RecordEdgeSelfRelation<T extends Recordable, K extends keyof T> = Recordable & {
    id: string
    parent: RecordWithSelfRelation
    edges: {
      children: RecordWithSelfRelation[]
    }
    [K]: T[K]
  }

  declare type RemoveReadonly<T> = {
    -readonly [P in keyof T]: T[P]
  }

  declare type ComponentRef<T> = InstanceType<T>

  declare type LocaleType = 'zh-CN' | 'en'

  declare type TimeoutHandle = ReturnType<typeof setTimeout>
  declare type IntervalHandle = ReturnType<typeof setInterval>

  declare type ElementPlusInfoType = 'success' | 'info' | 'warning' | 'danger'

  declare type LayoutType = 'classic' | 'topLeft' | 'top' | 'cutMenu'

  declare type AxiosContentType =
    | 'application/json'
    | 'application/x-www-form-urlencoded'
    | 'multipart/form-data'
    | 'text/plain'

  declare type AxiosMethod = 'get' | 'post' | 'delete' | 'put' | 'patch'

  declare type AxiosResponseType = 'arraybuffer' | 'blob' | 'document' | 'json' | 'text' | 'stream'

  declare type AxiosConfig = {
    params?: any
    data?: any
    url?: string
    method?: AxiosMethod
    headers?: RawAxiosRequestHeaders
    responseType?: AxiosResponseType
  }

  declare type ThemeTypes = {
    elColorPrimary?: string
    elBorderColor?: string
    leftMenuBorderColor?: string
    leftMenuBgColor?: string
    leftMenuBgLightColor?: string
    leftMenuBgActiveColor?: string
    leftMenuCollapseBgActiveColor?: string
    leftMenuTextColor?: string
    leftMenuTextActiveColor?: string
    logoTitleTextColor?: string
    logoBorderColor?: string
    topHeaderBgColor?: string
    topHeaderTextColor?: string
    topHeaderHoverColor?: string
    topToolBorderColor?: string
  }

  declare type APIResponse<Data = any, Meta = any> = {
    code: number
    message: string
    data: Data
    meta?: Meta
  }

  declare var api: API

  declare interface Window {
    api: API
  }

  declare type APIModule = {
    name: string
    resource: any
  }
}
