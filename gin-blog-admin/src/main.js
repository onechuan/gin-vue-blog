// 重置样式
import '@/styles/reset.css'
import '@/styles/global.scss'
import 'uno.css'
import 'virtual:svg-icons-register' // vite-plugin-svg-icons

import { createApp } from 'vue'
import { setupRouter } from '@/router'
import { setupStore } from '@/store'
import App from './App.vue'

async function setupApp() {
  const app = createApp(App)
  setupStore(app)
  await setupRouter(app)
  app.mount('#app')
}

setupApp()