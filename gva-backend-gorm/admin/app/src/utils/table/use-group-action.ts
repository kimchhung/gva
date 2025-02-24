import { computed, ref } from 'vue';

import { useElementSize } from '@vueuse/core';

export const useGroupActionWidth = () => {
  const actionGroupRef = ref<HTMLElement | null>(null);
  const { width: actionGroupWidth } = useElementSize(actionGroupRef);

  const actionWidth = computed(() => {
    const width = actionGroupWidth.value ? actionGroupWidth.value + 20 : 100;
    return Math.max(width, 100);
  });

  return { actionGroupRef, actionWidth };
};
