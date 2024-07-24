import { req } from '@/axios'
import { App } from 'vue'
import { APIRes } from './types'

const modules = import.meta.glob('./**/index.ts', { eager: true })

export type NowAPI = typeof now

const now = () => req.get<APIRes<string>>({ url: '/now' })

const createApi = () => {
  const api = { now }

  for (const path in modules) {
    const { module } = modules[path] as { module: APIModule }
    if (module) {
      api[module.name] = module.resource
    }
  }

  return api as API
}

/**
 * const { node, route } = inject('api')
 * route.getMany()
 */
export const setupAPI = (app: App<Element>) => {
  if (typeof window === 'undefined') {
    return
  }

  /* eslint no-var: */
  var api = createApi() as API
  globalThis.api = api
  app.provide('api', api)
}
