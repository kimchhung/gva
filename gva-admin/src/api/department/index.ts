import { ResourceAPI } from '../node'
import { Department } from './types'

export class DepartmentAPI extends ResourceAPI<Department> {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'department',
  resource: new DepartmentAPI('department')
}
