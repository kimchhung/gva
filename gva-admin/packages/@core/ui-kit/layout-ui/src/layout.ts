import type {
  ContentCompactType,
  LayoutHeaderModeType,
  LayoutType,
  ThemeModeType,
} from '@gva-core/typings';

interface LayoutProps {
  /**
   * Content area width
   * @default 'wide'
   */
  contentCompact?: ContentCompactType;
  /**
   * Ding width layout width
   * @default 1200
   */
  contentCompactWidth?: number;
  /**
   * padding
   * @default 16
   */
  contentPadding?: number;
  /**
   * paddingBottom
   * @default 16
   */
  contentPaddingBottom?: number;
  /**
   * paddingLeft
   * @default 16
   */
  contentPaddingLeft?: number;
  /**
   * paddingRight
   * @default 16
   */
  contentPaddingRight?: number;
  /**
   * paddingTop
   * @default 16
   */
  contentPaddingTop?: number;
  /**
   * footer Whether it is visible
   * @default false
   */
  footerEnable?: boolean;
  /**
   * footer Whether it is fixed
   * @default true
   */
  footerFixed?: boolean;
  /**
   * footer high
   * @default 32
   */
  footerHeight?: number;

  /**
   * header height
   * @default 48
   */
  headerHeight?: number;
  /**
   * Head height increases height
   * When there is navigation at the top, an additional height height
   * @default 10
   */
  headerHeightOffset?: number;
  /**
   * Whether the top bar is hidden
   * @default false
   */
  headerHidden?: boolean;
  /**
   * header 显示模式
   * @default 'fixed'
   */
  headerMode?: LayoutHeaderModeType;
  /**
   * Whether to display the header switching the sidebar button
   * @default
   */
  headerToggleSidebarButton?: boolean;
  /**
   * header是否显示
   * @default true
   */
  headerVisible?: boolean;
  /**
   * Whether to display the moving end
   * @default false
   */
  isMobile?: boolean;
  /**
  * Layout method
  * Sidebar-NAV side menu layout
  * Header-nav Top menu layout
  * Mixed-NAV side & top menu layout
  * Sidebar-Mixed-NAV side mixed menu layout
  * FULL-CONTENT full screen content layout
   * @default sidebar-nav
   */
  layout?: LayoutType;
  /**
   * Side menu folding status
   * @default false
   */
  sidebarCollapse?: boolean;
  /**
   * 侧边菜单是否折叠时，是否显示title
   * @default true
   */
  sidebarCollapseShowTitle?: boolean;
  /**
   * 侧边栏是否可见
   * @default true
   */
  sidebarEnable?: boolean;
  /**
   * 侧边菜单折叠额外宽度
   * @default 48
   */
  sidebarExtraCollapsedWidth?: number;
  /**
   * 侧边栏是否隐藏
   * @default false
   */
  sidebarHidden?: boolean;
  /**
   * 混合侧边栏宽度
   * @default 80
   */
  sidebarMixedWidth?: number;
  /**
   * 侧边栏是否半深色
   * @default false
   */
  sidebarSemiDark?: boolean;
  /**
   * 侧边栏
   * @default dark
   */
  sidebarTheme?: ThemeModeType;
  /**
   * 侧边栏宽度
   * @default 210
   */
  sidebarWidth?: number;
  /**
   *  侧边菜单折叠宽度
   * @default 48
   */
  sideCollapseWidth?: number;
  /**
   * tab是否可见
   * @default true
   */
  tabbarEnable?: boolean;
  /**
   * tab高度
   * @default 30
   */
  tabbarHeight?: number;
  /**
   * zIndex
   * @default 100
   */
  zIndex?: number;
}
export type { LayoutProps };
