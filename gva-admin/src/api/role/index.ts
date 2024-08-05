import { ResourceAPI } from '../node'

export class RoleAPI extends ResourceAPI {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'role',
  resource: new RoleAPI('menu')
}
