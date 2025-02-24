import type { TableColumnsType, TableColumnType } from 'ant-design-vue';

import type { ColumnConfig } from '../pagi/form';

import { computed, type MaybeRef, unref } from 'vue';

const getSorter = (
  field: string,
  sorts: string[],
): Pick<TableColumnType, 'sorter' | 'sortOrder'> => {
  const found = sorts.find((sort) => {
    const [key] = sort.split(' ');
    return key === field;
  });

  const sortOpts = {
    sorter: {
      compare: () => {
        return 0;
      },
      multiple: 1,
    },
    sortOrder: undefined,
  };

  if (!found) {
    return sortOpts;
  }

  const [, dir] = found.split(' ');
  const direction = dir === 'asc' ? 'ascend' : 'descend';
  return {
    ...sortOpts,
    sortOrder: direction,
  };
};
export const useSortableColumns = <T = any>(
  querier: any,
  columns: MaybeRef<TableColumnsType<T>>,
) => {
  const cols = computed<TableColumnsType<T>>(() => {
    return unref(columns).map((column: any) => {
      let key: string;
      if (column.key) {
        key = `${column.key}`;
      } else if (Array.isArray(column.dataIndex)) {
        key = column.dataIndex[0] as string;
      } else {
        key = column.dataIndex as string;
      }

      const inModelConfig = querier.modelConfig && key in querier.modelConfig;
      if (!inModelConfig) return column;

      const v = querier.modelConfig[key] as ColumnConfig<any>;
      if (!v.isSort) return column;

      const sort = querier.form.sorts || [];
      return {
        ...getSorter(key, sort),
        ...column,
      };
    });
  });

  return cols;
};
