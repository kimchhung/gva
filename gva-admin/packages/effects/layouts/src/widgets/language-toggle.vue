<script setup lang="ts">
import type { SupportedLanguagesType } from '@gva/locales';

import { SUPPORT_LANGUAGES } from '@gva/constants';
import { Languages } from '@gva/icons';
import { loadLocaleMessages } from '@gva/locales';
import { preferences, updatePreferences } from '@gva/preferences';
import { VbenDropdownRadioMenu, VbenIconButton } from '@gva-core/shadcn-ui';

defineOptions({
  name: 'LanguageToggle',
});

async function handleUpdate(value: string) {
  const locale = value as SupportedLanguagesType;
  updatePreferences({
    app: {
      locale,
    },
  });
  await loadLocaleMessages(locale);
}
</script>

<template>
  <div>
    <VbenDropdownRadioMenu
      :menus="SUPPORT_LANGUAGES"
      :model-value="preferences.app.locale"
      @update:model-value="handleUpdate"
    >
      <VbenIconButton>
        <Languages class="size-4" />
      </VbenIconButton>
    </VbenDropdownRadioMenu>
  </div>
</template>
