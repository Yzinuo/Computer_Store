<template>
  <el-container class="product-detail">
    <!-- Product Details -->
    <el-row :gutter="20" class="main-section">
      <!-- Left Side: Main Image and Thumbnail Gallery -->
      <el-col :xs="24" :md="12">
        <div class="image-gallery">
          <el-image :src="mainImage" alt="Main Product Image" fit="cover" class="main-image" />
        </div>
      </el-col>

      <!-- Right Side: Product Details -->
      <el-col :xs="24" :md="12">
        <div class="product-info">
          <h1 class="product-name">{{ product.name }}</h1>
          <p class="product-price">¥{{ product.price }}</p>
          <p class="product-description">{{ product.description }}</p>
          <el-divider />
          <!-- Configurations -->
          <ul class="configurations">
            <li>
              <Icon :icon="'mdi:cpu-64-bit'" class="icon" />
              <span>Processor: {{ productDetails.processor }}</span>
            </li>
            <li>
              <Icon :icon="'mdi:memory'" class="icon" />
              <span>Memory: {{ productDetails.memory }}</span>
            </li>
            <li>
              <Icon :icon="'mdi:harddisk'" class="icon" />
              <span>Hard Disk: {{ productDetails.hard_disk }}</span>
            </li>
            <li>
              <Icon :icon="'mdi:gpu'" class="icon" />
              <span>Graphics Card: {{ productDetails.graphics_card }}</span>
            </li>
          </ul>
          <el-button type="primary" @click="goToProductDetail(2)">Go to Product 2</el-button>
        </div>
      </el-col>
    </el-row>

    <!-- Below Main Section: Introduction Images -->
    <el-row class="product-introduction">
      <h2>Product Introduction</h2>
      <el-col
          v-for="(img, index) in introImages"
          :key="index"
          :xs="24"
          :md="8"
          class="intro-image"
      >
        <img :src="img.image_url" fit="cover" />
      </el-col>
    </el-row>
  </el-container>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Icon } from "@iconify/vue";
import api from "@/api"
const route = useRoute();
const router = useRouter();

const product = ref({});
const productDetails = ref({});
const mainImage = ref("");
const introImages = ref([]);

const fetchProductData = async () => {
  const productId = route.params.id || 1; // Default to product 1 if no ID is passed

  try {
    // Fetch product data
    const productRes = await api.getProductRe(productId);
    product.value = productRes.data;
    mainImage.value = productRes.data.image;

    // Fetch product details
    const detailsRes = await api.getProductDetail(productId);
    productDetails.value = detailsRes.data;

    const imagesRes = await api.getProductImages(productId);
    introImages.value = imagesRes.data;
  } catch (error) {
    console.error("Error fetching product data:", error);
  }
};

const setMainImage = (img) => {
  mainImage.value = img;
};

const goToProductDetail = (productId) => {
  router.push({
    name: "ProductDetail",
    params: { id: productId },
  });
};

onMounted(() => {
  fetchProductData();
});
</script>

<style lang="scss" scoped>
.product-detail {
  margin: 2rem auto;
  max-width: 1200px;

  .main-section {
    display: flex;
    align-items: center;

    .image-gallery {
      .main-image {
        width: 100%;
        height: 400px;
        border-radius: 8px;
        object-fit: cover;
      }

    }

    .product-info {
      .product-name {
        font-size: 1.8rem;
        font-weight: bold;
      }
      .product-price {
        font-size: 1.5rem;
        color: #e74c3c;
        margin: 1rem 0;
      }
      .product-description {
        margin-bottom: 1rem;
        color: #777;
      }
      .configurations {
        list-style: none;
        padding: 0;

        li {
          display: flex;
          align-items: center;
          margin-bottom: 0.75rem;
          font-size: 1rem;

          .icon {
            margin-right: 10px;
            color: #409eff;
            font-size: 1.5rem;
          }
        }
      }
    }
  }

  .product-introduction {
    margin-top: 3rem;
    max-width: 1200px; /* 确保父容器宽度足够 */

    h2 {
      font-size: 1.5rem;
      margin-bottom: 1.5rem;
    }

    .intro-image {
      margin-bottom: 0rem;

      img {
        width: 100%; /* 强制设置宽度 */
        height: auto; /* 强制设置高度 */
        object-fit: cover ; /* 强制设置图片填充方式 */
        display: block;
      }
    }
  }
}
</style>
