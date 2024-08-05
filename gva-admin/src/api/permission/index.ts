import { ResourceAPI } from '../node'
import { Permission } from './types'

export class PermissionAPI extends ResourceAPI<Permission> {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'permission',
  resource: new PermissionAPI('permission')
}
