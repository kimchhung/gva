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
    label: t('menu.type'),
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
    label: t('menu.parentName')
  },
  {
    field: 'meta.title',
    label: t('menu.title')
  },
  {
    field: 'component',
    label: t('menu.component'),
    slots: {
      default: (data) => {
        const component = data.component
        return (
          <>{component === '#' ? 'Top-level' : component === '##' ? 'Subdirectory' : component}</>
        )
      }
    }
  },
  {
    field: 'name',
    label: t('menu.name')
  },
  {
    field: 'meta.icon',
    label: t('menu.icon'),
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
    label: t('menu.path')
  },
  {
    field: 'meta.activeMenu',
    label: t('menu.activeMenu')
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
    label: t('menu.isEnable'),
    slots: {
      default: (data) => {
        return renderTag(data.isEnable)
      }
    }
  },
  {
    field: 'meta.hidden',
    label: t('menu.hidden'),
    slots: {
      default: (data) => {
        return renderTag(data.enableHidden)
      }
    }
  },
  {
    field: 'meta.alwaysShow',
    label: t('menu.alwaysShow'),
    slots: {
      default: (data) => {
        return renderTag(data.enableDisplay)
      }
    }
  },
  {
    field: 'meta.noCache',
    label: t('menu.noCache'),
    slots: {
      default: (data) => {
        return renderTag(data.enableCleanCache)
      }
    }
  },
  {
    field: 'meta.breadcrumb',
    label: t('menu.breadcrumb'),
    slots: {
      default: (data) => {
        return renderTag(data.enableShowCrumb)
      }
    }
  },
  {
    field: 'meta.affix',
    label: t('menu.affix'),
    slots: {
      default: (data) => {
        return renderTag(data.enablePinnedTab)
      }
    }
  },
  {
    field: 'meta.noTagsView',
    label: t('menu.noTagsView'),
    slots: {
      default: (data) => {
        return renderTag(data.enableHiddenTab)
      }
    }
  },
  {
    field: 'meta.canTo',
    label: t('menu.canTo'),
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
