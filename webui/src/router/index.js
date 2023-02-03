import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/login', component: LoginView},
		{path: '/#/', component: LoginView},
		{path: '/stream/:username', component: StreamView},
	]
})

export default router
