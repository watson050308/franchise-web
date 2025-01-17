<template>
  <div class="flex justify-center items-center min-h-screen flex-grow bg-gradient-to-r from-indigo-100 to-teal-100 font-poppins">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <div class="text-2xl font-bold text-center text-indigo-600 mb-4">ExpoVerse</div>
      <h2 class="text-2xl font-semibold text-center text-indigo-600 mb-6">{{ isLogin ? '登入' : '註冊' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div v-if="!isLogin" class="mb-4">
          <label for="name" class="block text-gray-700 mb-2">姓名</label>
          <input type="text" id="name" v-model="name" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div class="mb-4">
          <label for="email" class="block text-gray-700 mb-2">電子郵件</label>
          <input type="email" id="email" v-model="email" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div class="mb-4">
          <label for="password" class="block text-gray-700 mb-2">密碼</label>
          <input type="password" id="password" v-model="password" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div v-if="!isLogin" class="mb-6">
          <label for="confirmPassword" class="block text-gray-700 mb-2">確認密碼</label>
          <input type="password" id="confirmPassword" v-model="confirmPassword" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <button type="submit"
                class="w-full bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-700 transition duration-300">
          {{ isLogin ? '登入' : '註冊' }}
        </button>
      </form>
      <div class="text-center mt-4">
        <p>
          {{ isLogin ? '還沒有帳號？' : '已有帳號？' }}
          <a href="#" @click.prevent="toggleForm" class="text-indigo-600 hover:text-indigo-800 transition duration-300">
            {{ isLogin ? '立即註冊' : '返回登入' }}
          </a>
        </p>
      </div>
      <div class="mt-6 text-center">
        <p class="text-gray-600 mb-2">或使用以下方式登入</p>
        <div class="space-x-4">
          <a href="#" class="inline-block px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition duration-300">Google</a>
          <a href="#" class="inline-block px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition duration-300">Facebook</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isLogin = ref(true)
const name = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')

const toggleForm = () => {
  isLogin.value = !isLogin.value
  resetForm()
}

const resetForm = () => {
  name.value = ''
  email.value = ''
  password.value = ''
  confirmPassword.value = ''
}

const handleSubmit = () => {
  if (isLogin.value) {
    console.log('登入表單提交', { email: email.value, password: password.value })
    // 在這裡處理登入邏輯
  } else {
    console.log('註冊表單提交', {
      name: name.value,
      email: email.value,
      password: password.value,
      confirmPassword: confirmPassword.value
    })
    // 在這裡處理註冊邏輯
  }
}
</script>