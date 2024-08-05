import { initPreferences } from '@gva/preferences';
import { unmountGlobalLoading } from '@gva/utils';

import { overridesPreferences } from './preferences';

/**
 * After the application is initialized, the page load rendering
 */
async function initApplication() {
  //:
  // Used to distinguish the preference settings of different projects, the Key prefix of the storage data, and other data that needs to be isolation
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
