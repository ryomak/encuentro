<template>
  <div id="header">
    <b-navbar type="dark" variant="info">
      <b-navbar-brand href="#">管理画面</b-navbar-brand>
        <b-navbar-nav right>
          <router-link v-if="loginStatus" to="/" @click.native="logout" id="logout">
            logout
          </router-link>
        </b-navbar-nav>
    </b-navbar>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
export default {
  name: 'Header',
  computed: {
    ...mapState({
      loginStatus: state => state.login.loginStatus
    })
  },

  methods: {
    ...mapActions('login', {
      logout: 'logout'
    })
  },
  mounted: function() {
    this.$store.dispatch('login/loginCheck')
  },
}
</script>

<style lang="scss">
#header {
  a {
    text-decoration: none;
    color: #fffdf6;
  }
}
</style>
