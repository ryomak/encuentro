<template>
  <div>
    <ul v-for="plan in plans" :key="plan.id">
        <li>日付: {{ plan.date }}</li>
        <li>場所: {{ plan.place }}</li>
        <li>飲酒: {{ plan.drink }}</li>
        <li>コース: {{ plan.course }}</li>
        <li>ステータス: {{ plan.status }}</li>
    </ul>
  </div>
</template>
<script>
import { mapState } from 'vuex'

export default {
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
  },
}


</script>
