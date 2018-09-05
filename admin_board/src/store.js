import Vue from 'vue'
import Vuex from 'vuex'
import axios from '@/assets/axios';

Vue.use(Vuex)

axios.interceptors.request.use(config => {
  config.headers['Authorization'] = `Bearer ${localStorage.getItem('jwt-token')}`
  return config
})

export default new Vuex.Store({
  state: {
    auth: {
      email: "",
      password: ""
    },
    users: [],

    loginStatus: false
  },
  mutations: {
    setEmail: (state, email) => {
      state.auth.email = email
    },

    setPassword: (state, password) => {
      state.auth.password = password
    },

    setUsers: (state, users) => {
      state.users = users
    },

    setStatus: (state, check) => {
      state.loginStatus = check
    }
  },
  actions: {
    setEmail: ({ commit }, email) => {
      commit('setEmail', email)
    },

    setPassword: ({ commit }, password) => {
      commit('setPassword', password)
    },

    getUser: ({ commit }) => {
      axios.get('http://0.0.0.0:3000/api/v1/admin/users')
        .then((res) => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setUsers', res.data)
              break;
          }
        })
    },

    login: ({ state }) => {
      axios.axios.post('http://0.0.0.0:3000/api/v1/login',{auth:state.auth})
        .then((res) => {
          const status = res.status
          switch(status) {
            case 201:
              localStorage.setItem('jwt-token',res.data.jwt)
              break;
          }
        })
    },

    logout: () => {
      localStorage.removeItem('jwt-token')
    },

    loginCheck: ({ commit }) => {
      axios.get('http://0.0.0.0:3000/api/v1/ping')
        .then((res) => {
          const status = res.status
          console.log(res)
          switch(status) {
            case 200:
              commit('setUsers', true)
              break;
            case 401:
              commit('setStatus', false)
              break;
          }
        })
    }


  },
})
