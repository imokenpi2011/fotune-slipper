import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Main.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/top',
    name: 'Main',
    component: () => import(/* webpackChunkName: "about" */ '../views/Main.vue')
  },
  {
    path: '/result',
    name: 'Result',
    component: () => import(/* webpackChunkName: "about" */ '../views/Result.vue')
  },
  {
    path: '*',
    redirect: '/top'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
