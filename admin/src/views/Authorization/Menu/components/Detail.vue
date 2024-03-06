<script setup lang="tsx">
import { Descriptions, DescriptionsSchema } from '@/components/Descriptions'
import { Icon } from '@/components/Icon'

import { ElTag } from 'element-plus'
import { PropType, ref } from 'vue'
import { useI18n } from 'vue-i18n'

defineProps({
  currentRow: {
    type: Object as PropType<any>,
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
    field: 'type',
    label: t('meta.type'),
    span: 24,
    slots: {
      default: (data) => {
        const type = data.type
        return <>{type === 1 ? 'menu' : 'Table of contents'}</>
      }
    }
  },
  {
    field: 'parentName',
    label: t('meta.parentName')
  },
  {
    field: 'meta.title',
    label: t('meta.title')
  },
  {
    field: 'component',
    label: t('meta.component'),
    slots: {
      default: (data) => {
        const component = data.component
        return <>{component === '#' ? '顶级目录' : component === '##' ? '子目录' : component}</>
      }
    }
  },
  {
    field: 'name',
    label: t('meta.name')
  },
  {
    field: 'meta.icon',
    label: t('meta.icon'),
    slots: {
      default: (data) => {
        const icon = data.icon
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
  {
    field: 'path',
    label: t('meta.path')
  },
  {
    field: 'meta.activeMenu',
    label: t('meta.activeMenu')
  },
  {
    field: 'permissionList',
    label: 'permissionList',
    span: 24,
    slots: {
      default: (data: any) => (
        <>
          {data?.permissionList?.map((v) => {
            return (
              <ElTag class="mr-1" key={v.value}>
                {v.label}
              </ElTag>
            )
          })}
        </>
      )
    }
  },
  {
    field: 'isEnable',
    label: t('meta.isEnable'),
    slots: {
      default: (data) => {
        return renderTag(data.isEnable)
      }
    }
  },
  {
    field: 'meta.hidden',
    label: t('meta.hidden'),
    slots: {
      default: (data) => {
        return renderTag(data.enableHidden)
      }
    }
  },
  {
    field: 'meta.alwaysShow',
    label: t('meta.alwaysShow'),
    slots: {
      default: (data) => {
        return renderTag(data.enableDisplay)
      }
    }
  },
  {
    field: 'meta.noCache',
    label: t('meta.noCache'),
    slots: {
      default: (data) => {
        return renderTag(data.enableCleanCache)
      }
    }
  },
  {
    field: 'meta.breadcrumb',
    label: t('meta.breadcrumb'),
    slots: {
      default: (data) => {
        return renderTag(data.enableShowCrumb)
      }
    }
  },
  {
    field: 'meta.affix',
    label: t('meta.affix'),
    slots: {
      default: (data) => {
        return renderTag(data.enablePinnedTab)
      }
    }
  },
  {
    field: 'meta.noTagsView',
    label: t('meta.noTagsView'),
    slots: {
      default: (data) => {
        return renderTag(data.enableHiddenTab)
      }
    }
  },
  {
    field: 'meta.canTo',
    label: t('meta.canTo'),
    slots: {
      default: (data) => {
        return renderTag(data.enableSkip)
      }
    }
  }
])
</script>

<template>
  <Descriptions :schema="detailSchema" :data="currentRow || {}" />
</template>
