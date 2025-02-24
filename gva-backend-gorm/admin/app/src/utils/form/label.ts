import { $t } from '@vben/locales';

export const getFormInfos = <T extends { [k: string]: string }>(schema: T) => {
  return Object.entries(schema)?.reduce(
    (acc, [key, value]) => {
      acc[key as keyof T] = {
        label: value,
        placeholder: $t(`form.input.placeholder`, {
          field: value,
        }),
        name: key,
      };
      return acc;
    },
    {} as Record<
      keyof T,
      {
        label: string;
        name: string;
        placeholder: string;
      }
    >,
  );
};
