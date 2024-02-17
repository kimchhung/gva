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

    /**
     * Permission Interface, "ADMIN.VIEW" | "1234"
     */
    code?: number | string

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

/**
 * Interface for the settings configuration of the application.
 * @interface SettingsConfig
 */
export interface SettingsConfig {
  /**
   * The title of the application.
   * @type {string}
   */
  title: string

  /**
   * Whether to display the logo in the sidebar.
   * @type {boolean}
   */
  sidebarLogo: boolean

  /**
   * Whether to display the settings right-panel.
   * @type {boolean}
   */
  showLeftMenu: boolean

  /**
   * Whether to display the drop-down.
   * @type {boolean}
   */
  showDropDown: boolean

  /**
   * Whether to display the Hamburger menu.
   * @type {boolean}
   */
  showHamburger: boolean

  /**
   * Whether the application requires login.
   * @type {boolean}
   */
  isNeedLogin: boolean

  /**
   * Whether to enable nprogress for loading indicators.
   * @type {boolean}
   */
  isNeedNprogress: boolean

  /**
   * Whether to display the TagsView.
   * @type {boolean}
   */
  showTagsView: boolean

  /**
   * The number of tags to show in the TagsView.
   * @type {number}
   */
  tagsViewNum: number

  /**
   * Whether to open production mock.
   * @type {boolean}
   */
  openProdMock: boolean

  /**
   * The environment(s) in which to show the error logs component.
   * @type {string | Array<string>}
   */
  errorLog: string | Array<string>

  /**
   * The mode for handling permissions, can be 'rbac', 'roles', or 'code'.
   * @type {string}
   */
  permissionMode: string

  /**
   * The height of the table (100vh-delWindowHeight).
   * @type {string}
   */
  delWindowHeight: string

  /**
   * The development token to use when `isNeedLogin` is `false`.
   * @type {string}
   */
  tmpToken: string

  /**
   * Whether to display the title in the Navbar.
   * @type {boolean}
   */
  showNavbarTitle: boolean

  /**
   * Whether to display the top Navbar.
   * @type {boolean}
   */
  showTopNavbar: boolean

  /**
   * Whether to enable animation of the main area.
   * @type {boolean}
   */
  mainNeedAnimation: boolean

  /**
   * The base path for the vite configuration.
   * @type {string}
   */
  viteBasePath: string

  /**
   * The default language for i18n.
   * @type {string}
   */
  defaultLanguage: string

  /**
   * The default size for the application.
   * @type {string}
   */
  defaultSize: string

  /**
   * The default theme for the application.
   * @type {string}
   */
  defaultTheme: string

  /**
   * The platform ID for the application.
   * @type {number}
   */
  plateFormId: number
}

export {}
