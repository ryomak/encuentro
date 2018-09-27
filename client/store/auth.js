import store2 from 'store2';
let local_storage = store2.namespace('auth');

export default{
    state:{
        AccessToken:null,
    },
    getters: {
        getAccessToken:state=>()=>{
            state.AccessToken;
        },
    },
    mutations: {
        setToken(state,payload){
            if (payload.AccessToken){
                state.AccessToken = payload.AccessToken;
            }
            local_storage.setAll(state);
        },
        loadToken(state) {
            let data = local_storage.getAll();
            Object.keys(data).forEach(key => {
                if (key in state) {
                    Vue.set(state, key, data[key]);
                }
            });
        },
    },
    actions: {
        login(context, payload) {
            context.commit('setToken', {
                AccessToken: payload.jwt,
            }); 
        },
    }

}