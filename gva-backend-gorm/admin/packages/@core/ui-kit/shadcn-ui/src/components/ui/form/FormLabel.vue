<script lang="ts" setup>
import type { LabelProps } from 'radix-vue';

import type { HTMLAttributes } from 'vue';

import { Label } from '@vben-core/shadcn-ui/components/ui/label';
import { cn } from '@vben-core/shared';

import { useFormField } from './useFormField';

const props = defineProps<{ class?: HTMLAttributes['class'] } & LabelProps>();

const { error, formItemId } = useFormField();
</script>

<template>
  <Label
    :class="cn(error && 'text-destructive', props.class)"
    :for="formItemId"
  >
    <slot></slot>
  </Label>
</template>
