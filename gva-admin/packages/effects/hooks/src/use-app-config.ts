import type {
  ApplicationConfig,
  VbenAdminProAppConfigRaw,
} from '@vben/types/global';

/**
 * 由 vite-inject-app-config 注入的全局配置
 */
export function useAppConfig(
  env: Record<string, any>,
  isProduction: boolean,
): ApplicationConfig {
  // Under the production environment, use Window._vben_ADMIN_PRO_APP_CONF_ global variables
  const config = isProduction
    ? window._VBEN_ADMIN_PRO_APP_CONF_
    : (env as VbenAdminProAppConfigRaw);

  const { VITE_GLOB_API_URL } = config;

  return {
    apiURL: VITE_GLOB_API_URL,
  };
}
