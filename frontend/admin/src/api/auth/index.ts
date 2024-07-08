import req, { useAPI } from '@/axios'
import { Admin } from '../admin/types'
import { AuthLoginReq, AuthLoginResp, AuthMe } from './types'

export class AuthAPI {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}s`
  }

  me({ opt }: AuthMe) {
    return useAPI({
      fn: () => req.get<Admin>({ url: `${this.base}/me` }),
      opt
    })
  }

  login({ body, opt }: AuthLoginReq) {
    return useAPI({
      fn: () => req.post<AuthLoginResp>({ url: `${this.base}/login`, data: body }),
      opt
    })
  }
}

export const module: APIModule = {
  name: 'auth',
  resource: new AuthAPI('auth')
}
