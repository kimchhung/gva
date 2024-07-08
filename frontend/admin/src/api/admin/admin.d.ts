import { AdminAPI } from '.'

declare global {
  interface API {
    menu: AdminAPI
  }
}
