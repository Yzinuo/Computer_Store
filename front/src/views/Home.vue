<template>
  <p>什么原因呢</p>
  <div class="home">
    <Carousel />
    <ProductList />
  </div>
</template>

<script>
import Carousel from '../components/Carousel.vue';
import ProductList from '../components/ProductList.vue';

export default {
  components: {
    Carousel,
    ProductList,
  },
};
</script>

<style scoped>
.home {
  width: 1220px;
  margin: 0 auto;
}
</style>