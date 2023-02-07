import {createApp, reactive} from 'vue'
import sha256 from 'js-sha256'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Comment from './components/Comment.vue'
import LikeCounter from './components/LikeCounter.vue'
import Photo from './components/Photo.vue'
import PhotoPost from './components/PhotoPost.vue'
import Stream from './components/Stream.vue'
import CommentWriter from './components/CommentWriter.vue'

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

const months = [
    "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"
]

app.config.globalProperties.$views = views;
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$user_state = reactive(state);
app.config.globalProperties.$months = months;

app.config.globalProperties.$hasher = (password) => {
    return sha256(password);
}

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Comment", Comment);
app.component("LikeCounter", LikeCounter);
app.component("Photo", Photo);
app.component("PhotoPost", PhotoPost);
app.component("Stream", Stream);
app.component("CommentWriter", CommentWriter);

app.use(router)
app.mount('#app')
