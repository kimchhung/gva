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
   * User's Nickname
   */
  name: string;
  /**
   * User ID
   */
  id: string;
  /**
   * username
   */
  username: string;
}

export type { BasicOption, BasicUserInfo, SelectOption, TabOption };
