<script setup lang="ts">
import { computed } from 'vue';

import {
  SvgAPIIcon,
  SvgCakeIcon,
  SvgChartIcon,
  SvgUserIcon,
} from '@vben/icons';
import { $t } from '@vben/locales';
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  Spinner,
  VbenCountToAnimator,
  VbenIcon,
} from '@vben-core/shadcn-ui';

import { type AnalysisOverviewItem, AnalysisOverviewItemKey } from '../typing';

interface Props {
  // eslint-disable-next-line no-unused-vars
  items: { [key in AnalysisOverviewItemKey]: number };
  isLoading?: boolean;
}

defineOptions({
  name: 'AnalysisOverview',
});

defineProps<Props>();

const overviewItems = computed<AnalysisOverviewItem[]>(() => [
  {
    icon: SvgAPIIcon,
    title: $t('page.dashboard.analyse.totalRequest'),
    key: AnalysisOverviewItemKey.TOTAL_REQUEST,
    color: '#1F85DE',
  },
  {
    icon: SvgCakeIcon,
    title: $t('page.dashboard.analyse.requestSuccessRate'),
    suffix: '%',
    key: AnalysisOverviewItemKey.REQUEST_SUCCESS_RATE,
    color: '#6CC070',
  },
  {
    icon: SvgUserIcon,
    title: $t('page.dashboard.analyse.totalVisit'),
    key: AnalysisOverviewItemKey.TOTAL_VISIT,
    color: '#e458ff',
  },
  {
    icon: SvgChartIcon,
    title: $t('page.dashboard.analyse.totalUserGain'),
    prefix: '+',
    key: AnalysisOverviewItemKey.TOTAL_USER_GAIN,
    color: '#ffc798',
  },
]);
</script>

<template>
  <div class="grid gap-3 md:grid-cols-2 lg:grid-cols-4">
    <template v-for="item in overviewItems" :key="item.title">
      <Card
        :style="{ background: `${item.color}70`, borderColor: item.color }"
        :title="item.title"
        class="relative flex w-full flex-col overflow-hidden"
      >
        <CardHeader>
          <CardTitle class="text-lg">{{ item.title }}</CardTitle>
        </CardHeader>

        <CardContent class="mt-auto flex items-center justify-between">
          <VbenCountToAnimator
            :end-val="items[item.key]"
            :prefix="item.prefix"
            :start-val="0"
            :suffix="item.suffix"
            class="text-4xl font-semibold"
          />
          <VbenIcon :icon="item.icon" class="size-20 flex-shrink-0" />
        </CardContent>
        <Spinner :spinning="isLoading" class="bg-white/5" />
      </Card>
    </template>
  </div>
</template>
