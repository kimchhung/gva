import { APIRes, FailedRes, SuccessRes } from '@/api/types'
import { CONTENT_TYPE } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { ElNotification } from 'element-plus'
import service from './service'

const defaultOnError = (err: Error) => {
  ElNotification.error({
    message: err.message
  })

  if (err instanceof AxiosError) {
    if ([401, 402, 403].includes(err.response?.status ?? 0)) {
      useAdminStoreWithOut().logout()
    }
  }
}

const tranformError = (error: any) => {
  const err = error instanceof Error ? error : new Error(error as any)
  let resp: AxiosResponse<any, any> | undefined = undefined

  if (error instanceof AxiosError) {
    const axiosError = error
    if (axiosError.response) {
      resp = axiosError.response
      err.message = `Server error: ${[
        axiosError.response.status,
        axiosError.response.statusText,
        (axiosError.response as any)?.code,
        (axiosError.response as any)?.message
      ]
        .filter(Boolean)
        .join(' ')}`
    } else if (axiosError.request) {
      // The request was made but no response was received
      err.message = 'Network error. Please check your internet connection.'
    } else {
      err.message = 'An unexpected error occurred. Please try again later.'
    }
  }

  return [err, resp] as const
}

type RequestConfig = {
  onError?: (err: Error, onError: typeof defaultOnError) => void
  onFinally?: () => void
}

const requestConfig: RequestConfig = {}

const request = async <
  DataT,
  MetaT = any,
  FailedT extends Recordable<any, string> = any,
  T extends APIRes<DataT, MetaT, FailedT> = any,
  R = AxiosResponse<T>,
  D = any
>(
  config: AxiosRequestConfig<D>,
  opt = requestConfig
) => {
  const { headers, responseType, ...more } = config
  const userStore = useAdminStoreWithOut()
  try {
    const resp = await service.request<T, R, D>({
      responseType: responseType,
      headers: {
        'Content-Type': CONTENT_TYPE,
        [userStore.getTokenKey ?? 'Authorization']: `Bearer ${userStore.getToken}`,
        ...headers
      },
      ...more,
      validateStatus: (s) => s <= 500
    })

    const { data } = resp as AxiosResponse<T, R>

    if (data && data.code === 0) {
      ;(data as SuccessRes<DataT, MetaT>).success = true
    } else {
      ;(data as FailedRes<FailedT>).success = false
      throw new Error(data.message)
    }

    return [data as SuccessRes<DataT, MetaT>, null, resp] as const
  } catch (error) {
    const [err, resp] = tranformError(error)

    if (opt?.onError) {
      opt.onError(err, defaultOnError)
    } else {
      defaultOnError(err)
    }

    return [null, err, resp] as const
  }
}

export const req = {
  get: <D = any, M = any, F extends Recordable = any>(option: AxiosConfig) => {
    return request<D, M, F>({ method: 'get', ...option })
  },
  post: <D = any, M = any, F extends Recordable = any>(option: AxiosConfig) => {
    return request<D, M, F>({ method: 'post', ...option })
  },
  delete: <D = any, M = any, F extends Recordable = any>(option: AxiosConfig) => {
    return request<D, M, F>({ method: 'delete', ...option })
  },
  put: <D = any, M = any, F extends Recordable = any>(option: AxiosConfig) => {
    return request<D, M, F>({ method: 'put', ...option })
  },
  patch: <D = any, M = any, F extends Recordable = any>(option: AxiosConfig) => {
    return request<D, M, F>({ method: 'patch', ...option })
  },
  cancelRequest: (...urls: string[]) => {
    return service.cancelRequest(...urls)
  },
  cancelAllRequest: () => {
    return service.cancelAllRequest()
  }
}
