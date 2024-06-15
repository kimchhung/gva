import request from '@/axios'
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
    data: any[]
  }>({ url: '/mock/user/list', params })
}

export const getTestRoleApi = (params: RoleParams): Promise<IResponse<string[]>> => {
  return request.get({ url: '/mock/role/list2', params })
}
