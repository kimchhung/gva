import { onMounted } from 'vue';

import { dateUtil } from '#/utils/helper/date-util';

import { api } from '../api';

let ___serverDiffOffset = 0;

export const useTime = () => {
  const localTime = () => {
    return dateUtil();
  };

  // server time
  const now = () => {
    return dateUtil().add(___serverDiffOffset);
  };

  // want to check current localTime
  now.localTime = () => {
    dateUtil(now()).subtract(___serverDiffOffset);
  };

  onMounted(async () => {
    if (___serverDiffOffset !== 0) {
      return;
    }

    const syncTime = async () => {
      const [res, error] = await api().now();
      if (error) console.error(error);

      const serverTime = dateUtil(res?.data);
      ___serverDiffOffset = serverTime
        .subtract(localTime().millisecond(), 'millisecond')
        .millisecond();
    };

    syncTime();
    setInterval(syncTime, 15_000);
  });

  return { localTime, now };
};
