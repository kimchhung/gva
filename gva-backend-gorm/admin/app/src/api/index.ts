import type { App } from 'vue';

import { QueryClient } from '@tanstack/vue-query';

import { BaseAPI } from './base';

export const queryClient = new QueryClient();

const modules = import.meta.glob('./**/index.ts', { eager: true });

const createApi = () => {
  var api = new BaseAPI() as API;

  for (const path in modules) {
    const { module } = modules[path] as { module: APIModule };
    if (module) {
      (api as any)[module.name] = module.resource;
    }
  }

  return api as API;
};

/**
 * const { node, route } = inject('api')
 * route.getMany()
 */
export const setupAPI = (app: App<Element>) => {
  if (typeof window === 'undefined') {
    return;
  }

  /* eslint no-var: */

  const api = createApi() as API;
  app.provide('api', api);
  app.config.globalProperties.api = api;
  (window as any).api = api;
};
