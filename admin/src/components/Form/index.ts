import Form from './src/Form.vue'
import type { FormSchema, FormSetProps } from './src/types'
export type {
  AutocompleteComponentProps,
  CascaderComponentProps,
  CheckboxGroupComponentProps,
  CheckboxOption,
  ColProps,
  ColorPickerComponentProps,
  ComponentName,
  ComponentNameEnum,
  DatePickerComponentProps,
  DateTimePickerComponentProps,
  DividerComponentProps,
  FormItemProps,
  FormProps,
  FormSchema,
  FormSetProps,
  InputComponentProps,
  InputNumberComponentProps,
  InputPasswordComponentProps,
  PlaceholderModel,
  RadioButtonComponentProps,
  RadioGroupComponentProps,
  RadioOption,
  RateComponentProps,
  SelectComponentProps,
  SelectOption,
  SelectV2ComponentProps,
  SwitchComponentProps,
  TimePickerComponentProps,
  TimeSelectComponentProps,
  TransferComponentProps,
  TreeSelectComponentProps
} from './src/types'

export type FormExpose = {
  setValues: (data: Recordable) => void
  setProps: (props: Recordable) => void
  delSchema: (field: string) => void
  addSchema: (formSchema: FormSchema, index?: number) => void
  setSchema: (schemaProps: FormSetProps[]) => void
  formModel: Recordable
  getComponentExpose: (field: string) => any
  getFormItemExpose: (field: string) => any
}

export { Form }
