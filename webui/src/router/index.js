import {createRouter, createWebHashHistory} from 'vue-router'
import UserPageView from '../views/UserPageView.vue'
import HomePageView from '../views/HomePageView.vue'
import LoginView from '../views/LoginView.vue'
import ErrorPageView from '../views/ErrorPageView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', redirect: '/login' },
		{path: '/login', component: LoginView},
		{path: '/users/:id', component: UserPageView},
		{path: '/users/:id/feeds', component: HomePageView},
		{path: '/error', component: ErrorPageView},
	]
})

export default router
