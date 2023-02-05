import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Comment from './components/Comment.vue'
import LikeCounter from './components/LikeCounter.vue'
import Photo from './components/Photo.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)

const views = {
    LOGIN: "login",
    STREAM: "register",
    PROFILE: "profile"
}

var state = {
    headers: {
        Authorization: null
    },
    username: null,
    current_view: null

}


app.config.globalProperties.$views = views;
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$user_state = reactive(state);

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Comment", Comment);
app.component("LikeCounter", LikeCounter);
app.component("Photo", Photo);


app.use(router)
app.mount('#app')
