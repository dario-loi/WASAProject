import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/login', component: LoginView},
		{path: '/#/', component: LoginView},
		{path: '/stream/:username', component: StreamView},
		{path: '/profile/:username', component: ProfileView},
		// allow GET requests to /photos/... to be handled by the backend
		{path: '/photos/.*', redirect: '/'},
	]
})

export default router
