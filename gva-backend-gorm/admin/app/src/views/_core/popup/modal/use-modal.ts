import type { ExtendedModalApi, ModalApiOptions, ModalProps } from './modal';

import { defineComponent, h, inject, nextTick, provide, reactive } from 'vue';

import { useStore } from '@vben-core/shared';

import BaseModel from './modal.vue';
import { ModalApi } from './modal-api';

const USER_MODAL_INJECT_KEY = Symbol('BASE_MODAL_INJECT');

export function useBaseModal<TParentModalProps extends ModalProps = ModalProps>(
  options: ModalApiOptions = {},
) {
  const { connectedComponent } = options;
  if (connectedComponent) {
    const extendedApi = reactive({});
    const Modal = defineComponent(
      (props: TParentModalProps, { attrs, slots }) => {
        provide(USER_MODAL_INJECT_KEY, {
          extendApi(api: ExtendedModalApi) {
            Object.setPrototypeOf(extendedApi, api);
          },
          options,
        });
        checkProps(extendedApi as ExtendedModalApi, {
          ...props,
          ...attrs,
          ...slots,
        });

        return () => h(connectedComponent, { ...props, ...attrs }, slots);
      },
      {
        inheritAttrs: false,
        name: 'ParentBaseModal',
      },
    );

    return [Modal, extendedApi as ExtendedModalApi] as const;
  }

  const injectData = inject<any>(USER_MODAL_INJECT_KEY, {});

  const mergedOptions = {
    ...injectData.options,
    ...options,
  } as ModalApiOptions;

  const api = new ModalApi(mergedOptions);

  const extendedApi: ExtendedModalApi = api as never;

  extendedApi.useStore = (selector) => {
    return useStore(api.store, selector);
  };

  const Modal = defineComponent(
    (props: ModalProps, { attrs, slots }) => {
      return () =>
        h(BaseModel, { ...props, ...attrs, modalApi: extendedApi }, slots);
    },
    {
      inheritAttrs: false,
      name: 'BaseModel',
    },
  );
  injectData.extendApi?.(extendedApi);
  return [Modal, extendedApi] as const;
}

async function checkProps(api: ExtendedModalApi, attrs: Record<string, any>) {
  if (!attrs || Object.keys(attrs).length === 0) {
    return;
  }
  await nextTick();

  const state = api?.store?.state;

  if (!state) {
    return;
  }

  const stateKeys = new Set(Object.keys(state));

  for (const attr of Object.keys(attrs)) {
    if (stateKeys.has(attr)) {
      console.warn(
        `[Base Modal]: When 'connectedComponent' exists, do not set props or slots '${attr}', which will increase complexity. If you need to modify the props of Modal, please use useBaseModal or api.`,
      );
    }
  }
}
