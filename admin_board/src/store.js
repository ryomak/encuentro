import Vue from 'vue'
import Vuex from 'vuex'
import axios from '@/assets/axios';

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    auth: {
      email: "",
      password: ""
    },
    users: []
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
    }
  },
  actions: {
    setEmail: ({ commit }, email) => {
      commit('setEmail', email)
    },

    setPassword: ({ commit }, password) => {
      commit('setPassword', password)
    },

    login: ({ state }) => {
      axios.axios.post('api/v1/login',{auth:state.auth})
        .then((res) => {
          const status = res.status
          switch(status) {
            case 201:
              localStorage.setItem('jwt-token',res.data.jwt)
              break;
          }
        })
    },

    getUser: ({ commit, state }) => {
      axios.axios.get('/api/v1/admin/users')
        .then((res) => {
          const status = res.status
          console.log(res)
          switch(status) {
            case 200:
              commit('setUsers', res.data)
              break;
          }
        })
    }

  },
})
