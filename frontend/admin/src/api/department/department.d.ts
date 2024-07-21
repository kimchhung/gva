import { DepartmentAPI } from '.'

declare global {
  interface API {
    department: DepartmentAPI
  }
}
