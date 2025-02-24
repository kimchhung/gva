import type { App } from 'vue';

import { QueryClient } from '@tanstack/vue-query';

import { BaseAPI } from './base';

export const queryClient = new QueryClient();

const modules = import.meta.glob('./**/index.ts', { eager: true });

var _api: API;

export const api = () => {
  return _api;
};

const createApi = () => {
  _api = new BaseAPI() as API;

  for (const path in modules) {
    const { module } = modules[path] as { module: APIModule };
    if (module) {
      (_api as any)[module.name] = module.resource;
    }
  }

  // don't allow mutate
  Object.freeze(_api);
  return _api;
};

export const setupAPI = (app: App<Element>) => {
  if (typeof window === 'undefined') {
    return;
  }

  /* eslint no-var: */

  const apiInstance = createApi() as API;
  app.provide('api', apiInstance);
};
