import req, { useAPI } from '@/axios'
import { Admin } from '../admin/types'
import { AuthLoginReq, AuthLoginResp, AuthMe } from './types'

export type Resource = ReturnType<typeof resource>

export const resource = (base = '/auth') => {
  const me = ({ opt }: AuthMe) => {
    console.log('resource,auth')
    return useAPI({
      fn: () => req.get<Admin>({ url: `${base}/login` }),
      opt
    })
  }

  me.login = ({ body, opt }: AuthLoginReq) => {
    return useAPI({
      fn: () => req.get<AuthLoginResp>({ url: `${base}/login`, data: body }),
      opt
    })
  }

  return me
}

const module: APIModule = {
  name: 'auth',
  resource
}

export default module
