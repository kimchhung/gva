import { computed } from 'vue';

import { diff } from '@vben-core/shared';

import { preferencesManager } from './preferences';
import { isDarkTheme } from './update-css-variables';

function usePreferences() {
  const preferences = preferencesManager.getPreferences();
  const initialPreferences = preferencesManager.getInitialPreferences();
  /**
   * Calculation preference settings change
   */
  const diffPreference = computed(() => {
    return diff(initialPreferences, preferences);
  });

  const appPreferences = computed(() => preferences.app);

  const shortcutKeysPreferences = computed(() => preferences.shortcutKeys);

  /**
   * Determine whether it is a dark mode
   * @param Preferences -The current preference setting object, its theme value will be used to determine whether it is a dark mode.
   * @ReturnS If the theme is dark mode, return to TRUE, otherwise returns false.
   */
  const isDark = computed(() => {
    return isDarkTheme(preferences.theme.mode);
  });

  const isMobile = computed(() => {
    return appPreferences.value.isMobile;
  });

  const theme = computed(() => {
    return isDark.value ? 'dark' : 'light';
  });

  /**
   * Layout
   */
  const layout = computed(() =>
    isMobile.value ? 'sidebar-nav' : appPreferences.value.layout,
  );

  /**
   * Whether the full screen is displayed in Content, there is no need for side, bottom, top, TAB area
   */
  const isFullContent = computed(
    () => appPreferences.value.layout === 'full-content',
  );

  /**
   * Whether the side navigation mode
   */
  const isSideNav = computed(
    () => appPreferences.value.layout === 'sidebar-nav',
  );

  /**
   * Whether the side mixed mode
   */
  const isSideMixedNav = computed(
    () => appPreferences.value.layout === 'sidebar-mixed-nav',
  );

  /**
   * Whether it is a head navigation mode
   */
  const isHeaderNav = computed(
    () => appPreferences.value.layout === 'header-nav',
  );

  /**
   * Whether it is a hybrid navigation mode
   */
  const isMixedNav = computed(
    () => appPreferences.value.layout === 'mixed-nav',
  );

  /**
   * Does it include side navigation mode
   */
  const isSideMode = computed(() => {
    return isMixedNav.value || isSideMixedNav.value || isSideNav.value;
  });

  const sidebarCollapsed = computed(() => {
    return preferences.sidebar.collapsed;
  });

  /**
   * Whether to turn on Keep-Alive
   * Start when you can see it and start the Keep-Alive
   */
  const keepAlive = computed(
    () => preferences.tabbar.enable && preferences.tabbar.keepAlive,
  );

  /**
   * Whether the layout of the login registration page is the left side
   */
  const authPanelLeft = computed(() => {
    return appPreferences.value.authPageLayout === 'panel-left';
  });

  /**
   * 登录注册页面布局是否为左侧
   */
  const authPanelRight = computed(() => {
    return appPreferences.value.authPageLayout === 'panel-right';
  });

  /**
   * Whether the layout of the login registration page is the middle
   */
  const authPanelCenter = computed(() => {
    return appPreferences.value.authPageLayout === 'panel-center';
  });

  /**
   * Is the content maximized
   * Exclude Full-CONTENT mode
   */
  const contentIsMaximize = computed(() => {
    const headerIsHidden = preferences.header.hidden;
    const sidebarIsHidden = preferences.sidebar.hidden;
    return headerIsHidden && sidebarIsHidden && !isFullContent.value;
  });

  /**
   * Whether the global search shortcut is enabled
   */
  const globalSearchShortcutKey = computed(() => {
    const { enable, globalSearch } = shortcutKeysPreferences.value;
    return enable && globalSearch;
  });
  /**
   * Whether to enable global cancellation shortcut keys
   */
  const globalLogoutShortcutKey = computed(() => {
    const { enable, globalLogout } = shortcutKeysPreferences.value;
    return enable && globalLogout;
  });

  const globalLockScreenShortcutKey = computed(() => {
    const { enable, globalLockScreen } = shortcutKeysPreferences.value;
    return enable && globalLockScreen;
  });

  /**
   * Whether to enable global preference to set shortcut keys
   */
  const globalPreferencesShortcutKey = computed(() => {
    const { enable, globalPreferences } = shortcutKeysPreferences.value;
    return enable && globalPreferences;
  });

  return {
    authPanelCenter,
    authPanelLeft,
    authPanelRight,
    contentIsMaximize,
    diffPreference,
    globalLockScreenShortcutKey,
    globalLogoutShortcutKey,
    globalPreferencesShortcutKey,
    globalSearchShortcutKey,
    isDark,
    isFullContent,
    isHeaderNav,
    isMixedNav,
    isMobile,
    isSideMixedNav,
    isSideMode,
    isSideNav,
    keepAlive,
    layout,
    sidebarCollapsed,
    theme,
  };
}

export { usePreferences };
