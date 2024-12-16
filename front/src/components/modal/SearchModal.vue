<script setup>
import { computed, ref } from 'vue'
import { debouncedWatch } from '@vueuse/core'
import { useRouter } from 'vue-router'
import api from '@/api'
import { useAppStore } from '@/store'

const router = useRouter()
const appStore = useAppStore()

const searchFlag = computed({
  get: () => appStore.searchFlag,
  set: val => appStore.setSearchFlag(val),
})

// 搜索关键字
const keyword = ref('')

// 防抖 watch, 节流: throttledWatch
debouncedWatch(
    keyword,
    () => {
      if (keyword.value) {
        handleSearch()
      }
    },
    { debounce: 300 },
)

async function handleSearch() {
  // const resp = await api.searchArticles({ keyword: keyword.value })
  // 跳转到搜索结果页面，并传递搜索结果
  const resp = [
    {
      id: 1,
      imageUrl: 'https://via.placeholder.com/250x200?text=Product1',
      name: 'Product 1',
      description: 'This is the first product',
      price: 99.99,
    },
  ]
  router.push({
    path: '/search-results',
    query: {
      keyword: keyword.value,
      results: JSON.stringify(resp.data),
    },
  })
}
</script>

<template>
  <UModal v-model="searchFlag" :width="600">
    <div class="m-0">
      <div class="mb-4 text-xl font-bold">
        本地搜索
      </div>
      <div>
        <div class="relative rounded-md shadow-sm">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-2">
            <div class="i-mdi:flash text-xl text-yellow" />
          </div>
          <input
              v-model="keyword"
              class="block w-full border-0 rounded-full py-2 pl-8 pr-5 text-gray-900 outline-none ring-1 ring-gray-300 ring-inset placeholder:text-gray-400 focus:ring-2 focus:ring-green-300"
              placeholder="输入文章标题或内容..."
          >
        </div>
      </div>
    </div>
  </UModal>
</template>

<style scoped>
/* 省略文字最多 N 行 */
.ell-4 {
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical;
}
</style>
