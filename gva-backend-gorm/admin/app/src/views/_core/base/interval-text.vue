<script setup lang="ts">
defineProps<{
  intervals: {
    endTime: Date;
    everyInSec: number;
    startTime: Date;
  }[];
}>();
</script>

<template>
  <p>
    {{
      intervals
        ?.filter(Boolean)
        ?.map((interval) =>
          $t('common.intervalScheduleText', {
            second: interval?.everyInSec,
            startTime: interval?.startTime,
            endTime: interval?.endTime,
          }),
        )
        .join(',')
    }}
  </p>
</template>
