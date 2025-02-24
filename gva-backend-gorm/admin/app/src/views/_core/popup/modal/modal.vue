<script setup lang="ts">
import type { ExtendedModalApi, ModalProps } from './modal';

import { nextTick, ref, watch } from 'vue';

import { usePriorityValue } from '@vben/hooks';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  VisuallyHidden,
} from '@vben-core/shadcn-ui';
import { cn } from '@vben-core/shared';

import { Button } from 'ant-design-vue';

import { $t } from '#/locales';

interface Props extends ModalProps {
  class?: string;
  contentClass?: string;
  footerClass?: string;
  headerClass?: string;
  modalApi?: ExtendedModalApi;
}

const props = withDefaults(defineProps<Props>(), {
  class: '',
  contentClass: '',
  footerClass: '',
  headerClass: '',
  modalApi: undefined,
});

const contentRef = ref();
const wrapperRef = ref<HTMLElement>();
const dialogRef = ref();
const headerRef = ref();
const footerRef = ref();

const state = props.modalApi?.useStore?.();

const header = usePriorityValue('header', props, state);
const title = usePriorityValue('title', props, state);
const description = usePriorityValue('description', props, state);
const showFooter = usePriorityValue('footer', props, state);
const showLoading = usePriorityValue('loading', props, state);
const closable = usePriorityValue('closable', props, state);
const modal = usePriorityValue('modal', props, state);
const centered = usePriorityValue('centered', props, state);
const confirmLoading = usePriorityValue('confirmLoading', props, state);
const cancelText = usePriorityValue('cancelText', props, state);
const confirmText = usePriorityValue('confirmText', props, state);
const closeOnClickModal = usePriorityValue('closeOnClickModal', props, state);
const closeOnPressEscape = usePriorityValue('closeOnPressEscape', props, state);
const showCancelButton = usePriorityValue('showCancelButton', props, state);
const showConfirmButton = usePriorityValue('showConfirmButton', props, state);

watch(
  () => state?.value?.isOpen,
  async (v) => {
    if (v) {
      await nextTick();
      if (!contentRef.value) return;
      const innerContentRef = contentRef.value.getContentRef();
      dialogRef.value = innerContentRef.$el;
    }
  },
);

watch(
  () => showLoading.value,
  (v) => {
    if (v && wrapperRef.value) {
      wrapperRef.value.scrollTo({
        // behavior: 'smooth',
        top: 0,
      });
    }
  },
);

function interactOutside(e: Event) {
  if (!closeOnClickModal.value) {
    e.preventDefault();
  }
}
function escapeKeyDown(e: KeyboardEvent) {
  if (!closeOnPressEscape.value) {
    e.preventDefault();
  }
}
// pointer-down-outside
function pointerDownOutside(e: Event) {
  const target = e.target as HTMLElement;
  const isDismissableModal = !!target?.dataset.dismissableModal;
  if (!closeOnClickModal.value || !isDismissableModal) {
    e.preventDefault();
  }
}
</script>
<template>
  <Dialog
    :modal="modal"
    :open="state?.isOpen"
    @update:open="() => modalApi?.close()"
  >
    <DialogContent
      ref="contentRef"
      :class="
        cn(
          'border-border left-0 right-0 top-[10vh] mx-auto flex max-h-[80%] w-[520px] flex-col border p-0',
          props.class,
          {
            'top-1/2 !-translate-y-1/2': centered,
          },
        )
      "
      :show-close="closable"
      close-class="top-3"
      @escape-key-down="escapeKeyDown"
      @interact-outside="interactOutside"
      @pointer-down-outside="pointerDownOutside"
    >
      <DialogHeader
        ref="headerRef"
        :class="
          cn(
            'border-b px-5 py-4',
            {
              hidden: !header,
            },
            props.headerClass,
          )
        "
      >
        <DialogTitle v-if="title" class="text-left">
          <slot name="title">
            {{ title }}
          </slot>
        </DialogTitle>
        <DialogDescription v-if="description">
          <slot name="description">
            {{ description }}
          </slot>
        </DialogDescription>
        <VisuallyHidden v-if="!title || !description">
          <DialogTitle v-if="!title" />
          <DialogDescription v-if="!description" />
        </VisuallyHidden>
      </DialogHeader>
      <div
        ref="wrapperRef"
        :class="
          cn('relative min-h-24 flex-1 overflow-y-auto p-3', contentClass, {
            'overflow-hidden': showLoading,
          })
        "
      >
        <slot></slot>
      </div>

      <DialogFooter
        v-if="showFooter"
        ref="footerRef"
        :class="
          cn(
            'flex-row items-center justify-end border-t p-2',
            props.footerClass,
          )
        "
      >
        <slot name="prepend-footer"></slot>
        <slot name="footer">
          <Button
            v-if="showCancelButton"
            type="ghost"
            @click="() => modalApi?.onCancel()"
          >
            <slot name="cancelText">
              {{ cancelText || $t('common.cancel') }}
            </slot>
          </Button>
          <Button
            v-if="showConfirmButton"
            :loading="confirmLoading"
            type="primary"
            @click="() => modalApi?.onConfirm()"
          >
            <slot name="confirmText">
              {{ confirmText || $t('common.confirm') }}
            </slot>
          </Button>
        </slot>
        <slot name="append-footer"></slot>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
