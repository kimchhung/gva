import 'vue/jsx'

// Introduce Windi CSS
import '@/plugins/unocss'

// Import a global SVG icon
import '@/plugins/svgIcon'

// Initialize multi -language
import { setupI18n } from '@/plugins/vueI18n'

// Introduction status management
import { setupStore } from '@/store'

// global components
import { setupGlobCom } from '@/components'

// Introduce Element-Plus
import { setupElementPlus } from '@/plugins/elementPlus'

// Introduce the overall style
import '@/styles/index.less'

// Introduce animation
import '@/plugins/animate.css'

// route
import { setupRouter } from './router'

// permissions
import { setupPermission } from './directives'

import { createApp } from 'vue'

import App from './App.vue'

import { setupApi } from './api'
import './permission'

// Create instance
const setupAll = async () => {
  const app = createApp(App)

  await setupApi(app)

  await setupI18n(app)

  setupStore(app)

  setupGlobCom(app)

  setupElementPlus(app)

  setupRouter(app)

  setupPermission(app)

  app.mount('#app')
}

setupAll()
