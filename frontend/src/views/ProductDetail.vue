<template>
  <div class="content-detail">
    <div class="box-shadow">
      <div class="row">
        <div class="col-md-6">
          <el-card shadow="always">
            <el-image
              @click="showMultiple(0)"
              class="cursor-pointer img-first"
              :src="imgFirst ? imgFirst : ''"
            >
              <div slot="placeholder" class="image-slot">Loading<span class="dot">...</span></div>
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </el-card>
          <div class="box-image-small">
            <div
              class="el-image is-always-shadow img-small cursor-pointer"
              v-for="(img, index) in imgsPreview"
              :key="index"
              @click="showMultiple(index + 1)"
            >
              <img :src="img" class="el-image__inner" style="object-fit: contain;" />
            </div>
          </div>

          <vue-easy-lightbox
            escDisabled
            moveDisabled
            :visible="visible"
            :imgs="productDetail.images"
            :index="index"
            @hide="handleHide"
          ></vue-easy-lightbox>
        </div>
        <div style="padding:30px" class="col-md-6">
          <div class="product--title">
            <h1>{{ productDetail.title }}</h1>
          </div>
          <h4 class="type">Category: {{ productDetail.category }}</h4>
          <div class="description-content">
            <h3>Description:</h3>
            <p>
              {{ productDetail.description }}
            </p>
          </div>

          <div class="price-content" v-if="productDetail.minPrice ? productDetail.minPrice : ''">
            <h2><span>Price:</span> 125$</h2>
          </div>
        </div>
        <div class="button-buy">
          <el-button type="success">
            <i class="el-icon-shopping-cart-full" />
            Buy
          </el-button>
        </div>
        <div class="button-sell">
          <el-button type="danger" @click="dialogVisible = true">
            <i class="el-icon-sell" />
            Sell
          </el-button>
        </div>
        <div class="button-cancel">
          <el-button type="warning">
            <i class="el-icon-error" />
            Cancel Sell
          </el-button>
        </div>
      </div>

      <div class="orders-product">
        <h2>Orders</h2>
        <div class=" margin-top-20 el-alert el-alert--success is-light">
          <i class="el-alert__icon el-icon-s-order is-big"></i>
          <div class="el-alert__content">
            <h3>Alice</h3>
            <!---->
            <p class="el-alert__description">Price: 100</p>
          </div>
          <el-button type="primary" icon="el-icon-success" class="button-confirm-order" round
            >Sell</el-button
          >
        </div>

        <div class=" margin-top-20 el-alert el-alert--success is-light">
          <i class="el-alert__icon el-icon-s-order is-big"></i>
          <div class="el-alert__content">
            <h3>Alice</h3>
            <!---->
            <p class="el-alert__description">Price: 100</p>
          </div>
          <el-button type="primary" icon="el-icon-success" class="button-confirm-order" round
            >Sell</el-button
          >
        </div>

        <div class=" margin-top-20 el-alert el-alert--success is-light">
          <i class="el-alert__icon el-icon-s-order is-big"></i>
          <div class="el-alert__content">
            <h3>Alice</h3>
            <!---->
            <p class="el-alert__description">Price: 100</p>
          </div>
          <el-button type="primary" icon="el-icon-success" class="button-confirm-order" round
            >Sell</el-button
          >
        </div>
      </div>
    </div>
    <el-dialog title="Tips" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
      <span>This is a message</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import VueEasyLightbox from 'vue-easy-lightbox';
import { mapState, mapActions } from 'vuex';
export default {
  name: 'ProductDetail',
  components: {
    VueEasyLightbox
  },
  data() {
    return {
      imgsPreview: [],
      imgFirst: '',
      visible: false,
      index: 0,
      dialogVisible: false
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail'])
  },
  methods: {
    ...mapActions('cosmos', ['getDetailProduct']),
    showMultiple(index) {
      this.index = index;
      this.show();
    },
    show() {
      this.visible = true;
    },
    handleHide() {
      this.visible = false;
    },
    handleClose(done) {
      this.$confirm('Are you sure to close this dialog?')
        .then(() => {
          done();
        })
        .catch(() => {});
    }
  },
  async created() {
    await this.getDetailProduct(this.$route.params.productId);
    this.imgsPreview = this.productDetail.images.filter((e, i) => i !== 0);
    this.imgFirst = this.productDetail.images[0];
  }
};
</script>
<style lang="scss">
.content-detail {
  padding-top: 170px;
  padding-bottom: 100px;
  width: 60%;
  margin: 0 auto;
  min-height: 900px;
}

.box-card {
  width: 60%;
  margin: 150px auto 150px;
  font-family: serif;
}
.image {
  width: 100%;
  height: 100%;
  display: block;
}
.type {
  color: blue;
  margin: auto;
}
.product--title {
  margin-bottom: 10px;
  color: #008080;
  font-weight: bold;
  text-align: left;
}
p {
  margin: 20px auto 20px;
}

.cursor-pointer {
  cursor: pointer;
}

.el-image.is-always-shadow {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.2);
}

.el-image.img-first {
  width: 100%;
  img {
    width: 100%;
  }
}

.box-image-small {
  .img-small:first-child {
    margin-left: 0;
  }
  .img-small {
    width: 95px;
    height: 95px;
    float: left;
    margin: 5px;
    border-radius: 5px;
  }
}

.description-content {
  margin: 30px 0;
}

.price-content {
  span {
    color: black;
    font-size: 1.2rem !important;
  }
  h2 {
    color: #ff4a4a;
  }
}

.button-buy {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}

.button-sell {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}

.button-cancel {
  width: 100%;
  padding: 0 30px 20px 0;
  button {
    float: right;
    padding: 15px 45px;
    font-size: 1rem;
  }
}
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

.margin-top-20 {
  margin: 20px 0 0 !important;
}

@media (max-width: 768px) {
  .content-detail {
    padding-top: 170px;
    width: 85%;
    margin: 0 auto;
  }
}
</style>
