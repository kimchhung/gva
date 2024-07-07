<script setup lang="tsx">
import { BaseButton } from '@/components/Button'
import { ContentWrap } from '@/components/ContentWrap'
import { Dialog } from '@/components/Dialog'
import { Search } from '@/components/Search'
import { Table } from '@/components/Table'
import { CrudSchema, useCrudSchemas } from '@/hooks/web/useCrudSchemas'
import { useI18n } from '@/hooks/web/useI18n'
import { useTable } from '@/hooks/web/useTable'
import { ElDivider, ElInput, ElTree } from 'element-plus'
import { reactive, ref, unref, watch } from 'vue'
import Detail from './components/Detail.vue'
import Write from './components/Write.vue'

const { t } = useI18n()

const { tableRegister, tableState, tableMethods } = useTable({
  fetchDataApi: async () => {
    // const { pageSize, currentPage } = tableState
    // const res = await getUserByIdApi({
    //   id: unref(currentNodeKey),
    //   pageIndex: unref(currentPage),
    //   pageSize: unref(pageSize),
    //   ...unref(searchParams)
    // })
    return {
      list: [],
      total: 0
    }
  },
  fetchDelApi: async () => {
    // const res = await deleteUserByIdApi(unref(ids))
    return true
  }
})

const { getList } = tableMethods

const crudSchemas = reactive<CrudSchema[]>([
  {
    field: 'selection',
    search: {
      hidden: true
    },
    form: {
      hidden: true
    },
    detail: {
      hidden: true
    },
    table: {
      type: 'selection'
    }
  },
  {
    field: 'index',
    label: t('userDemo.index'),
    form: {
      hidden: true
    },
    search: {
      hidden: true
    },
    detail: {
      hidden: true
    },
    table: {
      type: 'index'
    }
  },
  {
    field: 'username',
    label: t('userDemo.username')
  },
  {
    field: 'account',
    label: t('userDemo.account')
  },
  {
    field: 'department.id',
    label: t('userDemo.department'),
    detail: {
      hidden: true
      // slots: {
      //   default: (data: DepartmentUserItem) => {
      //     return <>{data.department.departmentName}</>
      //   }
      // }
    },
    search: {
      hidden: true
    },
    table: {
      hidden: true
    }
  },
  {
    field: 'role',
    label: t('userDemo.role'),
    search: {
      hidden: true
    },
    form: {
      component: 'Select',
      value: [],
      componentProps: {
        multiple: true,
        collapseTags: true,
        maxCollapseTags: 1
      },
      optionApi: async () => {
        // const res = await getRoleListApi()
        return [{ label: 'test', value: 1 }]
      }
    }
  },
  {
    field: 'email',
    label: t('userDemo.email'),
    form: {
      component: 'Input'
    },
    search: {
      hidden: true
    }
  },
  {
    field: 'createTime',
    label: t('userDemo.createTime'),
    form: {
      component: 'Input'
    },
    search: {
      hidden: true
    }
  },
  {
    field: 'action',
    label: t('common.action'),
    form: {
      hidden: true
    },
    detail: {
      hidden: true
    },
    search: {
      hidden: true
    },
    table: {
      width: 240,
      slots: {
        default: (_: any) => {
          return (
            <>
              <BaseButton type="primary">{t('exampleDemo.edit')}</BaseButton>
              <BaseButton type="success">{t('exampleDemo.detail')}</BaseButton>
              <BaseButton type="danger">{t('exampleDemo.del')}</BaseButton>
            </>
          )
        }
      }
    }
  }
])

const { allSchemas } = useCrudSchemas(crudSchemas)

const searchParams = ref({})
const setSearchParams = (params: any) => {
  tableState.page = 1
  searchParams.value = params
  getList()
}

const treeEl = ref<typeof ElTree>()

// const currentNodeKey = ref('')

const currentDepartment = ref('')
watch(
  () => currentDepartment.value,
  (val) => {
    unref(treeEl)!.filter(val)
  }
)

// const currentChange = (data: any) => {
//   // if (data.children) return
//   currentNodeKey.value = data.id
//   currentPage.value = 1
//   getList()
// }

// const filterNode = (value: string, data: any) => {
//   if (!value) return true
//   return data.departmentName.includes(value)
// }

const dialogVisible = ref(false)
const dialogTitle = ref('')

const currentRow = ref<any>()
const actionType = ref('')

const AddAction = () => {
  dialogTitle.value = t('exampleDemo.add')
  currentRow.value = undefined
  dialogVisible.value = true
  actionType.value = ''
}

const delLoading = ref(false)
// const ids = ref<string[]>([])

const writeRef = ref<ComponentRef<typeof Write>>()

const saveLoading = ref(false)

const save = async () => {
  // const write = unref(writeRef)
  // const formData = await write?.submit()
  // if (formData) {
  //   saveLoading.value = true
  //   try {
  //     // const res = await saveUserApi(formData)
  //     if (res) {
  //       currentPage.value = 1
  //       getList()
  //     }
  //   } catch (error) {
  //     console.log(error)
  //   } finally {
  //     saveLoading.value = false
  //     dialogVisible.value = false
  //   }
  // }
}
</script>

<template>
  <div class="flex w-100% h-100%">
    <ContentWrap class="w-250px">
      <div class="flex justify-center items-center">
        <div class="flex-1">{{ t('userDemo.departmentList') }}</div>
        <ElInput
          v-model="currentDepartment"
          class="flex-[2]"
          :placeholder="t('userDemo.searchDepartment')"
          clearable
        />
      </div>
      <ElDivider />
      <!-- <ElTree
        ref="treeEl"
        :data="departmentList"
        default-expand-all
        :expand-on-click-node="false"
        node-key="id"
        :current-node-key="currentNodeKey"
        :props="{
          label: 'departmentName'
        }"
        :filter-node-method="filterNode"
        @current-change="currentChange"
      >
        <template #default="{ data }">
          <div
            :title="data.departmentName"
            class="whitespace-nowrap overflow-ellipsis overflow-hidden"
          >
            {{ data.departmentName }}
          </div>
        </template>
      </ElTree> -->
    </ContentWrap>
    <ContentWrap class="flex-[3] ml-20px">
      <Search
        :schema="allSchemas.searchSchema"
        @reset="setSearchParams"
        @search="setSearchParams"
      />

      <div class="mb-10px">
        <BaseButton type="primary" @click="AddAction">{{ t('exampleDemo.add') }}</BaseButton>
        <BaseButton :loading="delLoading" type="danger">
          {{ t('exampleDemo.del') }}
        </BaseButton>
      </div>
      <Table
        v-model:current-page="tableState.page"
        v-model:page-size="tableState.pageSize"
        :columns="allSchemas.tableColumns"
        :data="tableState.dataList"
        :loading="tableState.isLoading"
        @register="tableRegister"
        :pagination="{
          total: tableState.total
        }"
      />
    </ContentWrap>

    <Dialog v-model="dialogVisible" :title="dialogTitle">
      <Write
        v-if="actionType !== 'detail'"
        ref="writeRef"
        :form-schema="allSchemas.formSchema"
        :current-row="currentRow"
      />

      <Detail
        v-if="actionType === 'detail'"
        :detail-schema="allSchemas.detailSchema"
        :current-row="currentRow"
      />

      <template #footer>
        <BaseButton
          v-if="actionType !== 'detail'"
          type="primary"
          :loading="saveLoading"
          @click="save"
        >
          {{ t('button.save') }}
        </BaseButton>
        <BaseButton @click="dialogVisible = false">{{ t('button.close') }}</BaseButton>
      </template>
    </Dialog>
  </div>
</template>
