<template>
  <div>
    <ul v-for="user in users" :key="user.id">
      <li>名前: {{ user.name }}</li>
      <li>性別: {{ user.sex }}</li>
      <li>メールアドレス: {{ user.email }}</li>
      <li>誕生日: {{ user.birtyday }}</li>
      <li>大学: {{ user.university }}</li>
    </ul>
  </div>
</template>
<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState({
      users: state => state.users.users,
      loginStatus: state => state.login.loginStatus
    })
  },
  

  mounted: function() {
    if(this.loginStatus){
      this.$store.dispatch('users/getUser')
    }
  },
  watch: {
    loginStatus: function () {
      this.$store.dispatch('users/getUser')
    }
  },

}


</script>
