import { Resource } from '.'

declare global {
  interface API {
    node: Resource
  }
}
