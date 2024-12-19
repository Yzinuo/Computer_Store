<template>
  <div class="carousel">
    <div class="bg" ref="bg"
    @mouseover="bgOver($refs.bg)" @mousemove="bgMove($refs.bg,$event)" @mouseout="bgOut($refs.bg)" >
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
      carouselData: [], // 轮播图数据
      currentIndex: 0, // 当前显示的轮播图索引
      timer: null,     // 定时器
      bgOpt: {
        px: 0,
        py: 0,
        w: 0,
        h: 0,
      },
    };
  },
  methods: {
    async fetchCarouselData() {
      try {
        const response = await this.$axios.get('/api/front/carousel');
        if (response.data.code === 0) { // 假设 Code 为 0 表示成功
          this.carouselData = response.data.data; // 提取 Data 字段
        } else {
          console.error('Failed to fetch carousel data:', response.data.Message);
          // 这里可以根据需要显示错误信息框
          alert(response.data.Message);
        }
      } catch (error) {
        console.error('Failed to fetch carousel data:', error);
        // 这里可以根据需要显示网络错误或其他错误信息
        alert('Network error or server issue. Please try again later.');
      }
    },

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
    bgOver(dom) {
      this.bgOpt = {
        px: dom.offsetLeft,
        py: dom.offsetTop,
        w: dom.offsetWidth,
        h: dom.offsetHeight,
      };
    },
    bgMove(dom, eve) {
      let bgOpt = this.bgOpt;
      let X, Y;
      let mouseX = eve.pageX - bgOpt.px;
      let mouseY = eve.pageY - bgOpt.py;
      if (mouseX > bgOpt.w / 2) {
        X = mouseX - (bgOpt.w / 2);
      } else {
        X = mouseX - (bgOpt.w / 2);
      }
      if (mouseY > bgOpt.h / 2) {
        Y = bgOpt.h / 2 - mouseY;
      } else {
        Y = bgOpt.h / 2 - mouseY;
      }
      dom.style['transform'] = `rotateY(${X / 50}deg) rotateX(${Y / 50}deg)`;
      dom.style.transform = `rotateY(${X / 50}deg) rotateX(${Y / 50}deg)`;
    },
    bgOut(dom) {
      dom.style.transform = 'rotateY(0deg) rotateX(0deg)';
    }
  },
  mounted() {
    this.fetchCarouselData();
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
  height: 700px;
  overflow: hidden;
}

.bg {
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