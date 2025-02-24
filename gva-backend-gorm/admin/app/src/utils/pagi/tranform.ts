import type { CursorParam } from '#/api/types';

import { defaultQuery, type QueryPagi } from './query';

export const urlQueryStringify = (
  obj: Record<string, unknown>,
  prefix?: string,
) => {
  const pairs: string[] = [];
  for (const key in obj) {
    if (!Object.prototype.hasOwnProperty.call(obj, key)) {
      continue;
    }

    const value = obj[key];
    const enkey = encodeURIComponent(key);
    let pair: string | undefined;
    // skip null or undefined
    if (value === null || value === undefined) {
      continue;
    }

    if (typeof value === 'object') {
      pair =
        Array.isArray(value) && prefix === ''
          ? `${enkey}=${encodeURIComponent(value.join(','))}`
          : urlQueryStringify(
              value as any,
              prefix ? `${prefix}[${enkey}]` : enkey,
            );
    } else {
      pair = prefix
        ? `${prefix}[${enkey}]=${encodeURIComponent(value as any)}`
        : `${enkey}=${encodeURIComponent(value as any)}`;
    }
    if (pair) {
      pairs.push(pair);
    }
  }
  return pairs.join('&');
};

interface RawFilter {
  filters: Record<string, any>;
  limit: number;
  page: number;
  search: string;
  selects: string;
  sorts: string;
}

export const urlQueryParse = (query: string): RawFilter => {
  const result: Record<string, unknown> = {};

  // Remove leading '?' if it exists
  query = query.startsWith('?') ? query.slice(1) : query;

  // Split the query string into key-value pairs
  const pairs = query.split('&');

  pairs.forEach((pair) => {
    const [encodedKey, encodedValue] = pair.split('=');

    // Decode the key and value
    const key = decodeURIComponent(encodedKey || '');
    const value = decodeURIComponent(encodedValue || '');

    // Handle nested keys (e.g., "key[innerKey]=value")
    const keyParts = key.split('[').map((part) => part.replace(']', ''));

    let current = result;

    keyParts.forEach((part, index) => {
      if (index === keyParts.length - 1) {
        // Last part, assign the value
        current[part] = value;
      } else {
        // If the part doesn't exist, create an empty object or array
        if (!current[part]) {
          current[part] = Number.isNaN(Number(keyParts[index + 1])) ? {} : [];
        }
        current = current[part] as Record<string, unknown>;
      }
    });
  });

  return {
    filters: result.filters as Record<string, any>,
    limit: +(result?.limit as string),
    page: +(result?.page as string),
    search: (result?.search as string) ?? '',
    selects: (result?.selects as string) ?? '',
    sorts: result?.sorts ?? '',
  } as RawFilter;
};

export const parseQueryPagi = <T extends string>(
  partialQuery?: Partial<QueryPagi<T>>,
) => {
  const query = { ...defaultQuery<T>(), ...partialQuery };

  // payload that backend can understand
  const payload: RawFilter = {
    filters: {} as Record<any, any>,
    limit: query.limit,
    page: query.page,
    search: query.search,
    selects: query.selects.join(','),
    sorts: '',
  };

  if (query.filters) {
    for (const f of query.filters) {
      if (!payload.filters[f.column]) {
        payload.filters[String(f.column)] = {};
      }

      if (!payload.filters[f.operation]) {
        payload.filters[f.column][f.operation] = f.value;
      }
    }
  }

  if (query.sorts) {
    payload.sorts = query.sorts
      .map((s) => {
        const direction = s.direction === 'desc' ? '-' : '';
        const sort = `${direction}${s.column.toString()}`;
        return sort;
      })
      .join(',');
  }

  return payload;
};

// set partial value of default
export const defaultCursor = <Key extends string>(
  raw?: Partial<CursorParam<Key>>,
) => {
  const defaultQ: CursorParam<Key> = {
    filters: [],
    first: 10,
    sorts: [],
  };

  if (!raw) {
    return defaultQ;
  }

  const c: CursorParam<Key> = {
    filters: raw?.filters || defaultQ.filters,
    first: raw?.first ?? defaultQ.first,
    last: raw?.last ?? defaultQ.last,
    before: raw?.before ?? defaultQ.before,
    after: raw?.after ?? defaultQ.after,
    sorts: raw?.sorts ?? defaultQ.sorts,
  };

  return c;
};

export const parseQueryCursor = <T extends string>(
  partialCursor?: Partial<CursorParam<T>>,
) => {
  const cursor = { ...defaultCursor<T>(), ...partialCursor };

  const payload = {
    sorts: '',
    first: cursor.first,
    last: cursor.last,
    after: cursor.after,
    before: cursor.before,
    filters: {} as Record<any, any>,
  };

  if (cursor.filters) {
    for (const f of cursor.filters) {
      if (!payload.filters[String(f.column)]) {
        payload.filters[String(f.column)] = {};
      }

      if (!payload.filters[f.operation]) {
        payload.filters[f.column][f.operation] = f.value;
      }
    }
  }

  if (cursor.sorts) {
    payload.sorts = cursor.sorts
      .map((s) => {
        const direction = s.direction === 'desc' ? '-' : '';
        const sort = `${direction}${s.column.toString()}`?.toString();
        return sort;
      })
      .join(',');
  }
  return payload;
};

export const reverseParseQueryPagi = <T extends string>(payload: RawFilter) => {
  const query: Partial<QueryPagi<T>> = {
    limit: payload.limit,
    page: payload.page,
    search: payload.search,
    selects: payload.selects
      .split(',')
      .filter((e) => ['list', 'totalCount'].includes(e)) as any,
    sorts: payload.sorts
      ? payload.sorts.split(',').map((sort) => {
          const direction = sort.startsWith('-') ? 'desc' : 'asc';
          const column = sort.replace(/^-/, '') as T;
          return { column, direction };
        })
      : [],
  };

  // Rebuild filters
  if (payload.filters) {
    query.filters = Object.entries(payload.filters).flatMap(
      ([column, operations]) => {
        return Object.entries(operations).map(([operation, value]) => ({
          column,
          operation,
          value,
        }));
      },
    ) as any;
  }

  return query;
};

export const updateUrlQuerySearch = (queryValue: string) => {
  const hash = window.location.href.split('#')?.[1] || '';
  const newUrl = `${window.location.pathname}?${queryValue}#${hash}`;
  window.history.replaceState(window.history.state, '', newUrl);
};
