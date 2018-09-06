import axios from 'axios'
import router from '../../router.js'

export default {
  namespaced: true,
  state: {
    auth: {
      email: "",
      password: ""
    },
    loginStatus: false
  },

  actions: {
    setEmail: ({ commit }, email) => {
      commit('setEmail', email)
    },

    setPassword: ({ commit }, password) => {
      commit('setPassword', password)
    },

    login: ({ commit, state }) => {
      axios.post('/api/v1/login', { auth: state.auth })
        .then((res) => {
          const status = res.status
          switch(status) {
            case 201:
              localStorage.setItem('jwt-token', res.data.jwt)
              commit('setStatus', true)
              break;
          }
        }).catch( err => {
        console.log('err:', err);
      })
    },

    logout: ({ commit }) => {
      localStorage.removeItem('jwt-token')
      commit('setStatus', false)
    },

    loginCheck: ({ commit }) => {
      axios.get('/api/v1/ping')
        .then( res => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setStatus', true)
              break;
          }
        }).catch( err => {
          commit('setStatus', false)
          router.push({ path: '/' })
          console.log('err:', err);
      })
    }
  },

  mutations: {
    setEmail: (state, email) => {
      state.auth.email = email
    },

    setPassword: (state, password) => {
      state.auth.password = password
    },

    setStatus: (state, check) => {
      state.loginStatus = check
    }
  }
}
