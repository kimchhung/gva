import { SUCCESS_CODE, TRANSFORM_REQUEST_DATA } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { objToFormData } from '@/utils'
import { ElMessage } from 'element-plus'
import qs from 'qs'
import { AxiosResponse, InternalAxiosRequestConfig } from './types'

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
    // let url = config.url as string
    // url += '?'
    // const keys = Object.keys(config.params)
    // for (const key of keys) {
    //   if (config.params[key] !== void 0 && config.params[key] !== null) {
    //     url += `${key}=${encodeURIComponent(config.params[key])}&`
    //   }
    // }
    // url = url.substring(0, url.length - 1)
    // config.params = {}
    // config.url = url

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
  } else if (response.data.code === SUCCESS_CODE) {
    return response.data
  } else {
    ElMessage.error(response?.data?.message)
    if ([401, 402, 403].includes(response.status)) {
      const userStore = useAdminStoreWithOut()
      userStore.logout()
    }
  }
}

export { defaultRequestInterceptors, defaultResponseInterceptors }
