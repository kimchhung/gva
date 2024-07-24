import { APIRes } from '@/api/types'
import { CONTENT_TYPE } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { ElNotification } from 'element-plus'
import service from './service'

const defaultOnError = (err: Error) => {
  console.log(err)
  ElNotification.error({
    message: err.message
  })
}

type RequestConfig = {
  onError?: (err: Error, onError: typeof defaultOnError) => void
  onFinally?: () => void
}

const requestConfig: RequestConfig = {}

const request = async <T = any, R = AxiosResponse<T>, D = any>(
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
      validateStatus: (s) => s >= 500
    })

    const { data } = resp as AxiosResponse<T, R>
    if ((data as APIRes<any>)?.code < 0) {
      throw new Error((data as APIRes<any>).message)
    }

    return [data, null, resp] as const
  } catch (error) {
    //  validateStatus: (s) => s >= 500, when status >= 500
    let err = error instanceof Error ? error : new Error(error as any)
    let resp: AxiosResponse<any, any> | undefined = undefined

    if (error instanceof AxiosError) {
      const axiosError = error
      if (axiosError.response) {
        resp = axiosError.response
        err = new Error(
          `Server error: ${[
            axiosError.response.status,
            axiosError.response.statusText,
            (axiosError.response as any)?.code,
            (axiosError.response as any)?.message
          ]
            .filter(Boolean)
            .join(' ')}`
        )
      } else if (axiosError.request) {
        // The request was made but no response was received
        err = new Error('Network error. Please check your internet connection.')
      } else {
        err = new Error('An unexpected error occurred. Please try again later.')
      }
    }

    if (opt?.onError) {
      opt.onError(err, defaultOnError)
    } else {
      defaultOnError(err)
    }

    return [null, err, resp] as const
  }
}

export const req = {
  get: <T = any, R = AxiosResponse<T>, D = any>(option: AxiosConfig) => {
    return request<T, R, D>({ method: 'get', ...option })
  },
  post: <T = any, R = AxiosResponse<T>, D = any>(option: AxiosConfig) => {
    return request<T, R, D>({ method: 'post', ...option })
  },
  delete: <T = any, R = AxiosResponse<T>, D = any>(option: AxiosConfig) => {
    return request<T, R, D>({ method: 'delete', ...option })
  },
  put: <T = any, R = AxiosResponse<T>, D = any>(option: AxiosConfig) => {
    return request<T, R, D>({ method: 'put', ...option })
  },
  patch: <T = any, R = AxiosResponse<T>, D = any>(option: AxiosConfig) => {
    return request<T, R, D>({ method: 'patch', ...option })
  },
  cancelRequest: (...urls: string[]) => {
    return service.cancelRequest(...urls)
  },
  cancelAllRequest: () => {
    return service.cancelAllRequest()
  }
}
