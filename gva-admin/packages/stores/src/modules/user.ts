import { acceptHMRUpdate, defineStore } from 'pinia';

interface BasicUserInfo {
  /**
   * avatar
   */
  avatar: string;
  /**
   * User's Nickname
   */
  name: string;
  /**
   * User role
   */
  roles?: string[];
  /**
   * User ID
   */
  id: string;
  /**
   * username
   */
  username: string;
}

interface AccessState {
  /**
   * User Info
   */
  userInfo: BasicUserInfo | null;
  /**
   * User role
   */
  userRoles: string[];
}

/**
 *  User information related
 */
export const useUserStore = defineStore('core-user', {
  actions: {
    setUserInfo(userInfo: BasicUserInfo | null) {
      // Set user information
      this.userInfo = userInfo;
      // Set role information
      const roles = userInfo?.roles ?? [];
      this.setUserRoles(roles);
    },
    setUserRoles(roles: string[]) {
      this.userRoles = roles;
    },
  },
  state: (): AccessState => ({
    userInfo: null,
    userRoles: [],
  }),
});

// Solve the problem of hot updates
const hot = import.meta.hot;
if (hot) {
  hot.accept(acceptHMRUpdate(useUserStore, hot));
}
