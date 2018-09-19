<template>
    <b-container class="bg all" fluid>
        <h1 class="title">
          encuentro
        </h1>
        <div class="entry-form">
          Email
          <b-form-input 
            type="text"
            v-model="email"
            placeholder="Enter your email"></b-form-input>
          Password
          <b-form-input 
            type="text"
            v-model="password"
            placeholder="Enter your password"></b-form-input>
        </div>
        <div class="links">
          <ul>
            <li class="button--white"
                @click="login">Login</li>
            <li class="button--white">
                What is encuentro</li>
          </ul>
        </div>
      </b-container>
</template>

<script>
export default {
  data(){
    return{
      email:"",
      password:"",
    }
  },
  methods:{
    login(){
      const conf={
        auth:{
          email:this.email,
          password:this.password,
        }
      }
      this.$axios.$post('/api/v1/login',conf).then(res=>{
        this.$store.dispatch("login",res)
      }).then(res=>{
        window.location.href = "/main"
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.all{
  justify-content: center;
  text-align: center;
}

.bg{
  width: 100%;
  background-color: #151A37;
}


.title {
  display: block;
  font-weight: 100%;
  width: 100%;
  color: white;
  letter-spacing: 1px;
  margin-top: 5vh;
  margin-bottom: 15vh;
}

.entry-form {
  font-weight: 300%;
  width: 100%;
  color: white;
  list-style: none;
  
}

.links{
  
}

</style>
