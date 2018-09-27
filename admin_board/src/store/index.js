import Vue from 'vue'
import Vuex from 'vuex'
import login from './modules/login'
import users from './modules/users'
import plans from './modules/plans'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    login,
    users,
    plans
  }
})

/*
- store,js内では
import axios from "@/assets/axios"
..
axios.axios.get("....")
*.vue内では
import 入らない
this.$axios.get("fdfdf")
//
*/
