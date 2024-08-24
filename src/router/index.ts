import { createRouter, createWebHistory } from 'vue-router'
import HostMapView from '../views/HostMapView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'HostMapView',
      component: HostMapView
    }
  ]
})

export default router
