<template>
  <header class="bg-gradient-to-r from-indigo-500 to-teal-400 text-white py-4 px-4 sm:px-6 lg:px-8">
    <nav class="flex flex-col sm:flex-row justify-between items-center max-w-6xl mx-auto">
      <div class="w-full sm:w-auto flex justify-between items-center">
        <router-link to="/expoverse" class="text-2xl font-bold inline-block">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 300 50" width="300" height="50">
          <defs>
            <linearGradient id="grad1" x1="0%" y1="0%" x2="100%" y2="0%">
              <stop offset="0%" style="stop-color:#4a90e2;stop-opacity:1" />
              <stop offset="100%" style="stop-color:#5cb3ff;stop-opacity:1" />
            </linearGradient>
          </defs>
          <rect x="5" y="5" width="60" height="40" rx="20" ry="20" fill="url(#grad1)" />
          <circle cx="20" cy="25" r="8" fill="#ffffff" />
          <circle cx="50" cy="25" r="8" fill="#ffffff" />
          <path d="M20 25 Q35 15 50 25" fill="none" stroke="#ffffff" stroke-width="2" />
          <text x="75" y="22" font-family="Arial, sans-serif" font-size="20" fill="white">展覽宇宙</text>
          <text x="75" y="44" font-family="Arial, sans-serif" font-size="20" fill="white">ExpoVerse</text>
        </svg>
      </router-link>
        <button @click="toggleMenu" class="sm:hidden">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
          </svg>
        </button>
      </div>
      
      <div :class="['sm:flex items-center space-y-4 sm:space-y-0 sm:space-x-4 mt-4 sm:mt-0 w-full sm:w-auto', {'hidden': !isMenuOpen, 'flex flex-col sm:flex-row': isMenuOpen}]">
        <router-link to="/expoverse" class="hover:opacity-80 transition-opacity duration-300 block w-full sm:w-auto text-center">首頁</router-link>
        <router-link to="/exhibition" class="hover:opacity-80 transition-opacity duration-300 block w-full sm:w-auto text-center">尋找展館</router-link>
        <router-link to="/brandSupport" class="hover:opacity-80 transition-opacity duration-300 block w-full sm:w-auto text-center">展商支援</router-link>
        <router-link to="/about-us" class="hover:opacity-80 transition-opacity duration-300 block w-full sm:w-auto text-center">關於我們</router-link>
        <router-link to="/login"><BaseButton variant="outline" class="text-white border-white hover:opacity-80 transition-opacity duration-300 w-full sm:w-auto">
          <span class="hover:underline">登入</span>
        </BaseButton></router-link>
        
        <div class="relative flex items-center" @mouseenter="showSearch = true" @mouseleave="hideSearchIfNotFocused">
          <button class="text-white hover:opacity-80 transition-opacity duration-300">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </button>
          <div 
          class="overflow-hidden transition-all duration-300 ease-in-out flex items-center"
            :class="showSearch ? 'w-64 opacity-100 px-2' : 'w-0 opacity-0 px-0'"
          >
            <input
              type="text"
              v-model="searchQuery"
              @keyup.enter="performSearch"
              @focus="showSuggestions = true"
              @blur="hideSuggestionsDelayed"
              placeholder="搜索..."
              class="w-full px-4 py-2 text-gray-700 bg-white rounded-full focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
            <div v-if="showSuggestions && filteredBrands.length > 0" class="absolute top-full left-0 mt-2 w-64 bg-white rounded-lg shadow-lg overflow-hidden z-50">
              <ul>
                <li 
                  v-for="brand in filteredBrands" 
                  :key="brand.name" 
                  @mousedown="selectBrand(brand.name)"
                  class="px-4 py-2 hover:bg-gray-100 cursor-pointer text-gray-700 flex items-center"
                >
                  <img :src="brand.logo" :alt="brand.name" class="w-8 h-8 mr-3 object-contain" />
                  <span>{{ brand.name }}</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
import { ref, computed } from 'vue';
import BaseButton from './ui/BaseButton.vue';

export default {
  name: 'BaseHeader',
  components: {
    BaseButton
  },
  setup() {
    const isMenuOpen = ref(false);
    const searchQuery = ref('');
    const showSearch = ref(false);
    const showSuggestions = ref(false);

    // 品牌列表，包含名稱和 logo URL
    const brands = [
      { name: '紅橘子', logo: '../assets/category/redorange.png' },
      { name: '紅吉子', logo: '../assets/category/redorange.png' },
      { name: '紅柿子', logo: '../assets/category/redorange.png' },
    ];

    const filteredBrands = computed(() => {
      return brands.filter(brand => 
        brand.name.toLowerCase().includes(searchQuery.value.toLowerCase())
      );
    });

    const toggleMenu = () => {
      isMenuOpen.value = !isMenuOpen.value;
      if (!isMenuOpen.value) {
        showSearch.value = false;
      }
    };

    const toggleSearch = () => {
      showSearch.value = !showSearch.value;
      if (showSearch.value) {
        showSuggestions.value = true;
      }
    };

    const performSearch = () => {
      console.log('Searching for:', searchQuery.value);
      showSuggestions.value = false;
      showSearch.value = false;
    };

    const hideSearchIfNotFocused = () => {
      setTimeout(() => {
        if (!showSuggestions.value) {
          showSearch.value = false;
        }
      }, 200);
    };

    const hideSuggestionsDelayed = () => {
      setTimeout(() => {
        showSuggestions.value = false;
      }, 200);
    };

    const selectBrand = (brandName) => {
      searchQuery.value = brandName;
      showSuggestions.value = false;
      showSearch.value = false;
      performSearch();
    };

    return { 
      isMenuOpen, 
      toggleMenu, 
      searchQuery, 
      performSearch, 
      showSearch, 
      showSuggestions,
      filteredBrands,
      hideSearchIfNotFocused,
      hideSuggestionsDelayed,
      selectBrand,
      toggleSearch
    };
  }
}
</script>