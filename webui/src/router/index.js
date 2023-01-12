import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoggedIn from '../views/LoggedIn.vue'
import UsersView from '../views/UsersView.vue'
import UserView from '../views/UserView.vue'
import PostPhoto from  '../views/PostPhoto.vue'
import LogOut from '../views/LogOut.vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session',component:LoggedIn},
		{path: '/users',component:UsersView},
		{path: '/users/:username',component:UserView},
		{path: '/postPoto', component:PostPhoto},
		{path: '/logout', component:LogOut}
		]
})

export default router

