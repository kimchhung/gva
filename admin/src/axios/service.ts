import axios, { AxiosError } from 'axios'
import { defaultRequestInterceptors, defaultResponseInterceptors } from './config'

import { REQUEST_TIMEOUT } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { ElMessage } from 'element-plus'
import { AxiosInstance, AxiosResponse, InternalAxiosRequestConfig, RequestConfig } from './types'

export const PATH_URL = import.meta.env.VITE_API_BASE_PATH

const abortControllerMap: Map<string, AbortController> = new Map()

const axiosInstance: AxiosInstance = axios.create({
  timeout: REQUEST_TIMEOUT,
  baseURL: PATH_URL
})

axiosInstance.interceptors.request.use((res: InternalAxiosRequestConfig) => {
  const controller = new AbortController()
  const url = res.url || ''
  res.signal = controller.signal
  abortControllerMap.set(
    import.meta.env.VITE_USE_MOCK === 'true' ? url.replace('/mock', '') : url,
    controller
  )

  return res
})

axiosInstance.interceptors.response.use(
  (res: AxiosResponse) => {
    const url = res.config.url || ''
    abortControllerMap.delete(url)
    // Can't do any processing here, otherwise the interceptors in the back will not get a complete context

    return res
  },
  (error: AxiosError) => {
    if ([401, 402, 403].includes(error.response?.status ?? 0)) {
      ElMessage.error((error.response?.data as any).message)
      const userStore = useAdminStoreWithOut()
      userStore.logout()
    } else {
      console.log('err： ' + error) // for debug
      ElMessage.error(error.message)
    }

    return Promise.reject(error)
  }
)

axiosInstance.interceptors.request.use(defaultRequestInterceptors)
axiosInstance.interceptors.response.use(defaultResponseInterceptors)

const service = {
  request: (config: RequestConfig) => {
    return new Promise((resolve, reject) => {
      if (config.interceptors?.requestInterceptors) {
        config = config.interceptors.requestInterceptors(config as any)
      }

      axiosInstance
        .request(config)
        .then((res) => {
          resolve(res)
        })
        .catch((err: any) => {
          reject(err)
        })
    })
  },
  cancelRequest: (url: string | string[]) => {
    const urlList = Array.isArray(url) ? url : [url]
    for (const _url of urlList) {
      abortControllerMap.get(_url)?.abort()
      abortControllerMap.delete(_url)
    }
  },
  cancelAllRequest() {
    for (const [_, controller] of abortControllerMap) {
      controller.abort()
    }
    abortControllerMap.clear()
  }
}

export default service
