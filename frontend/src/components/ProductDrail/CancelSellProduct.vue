<template>
  <section v-loading.fullscreen.lock="loadingUpload">
    <div class="button-cancel">
      <el-button type="warning" @click="confirmCancel">
        <i class="el-icon-error" />
        Cancel Sell
      </el-button>
    </div>
  </section>
</template>

<script>
import { mapActions, mapState } from 'vuex';
export default {
  name: 'cancel-sell-product',
  data() {
    return {
      loadingUpload: false
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail'])
  },
  methods: {
    ...mapActions('cosmos', ['deleteSell', 'signTxt', 'getDetailProduct', 'setCosmosAccount']),
    async cancelSell() {
      await this.setCosmosAccount();
      this.loadingUpload = true;
      let resDeleteSell = await this.deleteSell({
        productID: this.productDetail.productID,
        sellID: this.productDetail.sellID
      });

      this.loadingUpload = false;
      await this.open(resDeleteSell);

      this.dialogOrder = false;
    },
    confirmCancel() {
      this.$confirm('You will cancel the product sell order. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true
      }).then(async () => {
        await this.cancelSell();
      });
    },

    open(resDeleteSell) {
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
        var reponseSign = await this.signTxt(resDeleteSell);
        console.log(reponseSign);
        await setTimeout(async () => {
          await this.getDetailProduct(this.$route.params.productId);
          await this.$message({
            type: 'success',
            message: 'Cancel sell success'
          });
          this.loadingUpload = false;
        }, 5000);
      });
    }
  }
};
</script>

<style lang="scss">
.button-cancel {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}
</style>
