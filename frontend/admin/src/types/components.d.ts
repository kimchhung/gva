declare module 'vue' {
  export type GlobalComponents = {
    Icon: (typeof import('../components/Icon/index'))['Icon']
    Permission: (typeof import('../components/Permission/index'))['Permission']
    BaseButton: (typeof import('../components/Button/index'))['BaseButton']
  }
}

export {}
