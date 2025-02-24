import type { DataNode } from 'ant-design-vue/es/tree';

import { $t } from '@vben/locales';

import { buildPermissionNode } from '#/utils/helper/permissions';

import {
  ADMIN_PERMISSION,
  ADMIN_ROLE_PERMISSION,
  CONFIGURATION_PERMISSION,
  DOCUMENT_PERMISSION,
  OPERATION_LOG_PERMISSION,
} from './permission';

export const getPermissionTree = (): DataNode[] => [
  {
    title: $t('menu.system.title'),
    key: 'system',
    children: [
      buildPermissionNode(
        $t('page.system.configurations.title'),
        CONFIGURATION_PERMISSION,
      ),
      {
        title: $t('page.admin.title'),
        key: 'admin',
        children: [
          buildPermissionNode($t('page.admin.users.title'), ADMIN_PERMISSION),
          buildPermissionNode(
            $t('page.admin.roles.title'),
            ADMIN_ROLE_PERMISSION,
          ),
        ],
      },

      buildPermissionNode(
        $t('page.system.operationLog.title'),
        OPERATION_LOG_PERMISSION,
      ),
    ],
  },
  buildPermissionNode($t('page.document.title'), DOCUMENT_PERMISSION),
];
