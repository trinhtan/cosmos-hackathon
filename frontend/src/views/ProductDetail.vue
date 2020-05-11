<template>
  <div class="content-detail">
    <div class="box-shadow">
      <div class="row">
        <div class="col-md-6">
          <el-card shadow="always">
            <el-image @click="showMultiple(0)" class="cursor-pointer img-first" :src="imgs[0]">
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
            :imgs="imgs"
            :index="index"
            @hide="handleHide"
          ></vue-easy-lightbox>
        </div>
        <div style="padding:30px" class="col-md-6">
          <div class="product--title">
            <h1>My Awesome Product</h1>
          </div>
          <h4 class="type">Category: Food</h4>
          <div class="description-content">
            <h3>Description:</h3>
            <p>
              The definition of a description is a statement that gives details about someone or
              something. An example of description is a story about the places visited on a family
              trip.
            </p>
          </div>

          <div class="price-content">
            <h2><span>Price:</span> 125$</h2>
          </div>
        </div>
        <div class="button-buy">
          <el-button type="primary ">
            <i class="el-icon-shopping-cart-full" />
            Buy
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import VueEasyLightbox from 'vue-easy-lightbox';
export default {
  name: 'ProductDetail',
  components: {
    VueEasyLightbox
  },
  data() {
    return {
      imgs: [
        'https://images.pexels.com/photos/70497/pexels-photo-70497.jpeg',
        'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
        'https://gateway.ipfs.io/ipfs/QmVL4qj7Dqe7THu2dH9ub23v4SDZCWTZR2ur5Mhc7cXaXd/logo-full.png',
        'https://gateway.ipfs.io/ipfs/QmVyiPL1yBr1fQJ5LNbdqzLpkYkdpN7QHHHPZxqQeJMwBc/word.png',
        'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'
      ],
      imgsPreview: [],
      imgFirst: '',
      visible: false,
      index: 0
    };
  },
  methods: {
    showMultiple(index) {
      this.index = index;
      this.show();
    },
    show() {
      this.visible = true;
    },
    handleHide() {
      this.visible = false;
    }
  },
  mounted() {
    this.imgsPreview = this.imgs.filter((e, i) => i !== 0);
    this.imgFirst = this.imgs[0];
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
    width: 100px;
    height: 100px;
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

@media (max-width: 768px) {
  .content-detail {
    padding-top: 170px;
    width: 85%;
    margin: 0 auto;
  }
}
</style>
