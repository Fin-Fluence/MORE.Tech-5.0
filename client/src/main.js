import { createApp } from 'vue'
import { createPinia } from 'pinia'

const pinia = createPinia()
import App from './App.vue'
import router from './router'

import "./assets/styles/reset.css";
import "./assets/styles/main.scss";
import "./assets/fonts/stylesheet.css";

createApp(App).use(pinia).use(router).mount('#app')
