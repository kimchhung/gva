import request from '@/axios'
import { createRouter, getRouters } from './authorization'

export const getServerTime = () => {
  return request.get<string>({ url: '/now' })
}

export const api = () => ({
  getRouters,
  createRouter,
  getServerTime
})
