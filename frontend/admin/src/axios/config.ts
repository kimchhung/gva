import { TRANSFORM_REQUEST_DATA } from '@/constants'
import { objToFormData } from '@/utils'
import { InternalAxiosRequestConfig } from 'axios'
import qs from 'qs'

const defaultRequestInterceptors = (config: InternalAxiosRequestConfig) => {
  if (
    config.method === 'post' &&
    config.headers['Content-Type'] === 'application/x-www-form-urlencoded'
  ) {
    config.data = qs.stringify(config.data)
  } else if (
    TRANSFORM_REQUEST_DATA &&
    config.method === 'post' &&
    config.headers['Content-Type'] === 'multipart/form-data'
  ) {
    config.data = objToFormData(config.data)
  }

  if (config.method === 'get' && config.params) {
    let url = config.url as string
    url += '?'
    url += qs.stringify(config.params)

    config.params = {}
    config.url = url
  }

  return config
}

export { defaultRequestInterceptors }
