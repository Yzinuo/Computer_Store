<template>
  <div class="carousel">
    <div class="carousel-inner">
      <div v-if="carouselData.length > 0" class="carousel-item-wrapper">
        <transition name="fade">
          <div
              v-if="currentIndex < carouselData.length"
              class="carousel-item"
              @click="goToLink(carouselData[currentIndex].linkUrl)"
              @mouseover="stopAutoPlay"
              @mouseout="startAutoPlay"
          >
            <img :src="carouselData[currentIndex].picUrl" alt="Carousel Image" />
          </div>
        </transition>
      </div>
    </div>
    <div class="carousel-dots">
      <span
          v-for="(item, i) in carouselData"
          :key="i"
          :class="{ active: i === currentIndex }"
          @click="changeSlide(i)"
      ></span>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      carouselData: [
        {
          picUrl: 'https://via.placeholder.com/1200x500?text=Slide1',
          linkUrl: 'https://example.com/slide1',
        },
        {
          picUrl: 'https://via.placeholder.com/1200x500?text=Slide2',
          linkUrl: 'https://example.com/slide2',
        },
        {
          picUrl: 'https://via.placeholder.com/1200x500?text=Slide3',
          linkUrl: 'https://example.com/slide3',
        },
      ], // 轮播图数据
      currentIndex: 0, // 当前显示的轮播图索引
      timer: null,     // 定时器
    };
  },
  methods: {
    // async fetchCarouselData() {
    //   try {
    //     const response = await this.$axios.get('/api/carousel');
    //     this.carouselData = response.data;
    //   } catch (error) {
    //     console.error('Failed to fetch carousel data:', error);
    //   }
    // },
    autoPlay() {
      this.currentIndex = (this.currentIndex + 1) % this.carouselData.length;
    },
    startAutoPlay() {
      this.timer = setInterval(this.autoPlay, 2000);
    },
    stopAutoPlay() {
      clearInterval(this.timer);
    },
    changeSlide(index) {
      this.currentIndex = index;
    },
    goToLink(url) {
      window.location.href = url;
    },
  },
  mounted() {
    //this.fetchCarouselData();
    console.log('Carousel mounted');
    this.startAutoPlay();
  },
  beforeUnmount() {
    this.stopAutoPlay();
  },
};
</script>

<style scoped>
.carousel {
  position: relative;
  width: 100%;
  height: 500px;
  overflow: hidden;
}

.carousel-inner {
  width: 100%;
  height: 100%;
}

.carousel-item {
  position: absolute;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.carousel-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.carousel-dots {
  position: absolute;
  bottom: 20px;
  width: 100%;
  display: flex;
  justify-content: center;
}

.carousel-dots span {
  display: inline-block;
  width: 10px;
  height: 10px;
  margin: 0 5px;
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
  cursor: pointer;
}

.carousel-dots span.active {
  background-color: #fff;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>