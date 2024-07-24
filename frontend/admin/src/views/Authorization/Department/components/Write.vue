<script setup lang="tsx">
import { Department } from '@/api/department/types'
import { Form, FormSchema, TreeSelectComponentProps } from '@/components/Form'
import { useForm } from '@/hooks/web/useForm'
import { useI18n } from '@/hooks/web/useI18n'
import { QueryUrl } from '@/hooks/web/usePagi'
import { useValidator } from '@/hooks/web/useValidator'
import { convertEdgeChildren, edgelistToTree } from '@/utils/edgeTree'
import { PropType, reactive } from 'vue'

defineProps({
  currentRow: {
    type: Object as PropType<Department>,
    default: () => null
  }
})

const { t } = useI18n()
const { required } = useValidator()
const formSchema = reactive<FormSchema[]>([
  {
    field: 'name',
    label: t('name'),
    component: 'Input'
  },
  {
    field: 'nameId',
    label: t('nameId'),
    component: 'Input'
  },
  {
    field: 'pid',
    label: t('common.parent'),
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
      clearable: true
    } as TreeSelectComponentProps,
    optionApi: async () => {
      const [res] = await api.department.getMany({ query: new QueryUrl(100) })
      if (!res?.success) {
        return []
      }

      const { data } = res
      return convertEdgeChildren(edgelistToTree(data))
    }
  },
  {
    field: 'isEnable',
    label: t('common.isEnable'),
    component: 'Switch',
    componentProps: {
      options: [
        {
          label: t('commmon.disable'),
          value: false
        },
        {
          label: t('commmon.enable'),
          value: true
        }
      ]
    }
  }
])

const rules = reactive({
  name: [required()],
  nameId: [required()]
})
const { formRegister, formMethods } = useForm()
const { getFormData, getElFormExpose } = formMethods

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

defineExpose({
  submit
})
</script>

<template>
  <Form :rules="rules" @register="formRegister" :schema="formSchema" />
</template>
