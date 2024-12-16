<template>


  <div id="app" class="h-full w-full flex flex-col">
    <!---顶部导航栏-->
    <Header />
    <!-- 路由视图 -->
    <router-view />
  </div>
</template>

<script>
import Header from './components/layout/Header.vue'
export default {
  name: 'App',
  components: {
    Header,
  },
};
</script>

<style scoped>
/* 全局样式 */
#app {
  font-family: Arial, sans-serif;
  text-align: center;
  margin: 0 auto;
  padding: 20px;
}
</style>