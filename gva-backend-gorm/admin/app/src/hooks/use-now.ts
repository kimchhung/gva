import type { IntervalHandle } from '@vben/types';

import { reactive, toRefs } from 'vue';

import { tryOnMounted, tryOnUnmounted } from '@vueuse/core';

import { dateUtil } from '#/utils/helper/date-util';

export const useNow = (immediate = true) => {
  let timer: IntervalHandle;

  const state = reactive({
    day: 0,
    hour: '',
    meridiem: '',
    minute: '',
    month: 0,
    second: 0,
    week: '',
    year: 0,
  });

  const update = () => {
    const now = dateUtil();
    const h = now.format('HH');
    const m = now.format('mm');
    const s = now.get('s');

    state.year = now.get('y');
    state.month = now.get('M') + 1;
    state.week = `星期${['日', '一', '二', '三', '四', '五', '六'][now.day()]}`;
    state.day = now.get('date');
    state.hour = h;
    state.minute = m;
    state.second = s;
    state.meridiem = now.format('A');
  };

  function start() {
    update();
    clearInterval(timer);
    timer = setInterval(() => update(), 1000);
  }

  function stop() {
    clearInterval(timer);
  }

  tryOnMounted(() => {
    immediate && start();
  });

  tryOnUnmounted(() => {
    stop();
  });

  return {
    ...toRefs(state),
    start,
    stop,
  };
};
