import { TRANSFORM_REQUEST_DATA } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { objToFormData } from '@/utils'
import { AxiosResponse, InternalAxiosRequestConfig } from 'axios'
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

const defaultResponseInterceptors = (response: AxiosResponse) => {
  if (response?.config?.responseType === 'blob') {
    // If it is a file flow, pass it directly
    return response
  }

  if ([401, 402, 403].includes(response.status)) {
    const userStore = useAdminStoreWithOut()
    userStore.logout()
  }

  return response
}

export { defaultRequestInterceptors, defaultResponseInterceptors }
