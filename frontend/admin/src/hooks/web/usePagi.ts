import { isArray } from 'lodash-es'
import * as qs from 'qs'

type ExtractStringType<T> = T extends `${infer U}` ? U : string
type PagiSort<T extends Object, S extends string = ExtractStringType<keyof T>> = `-${S}` | `${S}`
type PagiSelects = ('count' | 'list ')[]
type PagiSelect<T> = keyof T[]

export class QueryUrl<T extends Object = any> {
  offset?: number = 0
  limit?: number = 25
  filter?: Partial<Condition<T>> | ConditionOr<T>
  sorts?: PagiSort<T>[]
  search?: string
  select?: PagiSelect<T>
  selects?: PagiSelects;
  [key: string]: any

  constructor(query: number | Partial<QueryUrl<T>> = 25) {
    switch (typeof query) {
      case 'number':
        this.limit = query || 25
        break

      case 'object':
        const { limit, offset, ...rest } = query || {}
        this.limit = limit || 25
        this.offset = offset || 0
        Object.assign(this, rest)
        break
      default:
        break
    }
  }

  params() {
    const { filter, ...rest } = this
    return {
      ...rest,
      filter: filter ? useEncoded().encode(filter) : undefined
    }
  }

  string() {
    return qs.stringify({ ...this })
  }

  fromUrl(paramString: string): void {
    const parsedParams = qs.parse(paramString)
    if (typeof parsedParams === 'object') {
      if (typeof parsedParams.filter === 'string') {
        parsedParams.filter = useEncoded().decode(parsedParams.filter)
      }
      Object.assign(this, parsedParams)
    }
  }
}

export enum OpName {
  Eq = '$eq',
  Neq = '$neq',
  Gt = '$gt',
  Gte = '$gte',
  Lt = '$lt',
  Lte = '$lte',

  // custom value
  Between = '$between',
  In = '$in',

  Or = '$or',
  IsNull = '$isnull',
  IsNotNull = '$isnotnull'
}

export type OpTypeMap = {
  [OpName.IsNotNull]: boolean
  [OpName.IsNotNull]: boolean
  [OpName.In]: ColumnValue[]
  [OpName.Between]: [ColumnValue, ColumnValue]
  [OpName.Eq]: ColumnValue
  [OpName.Neq]: ColumnValue
  [OpName.Gt]: ColumnValue
  [OpName.Gte]: ColumnValue
  [OpName.Lt]: ColumnValue
  [OpName.Lte]: ColumnValue
}

export type ColumnValue = string | number | Date

export type ConditionOr<T extends object> = Record<OpName.Or, Partial<Condition<T>>[]>

// Adjusted Condition type for clarity and correctness
export type Condition<T extends object> = Record<keyof T, ColumnValue | Partial<OpTypeMap>>

export const defaultQuery = () => {
  return new QueryUrl()
}

const tranfromOperation = <T extends object>(query: QueryUrl<T>) => {
  const newQuery = new QueryUrl(query)
  if (newQuery.filter) {
    const filterColumn = newQuery.filter || {}
    Object.entries(newQuery.filter)?.forEach(([column, v]) => {
      switch (typeof v) {
        case 'object':
          Object.entries(v).forEach(([op, value]) => {
            if (!filterColumn[column]) {
              filterColumn[column] = {}
            }

            switch (true) {
              case op === OpName.Between:
                if (!isArray(value)) {
                  return
                }

                const [startDate, endDate] = value
                filterColumn[column][OpName.Gte] = startDate
                filterColumn[column][OpName.Lt] = endDate
                break
              case op === OpName.In:
                if (!isArray(value)) {
                  return
                }

                filterColumn[column][op] = value
                break

              default:
                filterColumn[column][op] = value
                break
            }
          })
          break

        default:
          filterColumn[column] = v
          break
      }
    })

    newQuery.filter = filterColumn
  }

  return newQuery
}

export const parseQuery = <T extends object = any>(query: QueryUrl<T>) => {
  const tranformed = tranfromOperation(query)
  return tranformed.params()
}

const safeUrl = (text: string) => text.replace(/\+/g, '-').replace(/\//g, '_')

const fromSafeUrl = (text: string) => text.replace(/-/g, '+').replace(/_/g, '/')

const useEncoded = (base = 'base64url') => {
  const isbase64Url = base === 'base64url'

  const decode = <T = any>(encoded: string) => {
    const decodedData = isbase64Url ? fromSafeUrl(encoded) : encoded
    return JSON.parse(atob(decodedData)) as T
  }

  const encode = (obj) => {
    const data: Recordable = {}
    if (typeof obj !== 'object') {
      data.value = obj
    } else {
      Object.assign(data, obj)
    }

    const jsonData = JSON.stringify(data)
    const encodedData = btoa(jsonData)
    return isbase64Url ? safeUrl(encodedData) : encodedData
  }

  return {
    encode,
    decode
  }
}
