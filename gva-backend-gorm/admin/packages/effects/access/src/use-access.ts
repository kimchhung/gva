import { computed } from 'vue';

import { preferences, updatePreferences } from '@vben/preferences';
import { useAccessStore, useUserStore } from '@vben/stores';

function useAccess() {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const accessMode = computed(() => {
    return preferences.app.accessMode;
  });

  /**
   * 基于角色判断是否有权限
   * @description: Determine whether there is permission，The role is judged by the user's role
   * @param roles
   */
  function hasAccessByRoles(roles: string[]) {
    if (userStore.userInfo?.isSuperAdmin) return true;
    const userRoleSet = new Set(userStore.userRoles);
    return roles.some((item) => userRoleSet.has(item));
  }

  /**
   * 基于权限码判断是否有权限
   * @description: Determine whether there is permission，The permission code is judged by the user's permission code
   * @param codes
   */
  function hasAccessByCodes(codes: string[]) {
    const userCodesSet = new Set(accessStore.accessCodes);

    const intersection = codes.filter((item) => userCodesSet.has(item));
    return intersection.length > 0;
  }

  async function toggleAccessMode() {
    updatePreferences({
      app: {
        accessMode:
          preferences.app.accessMode === 'frontend' ? 'backend' : 'frontend',
      },
    });
  }

  function hasAccessByPermissions(permissions: string[]) {
    if (userStore.userInfo?.isSuperAdmin) return true;
    const userPermissionSet = new Set(userStore.userPermissions);
    return permissions.some((value) => userPermissionSet.has(value));
  }

  return {
    accessMode,
    hasAccessByCodes,
    hasAccessByPermissions,
    hasAccessByRoles,
    toggleAccessMode,
  };
}

export { useAccess };
