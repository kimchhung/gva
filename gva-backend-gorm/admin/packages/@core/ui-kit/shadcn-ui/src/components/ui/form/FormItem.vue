<script lang="ts" setup>
import { type HTMLAttributes, provide } from 'vue';

import { cn } from '@vben-core/shared';

import { useId } from 'radix-vue';

import { FORM_ITEM_INJECTION_KEY } from './injectionKeys';

const props = defineProps<{
  class?: HTMLAttributes['class'];
}>();

const id = useId();
provide(FORM_ITEM_INJECTION_KEY, id);
</script>

<template>
  <div :class="cn('space-y-2', props.class)">
    <slot></slot>
  </div>
</template>
