import store2 from 'store2';
let local_storage = store2.namespace('auth');

export default function ({ $axios, redirect }) {
    $axios.onRequest(config => {
      console.log('Making request to ' + config.url)
      if (!config.headers['Authorization']){
          const loaded = local_storage.getAll()
        config.headers['Authorization'] =  'Bearer '+ loaded.AccessToken
      }
    })
    $axios.onError(error => {
      if (error.response.status === 401){
        window.location.href = '/login'
      }
    })
  }