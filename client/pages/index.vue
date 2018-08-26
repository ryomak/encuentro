<template>
  <div class="all">
    <b-container>
        <h1 class="title">
          encuentro
        </h1>
        <h2 class="subtitle">
          退屈な時間に新たな出会いを
        </h2>
        <div class="links">
          <a
            href="/login"
            target="_blank"
            class="button--white">Start</a>
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
  padding: 0px 0px;
  background-image: url("../assets/image.png");
  background-position: center center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  text-align: center;
  background-color: #151A37;
}



.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 100%;
  font-size: 10vw;
  width: 100%;
  color: white;
  letter-spacing: 1px;
  margin-bottom: 15vh;
}

.subtitle {
  font-weight: 300%;
  font-size: 5vw;
  width: 100%;
  color: white;
  word-spacing: 5px;
}

.links {
  padding-top: 15px;
  padding-bottom: 15px;
  margin-top:5vh;
}
</style>
