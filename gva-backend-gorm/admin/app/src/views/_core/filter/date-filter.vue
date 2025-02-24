<script setup lang="ts">
import type { PresetDate } from 'ant-design-vue/es/vc-picker/interface';
import type { Dayjs } from 'dayjs';

import { computed } from 'vue';

import { Search } from '@vben/icons';
import { $t } from '@vben/locales';

import { Button, Col, RangePicker } from 'ant-design-vue';

import { dateUtil, disableFutureDates } from '#/utils/helper/date-util';

defineOptions({ name: 'DateFilter' });
const props = defineProps<{ onSearch: () => void; value: [Dayjs, Dayjs] }>();

const emit = defineEmits(['update:value']);
const selectedDate = computed<[Dayjs, Dayjs]>({
  get() {
    return props.value;
  },
  set(value) {
    emit('update:value', value);
  },
});

const rangePresets = computed<PresetDate<[Dayjs, Dayjs]>[]>(() => [
  {
    label: $t('common.today'),
    value: [dateUtil().startOf('d'), dateUtil()],
  },
  {
    label: $t('common.yesterday'),
    value: [
      dateUtil().add(-1, 'd').startOf('d'),
      dateUtil().add(-1, 'd').endOf('d'),
    ],
  },
  {
    label: $t('common.last7Days'),
    value: [dateUtil().add(-7, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last14Days'),
    value: [dateUtil().add(-14, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last30Days'),
    value: [dateUtil().add(-30, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last90Days'),
    value: [dateUtil().add(-90, 'd').startOf('d'), dateUtil()],
  },
]);
</script>

<template>
  <div>
    <Col class="space-x-2">
      <RangePicker
        v-model:value="selectedDate"
        :disabled-date="disableFutureDates"
        :presets="rangePresets"
        format="YYYY-MM-DD HH:mm:ss"
        show-time
        style="width: 368px"
      />
      <Button style="width: 100px" type="primary" @click="onSearch">
        <template #icon><Search class="anticon size-4" /></template>
        {{ $t('common.search') }}
      </Button>
    </Col>
  </div>
</template>
