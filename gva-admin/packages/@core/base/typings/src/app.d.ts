type LayoutType =
  | 'full-content'
  | 'header-nav'
  | 'mixed-nav'
  | 'sidebar-mixed-nav'
  | 'sidebar-nav';

type ThemeModeType = 'auto' | 'dark' | 'light';

type BuiltinThemeType =
  | 'custom'
  | 'deep-blue'
  | 'deep-green'
  | 'default'
  | 'gray'
  | 'green'
  | 'neutral'
  | 'orange'
  | 'pink'
  | 'red'
  | 'rose'
  | 'sky-blue'
  | 'slate'
  | 'stone'
  | 'violet'
  | 'yellow'
  | 'zinc'
  | (Record<never, never> & string);

type ContentCompactType = 'compact' | 'wide';

type LayoutHeaderModeType = 'auto' | 'auto-scroll' | 'fixed' | 'static';

/**
 * Log in the expiration mode
 * modal Pop -up mode
 * page Page mode
 */
type LoginExpiredModeType = 'modal' | 'page';

/**
 * Bread dandruff
 * background background
 * normal default
 */
type BreadcrumbStyleType = 'background' | 'normal';

/**
 * Permissions
 * backend Back-end authority mode
 * frontend Front-end authority mode
 */
type AccessModeType = 'backend' | 'frontend';

/**
 * Navigation style
 * plain simple
 * rounded Round
 */
type NavigationStyleType = 'plain' | 'rounded';

/**
 * Tag bar style
 * brisk 轻快
 * card 卡片
 * chrome 谷歌
 * plain 朴素
 */
type TabsStyleType = 'brisk' | 'card' | 'chrome' | 'plain';

/**
 * Page switch animation
 */
type PageTransitionType = 'fade' | 'fade-down' | 'fade-slide' | 'fade-up';

/**
 * Page switch animation
 * panel-center 居中布局
 * panel-left 居左布局
 * panel-right 居右布局
 */
type AuthPageLayoutType = 'panel-center' | 'panel-left' | 'panel-right';

export type {
  AccessModeType,
  AuthPageLayoutType,
  BreadcrumbStyleType,
  BuiltinThemeType,
  ContentCompactType,
  LayoutHeaderModeType,
  LayoutType,
  LoginExpiredModeType,
  NavigationStyleType,
  PageTransitionType,
  TabsStyleType,
  ThemeModeType,
};
