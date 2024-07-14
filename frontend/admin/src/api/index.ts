import req, { useAPI, UseAPIOption } from '@/axios'
import { App } from 'vue'

const modules = import.meta.glob('./**/index.ts', { eager: true })

export type NowAPI = typeof now

const now = ({ opt }: UseAPIOption) =>
  useAPI({
    fn: () => req.get<string>({ url: '/now' }),
    opt
  })

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
