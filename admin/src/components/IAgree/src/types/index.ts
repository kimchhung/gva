export type LinkItem = {
  text: string
  url?: string
  onClick?: () => void
}

export type IAgreeProps = {
  text: string
  link: LinkItem[]
}
