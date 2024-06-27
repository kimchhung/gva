/// <reference types="vite/client" />

declare module '*.vue' {
  import { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any>
  export default component
}

type ImportMetaEnv = {
  readonly VITE_NODE_ENV: string
  readonly VITE_APP_TITLE: string
  readonly VITE_API_BASE_PATH: string
  readonly VITE_BASE_PATH: string
  readonly VITE_DROP_DEBUGGER: string
  readonly VITE_DROP_CONSOLE: string
  readonly VITE_SOURCEMAP: string
  readonly VITE_OUT_DIR: string
  readonly VITE_USE_BUNDLE_ANALYZER: string
  readonly VITE_USE_ALL_ELEMENT_PLUS_STYLE: string
  readonly VITE_USE_MOCK: string
  readonly VITE_USE_CSS_SPLIT: string
  readonly VITE_USE_ONLINE_ICON: string
}

declare global {
  type ImportMeta = {
    readonly env: ImportMetaEnv
  }
}