<script setup lang="ts">
import { useDesign } from '@/hooks/web/useDesign'
import { useLocale } from '@/hooks/web/useLocale'
import { useLocaleStore } from '@/store/modules/locale'
import { propTypes } from '@/utils/propTypes'
import { ElDropdown, ElDropdownItem, ElDropdownMenu } from 'element-plus'
import { computed, unref } from 'vue'

const { getPrefixCls } = useDesign()

const prefixCls = getPrefixCls('locale-dropdown')

defineProps({
  color: propTypes.string.def('')
})

const localeStore = useLocaleStore()

const langMap = computed(() => localeStore.getLocaleMap)

const currentLang = computed(() => localeStore.getCurrentLocale)

const setLang = (lang: LocaleType) => {
  if (lang === unref(currentLang).lang) return
  // Need to reload the page to make the entire language initialize more
  // window.location.reload()
  localeStore.setCurrentLocale({
    lang
  })
  const { changeLocale } = useLocale()
  changeLocale(lang)
}
</script>

<template>
  <ElDropdown :class="prefixCls" trigger="click" @command="setLang">
    <Icon
      :size="18"
      icon="ion:language-sharp"
      class="cursor-pointer !p-0"
      :class="$attrs.class"
      :color="color"
    />
    <template #dropdown>
      <ElDropdownMenu>
        <ElDropdownItem v-for="item in langMap" :key="item.lang" :command="item.lang">
          {{ item.name }}
        </ElDropdownItem>
      </ElDropdownMenu>
    </template>
  </ElDropdown>
</template>
