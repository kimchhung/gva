import { Table, TableColumn, TableExpose, TableProps, TableSetProps } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { ElMessage, ElMessageBox, ElTable } from 'element-plus'
import { nextTick, onMounted, ref, unref, watch } from 'vue'
import { QueryPagi } from './usePagi'

const { t } = useI18n()

type UseTableConfig = {
  /**
   * 是否初始化的时候请求一次
   */
  immediate?: boolean
  fetchDataApi: (props?: QueryPagi) => Promise<{
    list: any[]
    total?: number
  }>
  fetchDelApi?: () => Promise<boolean>
}

export const useTable = (config: UseTableConfig) => {
  const { immediate = true } = config

  const loading = ref(false)

  const fetchProps = ref({
    page: 1,
    limit: 20
  } as QueryPagi)

  const total = ref(0)
  const dataList = ref<any[]>([])

  watch(
    () => fetchProps.value,
    () => {
      methods.getList()
    }
  )

  watch(
    () => fetchProps.value.page,
    () => {
      // When the current page is not 1, after the number of modification pages, it will cause multiple times to call the getlist method
      if (unref(fetchProps.value.page) === 1) {
        methods.getList()
      } else {
        fetchProps.value.page = 1
        methods.getList()
      }
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
      loading.value = true
      try {
        const res = await config?.fetchDataApi(unref(fetchProps.value))
        if (res) {
          dataList.value = res.list
          total.value = res.total || 0
        }
      } catch (err) {
        console.log('fetchDataApi error')
      } finally {
        loading.value = false
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
    // 删除数据
    delList: async (idsLength: number) => {
      const { fetchDelApi } = config
      if (!fetchDelApi) {
        console.warn('fetchDelApi is undefined')
        return
      }
      ElMessageBox.confirm(t('common.delMessage'), t('common.delWarning'), {
        confirmButtonText: t('common.delOk'),
        cancelButtonText: t('common.delCancel'),
        type: 'warning'
      }).then(async () => {
        const res = await fetchDelApi()
        if (res) {
          ElMessage.success(t('common.delSuccess'))

          // 计算出临界点
          const current =
            unref(total) % unref(fetchProps.value.limit) === idsLength ||
            unref(fetchProps.value.limit) === 1
              ? unref(fetchProps.value.page) > 1
                ? unref(fetchProps.value.page) - 1
                : unref(fetchProps.value.page)
              : unref(fetchProps.value.page)

          fetchProps.value.page = current
          methods.getList()
        }
      })
    }
  }

  return {
    tableRegister: register,
    tableMethods: methods,
    tableState: {
      limit: fetchProps.value.limit,
      page: fetchProps.value.page,
      total,
      dataList,
      loading
    }
  }
}
