/* eslint-disable no-console */
import getWeb3 from '@/utils/getWeb3';

const state = {
  web3: null,
  account: null,
  balance: 0,
};

const mutations = {
  setWeb3(state, payload) {
    state.web3 = payload.web3;
    state.balance = payload.balance;
    state.account = payload.account;
  },
};

const actions = {
  async setWeb3({ commit }) {
    const web3 = await getWeb3();
    const accounts = await web3.eth.getAccounts();
    window.web3.version.getNetwork((e, netId) => {
      if (netId !== process.env.VUE_APP_NETWORK_ID) {
        alert('Unknown network');
        return;
      }
    });
    if (accounts.length > 0) {
      const account = accounts[0];
      let balance = await web3.eth.getBalance(account);
      balance = parseFloat(web3.utils.fromWei(balance)).toFixed(2);
      commit('setWeb3', { web3: getWeb3, balance, account });
    } else {
      console.log('Account not found');
    }
  },
};

const getters = {};

export const ethereum = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
