import {createRouter, createWebHashHistory} from 'vue-router'
import UserPageView from '../views/UserPageView.vue'
import HomePageView from '../views/HomePageView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/users/:id', component: UserPageView},
		{path: '/users/:id/feeds', component: HomePageView},
	]
})

export default router
