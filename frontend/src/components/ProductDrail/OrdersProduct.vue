<template>
  <div class="orders-product">
    <h2>Orders</h2>

    <div
      class=" margin-top-20 el-alert el-alert--success is-light"
      v-for="(order, index) in listOrdersOfSell"
      :key="index"
    >
      <i class="el-alert__icon el-icon-s-order is-big"></i>
      <div class="el-alert__content">
        <h3>
          Order Price : {{ order.price ? order.price[0].amount : '' }} -
          {{ order.price ? order.price[0].denom : '' }}
        </h3>
        <!---->
        <h4 class="el-alert__description">{{ order.buyer }}</h4>
      </div>
      <el-button type="primary" icon="el-icon-success" class="button-confirm-order">Sell</el-button>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';
export default {
  name: 'orders-product',
  data() {
    return {};
  },
  computed: {
    ...mapState('cosmos', ['productDetail', 'listOrdersOfSell', 'getDetailProduct'])
  },
  methods: {
    ...mapActions('cosmos', ['getOrderOfSell'])
  },
  async created() {
    // await this.getDetailProduct(this.$route.params.productId);
    await this.getOrderOfSell(this.productDetail.sellID);
  }
};
</script>

<style lang="scss">
.orders-product {
  margin-top: 30px;
  h2 {
    margin-bottom: 20px;
  }
  button.button-confirm-order {
    position: absolute;
    right: 0;
  }
}
</style>
