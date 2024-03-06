import { CONTENT_TYPE, SUCCESS_CODE } from '@/constants'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import service from './service'

type Loading = Record<string, boolean>

type FetchOption<L = Loading> = {
  loading?: L

  /**
   * need this case: form = reactive({isloading:true})
   * @description loading=form, loadingKey = "isloading"
   */
  loadingKey?: keyof L

  onFinally?: () => void
  onError?: (error) => void
}

type FetchFunc<T> = () => Promise<IResponse<T>>

/** api response tranformer
 * @returns [data,error,reponse]
 */
export const useApi = async <T, E = any, L = Loading>(fn: FetchFunc<T>, opt?: FetchOption<L>) => {
  const setIsLoading = (bool: boolean) => {
    if (!opt?.loading) return
    //@ts-ignore
    opt.loading[`${opt.loadingKey ?? 'value'}`] = bool
  }

  try {
    setIsLoading(true)
    const resp = await fn()
    if (resp.code !== SUCCESS_CODE) {
      throw new Error(resp.message, {
        cause: resp
      })
    }

    return [resp.data, null, resp] as const
  } catch (error) {
    opt?.onError?.(error)
    return [null, error as E, null, null] as const
  } finally {
    opt?.onFinally?.()
    setIsLoading(false)
  }
}

const request = async (option: AxiosConfig) => {
  const { headers, responseType, ...more } = option

  const userStore = useAdminStoreWithOut()
  const resp = await service.request({
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

export default {
  get: <T = any>(option: AxiosConfig) => {
    return request({ method: 'get', ...option }) as Promise<IResponse<T>>
  },
  post: <T = any>(option: AxiosConfig) => {
    return request({ method: 'post', ...option }) as Promise<IResponse<T>>
  },
  delete: <T = any>(option: AxiosConfig) => {
    return request({ method: 'delete', ...option }) as Promise<IResponse<T>>
  },
  put: <T = any>(option: AxiosConfig) => {
    return request({ method: 'put', ...option }) as Promise<IResponse<T>>
  },
  patch: <T = any>(option: AxiosConfig) => {
    return request({ method: 'patch', ...option }) as Promise<IResponse<T>>
  },
  cancelRequest: (url: string | string[]) => {
    return service.cancelRequest(url)
  },
  cancelAllRequest: () => {
    return service.cancelAllRequest()
  }
}
