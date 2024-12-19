<script setup>
import { onMounted, reactive, watch } from 'vue'
import { useRouter } from 'vue-router'
import { ElInput, ElButton, ElCard, ElMessage, ElRow, ElCol, ElForm, ElFormItem } from 'element-plus'

import UploadOne from './UploadOne.vue'
import { useUserStore } from '@/store'
import api from '@/api'

const userStore = useUserStore()
const router = useRouter()

// Form state
const form = reactive({
  avatar: '',
  nickname: '',
  email: '',
  intro: '',
})

// Fetch user data
async function fetchUserData() {
  await userStore.getUserInfo()
  const { nickname, email, avatar, intro } = userStore.userInfo
  form.nickname = nickname
  form.email = email
  form.avatar = avatar
  form.intro = intro
}

onMounted(async () => {
  await fetchUserData()
  if (!userStore.userId) {
    // router.push('/login') // Redirect if not logged in
  }
})

// Update user information
async function updateUserInfo() {
  try {
    await api.updateUser(form)
    ElMessage.success('修改成功!')
    await userStore.updateUserInfo(form)
    await fetchUserData()
  } catch (err) {
    console.error(err)
    ElMessage.error('修改失败，请重试！')
  }
}

// Watch for userInfo updates
watch(
    () => userStore.userInfo,
    (newInfo) => {
      form.nickname = newInfo.nickname
      form.email = newInfo.email
      form.avatar = newInfo.avatar
      form.intro = newInfo.intro
    }
)
</script>

<template>
  <div class="bg-container">
  <ElCard shadow="hover" class="max-w-screen-md mx-auto my-10 rounded-lg border">
    <h2 class="text-2xl font-semibold text-gray-800 mb-6 text-center">个人中心</h2>

    <ElRow :gutter="20" class="items-center">
      <!-- Profile Picture Section -->
      <ElCol :span="8" class="text-center mb-4">
        <UploadOne v-model:preview="form.avatar"/>
      </ElCol>

      <!-- Form Section -->
      <ElCol :span="16">
        <ElForm label-width="90px" :model="form" class="space-y-4">
          <ElFormItem label="昵称">
            <ElInput
                v-model="form.nickname"
                placeholder="请输入昵称"
                size="large"
                class="custom-input"
            />
          </ElFormItem>

          <ElFormItem label="简介">
            <ElInput
                v-model="form.intro"
                placeholder="请输入简介"
                size="large"
                class="custom-input"
            />
          </ElFormItem>

          <ElFormItem label="邮箱">
            <ElInput
                v-model="form.email"
                placeholder="请输入邮箱"
                size="large"
                class="custom-input"
            />
          </ElFormItem>
        </ElForm>

        <div class="text-center mt-6">
          <ElButton
              type="primary"
              @click="updateUserInfo"
              size="large"
              round
              class="custom-button"
          >
            修改
          </ElButton>
        </div>
      </ElCol>
    </ElRow>
  </ElCard></div>
</template>

<style scoped>
.custom-input {
  width: 100%;

}
.bg-container {
  background-image: url('https://gvbresource.oss-cn-hongkong.aliyuncs.com/home.jpg');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: relative;
  z-index: 0;
  min-height: 100vh;
}

.custom-button {
  background-color: #409eff;
  color: #fff;
  border: none;
  padding: 12px 24px;
  font-size: 16px;
  transition: all 0.3s ease;
}

.custom-button:hover {
  background-color: #66b1ff;
  transform: translateY(-2px);
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
}

h2 {
  color: #333;
}
</style>
