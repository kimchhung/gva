import { PermissionAPI } from '.'

declare global {
  interface API {
    permission: PermissionAPI
  }
}
