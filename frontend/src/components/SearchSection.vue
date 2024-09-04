<template>
  <section class="text-center mb-16">
    <!-- <h1 class="text-3xl sm:text-4xl font-bold mb-4">找到你的理想加盟機會</h1>
    <p class="text-lg sm:text-xl text-gray-600 mb-8">探索數千個優質加盟品牌，開啟你的創業之旅</p>
    <div class="flex flex-col sm:flex-row justify-center items-center">
      <div class="relative w-full sm:w-auto mb-4 sm:mb-0">
        <input
          type="text"
          placeholder="搜索加盟品牌或行業..."
          class="pl-10 pr-4 py-2 w-full sm:w-64 md:w-96 border rounded-full focus:outline-none focus:ring-2 focus:ring-blue-600"
        />
        <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size="20" />
      </div>
      <BaseButton class="w-full sm:w-auto sm:ml-2 bg-indigo-600 text-white px-6 py-3 rounded-full font-bold hover:bg-teal-400 transition-colors duration-300">
        搜索
      </BaseButton>
    </div> -->
    <SearchBar @search="handleSearch" />
  </section>
</template>
  
  <script>
  // import { Search } from 'lucide-vue-next'
  // import BaseButton from './ui/BaseButton.vue'
  import SearchBar from './ui/SearchBar.vue';
  
  export default {
    name: 'SearchSection',
    components: {
      // Search,
      SearchBar,
      // BaseButton
    }
  }
  </script>