import { ResourceAPI } from '../node'
import { Admin } from './types'

export class AdminAPI extends ResourceAPI<Admin> {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'admin',
  resource: new AdminAPI('admin')
}
