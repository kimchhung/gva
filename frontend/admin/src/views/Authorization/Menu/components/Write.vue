<script setup lang="tsx">
import { MenuRoute } from '@/api/menu/types'
import { permissionToTree } from '@/api/permission/tranform'
import { Form, FormSchema, TreeSelectComponentProps } from '@/components/Form'
import { MenuTypeEnum } from '@/constants/menuType'
import { useForm } from '@/hooks/web/useForm'
import { useI18n } from '@/hooks/web/useI18n'
import { useValidator } from '@/hooks/web/useValidator'
import { useAdminStoreWithOut } from '@/store/modules/admin'
import { cloneDeep } from 'lodash-es'
import { PropType, reactive, ref, unref, watch } from 'vue'

const { t } = useI18n()
const { required } = useValidator()

const props = defineProps({
  currentRow: {
    type: Object as PropType<MenuRoute>,
    default: () => null
  }
})

const formSchema = reactive<FormSchema[]>([
  {
    field: 'type',
    label: t('menu.type'),
    component: 'RadioButton',
    value: MenuTypeEnum.CATALOG,
    colProps: {
      span: 24
    },
    componentProps: {
      options: [
        {
          label: 'Table of contents',
          value: MenuTypeEnum.CATALOG
        },
        {
          label: 'MENU',
          value: MenuTypeEnum.MENU
        },
        {
          label: 'BUTTON',
          value: MenuTypeEnum.BUTTON
        },
        {
          label: 'LINK',
          value: MenuTypeEnum.LINK
        }
      ],
      on: {
        change: async (val: MenuTypeEnum) => {
          const formData = await getFormData()
          if (val === MenuTypeEnum.MENU) {
            setSchema([
              {
                field: 'component',
                path: 'componentProps.disabled',
                value: false
              }
            ])
            setValues({
              component: unref(cacheComponent)
            })
          } else {
            setSchema([
              {
                field: 'component',
                path: 'componentProps.disabled',
                value: true
              }
            ])

            if (formData.pid === void 0) {
              setValues({
                component: '#'
              })
            } else {
              setValues({
                component: '##'
              })
            }
          }
        }
      }
    }
  },
  {
    field: 'name',
    label: t('menu.name'),
    component: 'Input'
  },
  {
    field: 'pid',
    label: 'Parent menu',
    component: 'TreeSelect',
    componentProps: {
      nodeKey: 'id',
      props: {
        label: 'name',
        value: 'id',
        children: 'children'
      },
      highlightCurrent: true,
      expandOnClickNode: false,
      checkStrictly: true,
      checkOnClickNode: true,
      clearable: true,
      on: {
        change: async (val: number) => {
          const formData = await getFormData()

          if (val && formData.type === MenuTypeEnum.CATALOG) {
            setValues({
              component: '##'
            })
          } else if (!val && formData.type === MenuTypeEnum.CATALOG) {
            setValues({
              component: '#'
            })
          } else if (formData.type === MenuTypeEnum.MENU) {
            setValues({
              component: unref(cacheComponent) ?? ''
            })
          }
        }
      }
    } as TreeSelectComponentProps,
    optionApi: async () => {
      const list = useAdminStoreWithOut().getRoleRouters
      return list || []
    }
  },
  {
    field: 'path',
    label: t('menu.path'),
    component: 'Input'
  },
  {
    field: 'component',
    label: 'Component',
    component: 'Input',
    value: '#',
    componentProps: {
      disabled: true,
      placeholder: '#Is the top directory, ## is the child directory',
      on: {
        change: (val: string) => {
          cacheComponent.value = val
        }
      }
    }
  },
  {
    field: 'redirect',
    label: t('menu.redirect'),
    component: 'Input'
  },
  {
    field: 'order',
    label: t('common.order'),
    component: 'InputNumber'
  },
  {
    field: 'meta.title',
    label: t('menu.title'),
    component: 'Input'
  },
  {
    field: 'meta.permissions',
    label: t('menu.permission'),
    component: 'TreeSelect',
    componentProps: {
      nodeKey: 'id',
      props: {
        label: 'name',
        value: 'scope',
        children: 'children'
      },
      highlightCurrent: true,
      clearable: true,
      multiple: true,
      filterable: true,
      checkStrictly: true,
      tagType: 'primary'
    } as TreeSelectComponentProps,
    optionApi: async () => {
      const [res, err] = await api.permission.getMany({ query: { limit: 100 } })
      if (err) return []

      return permissionToTree(res.data)
    }
  },
  {
    field: 'meta.icon',
    label: t('menu.icon'),
    component: 'IconPicker'
  },
  {
    field: 'meta.activeMenu',
    label: t('menu.activeMenu'),
    component: 'Input'
  },

  {
    field: 'isEnable',
    label: t('menu.isEnable'),
    component: 'Switch',
    componentProps: {
      options: [
        {
          label: t('tagStatus.disable'),
          value: false
        },
        {
          label: t('tagStatus.enable'),
          value: true
        }
      ]
    }
  },

  {
    field: 'meta.hidden',
    label: t('menu.hidden'),
    component: 'Switch'
  },
  {
    field: 'meta.alwaysShow',
    label: t('menu.alwaysShow'),
    component: 'Switch'
  },
  {
    field: 'meta.noCache',
    label: t('menu.noCache'),
    component: 'Switch'
  },
  {
    field: 'meta.breadcrumb',
    label: t('menu.breadcrumb'),
    component: 'Switch'
  },
  {
    field: 'meta.affix',
    label: t('menu.affix'),
    component: 'Switch'
  },
  {
    field: 'meta.noTagsView',
    label: t('menu.noTagsView'),
    component: 'Switch'
  },
  {
    field: 'meta.canTo',
    label: t('menu.canTo'),
    component: 'Switch'
  }
])

const rules = reactive({
  name: [required()],
  type: [required()],
  component: [required()],
  path: [required()],
  'meta.title': [required()]
})

const { formRegister, formMethods } = useForm()
const { setValues, getFormData, getElFormExpose, setSchema } = formMethods

const submit = async () => {
  const elForm = await getElFormExpose()
  const valid = await elForm?.validate().catch((err) => {
    console.log(err)
  })
  if (valid) {
    const formData = await getFormData()
    return formData
  }
}

const cacheComponent = ref('')

watch(
  () => props.currentRow,
  (value) => {
    if (!value) return
    const currentRow = cloneDeep(value)
    console.log({ currentRow })

    cacheComponent.value = currentRow.type === MenuTypeEnum.MENU ? currentRow.component : ''
    if (!currentRow.pid) {
      setSchema([
        {
          field: 'component',
          path: 'componentProps.disabled',
          value: true
        }
      ])
    } else {
      setSchema([
        {
          field: 'component',
          path: 'componentProps.disabled',
          value: false
        }
      ])
    }
    if (currentRow.type === MenuTypeEnum.MENU) {
      setSchema([
        {
          field: 'component',
          path: 'componentProps.disabled',
          value: false
        }
      ])
    } else {
      setSchema([
        {
          field: 'component',
          path: 'componentProps.disabled',
          value: true
        }
      ])
    }
    setValues(currentRow)
  },
  {
    deep: true,
    immediate: true
  }
)

defineExpose({
  submit
})
</script>

<template>
  <Form :rules="rules" @register="formRegister" :schema="formSchema" />
</template>
