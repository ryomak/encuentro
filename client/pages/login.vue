<template>
  <div class="all">
    <b-container>
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
          <b-button 
            class="button--white"
            @click="login">Login</b-button>
          <b-button 
            class="button--white"
            >What is encuentro</b-button>
        </div>
      </b-container>
  </div>

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
        console.log(res.data.jwt)
      })
    }
  }
}
</script>

<style>

.all{
  margin:0px;
  padding:0px;
  background-image: url("../assets/image.png");
  background-size: cover;
  background-size:contain;
  background-repeat: no-repeat;
  background-attachment: scroll;
  background-position: center;
  justify-content: center;
  text-align: center;
}



.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 100%;
  font-size: 10vw;
  width: 100%;
  color: white;
  letter-spacing: 1px;
  margin-top: 5vh;
  margin-bottom: 15vh;
}

.entry-form {
  font-weight: 300%;
  font-size: 5vw;
  width: 100%;
  color: white;
  word-spacing: 5px;
  list-style: none;
  margin-bottom: 5vh;
}

.links {
  padding-top: 15px;
  padding-bottom: 15px;
  margin-top:5vh;
  margin-bottom:8vh;
}

li {
  margin-bottom:2vh;
}
</style>
