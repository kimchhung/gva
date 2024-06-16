import { isArray } from 'lodash-es'

export type QueryPagi = {
  page: number
  limit: number
  filters?: Record<string, Record<OpName, PagiValue> | {}>
  sorts?: string[]
  search?: string
}

export type SingleFilter<T = Record<string, any>> = {
  column: keyof T
  operation: OpName
  value: PagiValue
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
  In = 'in'
}

export const defaultQuery = () => {
  const q: QueryPagi = {
    page: 1,
    limit: 20
  }
  return q
}

/**
 * set value = null, to remove
 * */
export const setQueryfilter = (query: QueryPagi, fs: SingleFilter[]) => {
  if (!query.filters) {
    query.filters = {}
  }

  fs.forEach((f) => {
    if (!query.filters) {
      return
    }

    if (!query.filters[f.column]) {
      query.filters[f.column] = {}
    }

    if (f.value == null) {
      // remove;
      delete query.filters[f.column][f.operation]
      if ((query.filters[f.column] = {})) {
        delete query.filters[f.column]
      }

      return
    }

    query.filters[f.column][f.operation] = f.value
  })
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
}

export const createQueryPayload = (query: QueryPagi) => {
  const { limit, page, ...more } = query
  const payload = {
    offset: (page - 1) * limit,
    limit: limit * page,
    ...more
  } as Recordable

  if (query.search) {
    payload.search = query.search
  }

  if (query.filters) {
    Object.entries(query.filters).forEach(([column, v]) => {
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

            payload['filter'][column][op] = value
            break

          default:
            payload['filter'][column][op] = value
            break
        }
      })
    })
  }

  query.sorts?.forEach((s) => {
    const [column, direction] = s.split(' ')
    const syntax = direction === 'asc' ? '+' : '-'
    payload.sort = [...payload.sort(payload.sort ?? []), `${syntax}${column}`]
  })

  return payload
}
