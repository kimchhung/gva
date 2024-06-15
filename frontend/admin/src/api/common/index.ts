import request from '@/axios'

// Get all dictionaries
export const getDictApi = () => {
  return request.get({ url: '/mock/dict/list' })
}

// Simten to get a dictionary
export const getDictOneApi = async () => {
  return request.get({ url: '/mock/dict/one' })
}
