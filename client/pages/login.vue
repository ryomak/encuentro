<template>


  <div class="all">
    <b-container>
        <h1 class="title">
          encuentro
        </h1>
        <ul class="entry-form">
          <li>e-mail
            <input type=text>
          </li>
          <li>password
            <input type=text>
          </li>

        </ul>

        <div class="links">
          <a
              href="https://nuxtjs.org/"
            target="_blank"
            class="button--white">Login</a>
          <a
            href="https://github.com/nuxt/nuxt.js"
            target="_blank"
            class="button--white">What is encuentro?</a>
        </div>
      </b-container>
  </div>

</template>

<script>
import AppLogo from '~/components/AppLogo.vue'

export default {
  components: {
    AppLogo
  }
}
</script>

<style>

.all{
  margin:0px;
  padding:0px;
  background-image: url("../assets/image.png");
  background-repeat: no-repeat;
  background-attachment: scroll;
  background-position: left top;
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
