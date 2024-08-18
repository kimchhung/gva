import type { APIRes } from './types';

import type { App } from 'vue';

import { req } from '#/utils/axios';

const modules = import.meta.glob('./**/index.ts', { eager: true });

const now = () => req.get<APIRes<string>>({ url: '/now' });
export type NowAPI = typeof now;

const createApi = () => {
  const api = { now };

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

  (globalThis as any).api = createApi() as API;
  app.provide('api', (globalThis as any).api);
};
