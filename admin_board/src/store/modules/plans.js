import axios from '@/assets/axios'

export default {
  namespaced: true,
  state: {
    plans: [],
    user_plans: []
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
    },
    getUserPlans: ({ commit }, id) => {
      axios.axios.get(`/api/v1/admin/users/${id}/user_plans`)
        .then(res => {
          const status = res.status
          switch(status) {
            case 200:
              commit('setUserPlans', res.data)
              break;
          }
      });
    }
  },

  mutations: {
    setPlans: (state, plans) => {
      state.plans = plans
    },
    setUserPlans: (state, plan) => {
      state.user_plans = plan
    }
  }
}
