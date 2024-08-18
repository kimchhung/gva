import type { InternalAxiosRequestConfig } from 'axios';

const defaultRequestInterceptors = (config: InternalAxiosRequestConfig) => {
  // if (config.method === 'get' && config.params) {
  //   let url = config.url as string;
  //   url += '?';
  //   url += qs.stringify(config.params);

  //   config.params = {};
  //   config.url = url;
  // }

  return config;
};

export { defaultRequestInterceptors };
