export type Department = {
  id: string
  name: string
  nameId: string
  isEnable: boolean

  //relation
  edge: {
    children?: Department[]
  }
  parentId: string
}
