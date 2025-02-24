import type {
  ComponentRecordType,
  GenerateMenuAndRoutesOptions,
} from '@vben/types';

import { generateAccessible } from '@vben/access';
import { preferences } from '@vben/preferences';

import { BasicLayout, IFrameView } from '#/layouts';
import { mapConfigTree } from '#/utils/helper/configuration';

const forbiddenComponent = () => import('#/views/_core/fallback/forbidden.vue');

function convertMapRecursive(
  data: any[],
  transform: (item: any) => any,
): any[] {
  return data.map((item) => ({
    ...transform(item),
    children: item.children
      ? convertMapRecursive(item.children, transform)
      : [],
  }));
}

async function generateAccess(options: GenerateMenuAndRoutesOptions) {
  const pageMap: ComponentRecordType = import.meta.glob('../views/**/*.vue');

  const layoutMap: ComponentRecordType = {
    BasicLayout,
    IFrameView,
  };

  return await generateAccessible(preferences.app.accessMode, {
    ...options,
    fetchMenuListAsync: async () => {
      const [res, err] = await api.getDocs();
      if (err) {
        return [];
      }

      const data = mapConfigTree(
        res.data?.flatMap((item) => [item, ...(item.allChildren || [])]) ?? [],
      );

      const routes = data.map((item) => {
        return {
          component: 'BasicLayout',
          path: `/${item.key}`,
          name: item.key,
          meta: {
            title: 'menu.document.title',
            icon: item.metadata?.icon,
          },
          children: convertMapRecursive(item.children, (child) => ({
            component: 'IFrameView',
            meta: {
              iframeSrc: child.value || undefined,
              icon: child.metadata?.icon,
              title: child.key,
              titleEn: child.metadata?.labelEn,
              titleZh: child.metadata?.labelZh,
            },
            name: `${item.key}_${child.key}`,
            path: child.key,
          })),
        };
      });

      return routes;
    },
    // You can specify no permissions jump 403 page 403
    forbiddenComponent,
    // if route.meta.menuVisibleWithForbidden = true
    layoutMap,
    pageMap,
  });
}

export { generateAccess };
