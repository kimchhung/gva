import { defineOverridesPreferences } from '@vben/preferences';

/**
 * @description Project configuration file
 * Only part of the configuration in the project is needed. If the configuration is not needed, the default configuration will be automatically used.
 */
export const overridesPreferences = defineOverridesPreferences({
  // overrides
  app: {
    name: import.meta.env.VITE_APP_TITLE,
    locale: 'en-US',
  },
});
