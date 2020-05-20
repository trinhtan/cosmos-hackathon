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
        @click="openDialogChooseToke(order)"
        >Payment</el-button
      >
    </div>
    <el-dialog title="Choose to pay by token" :visible.sync="dialogChooseToken" width="30%">
      <div class="choose-token">
        <el-radio-group v-model="tabPosition">
          <el-radio-button label="producttoken">
            <div id="logo"></div>
            <p>Sun</p>
          </el-radio-button>
          <el-radio-button label="transfer/ruahosxkxc/uatom">
            <div id="logoATOM"></div>
            <p>Atom</p>
          </el-radio-button>
        </el-radio-group>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogChooseToken = false">Cancel</el-button>
        <el-button type="primary" @click="paymentSell">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import axios from 'axios';
export default {
  name: 'orders-product',
  data() {
    return {
      loadingUpload: false,
      decided: false,
      dialogChooseToken: false,
      orderDecide: {},
      tabPosition: 'producttoken'
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail', 'listOrdersOfSell', 'address', 'balance'])
  },
  methods: {
    ...mapActions('cosmos', [
      'getOrderOfSell',
      'decideSell',
      'signTxt',
      'getDetailProduct',
      'payReservation',
      'setCosmosAccount',
      'payReservationByAtom'
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
        await this.open(resDecideSell, 'decide');

        this.dialogOrder = false;
      });
    },
    paymentSellBySunToken(order) {
      this.$confirm('You will payment this order. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'success',
        center: true
      }).then(async () => {
        this.loadingUpload = true;
        let resPaymentSell = await this.payReservation(order.reservationID);

        this.loadingUpload = false;
        await this.open(resPaymentSell, 'payment');

        this.dialogOrder = false;
      });
    },
    paymentSellByATOM(order) {
      this.$confirm('You will payment this order. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'success',
        center: true
      }).then(async () => {
        this.loadingUpload = true;

        let resPaymentSellATOM = await this.payReservationByAtom(order.reservationID);

        this.loadingUpload = false;
        await this.open(resPaymentSellATOM, 'payment');

        this.dialogOrder = false;
      });
    },
    openDialogChooseToke(order) {
      this.orderDecide = order;
      this.dialogChooseToken = true;
    },
    async paymentSell() {
      if (this.tabPosition === 'producttoken') {
        let checkAmount = false;
        for (let i = 0; i < this.balance.length; i++) {
          const token = this.balance[i];
          if (
            token.denom === 'producttoken' &&
            parseInt(token.amount) >= parseInt(this.orderDecide.price[0].amount)
          ) {
            checkAmount = true;
          }
        }
        if (checkAmount) {
          await this.paymentSellBySunToken(this.orderDecide);
        } else {
          this.$message({
            type: 'error',
            message: 'Not enough token'
          });
        }
      } else if (this.tabPosition === 'transfer/ruahosxkxc/uatom') {
        this.loadingUpload = true;
        let respone = await axios.get(`https://api.coincap.io/v2/assets/cosmos`);
        this.loadingUpload = false;
        let price = respone.data.data.priceUsd;
        let amountPayment = (this.orderDecide.price[0].amount / price) * Math.pow(10, 6);
        let checkAmountAtom = false;
        for (let i = 0; i < this.balance.length; i++) {
          const token = this.balance[i];
          console.log(token);

          if (
            token.denom === 'transfer/ruahosxkxc/uatom' &&
            parseInt(token.amount) >= parseInt(amountPayment)
          ) {
            checkAmountAtom = true;
          }
        }
        console.log(checkAmountAtom);

        if (checkAmountAtom) {
          await this.paymentSellByATOM(this.orderDecide);
        } else {
          this.$message({
            type: 'error',
            message: 'Not enough token'
          });
        }
      }
    },
    open(reponseTxt, functionName) {
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
            message: `${functionName} success`
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

#logoATOM {
  width: 55px;
  height: 55px;
  background: #fff;
  border-radius: 50%;
  cursor: pointer;
  background-image: url('../../assets/images/cosmos.png');
  background-repeat: no-repeat;
  background-size: 100% 100%;
  box-shadow: 4px 8px 16px rgba(0, 0, 0, 0.3);
}

.choose-token {
  text-align: center;
  margin: 0 auto;
}
</style>
