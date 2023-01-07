import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoggedIn from '../views/LoggedIn.vue'
import UsersView from '../views/UsersView.vue'
import UserView from '../views/UserView.vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/link1', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
		{path: '/session',component:LoggedIn},
		{path: '/users',component:UsersView},
		{path: '/users/:username',component:UserView}

		]
})

export default router
