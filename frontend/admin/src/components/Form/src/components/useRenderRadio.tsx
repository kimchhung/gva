import { ElRadio, ElRadioButton } from 'element-plus'
import { defineComponent } from 'vue'
import { ComponentNameEnum, FormSchema, RadioGroupComponentProps } from '../types'

export const useRenderRadio = () => {
  const renderRadioOptions = (item: FormSchema) => {
    // If you have aliases, take aliases
    const componentProps = item?.componentProps as RadioGroupComponentProps
    const valueAlias = componentProps?.props?.value || 'value'
    const labelAlias = componentProps?.props?.label || 'label'
    const disabledAlias = componentProps?.props?.disabled || 'disabled'
    const Com = (
      item.component === ComponentNameEnum.RADIO_GROUP ? ElRadio : ElRadioButton
    ) as ReturnType<typeof defineComponent>
    return componentProps?.options?.map((option) => {
      const { value, ...other } = option
      return (
        <Com
          {...other}
          disabled={option[disabledAlias || 'disabled']}
          label={option[valueAlias || 'value']}
        >
          {option[labelAlias || 'label']}
        </Com>
      )
    })
  }

  return {
    renderRadioOptions
  }
}
