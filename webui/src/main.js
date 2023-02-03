import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Comment from './components/Comment.vue'
import LikeCounter from './components/LikeCounter.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)


var state = {
    headers: {
        Authorization: null
    },
    username: null
}

app.config.globalProperties.$axios = axios;
app.config.globalProperties.$user_state = state

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Comment", Comment);
app.component("LikeCounter", LikeCounter);


app.use(router)
app.mount('#app')
