import { IDomEditor } from '@wangeditor/editor'
import Editor from './src/Editor.vue'

export type EditorExpose = {
  getEditorRef: () => Promise<IDomEditor>
}

export { Editor }
