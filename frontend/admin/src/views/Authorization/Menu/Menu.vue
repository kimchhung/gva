<script setup lang="tsx">
import { convertEdgeChildren } from '@/api/admin/types'
import { deleteRouter, getRouters } from '@/api/route'
import { MenuRoute } from '@/api/route/types'
import { useApi } from '@/axios'
import { BaseButton } from '@/components/Button'
import { ContentWrap } from '@/components/ContentWrap'
import { Icon } from '@/components/Icon'
import { Table, TableColumn } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { useTable } from '@/hooks/web/useTable'
import { ElMessage, ElTag } from 'element-plus'
import { reactive } from 'vue'

import { useIcon } from '@/hooks/web/useIcon'
import { useRouter } from 'vue-router'

const { t } = useI18n()
const { push } = useRouter()

const { tableRegister, tableState } = useTable<MenuRoute>({
  fetchDataApi: async (props) => {
    const [data] = await useApi(() => getRouters(props))
    return { list: convertEdgeChildren(data || []) as MenuRoute[] }
  }
})

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
      default: ({ row }) => {
        const title = t(row.meta.title)
        return <>{title}</>
      }
    }
  },
  {
    field: 'meta.icon',
    label: t('meta.icon'),
    width: 80,
    slots: {
      default: ({ row }) => {
        const icon = row.meta.icon
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
      default: ({ row }) => {
        const component = row.component
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
      default: ({ row }) => {
        return (
          <>
            <ElTag type={row.isEnable ? 'success' : 'danger'}>
              {row.isEnable ? t('tagStatus.enable') : t('tagStatus.disable')}
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
      default: ({ row }) => {
        return (
          <>
            <BaseButton
              icon={useIcon({ icon: 'ep:view' })}
              type="success"
              onClick={() => action(row, 'detail')}
            >
              {t('button.detail')}
            </BaseButton>
            <BaseButton
              icon={useIcon({ icon: 'ep:edit' })}
              type="primary"
              onClick={() => action(row, 'edit')}
            />

            <BaseButton
              icon={useIcon({ icon: 'ep:delete' })}
              type="danger"
              onClick={() => action(row, 'delete')}
            />
          </>
        )
      }
    }
  }
])

const action = (row: Recordable, type: 'add' | 'edit' | 'detail' | 'delete') => {
  switch (type) {
    default:
      push({ path: `/authorization/menu/${type}`, query: { id: row?.id } })
      // push(`/authorization/menu-${type}?id=${row.id}`)
      break
    case 'delete':
      deleteRouter(Number(row?.id)).then(() => {
        ElMessage.success({
          message: 'delete successfully'
        })
      })
      break
  }
}
</script>

<template>
  <ContentWrap>
    <div class="mb-10px">
      <BaseButton type="primary" @click="action({}, 'add')"> {{ t('button.add') }}</BaseButton>
    </div>

    <Table
      v-model:pageSize="tableState.pageSize"
      v-model:currentPage="tableState.page"
      :columns="tableColumns"
      :data="tableState.dataList"
      :loading="tableState.isLoading"
      :pagination="{ total: tableState.total }"
      @register="tableRegister"
    />
  </ContentWrap>
</template>
