<script setup lang="ts">
import type { VbenLayoutProps } from './vben-layout';

import type { CSSProperties } from 'vue';
import { computed, ref, watch } from 'vue';

import { useMouse, useScroll, useThrottleFn } from '@vueuse/core';

import {
  LayoutContent,
  LayoutFooter,
  LayoutHeader,
  LayoutSidebar,
  LayoutTabbar,
} from './components';

interface Props extends VbenLayoutProps {}

defineOptions({
  name: 'VbenLayout',
});

const props = withDefaults(defineProps<Props>(), {
  contentCompact: 'wide',
  contentCompactWidth: 1200,
  contentPadding: 0,
  contentPaddingBottom: 0,
  contentPaddingLeft: 0,
  contentPaddingRight: 0,
  contentPaddingTop: 0,
  footerEnable: false,
  footerFixed: true,
  footerHeight: 32,
  headerHeight: 50,
  headerHeightOffset: 10,
  headerHidden: false,
  headerMode: 'fixed',
  headerToggleSidebarButton: true,
  headerVisible: true,
  isMobile: false,
  layout: 'sidebar-nav',
  sidebarCollapseShowTitle: false,
  sidebarExtraCollapsedWidth: 60,
  sidebarHidden: false,
  sidebarMixedWidth: 80,
  sidebarSemiDark: true,
  sidebarTheme: 'dark',
  sidebarWidth: 180,
  sideCollapseWidth: 60,
  tabbarEnable: true,
  tabbarHeight: 36,
  zIndex: 200,
});

const emit = defineEmits<{ sideMouseLeave: []; toggleSidebar: [] }>();
const sidebarCollapse = defineModel<boolean>('sidebarCollapse');
const sidebarExtraVisible = defineModel<boolean>('sidebarExtraVisible');
const sidebarExtraCollapse = defineModel<boolean>('sidebarExtraCollapse');
const sidebarExpandOnHover = defineModel<boolean>('sidebarExpandOnHover');
const sidebarEnable = defineModel<boolean>('sidebarEnable', { default: true });

const {
  arrivedState,
  directions,
  isScrolling,
  y: scrollY,
} = useScroll(document);

const { y: mouseY } = useMouse({ type: 'client' });

// Whether SIDE is in the hover state expansion menu
const sidebarExpandOnHovering = ref(false);
const headerIsHidden = ref(false);

const realLayout = computed(() =>
  props.isMobile ? 'sidebar-nav' : props.layout,
);

/**
 * Whether the full screen is displayed in Content, there is no need for side, bottom, top, TAB area
 */
const fullContent = computed(() => realLayout.value === 'full-content');

/**
 * Whether the side mixed mode
 */
const isSidebarMixedNav = computed(
  () => realLayout.value === 'sidebar-mixed-nav',
);

/**
 * Whether it is a head navigation mode
 */
const isHeaderNav = computed(() => realLayout.value === 'header-nav');

/**
 * Whether it is a hybrid navigation mode
 */
const isMixedNav = computed(() => realLayout.value === 'mixed-nav');

/**
 * Whether the top bar is automatically hidden
 */
const isHeaderAutoMode = computed(() => props.headerMode === 'auto');

/**
 * Header area height
 */
const getHeaderHeight = computed(() => {
  const { headerHeight, headerHeightOffset } = props;

  // if (!headerVisible) {
  //   return 0;
  // }

  // When there is navigation at the top, increase 10
  const offset = isMixedNav.value || isHeaderNav.value ? headerHeightOffset : 0;

  return headerHeight + offset;
});

const headerWrapperHeight = computed(() => {
  let height = 0;
  if (props.headerVisible && !props.headerHidden) {
    height += getHeaderHeight.value;
  }
  if (props.tabbarEnable) {
    height += props.tabbarHeight;
  }
  return height;
});

const getSideCollapseWidth = computed(() => {
  const { sidebarCollapseShowTitle, sidebarMixedWidth, sideCollapseWidth } =
    props;

  return sidebarCollapseShowTitle || isSidebarMixedNav.value
    ? sidebarMixedWidth
    : sideCollapseWidth;
});

/**
 * Whether it can be visible to the side area of dynamic obtaining
 */
const sidebarEnableState = computed(() => {
  return !isHeaderNav.value && sidebarEnable.value;
});

/**
 * High level from the side area from the top
 */
const sidebarMarginTop = computed(() => {
  const { isMobile } = props;
  return isMixedNav.value && !isMobile ? getHeaderHeight.value : 0;
});

/**
 * Dynamic obtaining side width
 */
const getSidebarWidth = computed(() => {
  const { isMobile, sidebarHidden, sidebarMixedWidth, sidebarWidth } = props;
  let width = 0;

  if (sidebarHidden) {
    return width;
  }

  if (
    !sidebarEnableState.value ||
    (sidebarHidden && !isSidebarMixedNav.value && !isMixedNav.value)
  ) {
    return width;
  }

  if (isSidebarMixedNav.value && !isMobile) {
    width = sidebarMixedWidth;
  } else if (sidebarCollapse.value) {
    width = isMobile ? 0 : getSideCollapseWidth.value;
  } else {
    width = sidebarWidth;
  }
  return width;
});

/**
 * Get the width of the extended area
 */
const sidebarExtraWidth = computed(() => {
  const { sidebarExtraCollapsedWidth, sidebarWidth } = props;

  return sidebarExtraCollapse.value ? sidebarExtraCollapsedWidth : sidebarWidth;
});

/**
 * Whether the sidebar mode contains the mixed side
 */
const isSideMode = computed(() =>
  ['mixed-nav', 'sidebar-mixed-nav', 'sidebar-nav'].includes(realLayout.value),
);

const showSidebar = computed(() => {
  return isSideMode.value && sidebarEnable.value;
});

const sidebarFace = computed(() => {
  const { sidebarSemiDark, sidebarTheme } = props;
  const isDark = sidebarTheme === 'dark' || sidebarSemiDark;
  return {
    theme: isDark ? 'dark' : 'light',
  };
});

/**
 * Cover visibility
 */
const maskVisible = computed(() => !sidebarCollapse.value && props.isMobile);

/**
 * headerFixed值
 */
const headerFixed = computed(() => {
  const { headerMode } = props;
  return (
    isMixedNav.value ||
    headerMode === 'fixed' ||
    headerMode === 'auto-scroll' ||
    headerMode === 'auto'
  );
});

const mainStyle = computed(() => {
  let width = '100%';
  let sidebarAndExtraWidth = 'unset';
  if (
    headerFixed.value &&
    realLayout.value !== 'header-nav' &&
    realLayout.value !== 'mixed-nav' &&
    showSidebar.value &&
    !props.isMobile
  ) {
    // Effective in FIXED mode
    const isSideNavEffective =
      isSidebarMixedNav.value &&
      sidebarExpandOnHover.value &&
      sidebarExtraVisible.value;

    if (isSideNavEffective) {
      const sideCollapseWidth = sidebarCollapse.value
        ? getSideCollapseWidth.value
        : props.sidebarMixedWidth;
      const sideWidth = sidebarExtraCollapse.value
        ? getSideCollapseWidth.value
        : props.sidebarWidth;

      // 100% - Side menu mixed width - menu width
      sidebarAndExtraWidth = `${sideCollapseWidth + sideWidth}px`;
      width = `calc(100% - ${sidebarAndExtraWidth})`;
    } else {
      sidebarAndExtraWidth =
        sidebarExpandOnHovering.value && !sidebarExpandOnHover.value
          ? `${getSideCollapseWidth.value}px`
          : `${getSidebarWidth.value}px`;
      width = `calc(100% - ${sidebarAndExtraWidth})`;
    }
  }
  return {
    sidebarAndExtraWidth,
    width,
  };
});

// Calculate the style of tabbar
const tabbarStyle = computed((): CSSProperties => {
  let width = '';
  let marginLeft = 0;

  // If it is not a mixed navigation, the width of Tabbar is 100%
  if (!isMixedNav.value) {
    width = '100%';
  } else if (sidebarEnable.value) {
    // When the mouse is on the sidebar, and the width of the sidebar when the sidebar is unfolded
    const onHoveringWidth = sidebarExpandOnHover.value
      ? props.sidebarWidth
      : getSideCollapseWidth.value;

    // Set MarginleFT and decide whether it is folded by the sidebar
    marginLeft = sidebarCollapse.value
      ? getSideCollapseWidth.value
      : onHoveringWidth;

    // Set the width of Tabbar, and the calculation method is 100% minus the width of the sidebar
    width = `calc(100% - ${sidebarCollapse.value ? getSidebarWidth.value : onHoveringWidth}px)`;
  } else {
    // By default, the width of tabbar is 100%
    width = '100%';
  }

  return {
    marginLeft: `${marginLeft}px`,
    width,
  };
});

const contentStyle = computed((): CSSProperties => {
  const fixed = headerFixed.value;

  const { footerEnable, footerFixed, footerHeight } = props;
  return {
    marginTop:
      fixed &&
      !fullContent.value &&
      !headerIsHidden.value &&
      (!isHeaderAutoMode.value || scrollY.value < headerWrapperHeight.value)
        ? `${headerWrapperHeight.value}px`
        : 0,
    paddingBottom: `${footerEnable && footerFixed ? footerHeight : 0}px`,
  };
});

const headerZIndex = computed(() => {
  const { zIndex } = props;
  const offset = isMixedNav.value ? 1 : 0;
  return zIndex + offset;
});

const headerWrapperStyle = computed((): CSSProperties => {
  const fixed = headerFixed.value;
  return {
    height: fullContent.value ? '0' : `${headerWrapperHeight.value}px`,
    left: isMixedNav.value ? 0 : mainStyle.value.sidebarAndExtraWidth,
    position: fixed ? 'fixed' : 'static',
    top:
      headerIsHidden.value || fullContent.value
        ? `-${headerWrapperHeight.value}px`
        : 0,
    width: mainStyle.value.width,
    'z-index': headerZIndex.value,
  };
});

/**
 * Sidebar Z-index
 */
const sidebarZIndex = computed(() => {
  const { isMobile, zIndex } = props;
  let offset = isMobile || isSideMode.value ? 1 : -1;

  if (isMixedNav.value) {
    offset += 1;
  }

  return zIndex + offset;
});

const footerWidth = computed(() => {
  if (!props.footerFixed) {
    return '100%';
  }

  return mainStyle.value.width;
});

const maskStyle = computed((): CSSProperties => {
  return { zIndex: props.zIndex };
});

const showHeaderToggleButton = computed(() => {
  return (
    props.headerToggleSidebarButton &&
    isSideMode.value &&
    !isSidebarMixedNav.value &&
    !isMixedNav.value &&
    !props.isMobile
  );
});

const showHeaderLogo = computed(() => {
  return !isSideMode.value || isMixedNav.value || props.isMobile;
});

watch(
  () => props.isMobile,
  (val) => {
    if (val) {
      sidebarCollapse.value = true;
    }
  },
  {
    immediate: true,
  },
);

{
  const mouseMove = () => {
    mouseY.value > headerWrapperHeight.value
      ? (headerIsHidden.value = true)
      : (headerIsHidden.value = false);
  };
  watch(
    [() => props.headerMode, () => mouseY.value],
    () => {
      if (!isHeaderAutoMode.value || isMixedNav.value || fullContent.value) {
        return;
      }
      headerIsHidden.value = true;
      mouseMove();
    },
    {
      immediate: true,
    },
  );
}

{
  const checkHeaderIsHidden = useThrottleFn((top, bottom, topArrived) => {
    if (scrollY.value < headerWrapperHeight.value) {
      headerIsHidden.value = false;
      return;
    }
    if (topArrived) {
      headerIsHidden.value = false;
      return;
    }

    if (top) {
      headerIsHidden.value = false;
    } else if (bottom) {
      headerIsHidden.value = true;
    }
  }, 300);

  watch(
    () => scrollY.value,
    () => {
      if (
        props.headerMode !== 'auto-scroll' ||
        isMixedNav.value ||
        fullContent.value
      ) {
        return;
      }
      if (isScrolling.value) {
        checkHeaderIsHidden(
          directions.top,
          directions.bottom,
          arrivedState.top,
        );
      }
    },
  );
}

function handleClickMask() {
  sidebarCollapse.value = true;
}

function handleOpenMenu() {
  sidebarCollapse.value = false;
}
</script>

<template>
  <div id="layout" class="relative flex min-h-full w-full">
    <slot name="preferences"></slot>
    <slot name="floating-groups"></slot>
    <LayoutSidebar
      v-if="sidebarEnableState"
      v-model:collapse="sidebarCollapse"
      v-model:expand-on-hover="sidebarExpandOnHover"
      v-model:expand-on-hovering="sidebarExpandOnHovering"
      v-model:extra-collapse="sidebarExtraCollapse"
      v-model:extra-visible="sidebarExtraVisible"
      :collapse-width="getSideCollapseWidth"
      :dom-visible="!isMobile"
      :extra-width="sidebarExtraWidth"
      :fixed-extra="sidebarExpandOnHover"
      :header-height="isMixedNav ? 0 : getHeaderHeight"
      :is-sidebar-mixed="isSidebarMixedNav"
      :margin-top="sidebarMarginTop"
      :mixed-width="sidebarMixedWidth"
      :show="showSidebar"
      :theme="sidebarFace.theme"
      :width="getSidebarWidth"
      :z-index="sidebarZIndex"
      @leave="() => emit('sideMouseLeave')"
    >
      <template v-if="isSideMode && !isMixedNav" #logo>
        <slot name="logo"></slot>
      </template>

      <template v-if="isSidebarMixedNav">
        <slot name="mixed-menu"></slot>
      </template>
      <template v-else>
        <slot name="menu"></slot>
      </template>

      <template #extra>
        <slot name="side-extra"></slot>
      </template>
      <template #extra-title>
        <slot name="side-extra-title"></slot>
      </template>
    </LayoutSidebar>

    <div
      class="flex flex-1 flex-col overflow-hidden transition-all duration-300 ease-in"
    >
      <div
        :style="headerWrapperStyle"
        class="overflow-hidden transition-all duration-200"
      >
        <LayoutHeader
          v-if="headerVisible"
          :full-width="!isSideMode"
          :height="getHeaderHeight"
          :is-mixed-nav="isMixedNav"
          :is-mobile="isMobile"
          :show="!fullContent && !headerHidden"
          :show-toggle-btn="showHeaderToggleButton"
          :sidebar-width="sidebarWidth"
          :width="mainStyle.width"
          :z-index="headerZIndex"
          @open-menu="handleOpenMenu"
          @toggle-sidebar="() => emit('toggleSidebar')"
        >
          <template v-if="showHeaderLogo" #logo>
            <slot name="logo"></slot>
          </template>
          <slot name="header"></slot>
        </LayoutHeader>

        <LayoutTabbar
          v-if="tabbarEnable"
          :height="tabbarHeight"
          :style="tabbarStyle"
        >
          <slot name="tabbar"></slot>
        </LayoutTabbar>
      </div>

      <!-- </div> -->
      <LayoutContent
        :content-compact="contentCompact"
        :content-compact-width="contentCompactWidth"
        :padding="contentPadding"
        :padding-bottom="contentPaddingBottom"
        :padding-left="contentPaddingLeft"
        :padding-right="contentPaddingRight"
        :padding-top="contentPaddingTop"
        :style="contentStyle"
        class="transition-[margin-top] duration-200"
      >
        <slot name="content"></slot>
      </LayoutContent>

      <LayoutFooter
        v-if="footerEnable"
        :fixed="footerFixed"
        :height="footerHeight"
        :show="!fullContent"
        :width="footerWidth"
        :z-index="zIndex"
      >
        <slot name="footer"></slot>
      </LayoutFooter>
    </div>
    <slot name="extra"></slot>
    <div
      v-if="maskVisible"
      :style="maskStyle"
      class="bg-overlay fixed left-0 top-0 h-full w-full transition-[background-color] duration-200"
      @click="handleClickMask"
    ></div>
  </div>
</template>
