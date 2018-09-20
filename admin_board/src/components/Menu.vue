<template>
  <b-col v-if="loginStatus" id="menu" cols="2">
    <ul>
      <router-link :to="{ name: 'users' }"><li>ユーザー</li></router-link>
      <router-link :to="{ name: 'plans' }"><li>プラン</li></router-link>
    </ul>
  </b-col>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState({
      loginStatus: state => state.login.loginStatus
    })
  }
}
</script>

<style lang="scss">
#menu {
  background-color: white;
  padding: 0;
  font-weight: bold;
  ul {
    li {
      background-color: #494949;
      padding: 15px;
    }
  }
  li:hover {
    opacity:0.9;
  }
  a {
    text-decoration: none;
    color: #fffdf6;
  }
}

</style>
