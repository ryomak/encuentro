import router from '../../router.js'
import axios from '@/assets/axios'

export default {
  namespaced: true,
  state: {
    users: []
  },

  actions: {
    getUser: ({ commit }) => {
      axios.axios.get('/api/v1/admin/users')
        .then(res => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setUsers', res.data)
              break;
          }
      });
    }
  },

  mutations: {
    setUsers: (state, users) => {
      state.users = users
    }
  }
}
