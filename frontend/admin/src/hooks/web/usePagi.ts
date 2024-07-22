import { isArray } from 'lodash-es'

export type QueryPagi<T extends Object = any> = {
  offset?: number
  limit: number
  filter?: Partial<Condition<T>> | ConditionOr<T>
  sorts?: string[]
  search?: string
  selects?: ('count' | 'list')[]
}

export type SingleSort<T = Record<string, any>> = {
  column: keyof T
  direction: PagiSortDirection
}

export type PagiValue = string[] | Date[] | string | null | number
export type PagiSortDirection = 'asc' | 'desc' | null

export enum OpName {
  Eq = '$eq',
  Neq = '$neq',
  Gt = '$gt',
  Gte = '$gte',
  Lt = '$lt',
  Lte = '$lte',

  // custom value
  Between = 'between',
  In = 'in',

  Or = '$or',
  IsNull = '$isnull',
  IsNotNull = '$isnotnull'
}

export type OpTypeMap = {
  [OpName.IsNotNull]: boolean
  [OpName.IsNotNull]: boolean
  [OpName.In]: string
  [OpName.Between]: [string, string]
  [OpName.Eq]: string
  [OpName.Neq]: string
  [OpName.Gt]: string
  [OpName.Gte]: string
  [OpName.Lt]: string
  [OpName.Lte]: string
}

export type ConditionOr<T extends object> = Record<OpName.Or, Partial<Condition<T>>[]>

// Adjusted Condition type for clarity and correctness
export type Condition<T extends object, K extends keyof OpTypeMap = any> = Record<
  keyof T,
  Record<K, OpTypeMap[K]> | string
>

export const defaultQuery = () => {
  const q: QueryPagi = {
    offset: 1,
    limit: 25
  }
  return q
}

/**
 * set direction = null, to remove sort
 * */
export const setQuerySort = <T extends object>(query: QueryPagi, msort: SingleSort<T>[]) => {
  const sorts = query?.sorts ?? []

  msort.forEach((sort) => {
    if (!sort.direction) {
      query.sorts = [...sorts.filter((str) => !str.includes(sort.column as any))]

      if (query.sorts?.length == 0) {
        delete query.sorts
      }
    }

    query.sorts = [
      ...sorts.filter((s) => !s.includes(sort.column as any)),
      `${sort.column as any} ${sort.direction}`
    ]
  })

  return query
}

export const createQueryPayload = <T extends object = any>(query: QueryPagi<T>) => {
  const { limit, offset, filter, selects, search, sorts, ...more } = query
  const payload: Recordable = {
    offset: offset,
    limit: limit,
    ...more
  }

  if (search) {
    payload.search = query.search
  }

  if (filter) {
    Object.entries(filter)?.forEach(([column, v]) => {
      if (typeof v == 'object') {
        Object.entries(v).forEach(([op, value]) => {
          if (!payload['filter']) {
            payload['filter'] = {}
          }

          if (!payload['filter'][column]) {
            payload['filter'][column] = {}
          }

          switch (true) {
            case op === OpName.Between:
              if (!isArray(value)) {
                return
              }

              const [startDate, endDate] = value
              payload['filter'][column][OpName.Gte] = startDate
              payload['filter'][column][OpName.Lt] = endDate
              break
            case op === OpName.In:
              if (!isArray(value)) {
                return
              }

              payload['filter'][column][op] = String(value)
              break

            default:
              payload['filter'][column][op] = String(value)
              break
          }
        })
      } else {
        payload['filter'][column] = String(v)
      }
    })
  }

  sorts?.forEach((s) => {
    const [column, direction] = s.split(' ')
    const syntax = direction === 'asc' ? '+' : '-'
    payload.sort = [...payload.sort(payload.sort ?? []), `${syntax}${column}`]
  })

  payload.selects = selects?.join(',') ?? 'list'

  return payload
}
