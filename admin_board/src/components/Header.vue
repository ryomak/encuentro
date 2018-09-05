<template>
  <div>
    <b-navbar type="dark" variant="info">
      <b-navbar-brand href="#">管理画面</b-navbar-brand>
        <b-navbar-nav right>
          <router-link v-if="loginStatus" to="/" @click="logout" id="logout">
            LogOut
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
    ...mapState([
      'loginStatus'
    ])
  },

  methods: {
    ...mapActions([
      'logout'
    ])
  },

  created: function() {
    this.$store.dispatch('loginCheck')
  }

}
</script>
