import axios from 'axios'
import router from '../../router.js'

export default {
  namespaced: true,
  state: {
    users: []
  },

  actions: {
    getUser: ({ commit }) => {
      axios.get('/api/v1/admin/users')
        .then((res) => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setUsers', res.data)
              break;
          }
        }).catch( err => {
          router.push({ path: '/' })
          console.log('err:', err);
      });
    }
  },

  mutations: {
    setUsers: (state, users) => {
      state.users = users
    }
  }
}
