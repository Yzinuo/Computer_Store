<template>
  <div class="product-list">
    <div
        v-for="product in products"
        :key="product.id"
        class="product-item"
        @click="goToProductDetail(product.id)"
    >
      <img :src="product.imageUrl" alt="Product Image" />
      <h3>{{ product.name }}</h3>
      <p>{{ product.description }}</p>
      <p class="price">￥{{ product.price }}</p>
    </div>
    <div v-if="loading" class="loading">Loading...</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      products: [],
      loading: false,
    }
  },
  methods: {
    goToProductDetail(productId) {
      this.$router.push({
        name: 'ProductDetail',
        params: { id: productId },
      })
    },
  },
  mounted() {
    // 从路由参数中获取搜索结果
    const results = this.$route.query.results
    if (results) {
      this.products = JSON.parse(results)
    }
  },
}
</script>

<style scoped>
.product-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
  padding: 20px;
}

.product-item {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 15px;
  cursor: pointer;
  transition: transform 0.2s;
}

.product-item:hover {
  transform: translateY(-5px);
}

.product-item img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
}

.product-item h3 {
  margin: 10px 0;
  font-size: 18px;
}

.product-item p {
  margin: 5px 0;
  font-size: 14px;
  color: #666;
}

.product-item .price {
  font-size: 16px;
  color: #e4393c;
  font-weight: bold;
}

.loading {
  text-align: center;
  margin: 20px 0;
  font-size: 16px;
  color: #999;
}
</style>