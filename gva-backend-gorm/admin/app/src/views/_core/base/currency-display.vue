<script setup lang="ts">
import { VbenIcon } from '@vben-core/shadcn-ui';

import { Tag } from 'ant-design-vue';

import { UserCurrencyName } from '#/api/user/types';

defineProps<{
  record: {
    amount: number;
    currencyNameId: UserCurrencyName;
  };
}>();
</script>

<template>
  <p class="inline-flex items-center">
    <span
      :class="record.amount > 0 ? 'text-green-500' : 'text-red-500'"
      class="font-semibold"
    >
      {{ `${record.amount > 0 ? '+' : ''}${record.amount}` }}
    </span>
    <template v-if="record.currencyNameId === UserCurrencyName.Gold">
      <VbenIcon
        :style="{ color: '#FF8800' }"
        class="ml-1 size-5"
        icon="ant-design:pay-circle-outlined"
      />
    </template>
    <template v-if="record.currencyNameId === UserCurrencyName.Exp">
      <VbenIcon
        :style="{ color: '#00ccbb' }"
        class="ml-1 size-5"
        icon="ant-design:crown-outlined"
      />
    </template>
    <template v-if="record.currencyNameId === UserCurrencyName.FreeChance">
      <Tag class="ml-2" color="orange">{{ $t('common.chance') }}</Tag>
    </template>
  </p>
</template>
