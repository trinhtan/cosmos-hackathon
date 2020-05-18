<template>
  <section v-loading.fullscreen.lock="loadingUpload">
    <div class="button-sell">
      <el-button type="danger" @click="dialogFormVisible = true">
        <i class="el-icon-sell" />
        Sell
      </el-button>
    </div>

    <el-dialog
      title="Set the selling price"
      :visible.sync="dialogFormVisible"
      class="dialog-set-sell"
    >
      <el-form :model="form" :rules="rulesSell" ref="rulesSell">
        <el-form-item label="Product Token" prop="price" class="input-price">
          <el-input-number
            v-model="form.price"
            controls-position="right"
            :min="0"
            placeholder="Please input price..."
          ></el-input-number>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="setSell">Confirm</el-button>
      </span>
    </el-dialog>
  </section>
</template>

<script>
import { mapActions, mapState } from 'vuex';
export default {
  name: 'sell-product',
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
      dialogFormVisible: false,
      form: {
        price: 0
      },
      rulesSell: {
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
    ...mapActions('cosmos', ['setPriceSell', 'signTxt', 'getDetailProduct', 'setCosmosAccount']),
    async setSell() {
      await this.setCosmosAccount();
      this.$refs['rulesSell'].validate(async (valid) => {
        if (valid) {
          this.loadingUpload = true;
          let resPriceSell = await this.setPriceSell({
            productID: this.productDetail.productID,
            minPrice: `${this.form.price.toString()}producttoken`
          });

          this.loadingUpload = false;
          await this.open(resPriceSell);

          this.form.price = 0;
          this.dialogFormVisible = false;
        } else {
          return false;
        }
      });
    },
    open(resSetSell) {
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
        var reponseSign = await this.signTxt(resSetSell);
        console.log(reponseSign);
        await setTimeout(async () => {
          await this.getDetailProduct(this.$route.params.productId);
          await this.$message({
            type: 'success',
            message: 'Set sell success'
          });
          this.loadingUpload = false;
          await this.$router.push({ name: 'Home' });
        }, 5000);
      });
    }
  }
};
</script>

<style lang="scss">
.button-sell {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}

.dialog-set-sell {
  .el-dialog {
    width: 30%;
  }
}

.input-price {
  width: 250px;
  margin: 0 auto;
}

@media (max-width: 768px) {
  .dialog-set-sell {
    .el-dialog {
      width: 95%;
    }
  }
}
</style>
