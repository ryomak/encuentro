<template>
  <section class="container">
    <table>
      <thead>
        <tr>
          <th>タイトル</th>
          <th>作者</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="r in reports" :key="r.id">
          <td><a :href="r.url">{{r.title}}</a></td>
          <td>{{r.user.name}}</td>
        </tr>        
      </tbody>
    </table>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  data(){
    return{
      reports:[],
    }
  },
  mounted(){
    axios.get('https://qiita.com/api/v2/items').then(res=>{
      this.reports = res.data
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

