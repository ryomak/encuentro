import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import BootstrapVue from 'bootstrap-vue'
import axios from 'axios'
Vue.use(BootstrapVue)

axios.interceptors.request.use(
	config => {
	  config.headers['Authorization']    = `Bearer ${localStorage.getItem('jwt-token')}`
	  return config
	},
	err =>{
		if (err.response.status == 401){
			redirect("/")
		}
	}
);

Vue.prototype.$axios = axios;

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
