<template>
  <div class="orders-product" v-loading.fullscreen.lock="loadingUpload">
    <h2>Orders</h2>

    <div
      class=" margin-top-20 el-alert el-alert--success is-light"
      v-for="(order, index) in listOrdersOfSell"
      :key="index"
    >
      <i class="el-alert__icon el-icon-s-order is-big color-darkcyan"></i>
      <div class="el-alert__content">
        <h3>
          <span class="price-order">{{ order.price ? order.price[0].amount : '' }}</span>
          -
          <span class="token-order">{{ order.price ? order.price[0].denom : '' }}</span>
        </h3>
        <!---->
        <h4 class="el-alert__description address-order">{{ order.buyer }}</h4>
      </div>
      <el-button
        type="primary"
        icon="el-icon-success"
        class="button-confirm-order"
        v-if="productDetail.owner === address && !order.decide && !decided"
        @click="chooseOrder(order)"
        >Sell</el-button
      >
      <el-button
        type="primary"
        icon="el-icon-success"
        class="button-confirm-order"
        v-if="productDetail.owner !== address && order.decide"
        @click="paymentSell(order)"
        >Payment</el-button
      >
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';
export default {
  name: 'orders-product',
  data() {
    return {
      loadingUpload: false,
      decided: false
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail', 'listOrdersOfSell', 'address'])
  },
  methods: {
    ...mapActions('cosmos', [
      'getOrderOfSell',
      'decideSell',
      'signTxt',
      'getDetailProduct',
      'payReservation',
      'setCosmosAccount'
    ]),
    chooseOrder(order) {
      this.$confirm('You will choose this order. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'success',
        center: true
      }).then(async () => {
        this.loadingUpload = true;
        let resDecideSell = await this.decideSell(order.reservationID);

        this.loadingUpload = false;
        await this.open(resDecideSell);

        this.dialogOrder = false;
      });
    },
    paymentSell(order) {
      this.$confirm('You will payment this order. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'success',
        center: true
      }).then(async () => {
        this.loadingUpload = true;
        let resPaymentSell = await this.payReservation(order.reservationID);

        this.loadingUpload = false;
        await this.open(resPaymentSell);

        this.dialogOrder = false;
      });
    },
    open(reponseTxt) {
      const h = this.$createElement;
      this.$msgbox({
        title: 'Confirm',
        message: h('p', null, [h('span', null, 'Do you want to broadcast the transaction?')]),
        showCancelButton: true,
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        beforeClose: (action, instance, done) => {
          if (action === 'confirm') {
            instance.confirmButtonLoading = true;
            instance.confirmButtonText = 'Loading...';
            done();
            instance.confirmButtonLoading = false;
          } else {
            done();
          }
        }
      }).then(async () => {
        this.loadingUpload = true;
        var reponseSign = await this.signTxt(reponseTxt);
        console.log(reponseSign);
        await setTimeout(async () => {
          await this.getDetailProduct(this.$route.params.productId);
          await this.$message({
            type: 'success',
            message: 'Decide success'
          });
          if (this.productDetail.selling) {
            await this.getOrderOfSell(this.productDetail.sellID);
          }
          await this.setCosmosAccount();
          this.loadingUpload = false;
        }, 5000);
      });
    }
  },
  async created() {
    await this.getDetailProduct(this.$route.params.productId);
    await this.getOrderOfSell(this.productDetail.sellID);
    if (this.listOrdersOfSell.length > 0) {
      for (let i = 0; i < this.listOrdersOfSell.length; i++) {
        if (this.listOrdersOfSell[i].decide) {
          this.decided = true;
          break;
        }
      }
    }
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

.color-darkcyan {
  color: darkcyan;
}

.price-order {
  margin-left: 5px;
  color: crimson;
}

.token-order {
  color: #007cff;
}

.address-order {
  color: #636363 !important;
}
</style>
