import { RawAxiosRequestHeaders } from 'axios'
declare global {

  declare var api: API

  declare interface Window {
    api: API
  }

  declare type APIModule = {
    name: string
    resource: any
  }
}
