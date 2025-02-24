import type { ModalApiOptions, ModalState } from './modal';

import { isFunction, Store } from '@vben-core/shared';

export class ModalApi {
  private api: Pick<
    ModalApiOptions,
    'onBeforeClose' | 'onCancel' | 'onConfirm' | 'onOpenChange'
  >;
  // private prevState!: ModalState;
  private state!: ModalState;

  public sharedData: Record<'payload', any> = {
    payload: {},
  };

  public store: Store<ModalState>;

  constructor(options: ModalApiOptions = {}) {
    const {
      connectedComponent: _,
      onBeforeClose,
      onCancel,
      onConfirm,
      onOpenChange,
      ...storeState
    } = options;

    const defaultState: ModalState = {
      centered: false,
      closeOnClickModal: true,
      closeOnPressEscape: true,
      confirmLoading: false,
      footer: true,
      header: true,
      isOpen: false,
      loading: false,
      modal: true,
      showCancelButton: true,
      showConfirmButton: true,
      title: '',
    };

    this.store = new Store<ModalState>(
      {
        ...defaultState,
        ...storeState,
      },
      {
        onUpdate: () => {
          const state = this.store.state;

          if (state?.isOpen === this.state?.isOpen) {
            this.state = state;
          } else {
            this.state = state;
            this.api.onOpenChange?.(!!state?.isOpen);
          }
        },
      },
    );

    this.api = {
      onBeforeClose,
      onCancel,
      onConfirm,
      onOpenChange,
    };
  }

  batchStore(cb: () => void) {
    this.store.batch(cb);
  }

  close() {
    const allowClose = this.api.onBeforeClose?.() ?? true;
    if (allowClose) {
      this.store.setState((prev) => ({ ...prev, isOpen: false }));
    }
  }

  getData<T extends object = Record<string, any>>() {
    return (this.sharedData?.payload ?? {}) as T;
  }

  onCancel() {
    if (this.api.onCancel) {
      this.api.onCancel?.();
    } else {
      this.close();
    }
  }

  onConfirm() {
    this.api.onConfirm?.();
  }

  open() {
    this.store.setState((prev) => ({ ...prev, isOpen: true }));
  }

  setData<T>(payload: T) {
    this.sharedData.payload = payload;
  }

  setState(
    stateOrFn:
      | ((prev: ModalState) => Partial<ModalState>)
      | Partial<ModalState>,
  ) {
    if (isFunction(stateOrFn)) {
      this.store.setState(stateOrFn);
    } else {
      this.store.setState((prev) => ({ ...prev, ...stateOrFn }));
    }
  }
}
