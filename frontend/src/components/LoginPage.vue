<template>
  <div class="flex justify-center items-center min-h-screen flex-grow bg-gradient-to-r from-indigo-100 to-teal-100 font-poppins">
    <!-- login success notify message -->
    <div v-if="successMessage" 
         class="fixed top-4 right-4 bg-green-500 text-white px-6 py-3 rounded-lg shadow-md">
      {{ successMessage }}
    </div>

    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <div class="text-2xl font-bold text-center text-indigo-600 mb-4">ExpoVerse</div>
      <h2 class="text-2xl font-semibold text-center text-indigo-600 mb-6">{{ isLogin ? '登入' : '註冊' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div v-if="!isLogin" class="mb-4">
          <label for="name" class="block text-gray-700 mb-2">姓名</label>
          <input type="text" id="name" v-model="formData.name" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div class="mb-4">
          <label for="email" class="block text-gray-700 mb-2">電子郵件</label>
          <input type="email" id="email" v-model="formData.email" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div class="mb-4">
          <label for="password" class="block text-gray-700 mb-2">密碼</label>
          <input type="password" id="password" v-model="formData.password" required
                 class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
        </div>
        <div v-if="!isLogin" class="mb-6">
          <label for="confirmPassword" class="block text-gray-700 mb-2">確認密碼</label>
          <input type="password" id="confirmPassword" v-model="formData.confirmPassword" required
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
      <!-- show error message -->
      <div v-if="error" class="mt-4 text-center text-red-500">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '@/composables/useAuth';
import type { LoginForm, RegisterForm } from '@/types/auth';

const router = useRouter();
const { login, register, loading, error } = useAuth();
const successMessage = ref('');

const isLogin = ref(true);
const formData = ref<RegisterForm>({
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
});

const resetForm = () => {
  formData.value = {
    name: '',
    email: '',
    password: '',
    confirmPassword: ''
  };
};

const toggleForm = () => {
  isLogin.value = !isLogin.value;
  resetForm();
};

const validate = (): boolean => {
  if (!isLogin.value && formData.value.password !== formData.value.confirmPassword) {
    error.value = '密碼與確認密碼不符';
    return false;
  }
  return true;
};

const handleSubmit = async () => {
  if (!validate()) return;

  try {
    if (isLogin.value) {
      const loginData: LoginForm = {
        email: formData.value.email,
        password: formData.value.password
      };
      await login(loginData);
      router.push('/');
    } else {
      await register(formData.value);
      isLogin.value = true;
      resetForm();
      successMessage.value = '註冊成功，請登入';
      setTimeout(() => {
        successMessage.value = '';
      }, 3000);
    }
  } catch (err) {
    console.error('Authentication failed:', err);
  }
};
</script>

<style scoped>
/* 如果需要額外的樣式可以在這裡添加 */
</style>