import axios, {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse,
  InternalAxiosRequestConfig
} from 'axios'
import { defaultRequestInterceptors, defaultResponseInterceptors } from './config'

import { REQUEST_TIMEOUT } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { ElMessage } from 'element-plus'

export const PATH_URL = import.meta.env.VITE_API_BASE_PATH

const abortControllerMap: Map<string, AbortController> = new Map()

const axiosInstance = axios.create({
  timeout: REQUEST_TIMEOUT,
  baseURL: PATH_URL
})

const interceptor = {
  request: (res: InternalAxiosRequestConfig) => {
    const controller = new AbortController()
    res.signal = controller.signal
    abortControllerMap.set(String(res.url), controller)
    return res
  },
  response: (res: AxiosResponse) => {
    const url = String(res.config.url)
    abortControllerMap.delete(url)
    return res
  },
  responseError: (error: AxiosError) => {
    const apiData = error.response?.data as {
      code: number
      data: any
      message: string
    }

    const text = apiData ? apiData.message : `msg: ${error.message}, code ${error.code}`

    if ([401, 402, 403].includes(error.response?.status ?? 0)) {
      ElMessage.error(text)
      useAdminStoreWithOut().logout()
    } else {
      ElMessage.error(text)
    }

    return Promise.reject(error)
  }
}

axiosInstance.interceptors.request.use(interceptor.request)
axiosInstance.interceptors.response.use(interceptor.response, interceptor.responseError)

axiosInstance.interceptors.request.use(defaultRequestInterceptors)
axiosInstance.interceptors.response.use(defaultResponseInterceptors)

const service = {
  request: <T = any, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>) => {
    return axiosInstance.request<T, R, D>(config)
  },
  cancelRequest: (...urls: string[]) => {
    for (const url of urls) {
      abortControllerMap.get(url)?.abort()
      abortControllerMap.delete(url)
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
