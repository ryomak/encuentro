import router from '../../router.js'
import axios from '@/assets/axios'

export default {
  namespaced: true,
  state: {
    plans: []
  },

  actions: {
    getPlans: ({ commit }) => {
      axios.axios.get('/api/v1/admin/plans')
        .then(res => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setPlans', res.data)
              break;
          }
      });
    }
  },

  mutations: {
    setPlans: (state, plans) => {
      state.plans = plans
    }
  }
}
