import type { LocaleSetupOptions, SupportedLanguagesType } from '@vben/locales';
import type { Locale } from 'ant-design-vue/es/locale';

import type { App } from 'vue';
import { ref } from 'vue';

import { $t, setupI18n as coreSetup, loadLocalesMap } from '@vben/locales';
import { preferences } from '@vben/preferences';

import { i18nChangeLanguage } from '@wangeditor/editor';
import antdEnLocale from 'ant-design-vue/es/locale/en_US';
import antdDefaultLocale from 'ant-design-vue/es/locale/zh_CN';
import dayjs from 'dayjs';

const antdLocale = ref<Locale>(antdDefaultLocale);

const modules = import.meta.glob('./langs/*.json');

const localesMap = loadLocalesMap(modules);

const timezonesMap = loadLocalesMap(import.meta.glob('./timezones/*.json'));

/**
 * Load the unique language package
 * Here you can also translate translation data from the server
 * @param lang
 */
async function loadMessages(lang: SupportedLanguagesType) {
  const [appLocaleMessages, timezones] = await Promise.all([
    localesMap[lang]?.(),
    timezonesMap[lang]?.(),
    loadThirdPartyMessage(lang),
  ]);

  return Object.assign(appLocaleMessages?.default ?? {}, {
    TIMEZONE: timezones?.default,
  });
}

/**
 * Load the language package of the third -party component library
 * @param lang
 */
async function loadThirdPartyMessage(lang: SupportedLanguagesType) {
  await Promise.all([
    loadAntdLocale(lang),
    loadDayjsLocale(lang),
    loadWangEditor(lang),
  ]);
}

/**
 * Load DAYJS language package
 * @param lang
 */
async function loadDayjsLocale(lang: SupportedLanguagesType) {
  let locale;
  switch (lang) {
    case 'en-US': {
      locale = await import('dayjs/locale/en');
      break;
    }
    case 'zh-CN': {
      locale = await import('dayjs/locale/zh-cn');
      break;
    }
    // 默认使用英语
    default: {
      locale = await import('dayjs/locale/en');
    }
  }
  if (locale) {
    dayjs.locale(locale);
  } else {
    console.error(`Failed to load dayjs locale for ${lang}`);
  }
}

/**
 * Load the ATD language package
 * @param lang
 */
async function loadAntdLocale(lang: SupportedLanguagesType) {
  switch (lang) {
    case 'en-US': {
      antdLocale.value = antdEnLocale;
      break;
    }
    case 'zh-CN': {
      antdLocale.value = antdDefaultLocale;
      break;
    }
  }
}

async function loadWangEditor(lang: SupportedLanguagesType) {
  switch (lang) {
    case 'en-US': {
      i18nChangeLanguage(lang);
      break;
    }
    case 'zh-CN': {
      i18nChangeLanguage(lang);
      break;
    }
  }
}

async function setupI18n(app: App, options: LocaleSetupOptions = {}) {
  await coreSetup(app, {
    defaultLocale: preferences.app.locale,
    loadMessages,
    missingWarn: !import.meta.env.PROD,
    ...options,
  });
}

export { $t, antdLocale, setupI18n };
