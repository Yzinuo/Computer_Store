<template>
  <div id="app">
    <!-- 路由视图 -->
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
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