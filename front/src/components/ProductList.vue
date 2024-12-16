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
      products: [ {
        id: 1,
        imageUrl: 'https://via.placeholder.com/250x200?text=Product1',
        name: 'Product 1',
        description: 'This is the first product',
        price: 99.99,
      },
        {
          id: 2,
          imageUrl: 'https://via.placeholder.com/250x200?text=Product2',
          name: 'Product 2',
          description: 'This is the second product',
          price: 199.99,
        },
        {
          id: 3,
          imageUrl: 'https://via.placeholder.com/250x200?text=Product3',
          name: 'Product 3',
          description: 'This is the third product',
          price: 299.99,
        },
      ],
      page: 1,
      size: 10,
      loading: false,
      hasMore: true,
    };
  },
  methods: {
    // async fetchProducts() {
    //   if (!this.hasMore || this.loading) return;
    //   this.loading = true;
    //   try {
    //     const response = await this.$axios.get('/api/products', {
    //       params: { page: this.page, size: this.size },
    //     });
    //     this.products = [...this.products, ...response.data.items];
    //     this.hasMore = response.data.items.length === this.size;
    //     this.page++;
    //   } catch (error) {
    //     console.error('Failed to fetch products:', error);
    //   } finally {
    //     this.loading = false;
    //   }
    // },
    goToProductDetail(productId) {
      this.$router.push({
        name: 'ProductDetail',
        params: { id: productId },
      });
    },
    handleScroll() {
      const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
      if (scrollTop + clientHeight >= scrollHeight - 100) {
        // this.fetchProducts();
      }
    },
  },
  mounted() {
    // this.fetchProducts();
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