import type { ComponentRecordType, GenerateMenuAndRoutesOptions } from '@gva/types';

import { generateAccessible } from '@gva/access';
import { preferences } from '@gva/preferences';

import { message } from 'ant-design-vue';

import { BasicLayout, IFrameView } from '#/layouts';
import { $t } from '#/locales';
import { menuToRoute } from '#/api/menu/tranform';

const forbiddenComponent = () => import('#/views/_core/fallback/forbidden.vue');

async function generateAccess(options: GenerateMenuAndRoutesOptions) {
  const pageMap: ComponentRecordType = import.meta.glob('../views/**/*.vue');

  const layoutMap: ComponentRecordType = {
    BasicLayout,
    IFrameView,
  };

  return await generateAccessible(preferences.app.accessMode, {
    ...options,
    fetchMenuListAsync: async () => {
      message.loading({
        content: `${$t('common.loadingMenu')}...`,
        duration: 1.5,
      });

      const [res, err] = await api.menu.enabledList();
      if (err) {
        return;
      }

      const list = menuToRoute(res.data);
      return list ?? [];
    },
    // You can specify no permissions jump 403 page 403
    forbiddenComponent,
    // if route.meta.menuVisibleWithForbidden = true
    layoutMap,
    pageMap,
  });
}

export { generateAccess };
