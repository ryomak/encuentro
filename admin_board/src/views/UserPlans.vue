<template>
  <div>
    <div v-for="plan in user_plans" :key="plan.id">
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
      user_plans: state => state.plans.user_plans,
    })
  },

  mounted: function() {
    this.$store.dispatch('plans/getUserPlans', this.$route.params.id)
  }
}


</script>
