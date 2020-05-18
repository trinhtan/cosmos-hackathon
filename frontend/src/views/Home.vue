<template>
  <div>
    <div id="hero-section">
      <form class="flex-form">
        <input
          type="search"
          placeholder="What assets do you need to find?"
          :value="searchTerm"
          @change="setSearchTerm"
        />
        <input type="submit" value="Search" />
      </form>
    </div>
    <section id="section00">
      <div class="content-home" v-loading="loading">
        <div class="row justify-content-center">
          <router-link
            :to="`/product/${product.productID}`"
            v-for="(product, index) in sellsProducts"
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
  </div>
</template>

<script>
// @ is an alias to /src
import { mapState, mapActions } from 'vuex';
export default {
  name: 'Home',
  data() {
    return {
      searchTerm: '',
      loading: true
    };
  },
  computed: {
    ...mapState('cosmos', ['sellsProducts'])
  },
  components: {},
  methods: {
    ...mapActions('cosmos', ['getSellsProducts']),
    setSearchTerm(e) {
      this.searchTerm = e.target.value;
    }
  },
  async created() {
    await this.getSellsProducts();
    this.loading = false;
  }
};
</script>

<style lang="scss">
// Input search
.flex-form input[type='submit'] {
  border-radius: 0 5px 5px 0;
  background: rgba(14, 96, 167, 0.95);
  border: 1px solid rgba(0, 0, 0, 0);
  color: #fff;
  padding: 0 30px;
  cursor: pointer;
  -webkit-transition: all 0.2s;
  -moz-transition: all 0.2s;
  transition: all 0.2s;
}

.flex-form input[type='submit']:hover {
  background: rgb(0, 82, 158);
  border: 1px solid rgb(0, 82, 158);
}

.flex-form input[type='search'] {
  border-radius: 5px 0 0 5px;
}

.flex-form {
  border-radius: 5px;
  display: -webkit-box;
  display: flex;
  position: relative;
  width: 500px;
  box-shadow: 4px 8px 16px rgba(0, 0, 0, 0.3);
}

.flex-form > * {
  border: 0;
  padding: 0 0 0 10px;
  background: #fff;
  line-height: 50px;
  font-size: 1rem;
  border-radius: 0;
  outline: 0;
  -webkit-appearance: none;
}

input[type='search'] {
  flex-basis: 500px;
}

.home-title-product {
  padding: 15px 0;
}

.home-title-product {
  h1 {
    color: cadetblue;
  }
}

.text-decoration-none {
  text-decoration: none;
}

@media all and (max-width: 800px) {
  .flex-form {
    width: 100%;
  }

  .flex-form input[type='search'] {
    flex-basis: 100%;
  }

  .flex-form > * {
    font-size: 0.9rem;
  }
}

.content-home {
  min-height: 570px;
  padding: 0px 50px 50px 50px;
  width: 75%;
  margin: 0 auto;
}

.col-card-home {
  margin: 20px 0;
  cursor: pointer;
}
.card-home {
  border-radius: 5px;
  padding: 20px;
  box-shadow: 0 10px 16px 0 rgba(0, 0, 0, 0.01), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  img {
    height: 150px;
  }
}

.col-card-home {
  margin: 20px 0;
  cursor: pointer;
}
.card-home {
  border-radius: 5px;
  padding: 20px;
  box-shadow: 0 10px 16px 0 rgba(0, 0, 0, 0.01), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  img {
    height: 150px;
  }
}

.home-image-product {
  height: 260px;
  width: 100%;
}

.text-title-product {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  height: 85px;
}

.text-description-product {
  display: -webkit-box;
  -webkit-line-clamp: 10;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  height: 230px;
}

@media all and (max-width: 768px) {
  .content-home {
    width: 100%;
    margin: 0 auto;
    padding: 30px;
  }
}

@media all and (max-width: 360px) {
  .flex-form {
    display: -webkit-box;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    flex-direction: column;
  }

  .flex-form input[type='search'] {
    flex-basis: 0;
  }
  .flex-form input[type='submit'] {
    border-radius: 5px;
  }

  .flex-form input[type='search'] {
    border-radius: 5px;
  }
}
</style>
