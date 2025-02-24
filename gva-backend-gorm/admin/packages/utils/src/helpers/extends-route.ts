export const translateMetaTitle = (
  meta: Record<string, any>,
  i18n: {
    $locale: any;
    $t: (key: string, options?: any) => string;
  },
) => {
  const { title, titleEn, titleZh, ...more } = meta;
  const newMeta = { ...more, title: '' };

  if (titleEn && titleZh) {
    const lang = i18n.$locale.value ?? i18n.$locale;
    newMeta.title = String(lang.includes('en') ? meta.titleEn : meta.titleZh);
  } else {
    newMeta.title =
      typeof title === 'function'
        ? title({ t: i18n.$t })
        : i18n.$t(String(title));
  }
  return newMeta;
};
