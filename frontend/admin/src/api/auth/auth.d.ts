import { Resource } from '.'

declare global {
  interface API {
    auth: Resource
  }
}
