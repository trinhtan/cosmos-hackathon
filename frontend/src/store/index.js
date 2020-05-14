import Vue from 'vue';
import Vuex from 'vuex';
import { ethereum } from './modules/ethereum';
import { cosmos } from './modules/cosmos';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    ethereum: ethereum,
    cosmos: cosmos
  }
});
