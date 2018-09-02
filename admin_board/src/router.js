import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Users from './views/Users.vue'
import axios from 'axios'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/users',
      name: 'users',
      component: Users
    }
  ]
})
