import req, { useAPI } from '@/axios'
import { Admin } from '../admin/types'
import { AuthLoginReq, AuthLoginResp, AuthMe } from './types'

export const authResource = (base = '/auth') => {
  const me = ({ opt }: AuthMe) => {
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
