import { req } from '@/axios'
import { ResourceAPI } from '../node'
import { MenuRoute } from './types'

export class MenuAPI extends ResourceAPI<MenuRoute> {
  constructor(name: string) {
    super(name)
  }

  enabledList() {
    return req.get<MenuRoute[]>({ url: `${this.base}/enabled-list` })
  }
}

export const module: APIModule = {
  name: 'menu',
  resource: new MenuAPI('menu')
}
