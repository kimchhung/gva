import type { Dayjs } from 'dayjs';

import {
  computed,
  type ComputedRef,
  isRef,
  type MaybeRef,
  reactive,
  ref,
  unref,
  type UnwrapRef,
  watch,
} from 'vue';

import { cloneDeep, merge } from 'lodash';

import { useTabbarStore } from '@vben/stores';
import { useRouter } from 'vue-router';
import { useTabbar } from '../../../../packages/effects/layouts/src/basic/tabbar';
import { DATE_TIME_FORMAT, dateUtil } from '../helper/date-util';
import {
  defaultQueryPagination,
  Operation,
  type QueryPagi,
  setQueryfilter,
  setQuerySearch,
  setQuerySort,
  type SingleFilter,
  type SingleSort,
} from './query';
import {
  parseQueryPagi,
  reverseParseQueryPagi,
  urlQueryParse,
  urlQueryStringify,
} from './tranform';

export enum FieldType {
  date = 'date',
  number = 'number',
  select = 'select',
  string = 'string',
}

export interface FieldSelect {
  operation: Operation;
  type: FieldType.select;
  value?: (number | string)[];
}

export interface FieldNumber {
  operation: Operation;
  type: FieldType.number;
  value?: number;
}

export interface FieldDate {
  type: FieldType.date;
  value?: [Dayjs, Dayjs] | [string, string];
}

export interface FieldBase<T = Record<string, any>> {
  column: Key<T>;
  isCompact?: boolean;
  isLock?: boolean;
}

export interface FieldString {
  operation: Operation;
  type: FieldType.string;
  value?: string;
}

export type FieldFilter<T = Record<string, any>> =
  | ({ column: UnwrapRef<keyof T> } & FieldBase<T> & FieldDate)
  | ({ column: UnwrapRef<keyof T> } & FieldBase<T> & FieldNumber)
  | ({ column: UnwrapRef<keyof T> } & FieldBase<T> & FieldSelect)
  | ({ column: UnwrapRef<keyof T> } & FieldBase<T> & FieldString);

interface SortSelect {
  label: ComputedRef<MaybeRef<string> | undefined> | MaybeRef<string>;
  value: string;
  direction: 'asc' | 'desc';
}

// convert type properly, eg -> type date to gt and lte
const parseFilter = <
  T extends Record<string, any>,
  K extends string = Extract<keyof T, string>,
>(
  forms: FieldFilter<T>[],
) => {
  const filters: SingleFilter<K>[] = [];

  forms?.forEach?.((form) => {
    switch (form.type) {
      case FieldType.date: {
        if (!form.value) break;
        const [startDate, endDate] = form.value;
        filters.push({
          column: form.column as any,
          operation: Operation.BetweenEq,
          value: [
            dateUtil(startDate.toString()).utc().format(DATE_TIME_FORMAT),
            dateUtil(endDate.toString()).utc().format(DATE_TIME_FORMAT),
          ],
        });

        break;
      }

      case FieldType.number: {
        filters.push({
          column: form.column as any,
          operation: form.operation as any as Operation,
          value: form?.value as any,
        });

        break;
      }
      default: {
        if (form.value) {
          filters.push({
            column: form.column as any,
            operation: form.operation as any as Operation,
            value: form.value as any,
          });
        }
      }
    }
  });

  return filters;
};

export interface Schema<T = any> {
  filters: FieldFilter<T>[];
  sorts: string[];
  columns: Key<T>[];
  search: FieldString;
}

export const StringOp = {
  Equal: Operation.Equal,
  'Has ': Operation.ContainsFold,
  'Has Prefix': Operation.HasPrefix,
  'Has Suffix': Operation.HasSuffix,
};

export const SelectOp = {
  In: Operation.In,
  'Not In ': Operation.NotIn,
};

export const NumberOp = {
  '<': Operation.Lt,
  '<=': Operation.Lte,
  '>': Operation.Gt,
  '>=': Operation.Gte,
  Equal: Operation.Equal,
  'Not Equal': Operation.NotEqual,
};

const initSchema = <T>(): Schema<T> => ({
  columns: [],
  filters: [],
  search: {
    operation: Operation.ContainsFold,
    type: FieldType.string,
    value: '',
  },
  sorts: [],
});

export type SetterFn<T> = (q: T) => T;
export type Setter<T> = SetterFn<T> | T;

export type GetterFn<T> = () => T;
export type Getter<T> = GetterFn<T> | T;

export const getGetter = <T>(v: Getter<T>): T => {
  if (typeof v === 'function') {
    return (v as GetterFn<T>)();
  }
  return v;
};

export interface SelectOption {
  label: string;
  value: number | string;
}
export type SetOptionFn = (options: SelectOption[]) => void;
export type SelectOptionFn = () => SelectOption[];

class TypeSelect_ {
  options: SelectOption[] | SelectOptionFn;
  value: SelectOption[] | undefined;
  constructor(
    options: SelectOption[] | SelectOptionFn,
    value?: SelectOption[],
  ) {
    this.options = options;
    if (value) this.value = value;
  }
}
export const TypeSelect = (
  options: SelectOption[] | SelectOptionFn,
  value?: SelectOption[],
) => new TypeSelect_(options, value);

export class TypeDate_ {
  range: [Dayjs, Dayjs] | undefined;
  constructor(range?: [Dayjs, Dayjs]) {
    if (range) this.range = range;
  }
  getRange() {
    return this.range;
  }
  hasValue() {
    return this.range?.length === 2;
  }
}

export const TypeDate = (range?: [Dayjs, Dayjs]) => new TypeDate_(range);

export const TypeString = (value?: string) => value || '';

export const TypeNumber = (value?: number) => value || 0;

type ModelValue = number | string | TypeDate_ | TypeSelect_;
export type BaseModel<T> = Record<Key<T>, ModelValue>;

export const isNumber = (value: any) =>
  value instanceof Number || typeof value === 'number';

export const isString = (value: any) =>
  value instanceof String || typeof value === 'string';

export interface Config {
  // enable disable filter button
  button?: {
    isFilter: boolean;
    isReset: boolean;
    isSearch: boolean;
  };
  disableUrlQuery?: boolean;
}

export interface ColumnConfig<T> {
  order?: number;
  label: MaybeRef<string>;

  // require default value, can not remove
  isFixed?: boolean;

  isHidden?: boolean;

  // default true
  isSort?: boolean;

  // default true
  isFilter?: boolean;

  compact?: {
    // default value
    value: FieldFilter<T>;
  };
}

export type GetLocalFunc<T> = (key: Key<T>) => string;

const defaultConfig = (): Config => {
  return {
    button: {
      isFilter: true,
      isReset: true,
      isSearch: true,
    },
    disableUrlQuery: false,
  };
};

export type ModelConfig<T> = Partial<Record<Key<T>, ColumnConfig<T>>>;
const setFieldDefaulConfig = <T extends BaseModel<T>>(
  modelConfig: MaybeRef<ModelConfig<T>> = {},
) => {
  const modelConfigValue = unref(modelConfig);

  for (const key in modelConfigValue) {
    const v = modelConfigValue[key];
    if (!v) continue;
    if (v.isFilter === undefined) v.isFilter = true;
    if (v.isSort === undefined) v.isSort = true;
  }

  if (isRef(modelConfig)) {
    modelConfig.value = modelConfigValue;
  }

  return modelConfig;
};

export interface UseQueryProp<
  T extends Record<string, any>,
  K extends string = Extract<keyof T, string>,
> {
  config?: Config;
  model: T;
  modelConfig?: MaybeRef<ModelConfig<T>>;
  query?: QueryPagi<K>;
  schema?: Schema<T>;
}

export const useQueryForm = <T extends BaseModel<T>>(
  props: UseQueryProp<T>,
) => {
  const model = props.model;
  const schema = Object.assign(initSchema<T>(), { ...props.schema });
  const initialQuery = defaultQueryPagination(props.query);
  const modelConfig = setFieldDefaulConfig(props.modelConfig);
  const config = merge(defaultConfig(), props.config);
  const initialForm = structuredClone(schema);
  const sortOption: SortSelect[] = [];
  const types = new Map<keyof T, FieldType>();
  const selectOption = new Map<keyof T, Getter<SelectOption[]>>();
  const getFieldConfig = (k: keyof T) =>
    isRef(modelConfig) ? modelConfig.value[k] : modelConfig[k];

  if (initialQuery.sorts && initialQuery.sorts?.length > 0) {
    initialForm.sorts = initialQuery.sorts.map(
      (q) => `${q.column} ${q.direction}`,
    );
  }

  for (const key in model) {
    schema?.columns?.push(key);
    initialForm?.columns?.push(key);

    const value = model[key];

    if (getFieldConfig(key)?.isSort) {
      const opt = [
        {
          direction: 'asc' as const,
          label: computed(() => getFieldConfig(key)?.label),
          value: `${key} asc`,
        },
        {
          direction: 'desc' as const,
          label: computed(() => getFieldConfig(key)?.label),
          value: `${key} desc`,
        },
      ];
      sortOption.push(...opt);
    }

    if (getFieldConfig(key)?.isFilter) {
      switch (true) {
        case value instanceof TypeSelect_: {
          selectOption.set(key, value.options);
          types.set(key, FieldType.select);

          const filter: FieldFilter<T> = {
            column: key as any,
            operation: Operation.In,
            type: FieldType.select,
          };

          schema.filters?.push(filter);
          if (value.value?.length) {
            initialForm.filters?.push({
              ...filter,
              value: value.value.map((e) => e.value),
            });
          }

          break;
        }

        case value instanceof TypeDate_: {
          const filter: FieldFilter<T> = {
            column: key as any,
            type: FieldType.date,
          };

          schema?.filters?.push(filter);
          types.set(key, FieldType.date);

          if (value && value.hasValue()) {
            initialForm.filters?.push({ ...filter, value: value.getRange() });
          }

          break;
        }
        case typeof value === 'number': {
          const filter: FieldFilter<T> = {
            column: key as any,
            operation: Operation.Equal,
            type: FieldType.number,
          };

          schema.filters?.push(filter);
          types.set(key, FieldType.number);

          if (Number(value)) {
            initialForm.filters?.push({ ...filter, value: Number(value) });
          }
          break;
        }
        case typeof value === 'string': {
          const filter: FieldFilter<T> = {
            column: key as any,
            operation: Operation.Equal,
            type: FieldType.string,
          };

          schema.filters?.push(filter);
          types.set(key, FieldType.string);

          if (value !== '') {
            initialForm.filters?.push({ ...filter, value: value as string });
          }

          break;
        }

        default: {
          break;
        }
      }
    }
  }

  const resetQuery = <T extends string>(raw?: Partial<QueryPagi<T>>) => {
    if (!initialQuery) return;
    Object.assign(initialQuery, defaultQueryPagination(raw));
  };

  const form = ref(cloneDeep(initialForm));

  const parseForm = (queryForm: typeof form.value) => {
    const copyQueryForm = cloneDeep(unref(queryForm));
    const copy = cloneDeep(initialQuery);
    const filters = parseFilter(copyQueryForm.filters as any);
    setQueryfilter(copy, filters);
    setQuerySearch(copy, copyQueryForm.search as any);

    const querySort = copyQueryForm.sorts.map((s) => {
      const [column, direction] = s.split(' ');
      return { column, direction } as SingleSort<string>;
    });
    setQuerySort(copy, querySort);
    return copy;
  };

  const query = ref(parseForm(form.value));

  const setQuery = (update: Setter<typeof query.value>) => {
    if (update) {
      query.value = typeof update === 'function' ? update(query.value) : update;
    }
  };

  const convertQueryToForm = (queryValue: any) => {
    const safeForm = unref(form);
    const formQ: any = {
      ...safeForm,
      filters: [...safeForm.filters],
      search: {
        operation: 'containsFold',
        type: 'string',
        value: queryValue.search.replace(/\*/g, '') || '',
      },
      sorts: queryValue.sorts.map(
        (sort: any) => `${sort.column} ${sort.direction}`,
      ),
    };

    // Map filters based on the types
    for (const filter of queryValue.filters || []) {
      const columnType = types.get(filter.column);

      if (columnType) {
        const filterValue = filter.value;
        const convertedFilter: FieldFilter = {
          column: filter.column,
          type: columnType,
          operation: filter.operation,
        };

        // Handle the filter value based on its type
        convertedFilter.value =
          columnType === 'date' && Array.isArray(filterValue)
            ? filterValue.map((dateStr: string) => {
                return dateUtil(dateStr.replace('+', ' ')).utc();
              })
            : filterValue;

        // if column exists in the form, update
        const formIndex = formQ.filters.findIndex(
          (f: FieldFilter) => f.column === convertedFilter.column,
        );

        if (formIndex === -1) {
          formQ.filters.push(convertedFilter);
        } else {
          formQ.filters[formIndex] = convertedFilter;
        }
      }
    }

    return formQ;
  };

  const setForm = (update: Setter<typeof form.value>) => {
    if (update) {
      form.value = typeof update === 'function' ? update(form.value) : update;
    }
  };

  const updateQuerySearchToForm = (querySearch: string) => {
    const queryValue = cloneDeep(
      reverseParseQueryPagi(urlQueryParse(querySearch.toString()) as any),
    );

    const newForm = cloneDeep(convertQueryToForm(queryValue));
    setForm(newForm);
    query.value = parseForm(newForm);
  };

  if (!config.disableUrlQuery && window.location.search.length > 0) {
    const queryUrls = new URLSearchParams(window.location.search);
    updateQuerySearchToForm(queryUrls.toString());
  }

  const resetForm = () => {
    form.value = cloneDeep(initialForm);
    query.value = cloneDeep(initialQuery);
  };

  const querySearch = useQuerySearchState({
    enable: !config.disableUrlQuery,
    onQuerySearchChange: updateQuerySearchToForm,
  });

  const getQuery = (): typeof query.value => {
    // Use queryValue if it exists and is not empty; otherwise, use form.value
    query.value = parseForm(form.value);

    if (!config.disableUrlQuery) {
      const payload = parseQueryPagi(query.value);
      const queryValue = urlQueryStringify(payload as any);
      querySearch.updateUrlQuerySearch(queryValue);
    }
    return query.value;
  };

  watch(
    form.value,
    () => {
      if (!config.disableUrlQuery) {
        const payload = parseQueryPagi(parseForm(form.value));
        const queryValue = urlQueryStringify(payload as any);
        querySearch.updateUrlQuerySearch(queryValue);
      }
    },
    { deep: true },
  );

  return reactive({
    config,
    form,
    query,
    setForm,
    resetForm,
    getFieldConfig,
    getQuery,
    modelConfig,
    resetQuery,
    defaultQuery: initialQuery,
    schema,
    selectOption,
    setQuery,
    sortOption,
    types,
  });
};

const querySearchStore = new Map<string, string>();

const useQuerySearchState = (opt: {
  enable: boolean;
  onQuerySearchChange?: (querySearch: string) => void;
}) => {
  const router = useRouter();
  const tab = useTabbar();
  const tabbarStore = useTabbarStore();

  const updateUrlQuerySearch = (queryValue: string) => {
    const newUrl = new URL(window.location.href);
    newUrl.search = queryValue;

    const isSameUrl = newUrl.toString() === window.location.toString();
    if (isSameUrl) return;

    router.currentRoute.value.meta.querySearch = newUrl.search;
    querySearchStore.set(tab.currentActive.value, queryValue);
    window.history.replaceState(window.history.state, '', newUrl.href);
  };

  const clearUrlQuerySearch = () => {
    const newUrl = new URL(window.location.href);
    newUrl.search = '';
    window.history.replaceState(window.history.state, '', newUrl.href);
  };
  if (!opt.enable) {
    return { updateUrlQuerySearch, clearUrlQuerySearch };
  }

  watch(
    tabbarStore.tabs,
    () => {
      const tabPath = tab.currentActive.value;
      let queryValue = querySearchStore.get(tabPath);
      if (!queryValue) return;

      const hasTab = tabbarStore.tabs.find((e) => e.path === tabPath);
      if (!hasTab) {
        querySearchStore.delete(tabPath);
        queryValue = '';
        clearUrlQuerySearch();
      }
    },
    { deep: true, immediate: true },
  );

  watch(
    tab.currentActive,
    (tabPath) => {
      let queryValue = querySearchStore.get(tabPath);
      if (!queryValue) return;

      updateUrlQuerySearch(queryValue);
      opt?.onQuerySearchChange?.(queryValue);
    },
    { deep: true, immediate: true },
  );

  return {
    updateUrlQuerySearch,
    clearUrlQuerySearch,
  };
};

export type Querier = ReturnType<typeof useQueryForm>;
