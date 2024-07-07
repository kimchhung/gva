export type Admin = {
  id: number
  createdAt: Date
  username: string
  isSuperAdmin?: boolean
  edges: {
    roles: any[]
  }

  [key: string]: any
}
