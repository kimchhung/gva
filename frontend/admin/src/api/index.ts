import req, { useAPI, UseAPIOption } from '@/axios'
import { App } from 'vue'

const modules = import.meta.glob<
  true,
  'string',
  {
    default: APIModule
  }
>('./**/index.ts', {
  eager: true
})

export type CreateApi = ReturnType<typeof createApi>

const getNow = ({ opt }: UseAPIOption) => {
  console.log({ modules })

  return useAPI({
    fn: () => req.get<string>({ url: '/now' }),
    opt
  })
}

const createApi = async () => {
  const api = {
    now: getNow
  }

  for (const key in modules) {
    const m = modules[key].default
    if (m?.name && m?.resource) {
      api[m.name] = m.resource
    }
  }
  const [data] = await api['auth']({})
  console.log({ data })
  return api
}

/**
 * const { node, route } = inject('api')
 * route.getMany()
 */
export const setupApi = async (app: App<Element>) => {
  if (typeof window === 'undefined') {
    return
  }

  if (!globalThis.api) {
    globalThis.api = createApi() as any
  }

  app.provide('api', globalThis.api)
  window.api = globalThis.api
}
