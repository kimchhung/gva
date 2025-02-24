<script setup lang="ts">
import { computed, useModel } from 'vue';

import { useI18n } from '@vben/locales';

import { Select, type SelectProps } from 'ant-design-vue';

const props = defineProps<SelectProps>();

const { messages, locale } = useI18n();

const value = useModel(props, 'value');

const listTimeZones = computed(() => {
  return Object.entries(messages.value?.[locale.value]?.TIMEZONE ?? {}).map(
    ([key, value]) => ({
      label: value,
      value: key,
    }),
  );
});
</script>

<template>
  <Select
    v-model:value="value"
    :dropdown-match-select-width="false"
    :options="listTimeZones"
    option-filter-prop="label"
    show-search
  />
</template>
