import type { App, Directive, DirectiveBinding } from 'vue';

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import DOMPurify from 'dompurify';

const safeHtmlDirective: Directive = {
  mounted: (el: Element, binding: DirectiveBinding<string | string[]>) => {
    const value = Array.isArray(binding.value)
      ? binding.value.join('')
      : binding.value;
    el.innerHTML = DOMPurify.sanitize(value);
  },
};

export function registerSafeHtmlDirective(app: App) {
  app.directive('safe-html', safeHtmlDirective);
}
