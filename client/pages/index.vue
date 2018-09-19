<template>
    <b-container class="bg" fluid>
        <h1 class="title">
          encuentro
        </h1>
        <h2 class="subtitle">
          退屈な時間に新たな出会いを
        </h2>
        <div class="links">
          <nuxt-link
            to="/login"
            class="button--white">Start</nuxt-link>
          <a
            href="https://github.com/nuxt/nuxt.js"
            target="_blank"
            class="button--white">What is encuentro?</a>
        </div>
    </b-container>
</template>

<script>

export default {
  components: {

  }
}
</script>

<style scoped>

.bg{
  background-image: url('~/assets/image2.jpg');
  background-color: #151A37;
  background-position: center center;
  width: 100%;
  height: 100%;
}



.title {
  color: white;
}

.subtitle {
  color: white;
}

.links {
  
}
</style>
