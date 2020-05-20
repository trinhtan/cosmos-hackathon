/* eslint-disable no-console */
/* eslint-disable no-useless-catch */
/* eslint-disable no-unused-vars */
import axios from 'axios';

const state = {
  name: 'alice',
  address: '',
  balance: [],
  products: [],
  public_key: '',
  account_number: 0,
  sequence: 0,
  productDetail: {},
  myProducts: [],
  sellsProducts: [],
  listOrdersOfSell: []
};

const mutations = {
  setCosmosAccount(state, payload) {
    state.address = payload.address;
    state.public_key = payload.public_key;
    state.account_number = payload.account_number;
    state.sequence = payload.sequence;
    state.balance = payload.balance;
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
  },
  setOrdersOfSell(state, orders) {
    state.listOrdersOfSell = orders;
  }
};

const actions = {
  async setCosmosAccount({ commit, state }) {
    let respone = await axios.get(
      `${process.env.VUE_APP_API_BACKEND}/sunchain/names/${state.name}/address`
    );
    let account = respone.data.result.value;
    let responeBalance = await axios.get(
      `${process.env.VUE_APP_API_BACKEND}/sunchain/names/${state.name}/balance`
    );
    let balance = JSON.parse(responeBalance.data.result);
    account['balance'] = balance;
    commit('setCosmosAccount', account);
  },
  async getAllProducts({ commit }) {
    try {
      let response = await axios.get(`${process.env.VUE_APP_API_BACKEND}/sunchain/products`);
      let products = response.data.result;
      if (products) {
        for (let index = 0; index < products.length; index++) {
          products[index].images = JSON.parse(products[index].images);
        }
      }

      commit('setAllProducts', products);
    } catch (error) {
      throw error;
    }
  },
  async createProduct({ commit, state }, product) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/sunchain/products`, {
        base_req: {
          from: state.address,
          chain_id: 'band-consumer'
        },
        title: product.asset.title,
        description: product.asset.description,
        category: product.asset.category,
        images: JSON.stringify(product.images)
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async signTxt({ commit, state }, sign) {
    try {
      let respone = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/tx/sign`,
        {
          base_req: {
            from: state.address,
            chain_id: 'band-consumer'
          },
          tx: JSON.stringify(sign.data),
          signer: state.name,
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
      `${process.env.VUE_APP_API_BACKEND}/sunchain/products/${productId}`
    );
    let product = response.data.result;
    product.images = JSON.parse(product.images);
    commit('setDetailProduct', product);
  },
  async getMyProduct({ commit, state, dispatch }) {
    await dispatch('setCosmosAccount');
    try {
      let response = await axios.get(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/names/${state.name}/products`
      );
      let products = response.data.result;
      if (products) {
        for (let index = 0; index < products.length; index++) {
          products[index].images = JSON.parse(products[index].images);
        }
      }

      commit('setMyProduct', products);
    } catch (error) {
      throw error;
    }
  },
  async getSellsProducts({ commit }) {
    try {
      let response = await axios.get(`${process.env.VUE_APP_API_BACKEND}/sunchain/sells`);
      let sellsProducts = response.data.result ? response.data.result : [];
      let listSellProducts = [];
      for (let i = 0; i < sellsProducts.length; i++) {
        const sell = sellsProducts[i];
        let response = await axios.get(
          `${process.env.VUE_APP_API_BACKEND}/sunchain/products/${sell.productID}`
        );
        let product = response.data.result;
        product.images = JSON.parse(product.images);
        listSellProducts.push(product);
      }
      commit('getAllSellsProducts', listSellProducts);
    } catch (error) {
      throw error;
    }
  },
  async setPriceSell({ state }, { productID, minPrice }) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/sunchain/sells`, {
        base_req: {
          from: state.address,
          chain_id: 'band-consumer'
        },
        productID,
        minPrice
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async orderProduct({ state }, { sellID, price }) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/sunchain/reservations`, {
        base_req: {
          from: state.address,
          chain_id: 'band-consumer'
        },
        sellID,
        price
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async deleteSell({ state }, { productID, sellID }) {
    try {
      let response = await axios.post(`${process.env.VUE_APP_API_BACKEND}/sunchain/cancelSell`, {
        base_req: {
          from: state.address,
          chain_id: 'band-consumer'
        },
        productID,
        sellID
      });
      return response;
    } catch (error) {
      throw error;
    }
  },
  async getOrderOfSell({ commit }, sellID) {
    try {
      let response = await axios.get(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/sells/${sellID}/reservations`
      );
      let orders = response.data.result ? response.data.result : [];
      commit('setOrdersOfSell', orders);
    } catch (error) {
      throw error;
    }
  },
  async decideSell({ commit }, reservationID) {
    try {
      let response = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/sells/decideSell`,
        {
          base_req: {
            from: state.address,
            chain_id: 'band-consumer'
          },
          reservationID
        }
      );
      return response;
    } catch (error) {
      throw error;
    }
  },
  async payReservation({ commit }, reservationID) {
    try {
      let response = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/reservations/payReservation`,
        {
          base_req: {
            from: state.address,
            chain_id: 'band-consumer'
          },
          reservationID
        }
      );
      return response;
    } catch (error) {
      throw error;
    }
  },
  async payReservationByAtom({ commit }, reservationID) {
    try {
      let response = await axios.post(
        `${process.env.VUE_APP_API_BACKEND}/sunchain/reservations/payReservationByAtom`,
        {
          base_req: {
            from: state.address,
            chain_id: 'band-consumer'
          },
          reservationID
        }
      );
      return response;
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
