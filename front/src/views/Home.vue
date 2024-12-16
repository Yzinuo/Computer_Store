<template>
  <div class="home">
    <Carousel />
    <ProductList />
    <span class="flex gap-2 text-gray">
      没有更多商品啦!
      <button class="flex items-center" @click="backTop">
        点击回到顶部 <span class="i-mdi:arrow-up-bold-box ml-1 inline-block text-xl" />
      </button>
    </span>
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