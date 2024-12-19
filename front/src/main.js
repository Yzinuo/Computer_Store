import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import pinia from './store'; // 导入 Pinia
import axios from 'axios'; // 引入 axios
import 'uno.css';
import ElementPlus from 'element-plus'
import Particles from 'particles.vue3'
// custom style
import './styles/index.css'
import './styles/common.css'
import './styles/animate.css'


const app = createApp(App);
app.use(ElementPlus )
app.use(Particles)
app.config.globalProperties.$axios = axios;
app.use(router).use(pinia).mount('#app'); // 注册 Pinia