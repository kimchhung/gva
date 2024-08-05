<!--
Access Control Component for Fine-Grained Access Control.
TODO: Can expand more complete features:
1. Support multiple permissions code, as long as one permissions code is satisfied or multiple permissions codes are satisfied
2. Support multiple characters, as long as one character is satisfied or multiple characters are satisfied
3. Support the logic of the judgment of the custom right code and the character
-->
<script lang="ts" setup>
import { computed } from 'vue';

import { useAccess } from './use-access';

interface Props {
  /**
   * Specified codes is visible
   * @default []
   */
  codes?: string[];

  /**
   * How can I control the component, if so role，Then pass the character, if so code，Then pass the entry authority code
   * @default 'role'
   */
  type?: 'code' | 'role';
}

defineOptions({
  name: 'AccessControl',
});

const props = withDefaults(defineProps<Props>(), {
  codes: () => [],
  type: 'role',
});

const { hasAccessByCodes, hasAccessByRoles } = useAccess();

const hasAuth = computed(() => {
  const { codes, type } = props;
  return type === 'role' ? hasAccessByRoles(codes) : hasAccessByCodes(codes);
});
</script>

<template>
  <slot v-if="!codes"></slot>
  <slot v-else-if="hasAuth"></slot>
</template>
