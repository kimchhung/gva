import req, { useAPI, UseAPIOption } from '@/axios'
import { authResource } from './auth'
import { nodeResource } from './node'
import { routeResource } from './route'

const getNow = ({ opt }: UseAPIOption) => {
  return useAPI({
    fn: () => req.get<string>({ url: '/now' }),
    opt
  })
}

export const api = {
  now: getNow,

  /* pattern */
  node: nodeResource(),
  // ---------------

  auth: authResource(),
  route: routeResource()
}
