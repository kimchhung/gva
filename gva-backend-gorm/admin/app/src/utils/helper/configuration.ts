import type { Configuration } from '#/api/configuration/types';

import { preferences } from '@vben/preferences';

export const mapConfigTree = (
  data: Configuration[],
  parentId?: number,
): Configuration[] => {
  return data
    .filter((item) => item.parentId === parentId)
    .map((item) => ({
      ...item,
      children: mapConfigTree(data, item.id),
    }));
};

export const findParents = (
  data: Configuration[],
  parentId?: number,
): Configuration[] => {
  const parent = data.find((item) => item.id === parentId);
  if (parent) {
    return [...findParents(data, parent.parentId), parent];
  }
  return [];
};

export const getConfigTitle = (config: Configuration): string => {
  if (config.metadata?.labelEn || config.metadata?.labelZh) {
    return preferences.app.locale === 'en-US'
      ? config.metadata?.labelEn
      : config.metadata?.labelZh;
  }

  return config.key;
};
