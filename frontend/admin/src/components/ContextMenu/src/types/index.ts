export type ContextMenuSchema = {
  disabled?: boolean
  divided?: boolean
  icon?: string
  label: string
  command?: (item: ContextMenuSchema) => void
}
