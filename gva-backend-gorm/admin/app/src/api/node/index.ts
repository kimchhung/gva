/* eslint-disable perfectionist/sort-classes */
import type { AxiosRequestConfig } from 'axios';

import type {
  CreateNode,
  GetManyMeta,
  GetManyNode,
  GetNode,
  Node,
  UpdateNode,
  UpdateNodePartial,
} from './types';

import { type ComputedRef, type MaybeRef, unref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { req, type RequestOption } from '#/utils/axios';
import { parseQueryPagi, urlQueryStringify } from '#/utils/pagi/tranform';

import { queryClient } from '..';

type GetFunc<T> = () => T;

/**
 * @property name shouold always start with '_'  to avoid conflict with external property
 */
export class ResourceAPI<
  T extends object = Node<any>,
  TCreate = Partial<T>,
  TUpdate = Partial<{ body: Partial<T>; id: number }>,
  TUpdatePartial = Partial<{ body: Partial<T>; id: number }>,
> {
  private _base: string;
  private _name: string;
  queryKey: string;

  get baseUrl() {
    return this._base;
  }

  protected _get = (() => {
    const get = ({
      id,
      config,
      opt,
    }: {
      config?: AxiosRequestConfig;
      id: number | string;
      opt?: RequestOption<T>;
    }) => {
      return req.get<Node<T>>({ url: `${this._base}/${id}`, ...config }, opt);
    };

    // ------------ cache management
    const queryKey = (
      id: number | string,
      params?: AxiosRequestConfig['params'],
    ) => [this.queryKey, get.name, id, params];

    get.invalidate = (
      id: number | string,
      params?: AxiosRequestConfig['params'],
      filters?: Parameters<typeof queryClient.invalidateQueries>[0],
      opt?: Parameters<typeof queryClient.invalidateQueries>[1],
    ) => {
      const filtersOpt = { queryKey: queryKey(id, params), ...filters };
      queryClient.invalidateQueries(filtersOpt, opt);
    };

    get.useQuery = ({
      id,
      opt,
      queryOpt,
      queryClient,
    }: {
      id: number | string;
      opt?: RequestOption<T>;
      queryClient?: Parameters<typeof useQuery>[1];
      queryOpt?: Partial<Parameters<typeof useQuery>[0]>;
    }) => {
      const optBase: Parameters<typeof useQuery>[0] = {
        queryFn: async () => {
          const [res, err] = await get({ id, opt });
          if (err) throw err;
          return res;
        },
        queryKey: queryKey(id),
        ...queryOpt,
      };
      return useQuery(optBase, queryClient);
    };
    // ------------ cache management end
    return get;
  })();

  protected _getMany = (() => {
    const getMany = ({
      opt,
      query,
    }: {
      opt?: RequestOption<Node<T[]>>;
      query?: GetManyNode<T>['query'];
    }) => {
      const payload = parseQueryPagi(query);
      const params = urlQueryStringify(payload as any);

      return req.get<Node<T>[], GetManyMeta>(
        { url: `${this._base}?${params}` },
        opt,
      );
    };

    const queryKey = (keys?: any[]) => [
      this._name,
      getMany.name,
      ...(keys ?? []),
    ];

    getMany.invalidate = (
      keys?: any[],
      filters?: Parameters<typeof queryClient.invalidateQueries>[0],
      opt?: Parameters<typeof queryClient.invalidateQueries>[1],
    ) => {
      const filtersOpt = { queryKey: queryKey(keys), ...filters };
      queryClient.invalidateQueries(filtersOpt, opt);
    };

    getMany.useQuery = ({
      query,
      opt,
      queryOpt,
      queryClient,
    }: {
      opt?: RequestOption<Node<T[]>>;
      query?:
        | ComputedRef<GetManyNode<T>['query']>
        | GetFunc<GetManyNode<T>['query']>
        | MaybeRef<GetManyNode<T>['query']>;
      queryClient?: Parameters<typeof useQuery>[1];
      queryOpt?: Omit<Parameters<typeof useQuery>[0], 'queryFn'>;
    } = {}) => {
      const valueQuery = typeof query === 'function' ? query() : query;

      return useQuery(
        {
          staleTime: 1000 * 60 * 5,
          queryKey: queryKey([valueQuery]),
          queryFn: async () => {
            const [res, err] = await getMany({
              query: unref(typeof query === 'function' ? query() : query),
              opt,
            });
            if (err) throw err;
            return res;
          },
          ...queryOpt,
        },
        queryClient,
      );
    };

    return getMany;
  })();

  constructor(name: string) {
    this._name = name;
    this._base = `/${this._name}`;
    this.queryKey = name;
  }

  protected _create({
    body,
    opt,
  }: { opt?: RequestOption<Node<T>> } & CreateNode<TCreate>) {
    return req.post<Node<T>>({ data: body, url: this._base }, opt);
  }

  protected _delete({ id, opt }: { opt?: RequestOption<Node<T>> } & GetNode) {
    return req.delete<Node<T>>({ url: `${this._base}/${id}` }, opt);
  }

  // whole data is required
  protected _update({
    body,
    id,
    opt,
  }: { opt?: RequestOption<Node<T>> } & UpdateNode<TUpdate>) {
    return req.put<Node<T>>({ data: body, url: `${this._base}/${id}` }, opt);
  }

  protected _updatePartial({
    body,
    id,
    opt,
  }: { opt?: RequestOption<Node<T>> } & UpdateNodePartial<TUpdatePartial>) {
    return req.patch<Node<T>>({ data: body, url: `${this._base}/${id}` }, opt);
  }
}
