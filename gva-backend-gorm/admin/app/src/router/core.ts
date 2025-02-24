import type { RouteRecordRaw } from 'vue-router';

import { DEFAULT_HOME_PATH } from '@vben/constants';

import { AuthPageLayout } from '#/layouts';
import { $t } from '#/locales';
import Login from '#/views/_core/authentication/login.vue';

/** Global 404 page */
const fallbackNotFoundRoute: RouteRecordRaw = {
  component: () => import('#/views/_core/fallback/not-found.vue'),
  meta: {
    hideInBreadcrumb: true,
    hideInMenu: true,
    hideInTab: true,
    title: '404',
  },
  name: 'FallbackNotFound',
  path: '/:path(.*)*',
};

/** Basic routing, these routes must exist */
const coreRoutes: RouteRecordRaw[] = [
  {
    meta: {
      title: 'Root',
    },
    name: 'Root',
    path: '/',
    redirect: DEFAULT_HOME_PATH,
  },
  {
    children: [
      {
        component: Login,
        meta: {
          title: $t('page.core.login'),
        },
        name: 'Login',
        path: 'login',
      },
      {
        component: () => import('#/views/_core/authentication/code-login.vue'),
        meta: {
          title: $t('page.core.codeLogin'),
        },
        name: 'CodeLogin',
        path: 'code-login',
      },
      {
        component: () =>
          import('#/views/_core/authentication/qrcode-login.vue'),
        meta: {
          title: $t('page.core.qrcodeLogin'),
        },
        name: 'QrCodeLogin',
        path: 'qrcode-login',
      },
      {
        component: () =>
          import('#/views/_core/authentication/forget-password.vue'),
        meta: {
          title: $t('page.core.forgetPassword'),
        },
        name: 'ForgetPassword',
        path: 'forget-password',
      },
      {
        component: () => import('#/views/_core/authentication/register.vue'),
        meta: {
          title: $t('page.core.register'),
        },
        name: 'Register',
        path: 'register',
      },
    ],
    component: AuthPageLayout,
    meta: {
      title: 'Authentication',
    },
    name: 'Authentication',
    path: '/auth',
  },
];

export { coreRoutes, fallbackNotFoundRoute };
