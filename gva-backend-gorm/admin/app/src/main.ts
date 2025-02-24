import { initPreferences } from '@vben/preferences';
import { unmountGlobalLoading } from '@vben/utils';

import { overridesPreferences } from './preferences';

/**
 * After the application is initialized, the page load rendering
 */
async function initApplication() {
  // name is used to specify the unique logo of the project
  // The preference settings for distinguishing different items and the storage data prefix, and other data that needs to be isolated
  const env = import.meta.env.PROD ? 'prod' : 'dev';
  const appVersion = import.meta.env.VITE_APP_VERSION;
  const namespace = `${import.meta.env.VITE_APP_NAMESPACE}-${appVersion}-${env}`;

  // APP preference settings initialization
  await initPreferences({
    namespace,
    overrides: overridesPreferences,
  });

  // Start the application and mount
  // Vue application main logic and view
  const { bootstrap } = await import('./bootstrap');
  await bootstrap(namespace);

  // Remove and destroy loading
  unmountGlobalLoading();
}

initApplication();
