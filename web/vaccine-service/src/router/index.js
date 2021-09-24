import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/vaccine/:id',
    name: 'product-detail',
    component: () => import('../views/ProductDetail.vue'),
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue'),
  },
  {
    path: '/customer',
    name: 'customer-home',
    component: () => import('../views/CustomerHome.vue'),
  },
  {
    path: '/customer/:id',
    name: 'customer-detail',
    component: () => import('../views/CustomerDetail.vue'),
  },
  {
    path: '/clinic',
    name: 'clinic-home',
    component: () => import('../views/ClinicHome.vue'),
  },
  {
    path: '/clinic/:id',
    name: 'clinic-detail',
    component: () => import('../views/ClinicDetail.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
