<template>
  <div>
    <div v-for="user in users" :key="user.id">
      <UserBox :user="user" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import UserBox from '../components/UserBox.vue'

export default {
  components: {
    UserBox
  },

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
