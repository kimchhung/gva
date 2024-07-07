import type { Form, FormExpose } from '@/components/Form'
import { FormProps, FormSchema, FormSetProps } from '@/components/Form'
import { isEmptyVal, isObject } from '@/utils/is'
import type { ElForm, ElFormItem } from 'element-plus'
import { nextTick, ref, unref } from 'vue'

export const useForm = () => {
  // From instance
  const formRef = ref<typeof Form & FormExpose>()

  // ElFormInstance
  const elFormRef = ref<ComponentRef<typeof ElForm>>()

  /**
   * @param ref Form instance
   * @param elRef Elform instance
   */
  const register = (ref: typeof Form & FormExpose, elRef: ComponentRef<typeof ElForm>) => {
    formRef.value = ref
    elFormRef.value = elRef
  }

  const getForm = async () => {
    await nextTick()
    const form = unref(formRef)
    if (!form) {
      console.error('The form is not registered. Please use the register method to register')
    }
    return form
  }

  // Some built -in methods
  const methods = {
    /**
     * @description Set the props of form component
     * @param props props of form component
     */
    setProps: async (props: FormProps = {}) => {
      const form = await getForm()
      form?.setProps(props)
      if (props.model) {
        form?.setValues(props.model)
      }
    },

    /**
     * @description Set the value of form
     * @param data Data needed to be set
     */
    setValues: async (data: Recordable) => {
      const form = await getForm()
      form?.setValues(data)
    },

    /**
     * @description Set SCHEMA
     * @param schemaProps SCHEMAPROPS that needs to be set
     */
    setSchema: async (schemaProps: FormSetProps[]) => {
      const form = await getForm()
      form?.setSchema(schemaProps)
    },

    /**
     * @description New SCHEMA
     * @param formSchema Need to add data
     * @param index Where to add
     */
    addSchema: async (formSchema: FormSchema, index?: number) => {
      const form = await getForm()
      form?.addSchema(formSchema, index)
    },

    /**
     * @description Delete SCHEMA
     * @param field Which data deletes
     */
    delSchema: async (field: string) => {
      const form = await getForm()
      form?.delSchema(field)
    },

    /**
     * @description Get the form data
     * @returns form data
     */
    getFormData: async <T = Recordable>(filterEmptyVal = true): Promise<T> => {
      const form = await getForm()
      const model = form?.formModel as any
      if (filterEmptyVal) {
        // Use Reduce to filter the empty value and return a new object
        return Object.keys(model).reduce((prev, next) => {
          const value = model[next]
          if (!isEmptyVal(value)) {
            if (isObject(value)) {
              if (Object.keys(value).length > 0) {
                prev[next] = value
              }
            } else {
              prev[next] = value
            }
          }
          return prev
        }, {}) as T
      } else {
        return model as T
      }
    },

    /**
     * @description Examination of form component of form components
     * @param field Form unique logo
     * @returns component instance
     */
    getComponentExpose: async (field: string) => {
      const form = await getForm()
      return form?.getComponentExpose(field)
    },

    /**
     * @description Example of obtaining formItem component
     * @param field Form unique logo
     * @returns formItem instance
     */
    getFormItemExpose: async (field: string) => {
      const form = await getForm()
      return form?.getFormItemExpose(field) as ComponentRef<typeof ElFormItem>
    },

    /**
     * @description Examples to obtain Elform components
     * @returns ElForm instance
     */
    getElFormExpose: async () => {
      await getForm()
      return unref(elFormRef)
    },

    getFormExpose: async () => {
      await getForm()
      return unref(formRef)
    }
  }

  return {
    formRegister: register,
    formMethods: methods
  }
}
