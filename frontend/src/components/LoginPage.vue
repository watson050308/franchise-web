<!-- LoginPage.vue -->
<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-blue-600 p-4">
    <div class="w-full max-w-md bg-white rounded-lg shadow-md p-8">
      <h2 class="text-3xl font-bold text-center text-gray-800 mb-6">登入 FranchiseConnect</h2>
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <input 
            type="email" 
            v-model="email" 
            placeholder="電子郵件" 
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <input 
            type="password" 
            v-model="password" 
            placeholder="密碼" 
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <button 
          type="submit" 
          class="w-full bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 transition duration-300"
        >
          登入
        </button>
      </form>
      <div class="mt-4 text-center">
        <p class="text-sm text-gray-600">
          還沒有帳號？ 
          <button @click="handleRegister" class="text-blue-600 hover:underline">
            立即註冊
          </button>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const email = ref('')
const password = ref('')

const handleSubmit = () => {
  // 處理登入邏輯
  console.log('Login attempt with:', email.value, password.value)
}

const handleRegister = () => {
  // 處理註冊邏輯
  console.log('Redirecting to registration page')
}
</script>