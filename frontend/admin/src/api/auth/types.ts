import { UseAPIOption } from '@/axios'
import { Admin } from '../admin/types'

export type AuthLoginBody = {
  username: string
  password: string
}
export type AuthLoginReq = UseAPIOption & {
  body: AuthLoginBody
}

export type AuthMe = UseAPIOption & {}

export type AuthLoginResp = {
  token: string
  admin: Admin
}
