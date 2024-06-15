<script setup lang="ts">
import { ThemeSwitch } from '@/components/ThemeSwitch'
import { useDesign } from '@/hooks/web/useDesign'
import { useI18n } from '@/hooks/web/useI18n'
import { useStorage } from '@/hooks/web/useStorage'
import { useAppStore } from '@/store/modules/app'
import { useServerTime } from '@/store/modules/time'
import { getCssVar, setCssVar, trim } from '@/utils'
import { useClipboard, useCssVar } from '@vueuse/core'
import { ElDivider, ElDrawer, ElMessage } from 'element-plus'
import { storeToRefs } from 'pinia'
import { computed, ref, unref } from 'vue'
import ColorRadioPicker from './components/ColorRadioPicker.vue'
import InterfaceDisplay from './components/InterfaceDisplay.vue'
import LayoutRadioPicker from './components/LayoutRadioPicker.vue'

const { clear: storageClear } = useStorage('localStorage')

const { getPrefixCls } = useDesign()

const prefixCls = getPrefixCls('setting')

const appStore = useAppStore()

const { t } = useI18n()

const drawer = ref(false)

// Theme -related
const systemTheme = ref(appStore.getTheme.elColorPrimary)

const setSystemTheme = (color: string) => {
  setCssVar('--el-color-primary', color)
  appStore.setTheme({ elColorPrimary: color })
  const leftMenuBgColor = useCssVar('--left-menu-bg-color', document.documentElement)
  setMenuTheme(trim(unref(leftMenuBgColor)))
}

const borderTheme = ref(appStore.getTheme.elBorderColor || '')

const setBorderTheme = (color: string) => {
  setCssVar('--el-border-color', color)
  appStore.setTheme({ elBorderColor: color })
}

// Head -themed theme
const headerTheme = ref(appStore.getTheme.topHeaderBgColor || '')

const setHeaderTheme = (color: string) => {
  appStore.setHeaderTheme(color)
}

// Menu theme related
const menuTheme = ref(appStore.getTheme.leftMenuBgColor || '')

const setMenuTheme = (color: string) => {
  appStore.setMenuTheme(color)
}

//Surveillance layout changes, reset some theme colors
// watch(
//   () => layout.value,
//   (n) => {
//     if (n === 'top' && !appStore.getIsDark) {
//       headerTheme.value = '#fff'
//       setHeaderTheme('#fff')
//     } else {
//       setMenuTheme(unref(menuTheme))
//     }
//   }
// )

// 拷贝
const copyConfig = async () => {
  const { copy, copied, isSupported } = useClipboard({
    source: `
      // Bread crumbs
      breadcrumb: ${appStore.getBreadcrumb},
      // Bread crumb icon
      breadcrumbIcon: ${appStore.getBreadcrumbIcon},
      // Folding icon
      hamburger: ${appStore.getHamburger},
      // Full -screen icon
      screenfull: ${appStore.getScreenfull},
      // Size icon
      size: ${appStore.getSize},
      // Multi -language icon
      locale: ${appStore.getLocale},
      // Bookmark page
      tagsView: ${appStore.getTagsView},
      // Tab icon
      getTagsViewIcon: ${appStore.getTagsViewIcon},
      // logo
      logo: ${appStore.getLogo},
      // Menukin piano
      uniqueOpened: ${appStore.getUniqueOpened},
      // Fixed header
      fixedHeader: ${appStore.getFixedHeader},
      // footer
      footer: ${appStore.getFooter},
      // Gray mode
      greyMode: ${appStore.getGreyMode},
      // layout布局
      layout: '${appStore.getLayout}',
      // Dark mode
      isDark: ${appStore.getIsDark},
      // Component size
      currentSize: '${appStore.getCurrentSize}',
      // Theme -related
      theme: {
        // Theme border color
        elBorderColor: '${appStore.getTheme.elBorderColor}',
        // Theme color
        elColorPrimary: '${appStore.getTheme.elColorPrimary}',
        // The color of the left menu border
        leftMenuBorderColor: '${appStore.getTheme.leftMenuBorderColor}',
        // The background color of the left menu
        leftMenuBgColor: '${appStore.getTheme.leftMenuBgColor}',
        // The left menu light background color
        leftMenuBgLightColor: '${appStore.getTheme.leftMenuBgLightColor}',
        // Left menu selected background color
        leftMenuBgActiveColor: '${appStore.getTheme.leftMenuBgActiveColor}',
        // The left menu put away the selected background color
        leftMenuCollapseBgActiveColor: '${appStore.getTheme.leftMenuCollapseBgActiveColor}',
        // The font color of the left menu
        leftMenuTextColor: '${appStore.getTheme.leftMenuTextColor}',
        // Select the font color of the left menu
        leftMenuTextActiveColor: '${appStore.getTheme.leftMenuTextActiveColor}',
        // logo font color
        logoTitleTextColor: '${appStore.getTheme.logoTitleTextColor}',
        // LOGO border color
        logoBorderColor: '${appStore.getTheme.logoBorderColor}',
        // Head background color
        topHeaderBgColor: '${appStore.getTheme.topHeaderBgColor}',
        // Head font color
        topHeaderTextColor: '${appStore.getTheme.topHeaderTextColor}',
        // Head suspension color
        topHeaderHoverColor: '${appStore.getTheme.topHeaderHoverColor}',
        // Head border color
        topToolBorderColor: '${appStore.getTheme.topToolBorderColor}'
      }
    `,
    legacy: true
  })
  if (!isSupported) {
    ElMessage.error(t('setting.copyFailed'))
  } else {
    await copy()
    if (unref(copied)) {
      ElMessage.success(t('setting.copySuccess'))
    }
  }
}

// Empty the cache
const clear = () => {
  storageClear()
  window.location.reload()
}

const themeChange = () => {
  const color = getCssVar('--el-bg-color')
  setMenuTheme(color)
  setHeaderTheme(color)
}

const { currentTime } = storeToRefs(useServerTime())
const time = computed(() => currentTime?.value?.format('HH:mm:ss Z'))
</script>

<template>
  <div
    :class="prefixCls"
    class="fixed top-[45%] right-0 w-40px h-40px flex items-center justify-center bg-[var(--el-color-primary)] cursor-pointer z-10"
    @click="drawer = true"
  >
    <Icon icon="ant-design:setting-outlined" color="#fff" />
  </div>

  <ElDrawer v-model="drawer" direction="rtl" size="350px" :z-index="4000">
    <template #header>
      <span class="text-16px font-700">{{ t('setting.projectSetting') }}</span>
    </template>

    <div class="text-center">
      <ElDivider>{{ t('common.time') }}</ElDivider>
      {{ time }}

      <ElDivider>{{ t('setting.theme') }}</ElDivider>
      <ThemeSwitch @change="themeChange" />

      <ElDivider>{{ t('setting.layout') }}</ElDivider>
      <LayoutRadioPicker />

      <ElDivider>{{ t('setting.systemTheme') }}</ElDivider>
      <ColorRadioPicker
        v-model="systemTheme"
        :schema="[
          '#409eff',
          '#14b8a6',
          '#06b6d4',
          '#6366f1',
          '#ec4899',
          '#0096c7',
          '#9c27b0',
          '#ff9800'
        ]"
        @change="setSystemTheme"
      />

      <ElDivider>{{ t('setting.headerTheme') }}</ElDivider>
      <ColorRadioPicker
        v-model="headerTheme"
        :schema="[
          '#fff',
          '#151515',
          '#5172dc',
          '#e74c3c',
          '#24292e',
          '#394664',
          '#009688',
          '#383f45'
        ]"
        @change="setHeaderTheme"
      />

      <ElDivider>{{ t('setting.menuTheme') }}</ElDivider>
      <ColorRadioPicker
        v-model="menuTheme"
        :schema="[
          '#fff',
          '#151515',
          '#5172dc',
          '#e74c3c',
          '#24292e',
          '#394664',
          '#009688',
          '#383f45'
        ]"
        @change="setMenuTheme"
      />

      <ElDivider>{{ t('setting.borderTheme') }}</ElDivider>
      <ColorRadioPicker
        v-model="borderTheme"
        :schema="[
          '#fff',
          '#f1f5f9',
          '#e2e8f0',
          '#cbd5e1',
          '#64748b',
          '#334155',
          '#1e293b',
          '#0f172a'
        ]"
        @change="setBorderTheme"
      />
    </div>

    <ElDivider>{{ t('setting.interfaceDisplay') }}</ElDivider>
    <InterfaceDisplay />

    <ElDivider />
    <div>
      <BaseButton type="primary" class="w-full" @click="copyConfig">{{
        t('setting.copy')
      }}</BaseButton>
    </div>
    <div class="mt-5px">
      <BaseButton type="danger" class="w-full" @click="clear">
        {{ t('setting.clearAndReset') }}
      </BaseButton>
    </div>
  </ElDrawer>
</template>

<style lang="less" scoped>
@prefix-cls: ~'@{namespace}-setting';

.@{prefix-cls} {
  border-radius: 6px 0 0 6px;
}
</style>
