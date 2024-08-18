<template>
    <section class="text-center mb-16">
      <h1 class="text-4xl font-bold mb-4">找到你的理想加盟機會</h1>
      <p class="text-xl text-gray-600 mb-8">探索數千個優質加盟品牌，開啟你的創業之旅</p>
      <div class="flex justify-center">
        <div class="relative">
          <input
            type="text"
            placeholder="搜索加盟品牌或行業..."
            class="pl-10 pr-4 py-2 w-96 border rounded-full"
          />
          <Search class="absolute left-3 top-2.5 text-gray-400" />
        </div>
        <BaseButton class="ml-2 bg-blue-600 hover:bg-blue-700">
          搜索
        </BaseButton>
      </div>
    </section>
  </template>
  
  <script>
  import { Search } from 'lucide-vue-next'
  import BaseButton from './ui/BaseButton.vue'
  
  export default {
    name: 'SearchSection',
    components: {
      Search,
      BaseButton
    }
  }
  </script>