import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'
var config={
    headers:{
        token: null 
        }
    };
    

var username={
    username: null
}
const app = createApp(App)
app.config.globalProperties.$username= username;
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$config = config;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')
