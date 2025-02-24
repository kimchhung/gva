import { ref, toRaw, useModel, watch } from 'vue';

import { $t } from '@vben/locales';

import { type FormInstance, notification } from 'ant-design-vue';
import { cloneDeep } from 'lodash';

interface BaseModalFormOptions<T> {
  onOpen?: (props: T, opts: { setInitialForm: (form: any) => void }) => void;
  onClose?: () => void;
}

export const useBaseModalForm = <T extends Record<string, any>>(
  props: T,
  opts: BaseModalFormOptions<T>,
) => {
  const formRef = ref<FormInstance>();
  const openRef = useModel(props, 'open');
  const loadingRef = ref(false);
  const initialFormState = ref<any>();

  const setInitialForm = (initial: any) => {
    initialFormState.value = cloneDeep(initial);
  };

  watch(openRef, () => {
    if (openRef.value) {
      const init = opts.onOpen?.(props, {
        setInitialForm,
      });
      initialFormState.value = cloneDeep(init);
    } else {
      setTimeout(() => {
        formRef.value?.resetFields();
      }, 300);
      opts.onClose?.();
    }
  });

  const getFormDirtyState = (form: any) => {
    if (!props.record) return true;
    return (
      JSON.stringify(toRaw(initialFormState.value)) !==
      JSON.stringify(toRaw(form))
    );
  };

  const showSuccessMessage = () => {
    notification.success({
      message: props.record
        ? $t('message.updateSuccess')
        : $t('message.createSuccess'),
    });
  };

  const closeModal = (opts?: { showSuccessMessage?: boolean }) => {
    (openRef.value as boolean) = false;
    if (opts?.showSuccessMessage) {
      showSuccessMessage();
    }
  };

  const handleSubmit = (callback: () => void) => {
    return async () => {
      if (!(await formRef.value?.validate())) return;
      callback();
    };
  };

  return {
    openRef,
    formRef,
    closeModal,
    loadingRef,
    handleSubmit,
    setInitialForm,
    getFormDirtyState,
  };
};
