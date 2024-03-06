<script setup lang="tsx">
import { convertEdgeChildren } from '@/api/admin/types'
import { getRouteListApi } from '@/api/menu'
import { Form, FormSchema } from '@/components/Form'
import { useForm } from '@/hooks/web/useForm'
import { useI18n } from '@/hooks/web/useI18n'
import { useValidator } from '@/hooks/web/useValidator'
import { cloneDeep } from 'lodash-es'
import { PropType, reactive, ref, unref, watch } from 'vue'
import AddButtonPermission from './AddButtonPermission.vue'

const { t } = useI18n()

const { required } = useValidator()

const props = defineProps({
  currentRow: {
    type: Object as PropType<any>,
    default: () => null
  }
})

const handleClose = async (tag: any) => {
  const formData = await getFormData()
  // Delete the corresponding permissions
  setValues({
    permissionList: formData?.permissionList?.filter((v: any) => v.value !== tag.value)
  })
}

const showDrawer = ref(false)

const formSchema = reactive<FormSchema[]>([
  {
    field: 'type',
    label: t('meta.type'),
    component: 'RadioButton',
    value: 0,
    colProps: {
      span: 24
    },
    componentProps: {
      options: [
        {
          label: 'Table of contents',
          value: 0
        },
        {
          label: 'menu',
          value: 1
        }
      ],
      on: {
        change: async (val: number) => {
          const formData = await getFormData()
          if (val === 1) {
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

            if (formData.parentId === void 0) {
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
    field: 'parentId',
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
          if (val && formData.type === 0) {
            setValues({
              component: '##'
            })
          } else if (!val && formData.type === 0) {
            setValues({
              component: '#'
            })
          } else if (formData.type === 1) {
            setValues({
              component: unref(cacheComponent) ?? ''
            })
          }
        }
      }
    },
    optionApi: async () => {
      const res = await getRouteListApi()
      const list = convertEdgeChildren(res.data.list)
      return list || []
    }
  },
  {
    field: 'meta.title',
    label: t('meta.title'),
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
    field: 'name',
    label: t('meta.name'),
    component: 'Input'
  },
  {
    field: 'meta.icon',
    label: t('meta.icon'),
    component: 'Input'
  },
  {
    field: 'path',
    label: t('meta.path'),
    component: 'Input'
  },
  {
    field: 'meta.activeMenu',
    label: t('meta.activeMenu'),
    component: 'Input'
  },
  {
    field: 'isEnable',
    label: t('meta.isEnable'),
    component: 'Select',
    componentProps: {
      options: [
        {
          label: t('tagStatus.disable'),
          value: 0
        },
        {
          label: t('tagStatus.enable'),
          value: 1
        }
      ]
    }
  },

  {
    field: 'meta.hidden',
    label: t('meta.hidden'),
    component: 'Switch'
  },
  {
    field: 'meta.alwaysShow',
    label: t('meta.alwaysShow'),
    component: 'Switch'
  },
  {
    field: 'meta.noCache',
    label: t('meta.noCache'),
    component: 'Switch'
  },
  {
    field: 'meta.breadcrumb',
    label: t('meta.breadcrumb'),
    component: 'Switch'
  },
  {
    field: 'meta.affix',
    label: t('meta.affix'),
    component: 'Switch'
  },
  {
    field: 'meta.noTagsView',
    label: t('meta.noTagsView'),
    component: 'Switch'
  },
  {
    field: 'meta.canTo',
    label: t('meta.canTo'),
    component: 'Switch'
  }
])

const rules = reactive({
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
    cacheComponent.value = currentRow.type === 1 ? currentRow.component : ''
    if (currentRow.parentId === 0) {
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
    if (currentRow.type === 1) {
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

const confirm = async (data: any) => {
  const formData = await getFormData()
  setValues({
    permissionList: [...(formData?.permissionList || []), data]
  })
}
</script>

<template>
  <Form :rules="rules" @register="formRegister" :schema="formSchema" />
  <AddButtonPermission v-model="showDrawer" @confirm="confirm" />
</template>
