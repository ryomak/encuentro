<template>
  <div>
    <div v-for="plan in plans" :key="plan.id">
      <PlanBox :plan="plan" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import PlanBox from '../components/PlanBox.vue'

export default {
  components: {
    PlanBox
  },

  computed: {
    ...mapState({
      plans: state => state.plans.plans,
      loginStatus: state => state.login.loginStatus
    })
  },

  mounted: function() {
    if(this.loginStatus){
      this.$store.dispatch('plans/getPlans')
    }
  },

  watch: {
    loginStatus: function () {
      this.$store.dispatch('plans/getPlans')
    }
  }
}


</script>
