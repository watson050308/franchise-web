import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import LoginPage from '../components/LoginPage.vue'
import ExhibitionPage from '../components/ExhibitionPage.vue'

const routes = [
  {
    path: '/',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/login',
    name: 'LoginPage',
    component: LoginPage
  },
  {
    path: '/exhibition',
    name: 'ExhibitionPage',
    component: ExhibitionPage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router