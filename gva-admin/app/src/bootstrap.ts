import { createApp } from 'vue';

import { registerAccessDirective } from '@vben/access';
import { initStores } from '@vben/stores';
import '@vben/styles';
import '@vben/styles/antd';

import { setupI18n } from '#/locales';

import { setupAPI } from './api';
import App from './app.vue';
import { router } from './router';

async function bootstrap(namespace: string) {
  const app = createApp(App);

  setupAPI(app);

  await setupI18n(app);

  await initStores(app, { namespace });

  registerAccessDirective(app);

  app.use(router);

  app.mount('#app');
}

export { bootstrap };
