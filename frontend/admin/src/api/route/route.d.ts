import { Resource } from '.'

declare global {
  interface API {
    route: Resource
  }
}
