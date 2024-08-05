import { RoleAPI } from '.'

declare global {
  type API = {
    node: RoleAPI
  }
}
