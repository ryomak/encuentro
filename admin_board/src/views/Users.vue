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
      'users'
    ])
  },
  methods: {
    ...mapActions([
      'getUser'
    ])
  },
  created: function() {
    this.$store.dispatch('getUser')
  }
  /*
  data(){
    return{
      users:[],
    }
  },
  created:function(){
    this.$axios('/api/v1/admin/users').then(res=>{
      console.log(res)
    })
  }
  */
}


</script>
