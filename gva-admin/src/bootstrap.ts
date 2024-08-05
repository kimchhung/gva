import { createApp } from 'vue';

import { registerAccessDirective } from '@gva/access';
import { initStores } from '@gva/stores';
import '@gva/styles';
import '@gva/styles/antd';

import { setupI18n } from '#/locales';

import App from './app.vue';
import { router } from './router';
import { setupAPI } from './api';

async function bootstrap(namespace: string) {
  const app = createApp(App);

  //Internationalization i18N configuration
  await setupI18n(app);

  // Configuration pinia-tore
  await initStores(app, { namespace });

  // setup api connector
  await setupAPI(app);

  // Installer instruction
  registerAccessDirective(app);

  // Configure route and route guard
  app.use(router);

  app.mount('#app');
}

export { bootstrap };
