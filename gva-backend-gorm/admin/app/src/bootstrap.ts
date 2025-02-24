import { createApp } from 'vue';

import { registerAccessDirective } from '@vben/access';
import { initStores } from '@vben/stores';
import '@vben/styles';
import '@vben/styles/antd';

import { autoAnimatePlugin } from '@formkit/auto-animate/vue';
import {
  VueQueryPlugin,
  type VueQueryPluginOptions,
} from '@tanstack/vue-query';
import AntDesignVue from 'ant-design-vue';

import { setupI18n } from '#/locales';

import { queryClient, setupAPI } from './api';
import App from './app.vue';
import { registerSafeHtmlDirective } from './directives/safe-html';
import { router } from './router';

async function bootstrap(namespace: string) {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  const Vue3Cron = await import('vue3-cron-antd')
    .then((m) => m.default)
    .catch(() => null);

  const app = createApp(App);

  setupAPI(app);

  await setupI18n(app);

  await initStores(app, { namespace });

  registerAccessDirective(app);
  registerSafeHtmlDirective(app);

  app.use(router);

  app.use(autoAnimatePlugin);
  app.use(AntDesignVue);
  app.use(Vue3Cron);

  const vueQueryPluginOptions: VueQueryPluginOptions = {
    queryClient,
    enableDevtoolsV6Plugin: false,
  };

  app.use(VueQueryPlugin, vueQueryPluginOptions);
  app.mount('#app');
}

export { bootstrap };
