<template>
  <section v-loading.fullscreen.lock="loadingUpload">
    <div class="button-buy">
      <el-button type="success" @click="dialogOrder = true">
        <i class="el-icon-shopping-cart-full" />
        Buy
      </el-button>
    </div>
    <el-dialog title="Order product" :visible.sync="dialogOrder" class="dialog-set-sell">
      <el-form :model="form" :rules="rulesBuy" ref="rulesBuy">
        <el-form-item label="Product Token" prop="price" class="input-price">
          <el-input-number
            v-model="form.price"
            controls-position="right"
            :min="minPrice"
            placeholder="Please input price..."
          ></el-input-number>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogOrder = false">Cancel</el-button>
        <el-button type="primary" @click="buyProduct">Confirm</el-button>
      </span>
    </el-dialog>
  </section>
</template>

<script>
import { mapActions, mapState } from 'vuex';
export default {
  name: 'buy-product',
  props: {
    minPrice: Number
  },
  data() {
    var checkPrice = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('Please input the price'));
      } else if (!Number.parseFloat(value)) {
        callback(new Error('Please input digits'));
      } else {
        if (value <= 0) {
          callback(new Error('Price must be greater than 0'));
        } else {
          callback();
        }
      }
    };
    return {
      dialogOrder: false,
      form: {
        price: this.minPrice
      },
      rulesBuy: {
        price: [{ validator: checkPrice, trigger: 'blur' }]
      },
      formLabelWidth: '50px',
      loadingUpload: false
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail'])
  },
  methods: {
    ...mapActions('cosmos', [
      'orderProduct',
      'signTxt',
      'getDetailProduct',
      'getOrderOfSell',
      'setCosmosAccount'
    ]),
    async buyProduct() {
      this.$refs['rulesBuy'].validate(async (valid) => {
        if (valid) {
          this.loadingUpload = true;
          let resorderProduct = await this.orderProduct({
            sellID: this.productDetail.sellID,
            price: `${this.form.price.toString()}producttoken`
          });

          this.loadingUpload = false;
          await this.open(resorderProduct);

          this.form.price = 0;
          this.dialogOrder = false;
        } else {
          return false;
        }
      });
    },
    open(resorderProduct) {
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
        var reponseSign = await this.signTxt(resorderProduct);
        console.log(reponseSign);
        await setTimeout(async () => {
          await this.getDetailProduct(this.$route.params.productId);
          await this.$message({
            type: 'success',
            message: 'Order product success'
          });
          await this.getOrderOfSell(this.productDetail.sellID);
          await this.setCosmosAccount();
          this.loadingUpload = false;
        }, 5000);
      });
    }
  }
};
</script>

<style lang="scss">
.button-buy {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}
</style>
