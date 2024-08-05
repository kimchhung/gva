import { Admin } from '../admin/types'

export type AuthLoginBody = {
  username: string
  password: string
}
export type AuthLoginReq = {
  body: AuthLoginBody
}

export type AuthLoginResp = {
  token: string
  admin: Admin
}
