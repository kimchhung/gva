import type { TabDefinition } from '@gva-core/typings';
import type { Router, RouteRecordNormalized } from 'vue-router';

import { toRaw } from 'vue';

import { openWindow, startProgress, stopProgress } from '@gva-core/shared';

import { acceptHMRUpdate, defineStore } from 'pinia';

interface TabbarState {
  /**
   * The currently opened tab list cache
   */
  cachedTabs: Set<string>;
  /**
   * The index of dragging and dragging
   */
  dragEndIndex: number;
  /**
   * The tab need to exclude the cache
   */
  excludeCachedTabs: Set<string>;
  /**
   * Whether to refresh
   */
  renderRouteView?: boolean;
  /**
   * List of the currently opened tab
   */
  tabs: TabDefinition[];
  /**
   * Update time, for some update scenarios, use Watch in -depth monitoring, will lose performance
   */
  updateTime?: number;
}

/**
 * View permissions related
 */
export const useTabbarStore = defineStore('core-tabbar', {
  actions: {
    /**
     * Close tabs in bulk
     */
    async _bulkCloseByPaths(paths: string[]) {
      this.tabs = this.tabs.filter((item) => {
        return !paths.includes(getTabPath(item));
      });

      this.updateCacheTab();
    },
    /**
     * Close tab pages
     * @param tab
     */
    _close(tab: TabDefinition) {
      const { fullPath } = tab;
      if (isAffixTab(tab)) {
        return;
      }
      const index = this.tabs.findIndex((item) => item.fullPath === fullPath);
      index !== -1 && this.tabs.splice(index, 1);
    },
    /**
     * Jump to the default tab page
     */
    async _goToDefaultTab(router: Router) {
      if (this.getTabs.length <= 0) {
        // TODO: Jump homepage
        return;
      }
      const firstTab = this.getTabs[0];
      await this._goToTab(firstTab, router);
    },
    /**
     * Jump to the tab page
     * @param tab
     */
    async _goToTab(tab: TabDefinition, router: Router) {
      const { params, path, query } = tab;
      const toParams = {
        params: params || {},
        path,
        query: query || {},
      };
      await router.replace(toParams);
    },
    /**
     * @zh_CN Add tab pages
     * @param routeTab
     */
    addTab(routeTab: TabDefinition) {
      const tab = cloneTab(routeTab);
      if (!isTabShown(tab)) {
        return;
      }

      const tabIndex = this.tabs.findIndex((tab) => {
        return getTabPath(tab) === getTabPath(routeTab);
      });

      if (tabIndex === -1) {
        //Get the dynamic route opening number, exceeding 0 means that the number of open numbers need to be controlled
        const maxNumOfOpenTab = (routeTab?.meta?.maxNumOfOpenTab ?? -1) as number;
        // If the dynamic routing level is greater than 0, then the number of openings of the route must be limited.
        // Get the number of dynamic routing that has been opened to determine whether it is greater than a certain value
        if (
          maxNumOfOpenTab > 0 &&
          this.tabs.filter((tab) => tab.name === routeTab.name).length >= maxNumOfOpenTab
        ) {
          // 关闭第一个
          const index = this.tabs.findIndex((item) => item.name === routeTab.name);
          index !== -1 && this.tabs.splice(index, 1);
        }

        this.tabs.push(tab);
      } else {
        //The page has already existed, and the tab is not repeated, only update the tab parameter
        const currentTab = toRaw(this.tabs)[tabIndex];
        const mergedTab = { ...currentTab, ...tab };
        if (Reflect.has(currentTab.meta, 'affixTab')) {
          mergedTab.meta.affixTab = currentTab.meta.affixTab;
        }
        this.tabs.splice(tabIndex, 1, mergedTab);
      }
      this.updateCacheTab();
    },
    /**
     * Turn off all tab pages
     */
    async closeAllTabs(router: Router) {
      this.tabs = this.tabs.filter((tab) => isAffixTab(tab));
      await this._goToDefaultTab(router);
      this.updateCacheTab();
    },
    /**
     * Turn off the left tab page
     * @param tab
     */
    async closeLeftTabs(tab: TabDefinition) {
      const index = this.tabs.findIndex((item) => getTabPath(item) === getTabPath(tab));

      if (index < 1) {
        return;
      }

      const leftTabs = this.tabs.slice(0, index);
      const paths: string[] = [];

      for (const item of leftTabs) {
        if (!isAffixTab(item)) {
          paths.push(getTabPath(item));
        }
      }
      await this._bulkCloseByPaths(paths);
    },
    /**
     * @zh_CN Turn off other tag pages
     * @param tab
     */
    async closeOtherTabs(tab: TabDefinition) {
      const closePaths = this.tabs.map((item) => getTabPath(item));

      const paths: string[] = [];

      for (const path of closePaths) {
        if (path !== tab.fullPath) {
          const closeTab = this.tabs.find((item) => getTabPath(item) === path);
          if (!closeTab) {
            continue;
          }
          if (!isAffixTab(closeTab)) {
            paths.push(getTabPath(closeTab));
          }
        }
      }
      await this._bulkCloseByPaths(paths);
    },
    /**
     * @zh_CN Turn off the right tab page
     * @param tab
     */
    async closeRightTabs(tab: TabDefinition) {
      const index = this.tabs.findIndex((item) => getTabPath(item) === getTabPath(tab));

      if (index >= 0 && index < this.tabs.length - 1) {
        const rightTabs = this.tabs.slice(index + 1);

        const paths: string[] = [];
        for (const item of rightTabs) {
          if (!isAffixTab(item)) {
            paths.push(getTabPath(item));
          }
        }
        await this._bulkCloseByPaths(paths);
      }
    },

    /**
     * @zh_CN Close tab pages
     * @param tab
     * @param router
     */
    async closeTab(tab: TabDefinition, router: Router) {
      const { currentRoute } = router;

      // Closing is not the activation tab
      if (getTabPath(currentRoute.value) !== getTabPath(tab)) {
        this._close(tab);
        this.updateCacheTab();
        return;
      }
      const index = this.getTabs.findIndex(
        (item) => getTabPath(item) === getTabPath(currentRoute.value)
      );

      const before = this.getTabs[index - 1];
      const after = this.getTabs[index + 1];

      // The next TAB exists, jump to the next one
      if (after) {
        this._close(currentRoute.value);
        await this._goToTab(after, router);
        // The previous TAB exists, jump to the previous one
      } else if (before) {
        this._close(currentRoute.value);
        await this._goToTab(before, router);
      } else {
        console.error('Failed to close the tab; only one tab remains open.');
      }
    },

    /**
     * Close the tab via Key
     * @param key
     */
    async closeTabByKey(key: string, router: Router) {
      const index = this.tabs.findIndex((item) => getTabPath(item) === key);
      if (index === -1) {
        return;
      }

      await this.closeTab(this.tabs[index], router);
    },

    /**
     * Get the tab page according to the path
     * @param path
     */
    getTabByPath(path: string) {
      return this.getTabs.find((item) => getTabPath(item) === path) as TabDefinition;
    },
    /**
     * @zh_CN New window open tab
     * @param tab
     */
    async openTabInNewWindow(tab: TabDefinition) {
      const { hash, origin } = location;
      const path = tab.fullPath;
      const fullPath = path.startsWith('/') ? path : `/${path}`;
      const url = `${origin}${hash ? '/#' : ''}${fullPath}`;
      openWindow(url, { target: '_blank' });
    },

    /**
     * @zh_CN Fixed tab page
     * @param tab
     */
    async pinTab(tab: TabDefinition) {
      const index = this.tabs.findIndex((item) => getTabPath(item) === getTabPath(tab));
      if (index !== -1) {
        tab.meta.affixTab = true;
        // this.addTab(tab);
        this.tabs.splice(index, 1, tab);
      }
    },

    /**
     * 刷新标签页
     */
    async refresh(router: Router) {
      const { currentRoute } = router;
      const { name } = currentRoute.value;

      this.excludeCachedTabs.add(name as string);
      this.renderRouteView = false;
      startProgress();

      await new Promise((resolve) => setTimeout(resolve, 200));

      this.excludeCachedTabs.delete(name as string);
      this.renderRouteView = true;
      stopProgress();
    },

    /**
     * @zh_CN 重置标签页标题
     */
    async resetTabTitle(tab: TabDefinition) {
      if (!tab?.meta?.newTabTitle) {
        return;
      }
      const findTab = this.tabs.find((item) => getTabPath(item) === getTabPath(tab));
      if (findTab) {
        findTab.meta.newTabTitle = undefined;
        await this.updateCacheTab();
      }
    },

    /**
     * 设置固定标签页
     * @param tabs
     */
    setAffixTabs(tabs: RouteRecordNormalized[]) {
      for (const tab of tabs) {
        tab.meta.affixTab = true;
        this.addTab(routeToTab(tab));
      }
    },

    /**
     * @zh_CN 设置标签页标题
     * @param tab
     * @param title
     */
    async setTabTitle(tab: TabDefinition, title: string) {
      const findTab = this.tabs.find((item) => getTabPath(item) === getTabPath(tab));

      if (findTab) {
        findTab.meta.newTabTitle = title;

        await this.updateCacheTab();
      }
    },

    setUpdateTime() {
      this.updateTime = Date.now();
    },
    /**
     * @zh_CN 设置标签页顺序
     * @param oldIndex
     * @param newIndex
     */
    async sortTabs(oldIndex: number, newIndex: number) {
      const currentTab = this.tabs[oldIndex];
      this.tabs.splice(oldIndex, 1);
      this.tabs.splice(newIndex, 0, currentTab);
      this.dragEndIndex = this.dragEndIndex + 1;
    },
    /**
     * @zh_CN 切换固定标签页
     * @param tab
     */
    async toggleTabPin(tab: TabDefinition) {
      const affixTab = tab?.meta?.affixTab ?? false;

      await (affixTab ? this.unpinTab(tab) : this.pinTab(tab));
    },

    /**
     * @zh_CN 取消固定标签页
     * @param tab
     */
    async unpinTab(tab: TabDefinition) {
      const index = this.tabs.findIndex((item) => getTabPath(item) === getTabPath(tab));

      if (index !== -1) {
        tab.meta.affixTab = false;
        // this.addTab(tab);
        this.tabs.splice(index, 1, tab);
      }
    },

    /**
     * 根据当前打开的选项卡更新缓存
     */
    async updateCacheTab() {
      const cacheMap = new Set<string>();

      for (const tab of this.tabs) {
        // 跳过不需要持久化的标签页
        const keepAlive = tab.meta?.keepAlive;
        if (!keepAlive) {
          continue;
        }
        tab.matched.forEach((t, i) => {
          if (i > 0) {
            cacheMap.add(t.name as string);
          }
        });

        const name = tab.name as string;
        cacheMap.add(name);
      }
      this.cachedTabs = cacheMap;
    },
  },
  getters: {
    affixTabs(): TabDefinition[] {
      const affixTabs = this.tabs.filter((tab) => isAffixTab(tab));

      return affixTabs.sort((a, b) => {
        const orderA = (a.meta?.affixTabOrder ?? 0) as number;
        const orderB = (b.meta?.affixTabOrder ?? 0) as number;
        return orderA - orderB;
      });
    },
    getCachedTabs(): string[] {
      return [...this.cachedTabs];
    },
    getExcludeCachedTabs(): string[] {
      return [...this.excludeCachedTabs];
    },
    getTabs(): TabDefinition[] {
      const normalTabs = this.tabs.filter((tab) => !isAffixTab(tab));
      return [...this.affixTabs, ...normalTabs].filter(Boolean);
    },
  },
  persist: [
    // tabs不需要保存在localStorage
    {
      paths: ['tabs'],
      storage: sessionStorage,
    },
  ],
  state: (): TabbarState => ({
    cachedTabs: new Set(),
    dragEndIndex: 0,
    excludeCachedTabs: new Set(),
    renderRouteView: true,
    tabs: [],
    updateTime: Date.now(),
  }),
});

// 解决热更新问题
const hot = import.meta.hot;
if (hot) {
  hot.accept(acceptHMRUpdate(useTabbarStore, hot));
}

/**
 * @zh_CN 克隆路由,防止路由被修改
 * @param route
 */
function cloneTab(route: TabDefinition): TabDefinition {
  if (!route) {
    return route;
  }
  const { matched, ...opt } = route;
  return {
    ...opt,
    matched: (matched
      ? matched.map((item) => ({
          meta: item.meta,
          name: item.name,
          path: item.path,
        }))
      : undefined) as RouteRecordNormalized[],
  };
}

/**
 * @zh_CN 是否是固定标签页
 * @param tab
 */
function isAffixTab(tab: TabDefinition) {
  return tab?.meta?.affixTab ?? false;
}

/**
 * @zh_CN 是否显示标签
 * @param tab
 */
function isTabShown(tab: TabDefinition) {
  return !tab.meta.hideInTab;
}

/**
 * @zh_CN 获取标签页路径
 * @param tab
 */
function getTabPath(tab: RouteRecordNormalized | TabDefinition) {
  return decodeURIComponent((tab as TabDefinition).fullPath || tab.path);
}

function routeToTab(route: RouteRecordNormalized) {
  return {
    meta: route.meta,
    name: route.name,
    path: route.path,
  } as TabDefinition;
}
