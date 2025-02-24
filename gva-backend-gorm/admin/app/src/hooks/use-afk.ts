import { watch } from 'vue';

import { useIdle } from '@vueuse/core';

const useAFK = (time: number, action: () => void) => {
  const { idle } = useIdle(time, { events: ['mousedown'] });

  watch(idle, () => {
    if (idle.value) {
      action();
    }
  });

  const lastCloseTimeKey = 'lastCloseTime';
  const setCloseTime = () => {
    localStorage.setItem(lastCloseTimeKey, Date.now().toString());
  };

  window.addEventListener('beforeunload', () => {
    setCloseTime();
  });

  const lastUnloadTime = localStorage.getItem(lastCloseTimeKey);
  if (lastUnloadTime) {
    const lastUnloadTimeNumber = Number(lastUnloadTime);
    const now = Date.now();
    const diff = now - lastUnloadTimeNumber;
    if (diff > time) {
      action();
    }
  }

  // handle incase when beforeunload not triggered
  setCloseTime();
  setInterval(
    () => {
      setCloseTime();
    },
    1000 * 60 * 5, // 5 minutes,
  );
};

export default useAFK;
