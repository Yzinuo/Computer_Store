import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import pinia from './store'; // 导入 Pinia
import axios from 'axios'; // 引入 axios

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.use(router).use(pinia).mount('#app'); // 注册 Pinia