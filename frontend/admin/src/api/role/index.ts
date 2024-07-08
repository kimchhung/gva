import { CRUD } from '../node'

export class RoleAPI extends CRUD {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'role',
  resource: new RoleAPI('menu')
}
