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

.nav {
  transition: all 0.8s;
  color: #000000;
  background: rgb(252, 252, 252) !important;
}

.nav-fixed {
  transition: all 0.8s;
  color: #000;
  background: rgba(255, 2, 2, 0.36) !important;
  box-shadow: 0 5px 6px -5px rgba(133, 133, 133, 0.6);
  & .menu-btn:hover {
    color: #49b1f5 !important;
  }
}

.menus-item {
  position: relative;
  display: inline-block;
  // margin: 0 0 0 1rem;
  a {
    transition: all 0.2s;
  }
  a::after {
    position: absolute;
    bottom: -5px;
    left: 0;
    z-index: -1;
    width: 0;
    height: 3px;
    background-color: #80c8f8;
    content: "";
    transition: all 0.3s ease-in-out;
  }
  .menu-btn {
    color: #000; /* 将文字颜色改为黑色 */
    cursor: pointer;
    &:hover::after {
      width: 100%;
    }
  }
}

.menus-item:hover .menus-submenu {
  display: block;
}

.menus-submenu {
  position: absolute;
  display: none;
  right: 0;
  width: max-content;
  margin-top: 8px;
  box-shadow: 0 5px 20px -4px rgba(0, 0, 0, 0.5);
  background-color: #fff;
  animation: submenu 0.3s 0.1s ease both;

  &::before {
    position: absolute;
    top: -8px;
    left: 0;
    width: 100%;
    height: 20px;
    content: "";
  }
  a {
    line-height: 2;
    color: #4c4948 !important;
    text-shadow: none;
    display: block;
    padding: 6px 14px;
  }
  a:hover {
    background: #4ab1f4;
  }
}

@keyframes submenu {
  0% {
    opacity: 0;
    filter: alpha(opacity=0);
    transform: translateY(10px);
  }

  100% {
    opacity: 1;
    filter: none;
    transform: translateY(0);
  }
}
</style>
