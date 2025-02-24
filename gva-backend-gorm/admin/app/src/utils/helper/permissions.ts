import type { DataNode } from 'ant-design-vue/es/tree';

import { $t } from '@vben/locales';

import { type PERMISSION } from '#/constants';

export const withSuper = (...permissions: PERMISSION[]) => {
  const [first] = permissions;
  if (!first) return [] as PERMISSION[];

  const [feature] = first.split(':');
  return [...permissions, `${feature}:super`] as PERMISSION[];
};

export const getSuper = (permission: PERMISSION) => {
  if (typeof permission !== 'string') return Date.now().toString();
  const [feature] = permission.split(':');
  return `${feature}:super` as PERMISSION;
};

export const buildPermissionNode = (
  title: string,
  permissions: any,
): DataNode => {
  if (Array.isArray(permissions)) {
    const [first] = permissions ?? [];
    if (!first)
      return {
        title,
        children: [],
        key: 0,
      };

    return {
      title,
      key: getSuper(first),
      children: permissions.map((permission) => {
        if (typeof permission === 'object') {
          return permission;
        }

        return {
          title: $t(`permission.${permission.split(':')[1]}`),
          key: permission,
        };
      }),
    };
  }

  if (typeof permissions === 'object') {
    const keys = Object.keys(permissions);
    return {
      title,
      key: permissions.SUPER,
      children: keys
        .filter((item) => !item.toLowerCase().includes('super'))
        .map((key) => {
          return {
            title: $t(`permission.${key.toLowerCase()}`),
            key: permissions[key],
          };
        }),
    };
  }

  return {
    title,
    children: [],
    key: 0,
  };
};
