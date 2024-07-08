import { MenuAPI } from '.'

declare global {
  interface API {
    menu: MenuAPI
  }
}
