import req, { useAPI } from '@/axios'
import { createQueryPayload } from '@/hooks/web/usePagi'
import { CRUD } from '../node'
import { GetManyRoute, MenuRoute } from './types'

export class MenuAPI extends CRUD<MenuRoute> {
  constructor(name: string) {
    super(name)
  }

  getMany({ query, opt }: GetManyRoute) {
    return useAPI({
      fn: () => req.get<MenuRoute[]>({ url: this.base, params: createQueryPayload(query) }),
      opt
    })
  }
}

export const module: APIModule = {
  name: 'menu',
  resource: new MenuAPI('menu')
}
