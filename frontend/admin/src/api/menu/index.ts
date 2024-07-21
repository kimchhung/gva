import { ResourceAPI } from '../node'
import { MenuRoute } from './types'

export class MenuAPI extends ResourceAPI<MenuRoute> {
  constructor(name: string) {
    super(name)
  }
}

export const module: APIModule = {
  name: 'menu',
  resource: new MenuAPI('menu')
}
