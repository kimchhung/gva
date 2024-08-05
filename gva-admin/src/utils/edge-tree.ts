
export type TreeHelperConfig = {
  id: string
  pid: string
  edges: {
    children: string
  }
}
export const DEFAULT_CONFIG: TreeHelperConfig = {
  id: 'id',
  pid: 'pid',
  edges: {
    children: 'children'
  }
}

const getConfig = (config: Partial<TreeHelperConfig>) => Object.assign({}, DEFAULT_CONFIG, config)


export const edgeTreeToList = <T = any>(tree: any, config: Partial<TreeHelperConfig> = {}): T => {
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  const result: any = [...tree]
  for (let i = 0; i < result.length; i++) {
    if (!result[i][children!]) continue
    result.splice(i + 1, 0, ...result[i][children!])
  }
  return result
}



// tree from list
export const edgelistToTree = <T = any>(
  list: any[],
  config: Partial<TreeHelperConfig> = {}
): T[] => {
  const conf = getConfig(config) as TreeHelperConfig
  const nodeMap = new Map()
  const result: T[] = []
  const { id, edges, pid } = conf
  const { children } = edges

  for (const node of list) {
    node[children] = node[children] || []
    nodeMap.set(node[id], node)
  }
  for (const node of list) {
    const parent = nodeMap.get(node[pid])
    ;(parent ? parent.children : result).push(node)
  }
  return result as T[]
}



export const convertEdgeChildren = <T extends object=any, R = Omit<T, 'edges'>>(
  list: T[],
  mapToChilden = (t: T) => (t as any)?.['edges']?.['children'],
  appRoutes: R[] = []
): R[] => {
  list.forEach((r) => {
    const { edges, ...more } = r as any
    const appRoute = { ...more } as any

    const children = mapToChilden(r)
    if (children && children?.length > 0) {
      appRoute.children = convertEdgeChildren(children) // Pass the list and result to the recursive call
    }

    appRoutes.push(appRoute)
  })

  return appRoutes
}
