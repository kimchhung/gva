import request from '@/axios'

export const getServerTime = () => {
  return request.get<string>({ url: '/now' })
}
