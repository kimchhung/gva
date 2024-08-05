import type {
  ImportLocaleFn,
  LoadMessageFn,
  LocaleSetupOptions,
  SupportedLanguagesType,
} from './typing';

import type { App } from 'vue';

import { i18n, loadI18nMessages, loadLocalesMap } from './i18n';

const $t = i18n.global.t;

let loadMessages: LoadMessageFn;

async function loadLocaleMessages(lang: SupportedLanguagesType) {
  const mergeMessage = await loadMessages(lang);
  await loadI18nMessages(lang);
  i18n.global.mergeLocaleMessage(lang, mergeMessage);
}

async function setupI18n(app: App, options: LocaleSetupOptions = {}) {
  const { defaultLocale = 'en-US' } = options;
  // APP can expand the internationalization of some third -party libraries and component libraries by itself
  loadMessages = options.loadMessages || (async () => ({}));
  app.use(i18n);
  await loadLocaleMessages(defaultLocale);

  // Print warning at the console
  i18n.global.setMissingHandler((locale, key) => {
    if (options.missingWarn && key.includes('.')) {
      console.warn(
        `[intlify] Not found '${key}' key in '${locale}' locale messages.`,
      );
    }
  });
}

export { $t, i18n, loadLocaleMessages, loadLocalesMap, setupI18n };
export type { CompileError } from '@intlify/core-base';
export { useI18n } from 'vue-i18n';
export type { ImportLocaleFn, LocaleSetupOptions, SupportedLanguagesType };
export type { Locale } from 'vue-i18n';
