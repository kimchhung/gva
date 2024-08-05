<script setup lang="ts">
import type { HTMLAttributes } from 'vue';

import { cn } from '@gva-core/shared';

const props = defineProps<{ class?: HTMLAttributes['class'] }>();
</script>

<template>
  <div :class="cn('flex flex-col text-center sm:text-left', props.class)">
    <slot></slot>
  </div>
</template>
