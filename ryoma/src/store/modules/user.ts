/*
import Vuex from 'vuex'
import {DefineActions, DefineGetters, DefineMutations} from "vuex-type-helper";

export interface UserState {
    users:any
}

export interface UserGetters {
    user:any
}

export interface UserMutations {
    name: {
        name: string
    }
}

export interface Actions {
    
}

const state: UserState = {
    users: []
}

const getters: DefineGetters<UserGetters, UserState> = {
    allUsers: state => state.users;
};

const mutations: DefineMutations<UserMutations, UserState> = {
    setUsers(state:any, users:any) {
        state.users = users;
    }
};

const actions: DefineActions<Actions, UserState, UserMutations, UserGetters> = {
    name({ commit }, payload) {
        commit('name', payload);
    }
};

export const {
    mapState,
    mapGetters,
    mapMutations,
    mapActions
} = Vuex.createNamespacedHelpers<State, Getters, Mutations, Actions>('app');

export const app = {
    namespaced: true,
    state: state,
    getters: getters,
    mutations: mutations,
    actions: actions
};

*/