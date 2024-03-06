import request from '@/axios'
import { RouteReponse } from '../role'
import { AdminInfo, LoginResponse } from './types'

type RoleParams = {
  roleName: string
}

export const loginApi = (data: {
  username: string
  password: string
}): Promise<IResponse<LoginResponse>> => {
  return request.post({ url: '/auth/login', data })
}

export const getAdminInfoReq = () =>
  request.get<AdminInfo>({
    url: '/auth/me',
    method: 'get'
  })

export const getAdminListApi = ({ params }: AxiosConfig) => {
  return request.get<{
    code: string
    data: {
      list: any[]
      total: number
    }
  }>({ url: '/mock/user/list', params })
}

export const getAdminRoleRouters = () => {
  return request.get<RouteReponse>({ url: '/admin/route' })
}

export const getTestRoleApi = (params: RoleParams): Promise<IResponse<string[]>> => {
  return request.get({ url: '/mock/role/list2', params })
}
