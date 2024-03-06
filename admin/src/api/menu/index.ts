import request from '@/axios'
import { AdminRoute } from '../role'

export const getRouteListApi = () => {
  return request.get<{
    list: AdminRoute[]
  }>({ url: '/route' })
}
