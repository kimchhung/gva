import { ElOption, ElOptionGroup } from 'element-plus'
import { FormSchema, SelectComponentProps, SelectOption } from '../types'

export const useRenderSelect = () => {
  // Rendering select options
  const renderSelectOptions = (item: FormSchema) => {
    const componentsProps = item?.componentProps as SelectComponentProps
    const optionGroupDefaultSlot = componentsProps?.slots?.optionGroupDefault
    // If you have aliases, take aliases
    const labelAlias = componentsProps?.props?.label
    const keyAlias = componentsProps?.props?.key
    return componentsProps?.options?.map((option) => {
      if (option?.options?.length) {
        return optionGroupDefaultSlot ? (
          optionGroupDefaultSlot(option)
        ) : (
          <ElOptionGroup label={option[labelAlias || 'label']} key={option[keyAlias || 'key']}>
            {{
              default: () =>
                option?.options?.map((v) => {
                  return renderSelectOptionItem(item, v)
                })
            }}
          </ElOptionGroup>
        )
      } else {
        return renderSelectOptionItem(item, option)
      }
    })
  }

  // Rendering select option item
  const renderSelectOptionItem = (item: FormSchema, option: SelectOption) => {
    // If you have aliases, take alias
    const componentsProps = item.componentProps as SelectComponentProps
    const labelAlias = componentsProps?.props?.label
    const valueAlias = componentsProps?.props?.value
    const keyAlias = componentsProps?.props?.key
    const optionDefaultSlot = componentsProps.slots?.optionDefault

    return (
      <ElOption
        {...option}
        key={option[keyAlias || 'key']}
        label={option[labelAlias || 'label']}
        value={option[valueAlias || 'value']}
      >
        {{
          default: () => (optionDefaultSlot ? optionDefaultSlot(option) : undefined)
        }}
      </ElOption>
    )
  }

  return {
    renderSelectOptions
  }
}
