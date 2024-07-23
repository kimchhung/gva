<script setup lang="tsx">
import { menuToNested } from '@/api/menu/tranform'
import { MenuRoute } from '@/api/menu/types'
import { BaseButton } from '@/components/Button'
import { ContentWrap } from '@/components/ContentWrap'
import { Icon } from '@/components/Icon'
import { Table, TableColumn } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { useTable } from '@/hooks/web/useTable'

import { ElButton, ElMessage, ElTag } from 'element-plus'
import { computed, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { renderPermisionTag } from './hooks/useRenderPermTag'

const { t } = useI18n()
const { push } = useRouter()

const { tableRegister, tableState } = useTable({
  fetchDataApi: async (query) => {
    const [res, err] = await api.menu.getMany({ query })
    if (err) return null
    res.data = menuToNested(res.data)
    return res
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
    label: t('menu.title'),
    slots: {
      default: ({ row }) => {
        const title = t(row.meta.title)
        return <>{title}</>
      }
    }
  },
  {
    field: 'meta.icon',
    label: t('menu.icon'),
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
  //   label: t('menu.permission'),
  //   slots: {
  //     default: (data) => {
  //       const permission = data.row.meta.permission
  //       return permission ? <>{permission.join(', ')}</> : null
  //     }
  //   }
  // },
  {
    field: 'component',
    label: t('menu.component'),
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
    label: t('menu.path')
  },

  {
    field: 'isEnable',
    label: t('menu.isEnable'),
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
    field: 'meta.permissions',
    label: t('menu.permission'),
    slots: {
      default: ({ row }) => {
        return renderPermisionTag(row.meta.permissions)
      }
    }
  },
  {
    field: 'action',
    label: t('common.action'),
    width: 250,

    headerAlign: 'center',
    align: 'center',
    slots: {
      default: ({ row }) => {
        return (
          <>
            <ElButton
              type="success"
              size="small"
              icon={<Icon icon="ep:view" />}
              onClick={() => action(row, 'detail')}
            />
            <ElButton
              type="primary"
              size="small"
              icon={<Icon icon="ep:edit-pen" />}
              onClick={() => action(row, 'edit')}
            />
            <BaseButton
              type="danger"
              size="small"
              icon={<Icon icon="ant-design:delete-outlined" />}
              onClick={() => action(row, 'delete')}
            />
          </>
        )
      }
    }
  }
])

const action = async (row: Recordable, type: 'add' | 'edit' | 'detail' | 'delete') => {
  switch (type) {
    default:
      push({ path: `/authorization/menu/${type}`, query: { id: row?.id } })

      break
    case 'delete':
      const [res] = await api.menu.delete({ id: row?.id })
      if (res) {
        ElMessage.success({
          message: 'delete successfully'
        })
      }

      break
  }
}

const expandRowKeys = computed(() => tableState?.data?.map((e) => e.id) || [])
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
      :data="tableState.data"
      :loading="tableState.isLoading"
      :pagination="{ total: tableState.meta?.total }"
      :expandRowKeys="expandRowKeys"
      @register="tableRegister"
    />
  </ContentWrap>
</template>
