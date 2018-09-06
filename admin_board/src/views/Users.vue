<template>
  <div>
    <ul v-for="user in users" v-bind:key="user.id">
      <li>名前: {{ user.name }}</li>
      <li>性別: {{ user.sex }}</li>
      <li>メールアドレス: {{ user.email }}</li>
      <li>誕生日: {{ user.birtyday }}</li>
      <li>大学: {{ user.university }}</li>
    </ul>
  </div>
</template>
<script>
import { mapActions, mapState } from 'vuex'

export default {
  computed: {
    ...mapState([
      'users',
      'loginStatus'
    ])
  },
  methods: {
    ...mapActions([
      'getUser'
    ])
  },

  created: function() {
    if(this.loginStatus){
      this.$store.dispatch('getUser')
    }
  },
  
  watch: {
    // この関数は question が変わるごとに実行されます。
    loginStatus: function () {
      this.$store.dispatch('getUser')
    }
  },

}


</script>
