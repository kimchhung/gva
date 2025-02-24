<script setup lang="tsx">
import { computed, onBeforeMount, type PropType } from 'vue';

import { useAccess } from '@vben/access';
import { VbenIcon } from '@vben-core/shadcn-ui';

import { useMediaQuery } from '@vueuse/core';
import { Button, Flex, Popover } from 'ant-design-vue';

import ActionButton, { type ActionButtonProps } from './action-button.vue';

export type GroupButtonProp = {
  slot?: string;
  vPermissions?: string[];
} & ActionButtonProps;

export type GroupButtonProps = {
  buttons: GroupButtonProp[];
  maxDisplay?: number;
  popoverProp?: PropType<typeof Popover>;
};

const props = defineProps<GroupButtonProps>();
const emit = defineEmits(['noPermissions']);

const isMobileScreen = useMediaQuery('(max-width: 820px)');

const maxDisplay = computed(() =>
  isMobileScreen.value ? 0 : (props.maxDisplay ?? 3),
);

const { hasAccessByPermissions } = useAccess();
const visibleButtons = computed(() =>
  props.buttons.filter(({ vPermissions }) =>
    vPermissions ? hasAccessByPermissions(vPermissions) : true,
  ),
);

const displayButtons = computed(() =>
  visibleButtons.value.filter((_, i) => i < maxDisplay.value),
);

const popoverButtons = computed(() =>
  visibleButtons.value.filter((_, i) => i >= maxDisplay.value),
);

const canUserAccessAnyPopoverButton = computed(() =>
  popoverButtons.value.some(({ vPermissions }) =>
    vPermissions ? hasAccessByPermissions(vPermissions) : true,
  ),
);

onBeforeMount(() => {
  const hasNoAccessEveryButton = props.buttons.every(({ vPermissions }) => {
    if (vPermissions) {
      return !hasAccessByPermissions(vPermissions);
    }
    return false;
  });

  if (hasNoAccessEveryButton) {
    emit('noPermissions');
  }
});
</script>

<template>
  <Flex
    :style="{
      width: 'fit-content',
      display: isMobileScreen ? 'inline-block' : 'flex',
    }"
    class="justify-center"
    gap="small"
    horizontal
  >
    <ActionButton
      v-for="({ vPermissions, ...buttonProps }, index) in displayButtons"
      v-permissions="vPermissions"
      v-bind="buttonProps"
      :key="index"
    />
    <template v-if="popoverButtons.length > 0 && canUserAccessAnyPopoverButton">
      <Popover placement="bottomRight" v-bind="props.popoverProp ?? {}">
        <Button class="px-2.5"><VbenIcon icon="lucide:ellipsis" /></Button>
        <template #content>
          <Flex gap="small" vertical>
            <ActionButton
              v-for="(
                { vPermissions, ...buttonProps }, index
              ) in popoverButtons"
              class="w-full"
              v-bind="buttonProps"
              :key="index"
              v-permissions="vPermissions"
            />
          </Flex>
        </template>
      </Popover>
    </template>
  </Flex>
</template>
