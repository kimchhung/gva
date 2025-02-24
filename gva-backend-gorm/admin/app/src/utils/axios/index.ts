import type { APIRes, FailedRes, SuccessRes } from '#/api/types';

import type { Ref } from 'vue';

import { preferencesManager } from '@vben/preferences';
import { useAccessStore } from '@vben/stores';

import { notification } from 'ant-design-vue';
import { AxiosError, type AxiosRequestConfig, type AxiosResponse } from 'axios';

import { CONTENT_TYPE } from '#/constants';
import { useAuthStore } from '#/store';

import service from './service';

const ignoreRoutes = new Set(['/auth/login']);

const defaultOnError = (err: Error) => {
  notification.error({
    message: err.message,
  });

  if (
    err instanceof AxiosError &&
    [401, 402, 403].includes(err.response?.status ?? 0)
  ) {
    if (ignoreRoutes.has(window.location.pathname)) return;

    useAuthStore().logout();
  }
};

const tranformError = (error: any) => {
  const err = error instanceof Error ? error : new Error(error as any);
  let resp: AxiosResponse<any, any> | undefined;

  if (error instanceof AxiosError) {
    const axiosError = error;
    if (axiosError.response) {
      resp = axiosError.response;
      const hasRespMsg = !!(resp as any)?.data?.message;
      err.message = hasRespMsg
        ? (resp as any)?.data?.message
        : `Server error: ${[
            axiosError.response.status,
            axiosError.response.statusText,
          ].join(' ')}`;
    } else if (axiosError.request) {
      // The request was made but no response was received
      err.message = 'Network error. Please check your internet connection.';
    } else {
      // what kind of error huh?
      err.message = 'An unexpected error occurred. Please try again later.';
    }
  }

  console.error(error);
  return [err, resp] as const;
};

export interface RequestOption<T> {
  /**
   *
   * @param err
   * @param handleError handle default error
   * @returns
   */
  onError?: (err: Error, handleError: typeof defaultOnError) => void;
  onSuccess?: (data: SuccessRes<T, any>) => void;
  onFinally?: () => void;
  loading?: Ref<boolean>;
}

const defaultRequestOption = (opt?: RequestOption<any>) => {
  const option = opt ?? {};
  option.onError = option.onError ?? defaultOnError;
  return opt;
};

const request = async <
  DataT,
  MetaT = any,
  FailedT extends Record<any, string> = any,
  T extends APIRes<DataT, MetaT, FailedT> = any,
  R = AxiosResponse<T>,
  D = any,
>(
  config: AxiosRequestConfig<D>,
  opt?: RequestOption<T>,
) => {
  opt = defaultRequestOption(opt);
  const { headers, responseType, ...more } = config;
  const { app } = preferencesManager.getPreferences();
  const accessStore = useAccessStore();

  const setLoading = (isLoading: boolean) => {
    if (opt?.loading) {
      opt.loading.value = isLoading;
    }
  };
  setLoading(true);
  try {
    const resp = await service.request<T, R, D>({
      headers: {
        'Accept-Language': app.locale,
        Authorization: `Bearer ${accessStore.accessToken}`,
        'Content-Type': CONTENT_TYPE,
        ...headers,
      },
      responseType,
      ...more,
      validateStatus: (s) => s < 500,
    });

    const { data } = resp as AxiosResponse<T, R>;
    if (!data) {
      const res = resp as AxiosResponse;
      throw new AxiosError(
        'invalid request',
        res.statusText,
        res.config,
        res.request,
        res,
      );
    }

    if (data.code === 0) {
      (data as SuccessRes<DataT, MetaT>).success = true;
      opt?.onSuccess?.(data as any);
    } else {
      (data as FailedRes<FailedT>).success = false;
      const res = resp as AxiosResponse;
      throw new AxiosError(
        data.message,
        res.statusText,
        res.config,
        res.request,
        res,
      );
    }

    return [data as SuccessRes<DataT, MetaT>, null, resp] as const;
  } catch (error) {
    const [err, resp] = tranformError(error);

    if (opt?.onError) {
      opt.onError(err, defaultOnError);
    } else {
      defaultOnError(err);
    }

    return [null, err, resp] as const;
  } finally {
    opt?.onFinally?.();
    setLoading(false);
  }
};

export const req = {
  cancelAllRequest: () => {
    return service.cancelAllRequest();
  },
  cancelRequest: (...urls: string[]) => {
    return service.cancelRequest(...urls);
  },
  delete: <D = any, M = any, F extends Record<string, any> = any>(
    config: AxiosRequestConfig,
    opt?: RequestOption<D>,
  ) => {
    return request<D, M, F>({ method: 'delete', ...config }, opt);
  },
  get: <D = any, M = any, F extends Record<string, any> = any>(
    config: AxiosRequestConfig,
    opt?: RequestOption<D>,
  ) => {
    return request<D, M, F>({ method: 'get', ...config }, opt);
  },
  patch: <D = any, M = any, F extends Record<string, any> = any>(
    config: AxiosRequestConfig,
    opt?: RequestOption<D>,
  ) => {
    return request<D, M, F>({ method: 'patch', ...config }, opt);
  },
  post: <D = any, M = any, F extends Record<string, any> = any>(
    config: AxiosRequestConfig,
    opt?: RequestOption<D>,
  ) => {
    return request<D, M, F>({ method: 'post', ...config }, opt);
  },
  put: <D = any, M = any, F extends Record<string, any> = any>(
    config: AxiosRequestConfig,
    opt?: RequestOption<D>,
  ) => {
    return request<D, M, F>({ method: 'put', ...config }, opt);
  },
};
