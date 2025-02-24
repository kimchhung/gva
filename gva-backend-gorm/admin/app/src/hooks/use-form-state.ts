import type { FormInstance } from 'ant-design-vue';

import { reactive, ref, toRaw, watch } from 'vue';

import { cloneDeep } from 'lodash';

export const useFormState = <T extends Record<string, any>>({
  defaultValue,
}: {
  defaultValue: T;
}) => {
  const form = reactive<T>(defaultValue);
  const initialFormState = reactive<T>(cloneDeep(defaultValue));
  const formRef = ref<FormInstance>();

  const isDirty = ref(false);

  watch(
    () => form,
    () => {
      isDirty.value =
        JSON.stringify(toRaw(initialFormState)) !== JSON.stringify(toRaw(form));
    },
    { deep: true },
  );

  const resetForm = () => {
    Object.assign(form, cloneDeep(initialFormState));
  };

  const setInitialForm = (initial: any) => {
    Object.assign(initialFormState, cloneDeep(initial));
    Object.assign(form, cloneDeep(initial));
  };

  const validateFields = async () => {
    return formRef.value?.validateFields();
  };

  return {
    form,
    formRef,
    isDirty,
    setInitialForm,
    resetForm,
    validateFields,
  };
};
