<template>
  <div>
    <ul v-for="plan in user_plans" :key="plan.id">
        <li>日付: {{ plan.date }}</li>
        <li>場所: {{ plan.place }}</li>
        <li>飲酒: {{ plan.drink }}</li>
        <li>コース: {{ plan.course }}</li>
        <li>ステータス: {{ plan.status }}</li>
    </ul>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
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
