import { CONTENT_TYPE, SUCCESS_CODE } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import service from './service'

type Loading = Record<string, boolean>

export type FetchOption<L extends Loading = any> = {
  loading?: L

  /**
   * need this case: form = reactive({isloading:true})
   * @description loading=form, loadingKey = "isloading"
   */
  loadingKey?: keyof L

  onFinally?: () => void
  onError?: (error: Error | AxiosError) => void
}

export type FetchFunc<T, M> = () => Promise<AxiosResponse<APIResponse<T, M>>>

export type UseAPIOption = {
  opt?: FetchOption
}

/** api response tranformer
 * @returns [data,error,reponse]
 */
export const useAPI = async <
  T = any,
  M = any,
  E extends Error | AxiosError = AxiosError,
  L extends Loading = any
>({
  fn,
  opt
}: {
  fn: FetchFunc<T, M>
  opt?: FetchOption<L>
}) => {
  const setIsLoading = (bool: boolean) => {
    if (!opt?.loading) return
    //@ts-ignore
    opt.loading[`${opt.loadingKey ?? 'value'}`] = bool
  }

  try {
    setIsLoading(true)
    const axiosResp = await fn()
    const resp = axiosResp.data

    if (resp?.code !== SUCCESS_CODE) {
      throw new AxiosError(resp.message, axiosResp.statusText, axiosResp.config, resp)
    }

    return [resp, null, axiosResp] as const
  } catch (error) {
    opt?.onError?.(error as E)

    return [null, error as E, null] as const
  } finally {
    opt?.onFinally?.()
    setIsLoading(false)
  }
}

const request = async <T = any, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>) => {
  const { headers, responseType, ...more } = config

  const userStore = useAdminStoreWithOut()
  const resp = await service.request<T, R, D>({
    responseType: responseType,
    headers: {
      'Content-Type': CONTENT_TYPE,
      [userStore.getTokenKey ?? 'Authorization']: `Bearer ${userStore.getToken}`,
      ...headers
    },
    ...more
  })

  return resp
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
