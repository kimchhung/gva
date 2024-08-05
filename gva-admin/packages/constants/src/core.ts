/**
 * @zh_CN Landing page URL address
 */
export const LOGIN_PATH = '/auth/login';

/**
 * @zh_CN Default homepage address
 */
export const DEFAULT_HOME_PATH = '/home';

export interface LanguageOption {
  label: string;
  value: 'en-US' | 'zh-CN';
}

/**
 * Supported languages
 */
export const SUPPORT_LANGUAGES: LanguageOption[] = [
  {
    label: '简体中文',
    value: 'zh-CN',
  },
  {
    label: 'English',
    value: 'en-US',
  },
];
