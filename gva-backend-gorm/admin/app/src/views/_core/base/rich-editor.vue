<script setup lang="ts">
import { computed, onBeforeUnmount, reactive, shallowRef, toRefs } from 'vue';

import {
  type IDomEditor,
  type IEditorConfig,
  type IToolbarConfig,
} from '@wangeditor/editor';
import { Editor, Toolbar } from '@wangeditor/editor-for-vue';

const props = defineProps<{
  class?: string;
  height?: string;
  modelValue: string;
}>();

const emit = defineEmits(['update:modelValue']);

const content = computed<string>({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit('update:modelValue', value);
  },
});

// editor instance, use `shallowRef`
const editorRef = shallowRef<IDomEditor>();
const editorState = reactive({
  editorConfig: {
    autoFocus: false,
    customAlert: () => {},
    hoverbarKeys: {},
    MENU_CONF: {
      fontSize: {
        fontSizeList: [
          '12px',
          '14px',
          '16px',
          '18px',
          '20px',
          '24px',
          '28px',
          '32px',
          '36px',
        ],
      },
      lineHeight: {
        lineHeightList: ['1', '1.5', '1.75', '2', '2.5', '3'],
      },
      uploadImage: {
        async customUpload(
          file: File,
          insertFn: (url: string, alt: string) => void,
        ) {
          const formData = new FormData();
          formData.append('file', file);
          const [res] = await api.uploadImage(formData);
          if (!res) return;
          insertFn(res.data.url, res.data.filename);
        },
      },
    },
    readOnly: false,
    scroll: true,
  } as IEditorConfig,
  mode: 'default',
  toolbarConfig: {
    excludeKeys: ['fullScreen', 'emotion', 'group-video'],
  } as IToolbarConfig,
});

const { editorConfig, mode, toolbarConfig } = toRefs(editorState);

const handleCreated = (editor: IDomEditor) => {
  editorRef.value = editor;
};

const handleEditorChange = (editor: IDomEditor) => {
  if (editor.isFocused()) {
    content.value = editor.getHtml();
  }
};

onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor) {
    editor.destroy();
  }
});
</script>

<template>
  <div :class="props.class" class="focus:border-primary custom-editor border">
    <Toolbar
      :default-config="toolbarConfig"
      :editor="editorRef"
      :mode="mode"
      class="focus:border-primary border-b"
    />
    <Editor
      v-model="content"
      :default-config="editorConfig"
      :mode="mode"
      :style="{ height: height ?? '500px', overflowY: 'hidden' }"
      @on-change="handleEditorChange"
      @on-created="handleCreated"
    />
  </div>
</template>

<style scoped>
.custom-editor {
  :deep(.w-e-text-container) {
    padding: 0 0 0 12px !important;
  }
}

/* wangEditor Modal */
:global(.w-e-modal) {
  position: fixed;
  top: 50% !important;
  left: 50% !important;
  transform: translate(-50%, -50%);
}
</style>
