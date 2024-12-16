<script setup>
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useWindowScroll, watchThrottled } from '@vueuse/core';
import { Icon } from '@iconify/vue';

import { convertImgUrl } from '@/utils';
import { useAppStore, useUserStore } from '@/store';
import api from '@/api';

const appStore = useAppStore();
const userStore = useUserStore();
const router = useRouter();
const route = useRoute();

const blogTitle = ref('子诺'); // 定义标题变量
const navClass = ref('nav');
const barShow = ref(true);

// * 监听 y 效果比添加 scroll 监听器效果更好
// * 节流操作, 效果很好
const { y } = useWindowScroll();
const preY = ref(0); // 记录上一次的 y 滚动距离
watchThrottled(
    y,
    () => {
      if (Math.abs(preY.value - y.value) >= 50) {
        // 小幅度滚动不进行操作
        barShow.value = y.value < preY.value;
        navClass.value = y.value > 60 ? 'nav-fixed' : 'nav';
        preY.value = y.value;
      }
    },
    { throttle: 100 }
);

function handleSearchClick() {
  console.log('搜索按钮被点击'); // 调试日志
  appStore.setSearchFlag(true); // 设置 searchFlag 为 true
  console.log('appStore.searchFlag 的值:', appStore.searchFlag); // 调试日志
}

async function logout() {
  await userStore.logout();
  if (route.name === 'User') {
    router.push('/');
  }
  window.$notify.success('退出登录成功!');
}

function goToLogin() {
  router.push('/login'); // 重定向到 /login 页面
}

</script>

<template>
  <!-- 电脑端顶部导航栏 -->
  <Transition name="slide-fade" appear>
    <div v-if="barShow" :class="navClass" class="fixed inset-x-0 top-0 z-11 hidden h-[60px] lg:block">
      <div class="h-full flex items-center justify-between px-9">
        <!-- 左上角标题 -->
        <RouterLink to="/" class="text-xl font-bold blog-title">
          {{ blogTitle }} <!-- 使用定义的 blogTitle -->
        </RouterLink>
        <!-- 右上角菜单 -->
        <div class="flex items-center space-x-4">
          <!-- 搜索 -->
          <div class="menus-item">
            <a class="menu-btn flex items-center" @click="handleSearchClick">
              <Icon icon="mdi:magnify" class="text-xl" />
              <span class="ml-1"> 搜索 </span>
            </a>
          </div>

          <!-- 登录 -->
          <div class="menus-item">
            <a v-if="!userStore.userId" class="menu-btn" @click="goToLogin">
              <div class="flex items-center">
                <Icon icon="ph:user-bold" class="text-xl" />
                <span class="ml-1"> 登录 </span>
              </div>
            </a>
            <template v-else>
              <img :src="convertImgUrl(userStore.avatar)" class="w-8 cursor-pointer rounded-full">
              <ul class="menus-submenu">
                <RouterLink to="/user">
                  <div class="flex items-center">
                    <Icon icon="mdi:account-circle" class="mr-1 text-xl" /> 个人中心
                  </div>
                </RouterLink>
                <a @click="logout">
                  <div class="flex items-center">
                    <Icon icon="mdi:logout" class="mr-1 text-xl" /> 退出登录
                  </div>
                </a>
              </ul>
            </template>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped lang="scss">
.nav {
  transition: all 0.8s ease;
  background-color: #111; /* 设置黑色背景 */
  color: #fff; /* 白色文字 */
  height: 60px; /* 固定高度 */
  display: flex;
  align-items: center;
  justify-content: center; /* 将内容居中 */
  padding: 0 20px; /* 左右内边距 */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* 增加阴影效果 */
  position: relative;
  z-index: 10; /* 确保导航栏位于页面顶部 */
}

.nav-fixed {
  background-color: rgba(0, 0, 0, 0.9); /* 滚动时背景色更深 */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.5); /* 强化阴影效果 */
}

/* 美化 Blog Title */
.blog-title {
  font-size: 1.8rem; /* 大字号 */
  font-weight: bold;
  color: #fff;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.6); /* 文字阴影 */
  letter-spacing: 2px; /* 增加字母间距 */
  transition: color 0.3s ease, transform 0.3s ease;
}

.blog-title:hover {
  color: #49b1f5; /* 浅蓝色 hover 效果 */
  transform: scale(1.05); /* 微妙的放大效果 */
}

/* 菜单项样式 */
.menus-item {
  position: relative;
  display: inline-block;
}

.menus-item a {
  color: #fff;
  font-size: 16px;
  transition: color 0.3s, transform 0.3s;
  padding: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menus-item a:hover {
  color: #49b1f5; /* 浅蓝色 hover 效果 */
  transform: translateY(-2px); /* 微妙的上升效果 */
}

.menus-item a::after {
  content: "";
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 0;
  height: 3px;
  background-color: #49b1f5; /* 下划线 hover 效果 */
  transition: width 0.3s ease;
}

.menus-item a:hover::after {
  width: 100%;
}

/* 子菜单样式 */
.menus-submenu {
  display: none;
  position: absolute;
  right: 0;
  top: 50px; /* 调整子菜单的位置 */
  background-color: #111; /* 黑色背景 */
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.4);
  width: 150px;
  border-radius: 8px;
  animation: submenu 0.3s ease-out;
}

.menus-item:hover .menus-submenu {
  display: block;
}

.menus-submenu a {
  padding: 10px 20px;
  color: #ddd; /* 灰色文字 */
  font-size: 14px;
  display: block;
  text-decoration: none;
  transition: background-color 0.3s;
}

.menus-submenu a:hover {
  background-color: #333; /* 子菜单项的 hover 效果 */
}

@keyframes submenu {
  0% {
    opacity: 0;
    transform: translateY(10px);
  }

  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 搜索框样式 */
.search-box {
  position: fixed;
  top: 80px; /* 距离导航栏 */
  left: 50%;
  transform: translateX(-50%);
  background-color: #fff;
  padding: 10px 20px;
  border-radius: 30px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  z-index: 20;
  width: 300px;
}

.search-box input {
  width: 100%;
  padding: 8px 12px;
  font-size: 16px;
  border: none;
  border-radius: 20px;
  outline: none;
}

.search-box input:focus {
  border: 1px solid #49b1f5;
}

@media (max-width: 1024px) {
  .nav {
    padding: 0 10px; /* 小屏幕时内边距缩小 */
  }
  .menus-item a {
    font-size: 14px; /* 在小屏幕上文字尺寸缩小 */
  }
}

/* 添加动画效果 */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: transform 0.5s ease, opacity 0.5s ease;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}
</style>