interface BasicOption {
  label: string;
  value: string;
}

type SelectOption = BasicOption;

type TabOption = BasicOption;

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

export type { BasicOption, BasicUserInfo, SelectOption, TabOption };
