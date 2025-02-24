import type { BasicUserInfo } from '@vben-core/typings';

/** 用户信息 */
interface UserInfo extends BasicUserInfo {
  /**
   * User description
   */
  desc: string;
  /**
   * Home address
   */
  homePath: string;

  /**
   * accessToken
   */
  token: string;
}

export type { UserInfo };
