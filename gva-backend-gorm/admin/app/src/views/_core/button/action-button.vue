<script setup lang="tsx">
import { computed, defineProps } from 'vue';

import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import {
  Button,
  type ButtonProps,
  Flex,
  Popconfirm,
  Popover,
  type PopoverProps,
  type TooltipProps,
} from 'ant-design-vue';
import { omit } from 'lodash';

export type MenuItemProps = {
  icon?: string;
  key?: string;
  value?: any;
  vPermissions?: string[];
} & ButtonProps;

export type ActionButtonProps = {
  actionType?:
    | 'confirmation'
    | 'create'
    | 'default'
    | 'delete'
    | 'detail'
    | 'edit'
    | 'menu'
    | 'more'
    | 'view';
  icon?: string;
  menuItems?: MenuItemProps[];
  menuProps?: PopoverProps;
  slot?: string;
  tooltip?: boolean;
  tooltipProps?: TooltipProps;
  value?: any;
} & ButtonProps;

const props = defineProps<ActionButtonProps>();

const iconName = computed(() => {
  if (props.icon) return props.icon;
  switch (props.actionType) {
    case 'create': {
      return 'streamline:add-1-solid';
    }
    case 'delete': {
      return 'lucide:trash-2';
    }
    case 'detail': {
      return 'ooui:view-details-ltr';
    }
    case 'edit': {
      return 'lucide:edit';
    }
    case 'more': {
      return 'lucide:ellipsis';
    }
    case 'view': {
      return 'ant-design:eye-outlined';
    }
    default: {
      return '';
    }
  }
});

const actionValue = computed(() => {
  if (props.value) return props.value;
  return $t(`common.${props.actionType || 'view'}`);
});

const isActionTypeDelete = computed(() => props.actionType === 'delete');
</script>

<template>
  <div>
    <Button
      v-if="actionType === 'view'"
      v-bind="props"
      class="flex items-center gap-2 px-2"
      type="primary"
    >
      <template #icon>
        <VbenIcon :icon="iconName" />
      </template>
      {{ actionValue }}
    </Button>
    <Button
      v-else-if="
        actionType === 'edit' ||
        actionType === 'detail' ||
        actionType === 'create'
      "
      v-bind="props"
      class="flex items-center gap-2 px-2"
      type="primary"
    >
      <template #icon>
        <VbenIcon :icon="iconName" />
      </template>
      {{ actionValue }}
    </Button>
    <template v-else-if="isActionTypeDelete || actionType === 'confirmation'">
      <Popconfirm
        :disabled="props.disabled"
        :ok-text="$t('common.confirm')"
        :title="$t('common.confirmation')"
        @confirm="
          (e) => {
            return typeof onClick === 'function' ? onClick(e) : null;
          }
        "
      >
        <Button
          v-bind="omit(props, 'onClick', isActionTypeDelete ? '' : 'icon')"
          :danger="isActionTypeDelete"
          :type="isActionTypeDelete ? 'primary' : 'default'"
          class="flex items-center gap-2 px-2"
        >
          <template #icon>
            <VbenIcon :icon="iconName" />
          </template>
          {{ actionValue }}
        </Button>
      </Popconfirm>
    </template>
    <template v-else-if="actionType === 'menu'">
      <Popover placement="bottomRight" v-bind="menuProps">
        <Button
          v-bind="omit(props, 'icon')"
          class="flex items-center gap-2 px-2"
        >
          <template #icon>
            <VbenIcon :icon="iconName" />
          </template>
          {{ actionValue }}
        </Button>
        <template #content>
          <Flex gap="small" vertical>
            <Button
              v-for="(
                {
                  vPermissions,
                  value: menuItemValue,
                  icon: menuItemIcon,
                  ...buttonProps
                },
                index
              ) in menuItems"
              v-bind="omit(buttonProps, 'icon')"
              :key="index"
              class="flex items-center gap-2 px-2"
              v-permissions="vPermissions"
            >
              <template #icon>
                <VbenIcon :icon="menuItemIcon" />
              </template>
              {{ menuItemValue }}
            </Button>
          </Flex>
        </template>
      </Popover>
    </template>
    <Button
      v-else
      v-bind="omit(props, 'icon')"
      class="flex w-full items-center gap-2 px-2"
    >
      <template #icon>
        <VbenIcon :icon="iconName" />
      </template>

      {{ actionValue }}
    </Button>
  </div>
</template>
