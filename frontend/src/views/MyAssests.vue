<template>
  <section id="section-myProducts">
    <div class="content-myProducts" v-loading="loading">
      <div class="row justify-content-center">
        <router-link
          :to="`/product/${product.productID}`"
          v-for="(product, index) in myProducts"
          :key="index"
          class="col-12 col-ms-6 col-md-6 col-lg-3 text-decoration-none"
        >
          <div class="col-card-home">
            <el-card>
              <img :src="product.images[0]" class="home-image-product" />
              <div class="home-title-product">
                <h1 class="text-title-product">{{ product.title }}</h1>
                <div class="bottom clearfix text-description-product">
                  <p>{{ product.description }}</p>
                </div>
              </div>
            </el-card>
          </div>
        </router-link>
      </div>
    </div>
  </section>
</template>

<script>
import { mapState, mapActions } from 'vuex';
export default {
  name: 'my-assets',
  data() {
    return {
      loading: true
    };
  },
  computed: {
    ...mapState('cosmos', ['myProducts'])
  },
  methods: {
    ...mapActions('cosmos', ['getMyProduct'])
  },
  async created() {
    await this.getMyProduct();
    this.loading = false;
  }
};
</script>

<style lang="scss">
#section-myProducts {
  margin-top: 120px;
  .content-myProducts {
    min-height: 750px;
    padding: 0px 50px 50px 50px;
    width: 75%;
    margin: 0 auto;
  }
}
</style>
