<template>
    <div class="min-h-screen bg-gradient-to-r from-purple-500 via-blue-500 to-green-500">
      <!-- 主要內容區 -->
      <div class="container mx-auto px-4 py-8">
        <h1 class="text-4xl font-bold text-white text-center mb-8">展商支援中心</h1>
  
        <!-- 搜索欄 -->
        <div class="mb-8 flex justify-center">
            <div class="relative w-full max-w-xl">
            <input
                type="text"
                placeholder="搜索支援主題..."
                class="w-full p-3 pl-10 rounded-full shadow-lg"
            />
            <Search class="absolute left-3 top-3 text-gray-400" />
            </div>
        </div>
  
        <!-- 支援類別卡片 -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <div 
          v-for="(category, index) in supportCategories" 
          :key="index" 
          class="bg-white rounded-lg shadow-lg p-6 hover:shadow-xl transition duration-300"
        >
          <div class="text-4xl mb-4">{{ category.icon }}</div>
          <h3 class="text-xl font-semibold mb-2">{{ category.title }}</h3>
          <p class="text-gray-600 mb-4">{{ category.description }}</p>
          <button class="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700 transition duration-300 w-full">
            了解更多
          </button>
        </div>
      </div>
  
        <!-- 聯繫支援按鈕 -->
        <div class="mt-12 text-center">
            <button class="bg-green-500 text-white px-6 py-3 rounded-full text-lg font-semibold hover:bg-green-600 transition duration-300 flex items-center justify-center mx-auto">
            <MessageCircle class="mr-2" />
            聯繫客服
            </button>
        </div>
      </div>
    </div>
  </template>
  
<script setup>
import { ref } from 'vue';
import { Search, MessageCircle } from 'lucide-vue-next';

const supportCategories = ref([
	{ title: '開始設置', icon: '🚀', description: '學習如何快速設置您的線上展位' },
	{ title: '管理您的展位', icon: '🛠️', description: '更新和優化您的展位內容' },
	{ title: '與訪客互動', icon: '💬', description: '了解如何與潛在客戶進行有效溝通' },
	{ title: '數據分析', icon: '📊', description: '追蹤和分析您的展位表現' },
]);
</script>
  
<style scoped>
/* 如果需要，您可以在這裡添加任何組件特定的樣式 */
@media (max-width: 640px) {
	.container {
			padding-left: 1rem;
			padding-right: 1rem;
	}
}
</style>