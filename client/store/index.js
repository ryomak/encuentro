import authStore from './auth.js';
import Vuex from 'vuex';
import Vue from 'vue';
Vue.use(Vuex);

const store = () => new Vuex.Store({
    modules:{
        authStore:authStore,
    }
})

export default store;