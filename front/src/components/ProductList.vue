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
      page: 1,
      size: 10,
      loading: false,
      hasMore: true,
    };
  },
  methods: {
    async fetchProducts() {
      if (!this.hasMore || this.loading) return;
      this.loading = true;
      try {
        const response = await this.$axios.get('/api/products', {
          params: { page_num: this.page, page_size: this.size },
        });
        if (response.data.code === 0) { // 假设 Code 为 0 表示成功
          const newProducts = response.data.data || []; // 提取 Data 字段中的 items
          this.products = [...this.products, ...newProducts];
          this.hasMore = newProducts.length === this.size;
          this.page++;
        } else {
          console.error('Failed to fetch products:', response.data.Message);
          // 这里可以根据需要显示错误信息框
          alert(response.data.Message);
        }
      } catch (error) {
        console.error('Failed to fetch products:', error);
        // 这里可以根据需要显示网络错误或其他错误信息
        alert('Network error or server issue. Please try again later.');
      } finally {
        this.loading = false;
      }
    },
    goToProductDetail(productId) {
      this.$router.push({
        name: 'ProductDetail',
        params: { id: productId },
      });
    },
    handleScroll() {
      const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
      if (scrollTop + clientHeight >= scrollHeight - 100) {
        this.fetchProducts();
      }
    },
  },
  mounted() {
    this.fetchProducts();
    window.addEventListener('scroll', this.handleScroll);
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  },
};
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