import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';
import ClipboardJS from 'clipboard';

export const copyToClipboard = (val: string) => {
  notification.close('copy');
  const isSupported = ClipboardJS.isSupported();
  if (isSupported) {
    const clipboard = new ClipboardJS('.copybutton');
    const button = document.createElement('button');
    document.body.append(button);
    button.setAttribute('class', 'copybutton');
    button.dataset.clipboardText = val;
    button.click();
    clipboard.destroy();
    button.remove();
    notification.success({
      key: 'copy',
      message: $t('common.copied', {
        value: val,
      }),
    });
  }
  return isSupported;
};
