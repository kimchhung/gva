import type { SorterResult } from 'ant-design-vue/es/table/interface';

import type { Querier } from '../pagi/form';

export const useTableSort = (querier: Querier) => {
  const handleSort = (sorts: SorterResult<any> | SorterResult<any>[]) => {
    const _sorts = Array.isArray(sorts) ? sorts : [sorts];

    const sort = _sorts.length > 0 ? _sorts.at(-1) : null;

    if (!sort || !sort.order) {
      querier.form.sorts = [];

      return;
    }

    const order = sort.order === 'ascend' ? 'asc' : 'desc';
    const key =
      sort.columnKey ||
      (Array.isArray(sort.column?.dataIndex) && sort.column.dataIndex.length > 0
        ? sort.column?.dataIndex?.[0]
        : sort.column?.dataIndex) ||
      sort.field;
    querier.form.sorts = [`${key} ${order}`];
    querier.setQuery(() => querier.getQuery());
  };

  return { handleSort };
};
