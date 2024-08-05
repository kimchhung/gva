import { WithTree } from '#/types/edges'
import { groupBy } from 'lodash-es'
import { Permission } from './types'

export const permissionToTree = (list: Permission[]) => {
  const result = Object.values(groupBy(list, (e) => e.group)).reduce((result, v) => {
    const superPermission = v.find((p) => p.scope.toLowerCase().includes('super'))
    const parent = superPermission as WithTree<Permission>

    if (superPermission) {
      parent.children = v
        .filter((p) => p.id !== superPermission.id)
        .map((p) => ({
          pid: superPermission.id,
          ...p
        }))
      result.push(parent)
    } else {
      result.push(...v)
    }

    return result
  }, [] as WithTree<Permission>[])
  return result
}
