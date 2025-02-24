<script setup lang="ts">
interface Props {
  title?: string;
}

defineOptions({
  name: 'Block',
});

withDefaults(defineProps<Props>(), {
  title: '',
});
</script>

<template>
  <div class="flex flex-col py-2 first:pt-0">
    <h3 class="mb-3 font-medium leading-none tracking-tight">
      {{ title }}
    </h3>
    <slot></slot>
  </div>
</template>
