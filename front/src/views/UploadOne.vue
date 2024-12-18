<script setup>
import { ref, watch } from 'vue'
import api from '@/api'

const props = defineProps({
  preview: {
    type: String,
    required: true,
  },
})
const emit = defineEmits(['update:preview'])

const preview = ref(props.preview)

// Upload new image
async function uploadFile(file) {
  if (!file) return
  const formData = new FormData()
  formData.append('file', file)

  try {
    const response = await api.upload(formData)
    preview.value = response.data.data.url
    emit('update:preview', response.data.data.url)
  } catch (error) {
    console.error('Failed to upload image:', error)
  }
}

// Watch for changes in props
watch(
    () => props.preview,
    (newVal) => {
      preview.value = newVal
    }
)
</script>

<template>
  <div class="relative">
    <!-- Profile Picture -->
    <img
        :src="preview || 'https://via.placeholder.com/150'"
        alt="Avatar"
        class="w-32 h-32 rounded-full object-cover border-2 border-gray-200 shadow-sm"
    />
    <!-- File Input -->
    <label class="absolute bottom-0 right-0 bg-white border rounded-full cursor-pointer shadow-md">
      <input
          type="file"
          class="hidden"
          @change="uploadFile($event.target.files[0])"
      />
      <span class="text-xs px-2 py-1">选择头像</span>
    </label>
  </div>
</template>

<style scoped>
img {
  transition: transform 0.3s ease;
}
img:hover {
  transform: scale(1.05);
}
</style>
