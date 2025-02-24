import {
  ADMIN_PERMISSION,
  ADMIN_ROLE_PERMISSION,
  CONFIGURATION_PERMISSION,
  OPERATION_LOG_PERMISSION,
} from '#/constants';
import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';
import { withSuper } from '#/utils/helper/permissions';

import { defineRoute } from '../util';

export default defineRoute([
  {
    component: BasicLayout,
    name: 'System',
    path: '/system',
    meta: {
      icon: 'lucide:monitor-cog',
      keepAlive: true,
      order: 7,
      title: $t('menu.system.title'),
    },
    children: [
      {
        name: 'Configurations',
        path: 'configurations',
        component: () => import('#/views/system/configurations/index.vue'),
        meta: {
          keepAlive: true,
          icon: 'lucide:settings',
          title: $t('page.system.configurations.title'),
          permissions: withSuper(CONFIGURATION_PERMISSION.VIEW),
        },
      },
      {
        meta: {
          keepAlive: true,
          icon: 'lucide:user-cog',
          title: 'page.admin.title',
        },
        name: 'Admin',
        path: 'admin',
        children: [
          {
            component: () => import('#/views/system/users/index.vue'),
            meta: {
              keepAlive: true,
              icon: 'lucide:users',
              title: 'page.admin.users.title',
              permissions: withSuper(ADMIN_PERMISSION.VIEW),
            },
            name: 'Users',
            path: 'users',
          },
          {
            component: () => import('#/views/system/roles/index.vue'),
            meta: {
              keepAlive: true,
              icon: 'lucide:shield-check',
              title: 'page.admin.roles.title',
              permissions: withSuper(ADMIN_ROLE_PERMISSION.VIEW),
            },
            name: 'Roles',
            path: 'roles',
          },
          {
            component: () => import('#/views/system/roles/form.vue'),
            meta: {
              keepAlive: true,
              icon: 'lucide:shield',
              title: ({ t }) =>
                t('common.create', {
                  name: t('page.admin.roles.title'),
                }),
              hideInMenu: true,
              permissions: withSuper(ADMIN_ROLE_PERMISSION.ADD),
              activePath: '/system/admin/roles',
            },
            name: 'Create Role',
            path: 'roles/create',
          },
          {
            component: () => import('#/views/system/roles/form.vue'),
            meta: {
              keepAlive: true,
              icon: 'lucide:shield',
              title: ({ t }) =>
                t('common.edit', {
                  name: t('page.admin.roles.title'),
                }),
              permissions: withSuper(ADMIN_ROLE_PERMISSION.EDIT),
              hideInMenu: true,
            },
            name: 'Update Role',
            path: 'roles/:id',
          },
        ],
      },
      {
        name: 'Operation Logs',
        path: 'operation-logs',
        component: () => import('#/views/system/operation-log/index.vue'),
        meta: {
          keepAlive: true,
          icon: 'ant-design:file-done-outlined',
          title: $t('page.system.operationLog.title'),
          permissions: withSuper(OPERATION_LOG_PERMISSION.VIEW),
        },
      },
    ],
  },
]);
