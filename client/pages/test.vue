<template>
  <section class="container">
  </section>
</template>

<script>
export default {
  data(){
    return{
      reports:[],
    }
  },
  mounted(){
    this.$axios.post('/api/v1/login',{auth: {email: "yuiyui1@sample.com", password: "yuiyui"}}).then(res=>{
      console.log(res.data)
    })
  }, 
}
</script>

<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
thead {
    background-color: #3f87a6;
    color: #fff;
}

tbody {
    background-color: #e4f0f5;
}

caption {
    padding: 10px;
    caption-side: bottom;
}

table {
    border-collapse: collapse;
    border: 2px solid rgb(200, 200, 200);
    letter-spacing: 1px;
    font-family: sans-serif;
    font-size: .8rem;
}

td,
th {
    border: 1px solid rgb(190, 190, 190);
    padding: 5px 10px;
}

td {
    text-align: center;
}


</style>

