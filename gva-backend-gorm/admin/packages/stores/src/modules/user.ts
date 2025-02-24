import { acceptHMRUpdate, defineStore } from 'pinia';

interface BasicUserInfo {
  /**
   * avatar
   */
  avatar: string;
  /**
   * Super Admin
   */
  isSuperAdmin?: boolean;
  /**
   * User permissions
   */
  permissions?: string[];
  /**
   * User nickname
   */
  realName: string;
  /**
   * User role
   */
  roles?: string[];
  /**
   * User ID
   */
  userId: number;
  /**
   * username
   */
  username: string;
}

interface AccessState {
  /**
   * User information
   */
  userInfo: BasicUserInfo | null;
  /**
   * User permissions
   */
  userPermissions: string[];
  /**
   * User role
   */
  userRoles: string[];
}

/**
 * @zh_CN User information related
 */
export const useUserStore = defineStore('core-user', {
  actions: {
    setUserInfo(userInfo: BasicUserInfo | null) {
      // Set user information
      this.userInfo = userInfo;
      // Set role information
      this.setUserRoles(userInfo?.roles ?? []);
      this.setUserPermissions(userInfo?.permissions ?? []);
    },
    setUserPermissions(permissions: string[]) {
      this.userPermissions = permissions;
    },
    setUserRoles(roles: string[]) {
      this.userRoles = roles;
    },
  },
  state: (): AccessState => ({
    userInfo: null,
    userPermissions: [],
    userRoles: [],
  }),
});

// Solve the problem of hot updates
const hot = import.meta.hot;
if (hot) {
  hot.accept(acceptHMRUpdate(useUserStore, hot));
}
