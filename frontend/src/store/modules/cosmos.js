/* eslint-disable no-console */
/* eslint-disable no-useless-catch */
/* eslint-disable no-unused-vars */
import axios from 'axios';

const state = {
  address: '',
  balance: '',
  products: [],
  coins: [],
  public_key: '',
  account_number: 0,
  sequence: 0,
  productDetail: {},
  myProducts: [],
  sellsProducts: []
};

const mutations = {
  setCosmosAccount(state, payload) {
    state.address = payload.address;
    state.coins = payload.coins;
    state.public_key = payload.public_key;
    state.account_number = payload.account_number;
    state.sequence = payload.sequence;
  },
  setAllProducts(state, products) {
    state.products = products;
  },
  setDetailProduct(state, product) {
    state.productDetail = product;
  },
  setMyProduct(state, products) {
    state.myProducts = products;
  },
  getAllSellsProducts(state, sellsProducts) {
    state.sellsProducts = sellsProducts;
  }
};

const actions = {
  async setCosmosAccount({ commit }) {
    let respone = await axios.get(`${process.env.VUE_APP_API_BACKEND}/nameservice/accAddress/jack`);
    let account = respone.data.result.value;
    commit('setCosmosAccount', account);
  },
  async getAllProducts({ commit }) {
    try {
      let response = await axios.get(`${process.env.VUE_APP_API_BACKEND}/nameservice/products`);
      let products = response.data.result;
      for (let index = 0; index < products.length; index++) {
        products[index].images = JSON.parse(products[index].images);
      }
      commit('setAllProducts', products);
    } catch (error) {
      throw error;
    }
  },
  async createProduct({ commit, state }, product) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/nameservice/products`, {
        base_req: {
          from: state.address,
          chain_id: 'namechain'
        },
        title: product.asset.title,
        description: product.asset.description,
        category: product.asset.category,
        images: JSON.stringify(product.images),
        signer: state.address
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async signTxt({ commit, state }, sign) {
    try {
      let respone = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/nameservice/tx/sign`,
        {
          base_req: {
            from: state.address,
            chain_id: 'namechain'
          },
          tx: JSON.stringify(sign.data),
          signer: state.address,
          sequence: state.sequence.toString(),
          accountNumber: state.account_number.toString()
        },
        null
      );
      return respone.data.result;
    } catch (error) {
      throw error;
    }
  },
  async getDetailProduct({ commit }, productId) {
    let response = await axios.get(
      `${process.env.VUE_APP_API_BACKEND}/nameservice/products/${productId}`
    );
    let product = response.data.result;
    product.images = JSON.parse(product.images);
    commit('setDetailProduct', product);
  },
  async getMyProduct({ commit, state, dispatch }) {
    await dispatch('setCosmosAccount');
    try {
      let response = await axios.get(
        `${process.env.VUE_APP_API_BACKEND}/nameservice/accAddress/${state.address}/products`
      );
      let products = response.data.result;
      for (let index = 0; index < products.length; index++) {
        products[index].images = JSON.parse(products[index].images);
      }
      commit('setMyProduct', products);
    } catch (error) {
      throw error;
    }
  },
  async getSellsProducts({ commit }) {
    try {
      let response = await axios.get(`${process.env.VUE_APP_API_BACKEND}/nameservice/sells`);
      let sellsProducts = response.data.result ? response.data.value.msg : [];
      let listSellProducts = [];
      for (let i = 0; i < sellsProducts.length; i++) {
        const sell = sellsProducts[i];
        let response = await axios.get(
          `${process.env.VUE_APP_API_BACKEND}/nameservice/products/${sell.productID}`
        );
        let product = response.data.result;
        product.images = JSON.parse(product.images);
        listSellProducts.push(product);
      }
      commit('getAllSellsProducts', listSellProducts);
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
