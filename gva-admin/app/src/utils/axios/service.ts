import { useAppConfig } from '@vben/hooks';

import axios, {
  AxiosError,
  type AxiosRequestConfig,
  type AxiosResponse,
  type InternalAxiosRequestConfig,
} from 'axios';

import { REQUEST_TIMEOUT } from '#/constants';
import { useAuthStore } from '#/store';

import { defaultRequestInterceptors } from './config';

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);
const abortControllerMap: Map<string, AbortController> = new Map();

const axiosInstance = axios.create({
  baseURL: apiURL,
  timeout: REQUEST_TIMEOUT,
});

const interceptor = {
  request: (res: InternalAxiosRequestConfig) => {
    const controller = new AbortController();
    res.signal = controller.signal;
    abortControllerMap.set(String(res.url), controller);
    return res;
  },
  response: (res: AxiosResponse) => {
    const url = String(res.config.url);
    abortControllerMap.delete(url);
    return res;
  },
  responseError: (error: AxiosError) => {
    if ([401, 402, 403].includes(error.response?.status ?? 0)) {
      useAuthStore().logout();
    }

    return Promise.reject(error);
  },
};

axiosInstance.interceptors.request.use(interceptor.request);
axiosInstance.interceptors.response.use(
  interceptor.response,
  interceptor.responseError,
);
axiosInstance.interceptors.request.use(defaultRequestInterceptors);

const service = {
  cancelAllRequest() {
    for (const [_, controller] of abortControllerMap) {
      controller.abort();
    }
    abortControllerMap.clear();
  },
  cancelRequest: (...urls: string[]) => {
    for (const url of urls) {
      abortControllerMap.get(url)?.abort();
      abortControllerMap.delete(url);
    }
  },
  request: <T = any, R = AxiosResponse<T>, D = any>(
    config: AxiosRequestConfig<D>,
  ) => {
    return axiosInstance.request<T, R, D>(config);
  },
};

export default service;
