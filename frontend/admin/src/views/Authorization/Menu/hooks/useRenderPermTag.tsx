import { ElTag } from 'element-plus'

export const renderPermisionTag = (permissions: string[]) => {
  return (
    <>
      {permissions?.map((scope) => {
        const [, action] = scope.split(':')
        const tagTypes = {
          super: 'primary',
          add: 'warning',
          delete: 'danger',
          edit: 'warning',
          view: 'info'
        }
        const typeColor = tagTypes[action]
        return (
          <ElTag class="mr-1" type={typeColor}>
            {action}
          </ElTag>
        )
      })}
    </>
  )
}
