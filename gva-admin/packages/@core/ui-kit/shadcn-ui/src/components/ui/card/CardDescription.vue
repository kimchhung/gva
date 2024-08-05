<script setup lang="ts">
import type { HTMLAttributes } from 'vue';

import { cn } from '@gva-core/shared';

const props = defineProps<{
  class?: HTMLAttributes['class'];
}>();
</script>

<template>
  <p :class="cn('text-muted-foreground text-sm', props.class)">
    <slot></slot>
  </p>
</template>
