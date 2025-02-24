import type { TableColumnsType } from 'ant-design-vue';
import type { ColumnType } from 'ant-design-vue/es/table';

import type { Querier } from '../pagi/form';

import { computed, type MaybeRef, reactive, ref, unref } from 'vue';

import { api } from '../../api';
import { useSortableColumns } from './use-column-sort';
import { useTableSort } from './use-table-sort';

interface UseTableProps<Data = any> {
  querier: Querier;
  fetcher: (globalAPi: API) => any;
  columns: MaybeRef<TableColumnsType<Data>>;
  table: {
    getHasNext: (data: { data: Data[]; meta: Record<string, any> }) => boolean;
    getList: (data: { data: Data[]; meta: Record<string, any> }) => Data[];
    getTotal: (data: { data: Data[]; meta: Record<string, any> }) => number;
  };
}

const uniqColumn = (col: any) => {
  if (col.key === 'undefined' || !col.key) return col.dataIndex;
  return col.key;
};

export const useTableData = <T = any>({
  querier,
  fetcher,
  table,
  columns,
}: UseTableProps<T>) => {
  const getMany = fetcher(api());
  const query = computed(() => querier.query);

  const {
    data: resp,
    isFetching: isLoading,
    refetch,
  } = getMany.useQuery({ query });
  const { handleSort } = useTableSort(querier);
  const sortableColumns = useSortableColumns<T>(querier, columns);

  const forceRefetch = () => {
    const queryNow = querier.getQuery();
    querier.setQuery(() => queryNow);
    refetch();
  };

  const goToPage = (page: number) => {
    querier.setQuery((q) => {
      q.page = page;
      return q;
    });
  };

  const setLimit = (limit: number) => {
    querier.setQuery((q) => {
      q.page = 1;
      q.limit = limit;
      return q;
    });
  };

  const nextPage = () => {
    const page = querier.query.page + 1;
    goToPage(page);
  };

  const previousPage = () => {
    const page = querier.query.page - 1;
    goToPage(page);
  };

  const viewColumnOptions = computed(() =>
    sortableColumns.value.map((c: ColumnType) => {
      return {
        value: String(uniqColumn(c)),
        label: c.title,
      };
    }),
  );

  const selectedViewColumns = ref<(number | string)[]>(
    viewColumnOptions?.value?.map((c) => String(c.value)) || [],
  );

  const setSelectedViewColumns = (selects: (number | string)[]) => {
    selectedViewColumns.value = selects;
  };

  const viewColumns = () =>
    unref(sortableColumns).filter((c) => {
      return (
        Array.isArray(unref(selectedViewColumns)) &&
        unref(selectedViewColumns.value)?.includes(uniqColumn(c))
      );
    });

  const hideColumn = (column: string) => {
    const columns =
      unref(viewColumnOptions)
        .filter((c) => column !== c.value)
        ?.map((c) => String(c.value)) || [];
    setSelectedViewColumns(columns);
  };

  const tableFunction = {
    handleSort,
    setLimit,
    goToPage,
    forceRefetch,
    nextPage,
    previousPage,
    hideColumn,
  };

  const total = ref(0);
  const tableState = computed(() => {
    return {
      hasNext: () => table.getHasNext(resp?.value) || false,
      total: () => {
        if (resp?.value) {
          // avoid total reset to 0 while fetching
          total.value = table.getTotal(resp?.value);
        }
        return unref(total);
      },
      list: () => table.getList(resp?.value) || [],
      columns: viewColumns,
      isLoading: () => unref(isLoading),
    };
  });

  const tableTool = reactive({
    viewColumnOptions,
    selectedViewColumns,
    setSelectedViewColumns,
  });

  return {
    tableTool,
    tableFunction,
    tableState,
  };
};

export type TableTool<T = any> = ReturnType<
  typeof useTableData<T>
>['tableTool'];
