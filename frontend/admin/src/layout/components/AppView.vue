<script setup lang="ts">
import { Footer } from '@/components/Footer'
import { useAppStore } from '@/store/modules/app'
import { useTagsViewStore } from '@/store/modules/tagsView'
import { computed } from 'vue'

const appStore = useAppStore()

const footer = computed(() => appStore.getFooter)

const tagsViewStore = useTagsViewStore()

const getCaches = computed((): string[] => {
  return tagsViewStore.getCachedViews
})

// router.afterEach((to, from) => {
//   const toDepth = to.path.split('/').length
//   const fromDepth = from.path.split('/').length
//   to.meta.transition = toDepth < fromDepth ? 'slide-right' : 'slide-left'
//   if (toDepth === fromDepth) {
//     to.meta.transition = 'fade'
//   }
// })
</script>

<template>
  <section
    :class="[
      'flex-1 p-[var(--app-content-padding)] w-[calc(100%-var(--app-content-padding)-var(--app-content-padding))] bg-[var(--app-content-bg-color)] dark:bg-[var(--el-bg-color)]'
    ]"
  >
    <router-view v-slot="{ Component, route }">
      <keep-alive :include="getCaches">
        <component :is="Component" :key="route.fullPath" />
      </keep-alive>
    </router-view>
  </section>
  <Footer v-if="footer" />
</template>
