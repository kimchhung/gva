import { defineStore } from 'pinia';

interface AppState {
  /**
   * Whether to lock the screen status
   */
  isLockScreen: boolean;
  /**
   * Lock screen password
   */
  lockScreenPassword?: string;
}

export const useLockStore = defineStore('core-lock', {
  actions: {
    lockScreen(password: string) {
      this.isLockScreen = true;
      this.lockScreenPassword = password;
    },

    unlockScreen() {
      this.isLockScreen = false;
      this.lockScreenPassword = undefined;
    },
  },
  persist: {
    paths: ['isLockScreen', 'lockScreenPassword'],
  },
  state: (): AppState => ({
    isLockScreen: false,
    lockScreenPassword: undefined,
  }),
});
