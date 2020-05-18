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

          <div class="price-content" v-if="productDetail.selling">
            <h2>
              <span>Min Price:</span> {{ sell.minPrice ? sell.minPrice[0].amount : '' }} -{{
                sell.minPrice ? sell.minPrice[0].denom : ''
              }}
            </h2>
          </div>
        </div>
        <buy-product
          v-if="productDetail.selling && productDetail.owner !== address"
          :minPrice="sell.minPrice ? parseInt(sell.minPrice[0].amount) : 0"
        ></buy-product>
        <sell-product
          v-if="!productDetail.selling && productDetail.owner === address"
        ></sell-product>
        <cancel-sell-product
          v-if="productDetail.selling && productDetail.owner === address"
        ></cancel-sell-product>
      </div>
      <orders-product v-if="productDetail.selling"></orders-product>
    </div>
  </div>
</template>
<script>
import VueEasyLightbox from 'vue-easy-lightbox';
import { mapState, mapActions } from 'vuex';
import SellProduct from '@/components/ProductDrail/SellProduct';
import BuyProduct from '@/components/ProductDrail/BuyProduct';
import CancelSellProduct from '@/components/ProductDrail/CancelSellProduct';
import OrdersProduct from '@/components/ProductDrail/OrdersProduct';
import axios from 'axios';
export default {
  name: 'ProductDetail',
  components: {
    VueEasyLightbox,
    SellProduct,
    BuyProduct,
    CancelSellProduct,
    OrdersProduct
  },
  data() {
    return {
      imgsPreview: [],
      imgFirst: '',
      visible: false,
      index: 0,
      sell: {}
    };
  },
  computed: {
    ...mapState('cosmos', ['productDetail', 'address'])
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
    },
    async getMinPrice() {
      if (this.productDetail.selling) {
        let response = await axios.get(
          `${process.env.VUE_APP_API_BACKEND}/sunchain/sells/${this.productDetail.sellID}`
        );
        this.sell = response.data.result;
      }
    }
  },
  async created() {
    await this.getDetailProduct(this.$route.params.productId);
    if (this.productDetail) {
      this.imgsPreview = this.productDetail.images.filter((e, i) => i !== 0);
      this.imgFirst = this.productDetail.images[0];
    }
    await this.getMinPrice();
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
