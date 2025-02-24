import type { ModalApi } from './modal-api';

import type { Component, Ref } from 'vue';

export interface ModalProps {
  cancelText?: string;
  centered?: boolean;
  closable?: boolean;
  closeOnClickModal?: boolean;
  closeOnPressEscape?: boolean;
  confirmLoading?: boolean;
  confirmText?: string;
  description?: string;
  footer?: boolean;
  header?: boolean;
  loading?: boolean;
  modal?: boolean;
  showCancelButton?: boolean;
  showConfirmButton?: boolean;
  title?: string;
}

export interface ModalState extends ModalProps {
  isOpen?: boolean;
  sharedData?: Record<string, any>;
}

export type ExtendedModalApi = {
  useStore: <T = NoInfer<ModalState>>(
    selector?: (state: NoInfer<ModalState>) => T,
  ) => Readonly<Ref<T>>;
} & ModalApi;

export interface ModalApiOptions extends ModalState {
  connectedComponent?: Component;
  onBeforeClose?: () => void;
  onCancel?: () => void;
  onConfirm?: () => void;
  onOpenChange?: (isOpen: boolean) => void;
}
