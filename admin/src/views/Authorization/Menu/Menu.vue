<script setup lang="tsx">
import { convertEdgeChildren } from '@/api/admin/types'

import { api } from '@/api'

import { MenuRoute } from '@/api/authorization/types'
import { useApi } from '@/axios'
import { BaseButton } from '@/components/Button'
import { ContentWrap } from '@/components/ContentWrap'
import { Dialog } from '@/components/Dialog'
import { FormSchema } from '@/components/Form'
import { Icon } from '@/components/Icon'
import { Search } from '@/components/Search'
import { Table, TableColumn } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { useTable } from '@/hooks/web/useTable'
import { ElTag } from 'element-plus'
import { reactive, ref, unref } from 'vue'
import Detail from './components/Detail.vue'
import Write from './components/Write.vue'

const { t } = useI18n()

const { tableRegister, tableState, tableMethods } = useTable({
  fetchDataApi: async (props) => {
    const [data] = await useApi(() => api().getRouters(props))
    if (!data) return { list: [] }

    return {
      list: convertEdgeChildren(data)
    }
  }
})

const { dataList, loading } = tableState
const { getList } = tableMethods

const tableColumns = reactive<TableColumn<MenuRoute>[]>([
  {
    field: 'index',
    label: t('common.index'),
    type: 'index'
  },

  {
    field: 'meta.title',
    label: t('meta.title'),
    slots: {
      default: (data) => {
        const title = t(data.row.meta.title)
        return <>{title}</>
      }
    }
  },
  {
    field: 'meta.icon',
    label: t('meta.icon'),
    width: 80,
    slots: {
      default: (data) => {
        const icon = data.row.meta.icon
        if (icon) {
          return (
            <>
              <Icon icon={icon} />
            </>
          )
        } else {
          return null
        }
      }
    }
  },

  // {
  //   field: 'meta.permission',
  //   label: t('meta.permission'),
  //   slots: {
  //     default: (data) => {
  //       const permission = data.row.meta.permission
  //       return permission ? <>{permission.join(', ')}</> : null
  //     }
  //   }
  // },
  {
    field: 'component',
    label: t('meta.component'),
    slots: {
      default: (data) => {
        const component = data.row.component
        return (
          <>
            {component === '#' ? 'Top directory' : component === '##' ? 'Subdirectory' : component}
          </>
        )
      }
    }
  },
  {
    field: 'path',
    label: t('meta.path')
  },
  {
    field: 'isEnable',
    label: t('meta.isEnable'),
    slots: {
      default: (data) => {
        return (
          <>
            <ElTag type={data.row.isEnable ? 'success' : 'danger'}>
              {data.row.isEnable ? t('tagStatus.enable') : t('tagStatus.disable')}
            </ElTag>
          </>
        )
      }
    }
  },
  {
    field: 'action',
    label: t('common.action'),
    width: 250,
    slots: {
      default: (data) => {
        const row = data.row
        return (
          <>
            <BaseButton type="primary" onClick={() => action(row, 'edit')}>
              {t('button.edit')}
            </BaseButton>
            <BaseButton type="success" onClick={() => action(row, 'detail')}>
              {t('button.detail')}
            </BaseButton>
            <BaseButton type="danger">{t('button.del')}</BaseButton>
          </>
        )
      }
    }
  }
])

const searchSchema = reactive<FormSchema[]>([
  {
    field: 'meta.title',
    label: t('meta.menuName'),
    component: 'Input'
  }
])

const searchParams = ref({})
const setSearchParams = (data) => {
  searchParams.value = data
  getList()
}

const dialogVisible = ref(false)
const dialogTitle = ref('')

const currentRow = ref()
const actionType = ref('')

const writeRef = ref<ComponentRef<typeof Write>>()

const saveLoading = ref(false)

const action = (row, type: string) => {
  dialogTitle.value = t(type === 'edit' ? 'button.edit' : 'button.detail')
  actionType.value = type
  currentRow.value = row
  dialogVisible.value = true
}

const AddAction = () => {
  dialogTitle.value = t('button.add')
  currentRow.value = undefined
  dialogVisible.value = true
  actionType.value = ''
}

const save = async () => {
  const write = unref(writeRef)
  const formData = await write?.submit()

  if (formData) {
    const [res] = await useApi(() => api().createRouter(formData), { loading: saveLoading })
    console.log({ formData, res })
  }
}
</script>

<template>
  <ContentWrap>
    <Search :schema="searchSchema" @reset="setSearchParams" @search="setSearchParams" />
    <div class="mb-10px">
      <BaseButton type="primary" @click="AddAction">{{ t('button.add') }}</BaseButton>
    </div>
    <Table
      :columns="tableColumns"
      default-expand-all
      node-key="id"
      :data="dataList"
      :loading="loading"
      @register="tableRegister"
    />
  </ContentWrap>

  <Dialog v-model="dialogVisible" :title="dialogTitle">
    <Write v-if="actionType !== 'detail'" ref="writeRef" :current-row="currentRow" />

    <Detail v-if="actionType === 'detail'" :current-row="currentRow" />

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
</template>
