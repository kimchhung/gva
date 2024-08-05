import type { BasicUserInfo } from '@gva-core/typings';

/** User Info */
interface UserInfo extends BasicUserInfo {
  id: string;

  // default redirect
  homePath?: string;

  /**
   * accessToken
   */
  token?: string;
  [key: string]: any;
}

export type { UserInfo };
