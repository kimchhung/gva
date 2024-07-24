import { req } from '@/axios'
import { Admin } from '../admin/types'
import { AuthLoginReq, AuthLoginResp } from './types'

export class AuthAPI {
  name: string
  base: string

  constructor(name: string) {
    this.name = name
    this.base = `/${this.name.replace('/', '')}`
  }

  me() {
    return req.get<Admin>({ url: `${this.base}/me` })
  }

  login({ body }: AuthLoginReq) {
    return req.post<AuthLoginResp>({ url: `${this.base}/login`, data: body })
  }
}

export const module: APIModule = {
  name: 'auth',
  resource: new AuthAPI('auth')
}
