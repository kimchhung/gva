<script setup lang="tsx">
import { Department } from '@/api/department/types'
import { BaseButton } from '@/components/Button'
import { ContentWrap } from '@/components/ContentWrap'
import { Table, TableColumn } from '@/components/Table'
import { useI18n } from '@/hooks/web/useI18n'
import { useIcon } from '@/hooks/web/useIcon'
import { useTable } from '@/hooks/web/useTable'
import { edgelistToTree } from '@/utils/edgeTree'
import { ElMessage, ElTag } from 'element-plus'
import { reactive } from 'vue'
import { useRouter } from 'vue-router'

const { t } = useI18n()
const { push } = useRouter()

const { tableRegister, tableState } = useTable<Department>({
  fetchDataApi: async (query) => {
    const [res, err] = await api.department.getMany({ query })
    if (err) return null
    res.data = edgelistToTree(res.data)
    return res
  }
})

const action = async (row: Recordable, type: 'add' | 'edit' | 'detail' | 'delete') => {
  switch (type) {
    default:
      push({ path: `/authorization/department/${type}`, query: { id: row?.id } })

      break
    case 'delete':
      const [res] = await api.department.delete({ id: row?.id })
      if (res) {
        ElMessage.success({
          message: 'delete successfully'
        })
      }
      break
  }
}

const renderTag = (enable?: boolean) => {
  return (
    <ElTag type={!enable ? 'danger' : 'success'}>
      {enable ? t('tagStatus.enable') : t('tagStatus.disable')}
    </ElTag>
  )
}

const tableColumns = reactive<TableColumn<Department>[]>([
  {
    field: 'index',
    label: t('common.index'),
    type: 'index'
  },
  {
    field: 'name',
    label: t('common.name')
  },
  {
    field: 'isEnable',
    label: t('common.isEnable'),
    slots: {
      default: ({ row }) => {
        return renderTag(row.isEnable)
      }
    }
  },
  {
    field: 'nameId',
    label: t('common.nameId')
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
      @register="tableRegister"
    />
  </ContentWrap>
</template>
