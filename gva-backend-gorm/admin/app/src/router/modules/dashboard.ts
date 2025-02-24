import { $t } from '@vben/locales';

import { BasicLayout } from '#/layouts';

import { defineRoute } from '../util';

export default defineRoute([
  {
    component: BasicLayout,
    meta: {
      keepAlive: true,
      icon: 'lucide:layout-dashboard',
      order: -1,
      title: $t('page.dashboard.title'),
    },
    name: 'Dashboard',
    path: '/dashboard',
    children: [
      {
        component: () => import('#/views/dashboard/analytics/index.vue'),
        meta: {
          keepAlive: true,
          affixTab: true,
          icon: 'lucide:area-chart',
          title: $t('page.dashboard.analytics'),
        },
        name: 'Analytics',
        path: '/analytics',
      },
    ],
  },
]);
