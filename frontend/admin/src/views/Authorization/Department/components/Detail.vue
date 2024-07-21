<script setup lang="tsx">
import { Department } from '@/api/department/types'
import { Descriptions, DescriptionsSchema } from '@/components/Descriptions'
import { ElTag } from 'element-plus'
import { PropType, ref } from 'vue'
import { useI18n } from 'vue-i18n'

defineProps({
  currentRow: {
    type: Object as PropType<Department>,
    default: () => undefined
  }
})

const { t } = useI18n()

const renderTag = (enable?: boolean) => {
  return (
    <ElTag type={!enable ? 'danger' : 'success'}>
      {enable ? t('tagStatus.enable') : t('tagStatus.disable')}
    </ElTag>
  )
}

const detailSchema = ref<DescriptionsSchema[]>([
  {
    field: 'name',
    label: t('common.name')
  },
  {
    field: 'nameId',
    label: t('common.nameId')
  },
  {
    field: 'isEnable',
    label: t('common.isEnable'),
    slots: {
      default: (data) => {
        return renderTag(data.isEnable)
      }
    }
  }
])
</script>

<template>
  <Descriptions :schema="detailSchema" :data="currentRow || {}" />
</template>
