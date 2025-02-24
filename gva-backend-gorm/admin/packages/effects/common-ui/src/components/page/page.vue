<script setup lang="ts">
import { Badge, VbenIcon } from '@vben-core/shadcn-ui';

interface Props {
  title?: string;
  description?: string;
  contentClass?: string;
  showFooter?: boolean;
  fullHeight?: boolean;
  noScroll?: boolean;
  onBack?: () => void;
  badgeText?: string;
}

defineOptions({
  name: 'Page',
});

const props = withDefaults(defineProps<Props>(), {
  fullHeight: true,
  title: '',
  description: '',
  contentClass: '',
  onBack: undefined,
  badgeText: '',
});
</script>

<template>
  <div
    id="page"
    :class="{ 'flex flex-col': fullHeight }"
    class="relative h-full"
  >
    <div
      v-if="
        $slots.title ||
        title ||
        description ||
        $slots.description ||
        $slots.appendHeader
      "
      class="bg-card z-10 p-4 shadow-sm"
    >
      <div class="flex w-full flex-wrap justify-between gap-y-3">
        <slot name="title">
          <div
            v-if="title"
            class="flex items-center justify-between space-x-3 text-lg font-semibold"
          >
            <button v-if="onBack" @click="onBack">
              <VbenIcon
                class="transition-all duration-300 hover:scale-125"
                icon="lucide:arrow-left"
              />
            </button>
            <div class="flex gap-2">
              <span>{{ title }}</span>
              <Badge v-if="props.badgeText" variant="secondary">
                {{ props.badgeText }}
              </Badge>
            </div>
          </div>
        </slot>

        <slot v-if="description" name="description">
          <p class="text-muted-foreground">
            {{ description }}
          </p>
        </slot>
        <slot name="appendHeader"></slot>
      </div>
    </div>
    <div
      :class="[
        {
          'h-0 flex-grow': fullHeight,
          'overflow-auto': props.noScroll,
        },
        contentClass,
      ]"
      class="m-4"
    >
      <div
        :class="[contentClass]"
        class="bg-card flex h-full flex-col rounded-lg p-3"
      >
        <slot></slot>
      </div>
    </div>

    <div
      v-if="props.showFooter"
      class="bg-card align-center shadow-float z-10 flex items-center px-4 pb-2"
    >
      <slot name="footer"></slot>
    </div>
  </div>
</template>
