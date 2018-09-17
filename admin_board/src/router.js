import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Users from './views/Users.vue'
import Plans from './views/Plans.vue'
import PlanDetail from './views/PlanDetail.vue'


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
    },
    {
      path: '/plans',
      name: 'plans',
      component: Plans
    },
    {
      path: '/users/:id/plans',
      name: 'plans_detail',
      component: PlanDetail
    }
  ]
})
