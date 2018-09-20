<template>
  <div id="app">
    <Header />
    <b-container fluid>
      <b-row >
        <Menu />
        <b-col>
          <router-view/>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
// @ is an alias to /src
import Header from './components/Header.vue'
import Menu from './components/Menu.vue'
import { mapState } from 'vuex'

export default {
  name: 'app',
  components : {
    Header,
    Menu
  },

  computed: {
    ...mapState({
      loginStatus: state => state.login.loginStatus
    })
  },
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #ece8d9
}
ul {
  list-style: none;
  padding:0;
}
a {
  text-decoration: none;
}
</style>
