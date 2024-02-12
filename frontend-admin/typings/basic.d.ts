/*
 * 声明.d.ts文件规范
 * 导出的类型以大写开头
 * 对象：config
 * 数组：options
 * 枚举：emu
 * 函数：Fn
 * 属性：props
 * 实例：instance
 * */

/*router*/

declare module 'vue-router' {
  interface RouteRecordSingleViewWithChildren {
    hidden?: boolean
    alwaysShow?: boolean
  }

  interface RouteLocationPathRaw {
    fullPath?: string
  }

  interface RouteRecordRaw {
    fullPath?: string
    hidden?: boolean
    alwaysShow?: boolean
    code?: number
    name?: string
    path?: string
    children?: RouteRecordRaw[]
    redirect?: string
  }

  interface RouteMeta {
    title: string
    icon?: string
    affix?: boolean
    activeMenu?: string
    breadcrumb?: boolean
    roles?: Array<string>
    elSvgIcon?: string
    code?: number

    /**
     * cachePage when page enter, default false
     */
    cachePage?: boolean

    /**
     *  remove cachePage when page leave, default false
     */
    leaveRmCachePage?: boolean

    /**
     *  closeTabRmCache: remove cachePage when tabs close, default false
     */
    closeTabRmCache?: boolean
  }
}

/*settings*/
export interface SettingsConfig {
  title: string
  sidebarLogo: boolean
  showLeftMenu: boolean
  ShowDropDown: boolean
  showHamburger: boolean
  isNeedLogin: boolean
  isNeedNprogress: boolean
  showTagsView: boolean
  tagsViewNum: number
  openProdMock: boolean
  errorLog: string | Array<string>
  permissionMode: string
  delWindowHeight: string
  tmpToken: string
  showNavbarTitle: boolean
  showTopNavbar: boolean
  mainNeedAnimation: boolean
  viteBasePath: string
  defaultLanguage: string
  defaultSize: string
  defaultTheme: string
  plateFormId: number
}

export {}
