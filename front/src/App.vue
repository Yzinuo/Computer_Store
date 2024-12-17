<template>
  <UToast ref="messageRef" position="top" align="center" :timeout="3000" closeable />
  <!-- 右上方的消息通知 -->
  <UToast ref="notifyRef" position="top" align="right" :timeout="3000" closeable />

  <div id="app" class="h-full w-full flex flex-col">
    <!---顶部导航栏-->
    <Header />
    <!-- 路由视图 -->
    <router-view />
  </div>

  <GlobalModal />
</template>

<script setup>
import Header from './components/layout/Header.vue'
import GlobalModal from '@/components/modal/index.vue'
import UToast from '@/components/ui/UToast.vue'
import { useAppStore, useUserStore } from '@/store'
import { onMounted, ref } from 'vue'

const messageRef = ref(null)
const notifyRef = ref(null)
const appStore = useAppStore()
const userStore = useUserStore()


onMounted(() => {
  // appStore.getPageList()

  userStore.getUserInfo()

  // 挂载全局提示
  window.$message = messageRef.value
  window.$notify = notifyRef.value
})

</script>

<style scoped>
.spacer {
  height: 80px; /* 根据需要调整高度 */
  width: 100%;
  background: transparent; /* 透明背景 */
  transition: height 0.3s ease; /* 添加过渡效果 */
}

/* 全局样式 */
#app {
  font-family: Arial, sans-serif;
  text-align: center;
  margin: 0 auto;
  padding: 20px;
}
</style>