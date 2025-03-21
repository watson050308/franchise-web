import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import LoginPage from '../components/LoginPage.vue'
import AppManagement from '../AppManagement.vue'
import ExhibitionPage from '../components/ExhibitionPage.vue'
import BrandIntroduce from '../components/BrandIntroduce.vue'
import ShowBrand from '../components/ShowBrand.vue'
import BrandSupport from '../components/BrandSupport.vue'
import AboutUsPage from '../components/AboutUsPage.vue'

const base = process.env.NODE_ENV === 'production' ? '/franchise-web/' : '/'

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
    path: '/app-management',
    name: 'AppManagement',
    component: AppManagement
  },
  {
    path: '/brandSupport',
    name: 'BrandSupport',
    component: BrandSupport
  },
  {
    path: '/exhibition',
    name: 'ExhibitionPage',
    component: ExhibitionPage
  },
  {
    path: '/vendor',
    name: 'BrandIntroduce',
    component: BrandIntroduce
  },
  {
    path: '/showbrand/:category?',
    name: 'ShowBrand',
    component: ShowBrand
  },
  {
    path: '/about-us',
    name: 'AboutUsPage',
    component: AboutUsPage
  }
]

const router = createRouter({
  history: createWebHistory(base),
  routes
})

export default router