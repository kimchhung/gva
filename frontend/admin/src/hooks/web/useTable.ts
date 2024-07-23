import { Table, TableColumn, TableExpose, TableProps, TableSetProps } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { ElMessage, ElMessageBox, ElTable } from 'element-plus'
import { nextTick, onMounted, reactive, ref, unref, watch } from 'vue'
import { QueryUrl } from './usePagi'

const { t } = useI18n()

type UseTableConfig<T extends Object = any> = {
  /**
   * Do you request once when you initialize
   */
  immediate?: boolean
  onFetchData: (props: QueryUrl<T>) => Promise<{
    data: T[]
    meta?: {
      total?: number
    }
  } | null>
  onDeleteData?: () => Promise<boolean>
}

const getTablePageAndSize = (v: { offset?: number; limit?: number }) => {
  return { page: Math.floor(Number(v?.offset) / Number(v?.limit)) + 1, pageSize: Number(v?.limit) }
}

export type TableState<T> = {
  page: number
  pageSize: number
  isLoading: boolean
  total: number

  data: T[]
  meta?: {
    total?: number
  }
}

export const useTable = <T extends RecordWithID>(config: UseTableConfig<T>) => {
  const { immediate = true } = config
  const fetchProps = ref(new QueryUrl<T>())
  const { page, pageSize } = getTablePageAndSize(fetchProps.value)
  const tableState = reactive<TableState<T>>({
    page: page,
    pageSize: pageSize,
    isLoading: false,
    total: 0,
    data: [],
    meta: {
      total: undefined
    }
  })

  watch(
    () => [fetchProps.value.offset, fetchProps.value.limit],
    () => {
      const { page, pageSize } = getTablePageAndSize(fetchProps.value)
      tableState.page = page
      tableState.pageSize = pageSize

      methods.getList()
    }
  )

  onMounted(() => {
    if (immediate) {
      methods.getList()
    }
  })

  // Table instance
  const tableRef = ref<typeof Table & TableExpose>()

  // Eltable instance
  const elTableRef = ref<ComponentRef<typeof ElTable>>()

  const register = (ref: typeof Table & TableExpose, elRef: ComponentRef<typeof ElTable>) => {
    tableRef.value = ref
    elTableRef.value = unref(elRef)
  }

  const getTable = async () => {
    await nextTick()
    const table = unref(tableRef)
    if (!table) {
      console.error('The table is not registered. Please use the register method to register')
    }
    return table
  }

  const methods = {
    /**
     * Get the form data
     */
    getList: async () => {
      tableState.isLoading = true
      try {
        const res = await config?.onFetchData(new QueryUrl(unref(fetchProps as any)))
        if (res) {
          tableState.data = (res.data as any) ?? []
          tableState.meta = res.meta as any
        }
      } catch (err) {
        console.error('fetchDataApi error', err)
      } finally {
        tableState.isLoading = false
      }
    },

    /**
     * @description Set the props of the table component
     * @param props PROPS of table components
     */
    setProps: async (props: TableProps = {}) => {
      const table = await getTable()
      table?.setProps(props)
    },

    /**
     * @description Set column
     * @param columnProps Columns that need to be set
     */
    setColumn: async (columnProps: TableSetProps[]) => {
      const table = await getTable()
      table?.setColumn(columnProps)
    },

    /**
     * @description New COLUMN
     * @param tableColumn Need to add data
     * @param index Where to add
     */
    addColumn: async (tableColumn: TableColumn, index?: number) => {
      const table = await getTable()
      table?.addColumn(tableColumn, index)
    },

    /**
     * @description Delete column
     * @param field Which data deletes
     */
    delColumn: async (field: string) => {
      const table = await getTable()
      table?.delColumn(field)
    },

    /**
     * @description Examples to obtain ELTable components
     * @returns ElTable instance
     */
    getElTableExpose: async () => {
      await getTable()
      return unref(elTableRef)
    },

    refresh: () => {
      methods.getList()
    },

    // sortableChange: (e: any) => {
    //   console.log('sortableChange', e)
    //   const { oldIndex, newIndex } = e
    //   dataList.value.splice(newIndex, 0, dataList.value.splice(oldIndex, 1)[0])
    //   // to do something
    // }
    // delete data
    delList: async (idsLength: number) => {
      const { onDeleteData: fetchDelete } = config
      if (!fetchDelete) {
        console.warn('fetchDelApi is undefined')
        return
      }
      ElMessageBox.confirm(t('common.delMessage'), t('common.delWarning'), {
        confirmButtonText: t('common.delOk'),
        cancelButtonText: t('common.delCancel'),
        type: 'warning'
      }).then(async () => {
        const res = await fetchDelete()
        if (res) {
          ElMessage.success(t('common.delSuccess'))

          // Calculate the critical point
          const current =
            unref(tableState.total) % unref(Number(fetchProps.value.limit)) === idsLength ||
            unref(fetchProps.value.limit) === 1
              ? Number(unref(fetchProps.value.offset)) > 1
                ? Number(unref(fetchProps.value.offset)) - 1
                : Number(unref(fetchProps.value.offset))
              : Number(unref(fetchProps.value.offset))

          fetchProps.value.offset = current
          methods.getList()
        }
      })
    }
  }

  return {
    tableRegister: register,
    tableMethods: methods,
    tableState
  }
}
