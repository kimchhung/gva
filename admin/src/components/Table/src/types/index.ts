import { TableProps as ElTableProps } from 'element-plus'
export type TableColumn<T extends object = {}> = {
  field: string
  label?: string
  type?: string
  /**
   * 是否隐藏
   */
  hidden?: boolean
  children?: TableColumn<T>[]
  slots?: {
    default?: (data: TableSlotDefault<T>) => JSX.Element | JSX.Element[] | null
    header?: (...args: any[]) => JSX.Element | null
  }
  index?: number | ((index: number) => number)
  columnKey?: string
  width?: string | number
  minWidth?: string | number
  fixed?: boolean | 'left' | 'right'
  renderHeader?: (...args: any[]) => JSX.Element | null
  // sortable?: boolean
  sortMethod?: (...args: any[]) => number
  sortBy?: string | string[] | ((...args: any[]) => string | string[])
  sortOrders?: (string | null)[]
  resizable?: boolean
  formatter?: (...args: any[]) => any
  showOverflowTooltip?: boolean
  align?: 'left' | 'center' | 'right'
  headerAlign?: 'left' | 'center' | 'right'
  className?: string
  labelClassName?: string
  selectable?: (...args: any[]) => boolean
  reserveSelection?: boolean
  filters?: Array<{ text: string; value: string }>
  filterPlacement?: string
  filterMultiple?: boolean
  filterMethod?: (...args: any[]) => boolean
  filteredValue?: string[]
  [key: string]: any
}

export type TableSlotDefault<T extends object> = {
  row: T
  column: TableColumn<T>
  $index: number
  [key: string]: any
}

export type Pagination = {
  small?: boolean
  background?: boolean
  limit?: number
  defaultLimit?: number
  total?: number
  pageCount?: number
  pagerCount?: number
  currentPage?: number
  defaultCurrentPage?: number
  layout?: string
  pageSizes?: number[]
  popperClass?: string
  prevText?: string
  nextText?: string
  disabled?: boolean
  hideOnSinglePage?: boolean
}

export type TableSetProps = {
  field: string
  path: string
  value: any
}

export type TableProps<T extends object = {}> = {
  limit?: number
  page?: number
  showAction?: boolean
  // Whether all the excess is exceeded, priority is lower than the showflowtooltip in SCHEMA,
  showOverflowTooltip?: boolean
  // header
  columns?: TableColumn<T>[]
  // Whether to show paging
  pagination?: Pagination | undefined
  // Only the column of Type = Selection is valid, and the type is Boolean. For True, the data selected before the data will be retained after the data is updated (requiring Row-Key)
  reserveSelection?: boolean
  // Load status
  loading?: boolean
  // Whether to overlap the index
  reserveIndex?: boolean
  // Alignment method
  align?: 'left' | 'center' | 'right'
  // Models on heads
  headerAlign?: 'left' | 'center' | 'right'
  imagePreview?: string[]
  videoPreview?: string[]
  sortable?: boolean
  data?: T
} & Omit<Partial<ElTableProps<any[]>>, 'data'>
