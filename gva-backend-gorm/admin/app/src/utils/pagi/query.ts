import type { FieldString } from './form';

import { cloneDeep, uniqBy } from 'lodash';

export enum Operation {
  Between = 'between',
  BetweenEq = 'betweenEq',
  Contains = 'contains',
  ContainsFold = 'containsFold',
  Equal = 'equal',
  Gt = 'gt',
  Gte = 'gte',
  HasPrefix = 'hasPrefix',
  HasSuffix = 'hasSuffix',
  In = 'in',
  IsNull = 'isNull',
  Lt = 'lt',
  Lte = 'lte',
  NotEqual = 'notEqual',
  NotIn = 'notIn',
}

export type StringValueOperator =
  | Operation.Contains
  | Operation.ContainsFold
  | Operation.Equal
  | Operation.Gt
  | Operation.Gte
  | Operation.HasPrefix
  | Operation.HasSuffix
  | Operation.Lt
  | Operation.Lte
  | Operation.NotEqual;

export type InValueOperator = Operation.In | Operation.NotIn;
export type BooleanValueOperator = Operation.IsNull;

export type PagiValue = boolean | Date[] | null | number | string | string[];
export type PagiSortDirection = 'asc' | 'desc' | null;

export interface StringFilter<Key = string> {
  column: Key;
  operation: `${StringValueOperator}`;
  value: string;
}

export interface InFilter<Key = string> {
  column: Key;
  operation: `${InValueOperator}`;
  value: string[];
}

export interface DateFilter<Key = string> {
  column: Key;
  operation: `${Operation.Between}` | `${Operation.BetweenEq}`;
  value: string[];
}

export interface BoolFilter<Key = string> {
  column: Key;
  operation: `${BooleanValueOperator}`;
  value: boolean;
}

export type SingleFilter<Key = string> =
  | BoolFilter<Key>
  | DateFilter<Key>
  | InFilter<Key>
  | StringFilter<Key>;

export interface SingleSort<T extends string> {
  column: string & T;
  direction: PagiSortDirection;
}

export interface QueryPagi<
  FilterableKey extends string,
  SortableKey extends string = FilterableKey,
> {
  page: number;
  limit: number;
  filters: SingleFilter<FilterableKey>[];
  sorts: SingleSort<SortableKey>[];
  search: string;
  selects: ('list' | 'totalCount')[];
}

// set partial value of default
export const defaultQuery = <T extends string>(raw?: Partial<QueryPagi<T>>) => {
  const defaultQ: QueryPagi<T> = {
    filters: [],
    limit: 200,
    page: 1,
    search: '',
    selects: ['list'],
    sorts: [],
  };

  if (!raw) {
    return defaultQ;
  }

  const q: QueryPagi<T> = {
    filters: raw?.filters || defaultQ.filters,
    limit: raw?.limit ?? defaultQ.limit,
    page: raw?.page ?? defaultQ.page,
    search: raw?.search ?? defaultQ.search,
    selects: raw?.selects ?? defaultQ.selects,
    sorts: raw?.sorts ?? defaultQ.sorts,
  };

  return q;
};

export const defaultQueryPagination = <T extends string>(
  raw?: Partial<QueryPagi<T>>,
) => {
  const defaultQ: QueryPagi<T> = {
    filters: [],
    limit: 20,
    page: 1,
    search: '',
    selects: ['list', 'totalCount'],
    sorts: [],
  };

  if (!raw) {
    return defaultQ;
  }

  const q: QueryPagi<T> = {
    filters: raw?.filters || defaultQ.filters,
    limit: raw?.limit ?? defaultQ.limit,
    page: raw?.page ?? defaultQ.page,
    search: raw?.search ?? defaultQ.search,
    selects: raw?.selects ?? defaultQ.selects,
    sorts: raw?.sorts ?? defaultQ.sorts,
  };

  return q;
};

/**
 * set value = null, to remove
 */
export const setQueryfilter = <
  FilterableKey extends string,
  SortableKey extends string = FilterableKey,
>(
  queryRaw: QueryPagi<FilterableKey, SortableKey>,
  fs: SingleFilter<FilterableKey>[],
) => {
  const query = cloneDeep(queryRaw);
  if (!query.filters || !fs?.length) {
    query.filters = [];
  }

  const filterMap: Record<string, SingleFilter<FilterableKey>> = {};
  fs.forEach((f) => {
    const uniq = `${f.column}-${f.operation}`;
    filterMap[uniq] = f;
  });

  queryRaw.filters = Object.values(filterMap);
  return queryRaw;
};

/**
 * set direction = null, to remove sort
 */
export const setQuerySort = <
  FilterableKey extends string,
  SortableKey extends string = FilterableKey,
>(
  query: QueryPagi<FilterableKey, SortableKey>,
  msort: SingleSort<SortableKey>[],
) => {
  if (!query.sorts) {
    query.sorts = [];
  }
  query.sorts = uniqBy([...msort], 'column');
  return query;
};

export const setQuerySearch = <
  FilterableKey extends string,
  SortableKey extends string = FilterableKey,
>(
  query: QueryPagi<FilterableKey, SortableKey>,
  field: FieldString,
): string => {
  if (!field.value) {
    query.search = '';
    return '';
  }

  const v = field.value.split(',').map((e) => e.trim());
  switch (field.operation) {
    case Operation.ContainsFold: {
      query.search = v.map((e) => `*${e}*`).join(',');
      break;
    }
    case Operation.HasPrefix: {
      query.search = v.map((e) => `${e}*`).join(',');
      break;
    }
    case Operation.HasSuffix: {
      query.search = v.map((e) => `*${e}`).join(',');
      break;
    }
    default: {
      query.search = field.value;
      break;
    }
  }

  return query.search;
};
