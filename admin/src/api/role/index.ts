import request from '@/axios'

export const getRoleListApi = () => {
  return request.get<any>({ url: '/mock/role/table' })
}
