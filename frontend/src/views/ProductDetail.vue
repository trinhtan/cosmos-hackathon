<template>
  <div class="container">
    <el-card shadow="always" class="box-card">
      <el-row>
        <el-col :span="12">
          <el-carousel :interval="5000" arrow="never">
            <el-carousel-item v-for="item in 4" :key="item">
              <img
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                class="image"
              />
            </el-carousel-item>
          </el-carousel>
        </el-col>
        <el-col style="padding:20px" :span="12">
          <div class="product--title">
            <h1>My Awesome Product</h1>
          </div>
          <div>
            <h3 class="type">Type: Food</h3>
          </div>
          <p>Description</p>
          <h2>Price: 125$</h2>
          <div class="button">
            <el-button type="primary">Buy</el-button>
          </div>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>
<script>
export default {
  name: 'ProductDetail',
  data() {
    return {};
  }
};
</script>
<style scoped>
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
  color: #008080;
  font-weight: bold;
  font-size: 1.5em;
  text-align: center;
}
p {
  margin: 20px auto 20px;
}

.button {
  padding-top: 60px;
  float: right;
}
</style>
