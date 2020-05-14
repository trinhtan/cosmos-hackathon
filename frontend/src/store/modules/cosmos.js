/* eslint-disable no-console */
/* eslint-disable no-useless-catch */
/* eslint-disable no-unused-vars */
import axios from 'axios';

const state = {
  address: '',
  balance: '',
  products: []
};

const mutations = {
  setCosmosAccount(state, payload) {
    state.address = payload.address;
    state.balance = payload.balance;
  },
  setAllProducts(state, products) {
    state.products = products;
  }
};

const actions = {
  async setCosmosAccount() {},
  async getAllProducts({ commit }) {
    try {
      let response = await axios.get(`${process.env.VUE_APP_API_BACKEND}/nameservice/products`);
      commit('setAllProducts', response.data.result);
    } catch (error) {
      throw error;
    }
  },
  async createProduct({ commit }, product) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/nameservice/products`, {
        base_req: {
          from: 'cosmos10n87ewwqmxw9dyvlhhhp54upnhq2h2rdglch7h',
          chain_id: 'namechain'
        },
        productID: '04',
        title: product.title,
        description: product.description,
        // category: product.category,
        // images: product.images,
        signer: 'cosmos10n87ewwqmxw9dyvlhhhp54upnhq2h2rdglch7h'
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async signCreateProduct({ commit }, sign) {
    try {
      let respone = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/nameservice/tx/sign`,
        {
          base_req: {
            from: 'cosmos10n87ewwqmxw9dyvlhhhp54upnhq2h2rdglch7h',
            chain_id: 'namechain'
          },
          tx: JSON.stringify(sign.data),
          signer: 'cosmos10n87ewwqmxw9dyvlhhhp54upnhq2h2rdglch7h',
          sequence: '3',
          accountNumber: '3'
        },
        null
      );
      return respone.data.result;
    } catch (error) {
      throw error;
    }
  },
  async broadcastProduct({ commit }, signed) {
    try {
      let respone = await axios.post(`${process.env.VUE_APP_API_BACKEND}/txs`, {
        tx: JSON.parse(signed).value,
        mode: 'block'
      });
      return respone;
    } catch (error) {
      throw error;
    }
  }
};

const getters = {};

export const cosmos = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
