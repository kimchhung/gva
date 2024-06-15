<script lang="tsx">
import { Breadcrumb } from '@/components/Breadcrumb'
import { Collapse } from '@/components/Collapse'
import { LocaleDropdown } from '@/components/LocaleDropdown'
import { Screenfull } from '@/components/Screenfull'
import { SizeDropdown } from '@/components/SizeDropdown'
import { UserInfo } from '@/components/UserInfo'
import { useDesign } from '@/hooks/web/useDesign'
import { useAppStore } from '@/store/modules/app'
import { computed, defineComponent } from 'vue'

const { getPrefixCls, variables } = useDesign()

const prefixCls = getPrefixCls('tool-header')

const appStore = useAppStore()

// Bread crumbs
const breadcrumb = computed(() => appStore.getBreadcrumb)

// Folding icon
const hamburger = computed(() => appStore.getHamburger)

// Full -screen icon
const screenfull = computed(() => appStore.getScreenfull)

// Size icon
const size = computed(() => appStore.getSize)

// layout
const layout = computed(() => appStore.getLayout)

// Multi -language icon
const locale = computed(() => appStore.getLocale)

export default defineComponent({
  name: 'ToolHeader',
  setup() {
    return () => (
      <div
        id={`${variables.namespace}-tool-header`}
        class={[
          prefixCls,
          'h-[var(--top-tool-height)] relative px-[var(--top-tool-p-x)] flex items-center justify-between'
        ]}
      >
        {layout.value !== 'top' ? (
          <div class="h-full flex items-center">
            {hamburger.value && layout.value !== 'cutMenu' ? (
              <Collapse class="custom-hover" color="var(--top-header-text-color)"></Collapse>
            ) : undefined}
            {breadcrumb.value ? <Breadcrumb class="<md:hidden"></Breadcrumb> : undefined}
          </div>
        ) : undefined}
        <div class="h-full flex items-center">
          {screenfull.value ? (
            <Screenfull class="custom-hover" color="var(--top-header-text-color)"></Screenfull>
          ) : undefined}
          {size.value ? (
            <SizeDropdown class="custom-hover" color="var(--top-header-text-color)"></SizeDropdown>
          ) : undefined}
          {locale.value ? (
            <LocaleDropdown
              class="custom-hover"
              color="var(--top-header-text-color)"
            ></LocaleDropdown>
          ) : undefined}
          <UserInfo></UserInfo>
        </div>
      </div>
    )
  }
})
</script>

<style lang="less" scoped>
@prefix-cls: ~'@{namespace}-tool-header';

.@{prefix-cls} {
  transition: left var(--transition-time-02);
}
</style>
