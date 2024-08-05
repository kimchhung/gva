export type Admin = {
  id: string
  createdAt: Date
  username: string
  isSuperAdmin?: boolean
  edges: {
    roles: any[]
  }

  [key: string]: any
}
