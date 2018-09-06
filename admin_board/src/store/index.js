import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import login from './modules/login'
import users from './modules/users'

Vue.use(Vuex)
axios.interceptors.request.use(
  config => {
    config.headers['Authorization'] = `Bearer ${localStorage.getItem('jwt-token')}`
    return config
  }
)
axios.defaults.baseURL = 'http://0.0.0.0:3000'

export default new Vuex.Store({
  modules: {
    login,
    users
  }
})
