<template>
  <el-container class="product-detail">
    <!-- Product Details -->
    <el-row :gutter="20" class="main-section">
      <!-- Left Side: Main Image and Thumbnail Gallery -->
      <el-col :xs="24" :md="12" class="image-gallery-col">
        <div class="image-gallery">
          <img :src="mainImage" alt="Main Product Image" fit="cover" class="main-image" />
        </div>
      </el-col>

      <!-- Right Side: Product Details -->
      <el-col :xs="24" :md="12" class="product-info-col">
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
    const productRes = await api.getProductRe({ id: productId });
    product.value = productRes.data;
    mainImage.value = productRes.data.imageUrl;

    // Fetch product details
    const detailsRes = await api.getProductDetail({ id: productId });
    productDetails.value = detailsRes.data;

    // Fetch product images
    const imagesRes = await api.getProductImages({ id: productId });
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
    align-items: flex-start; /* 确保内容从顶部对齐 */

    .image-gallery-col {
      display: flex;
      justify-content: flex-start; /* 确保图片部分左对齐 */
      align-items: flex-start; /* 确保图片部分顶部对齐 */
      flex: 1; /* 确保图片部分占据左侧空间 */
      max-width: 50%; /* 限制图片部分的最大宽度为 50% */

      .image-gallery {
        width: 100%; /* 图片宽度占满父容器 */

        .main-image {
          width: 100%; /* 图片宽度占满父容器 */
          height: 400px; /* 固定高度 */
          max-height: 400px; /* 限制最大高度 */
          border-radius: 8px;
          object-fit: cover; /* 图片填充方式 */
        }
      }
    }

    .product-info-col {
      flex: 1; /* 确保右侧信息部分占据剩余空间 */
      display: flex;
      justify-content: flex-start; /* 确保信息部分左对齐 */
      align-items: flex-start; /* 确保信息部分顶部对齐 */

      .product-info {
        width: 100%; /* 确保信息部分宽度占满父容器 */

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
        object-fit: cover; /* 强制设置图片填充方式 */
        display: block;
      }
    }
  }
}
</style>