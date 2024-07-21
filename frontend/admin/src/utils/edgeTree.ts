type TreeHelperConfig = {
  id: string
  pid: string
  edges: {
    children: string
  }
}
const DEFAULT_CONFIG: TreeHelperConfig = {
  id: 'id',
  pid: 'pid',
  edges: {
    children: 'children'
  }
}

const getConfig = (config: Partial<TreeHelperConfig>) => Object.assign({}, DEFAULT_CONFIG, config)

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

export const edgeFindNode = <T = any>(
  tree: any,
  func: Fn,
  config: Partial<TreeHelperConfig> = {}
): T | null => {
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  const list = [...tree]
  for (const node of list) {
    if (func(node)) return node
    node[children!] && list.push(...node[children!])
  }
  return null
}

export const edgeFindNodeAll = <T = any>(
  tree: any,
  func: Fn,
  config: Partial<TreeHelperConfig> = {}
): T[] => {
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  const list = [...tree]
  const result: T[] = []
  for (const node of list) {
    func(node) && result.push(node)
    node[children!] && list.push(...node[children!])
  }
  return result
}

export const findPath = <T = any>(
  tree: any,
  func: Fn,
  config: Partial<TreeHelperConfig> = {}
): T | T[] | null => {
  config = getConfig(config)
  const path: T[] = []
  const list = [...tree]
  const visitedSet = new Set()
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  while (list.length) {
    const node = list[0]
    if (visitedSet.has(node)) {
      path.pop()
      list.shift()
    } else {
      visitedSet.add(node)
      node[children!] && list.unshift(...node[children!])
      path.push(node)
      if (func(node)) {
        return path
      }
    }
  }
  return null
}

export const edgeFindPathAll = (tree: any, func: Fn, config: Partial<TreeHelperConfig> = {}) => {
  config = getConfig(config)
  const path: any[] = []
  const list = [...tree]
  const result: any[] = []
  const visitedSet = new Set()
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  while (list.length) {
    const node = list[0]
    if (visitedSet.has(node)) {
      path.pop()
      list.shift()
    } else {
      visitedSet.add(node)
      node[children!] && list.unshift(...node[children!])
      path.push(node)
      func(node) && result.push([...path])
    }
  }
  return result
}

export const edgeFilter = <T = any>(
  tree: T[],
  func: (n: T) => boolean,
  config: Partial<TreeHelperConfig> = {}
): T[] => {
  const conf = getConfig(config)
  const { edges } = conf
  const { children } = edges
  function listFilter(list: T[]) {
    return list
      .map((node: any) => ({ ...node }))
      .filter((node) => {
        node[children] = node[children] && listFilter(node[children])
        return func(node) || (node[children] && node[children].length)
      })
  }
  return listFilter(tree)
}

export const convertEdgeChildren = <T = any, R = Omit<T, 'edges'>>(
  list: T[],
  mapToChilden = (t: T) => t?.['edges']?.['children'],
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
