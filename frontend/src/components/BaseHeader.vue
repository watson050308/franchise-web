<template>
  <header class="bg-white border-b">
    <nav class="flex justify-between items-center max-w-7xl mx-auto px-4 py-3">
      <div class="w-full sm:w-auto flex justify-between items-center">
        <!-- Logo -->
        <router-link to="/" class="text-2xl font-bold inline-block">
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
            <text x="75" y="22" font-family="Arial, sans-serif" font-size="20" fill="black">展覽宇宙</text>
            <text x="75" y="44" font-family="Arial, sans-serif" font-size="20" fill="black">ExpoVerse</text>
          </svg>
        </router-link>

        <!-- hamburger menu -->
        <button 
          @click="toggleMenu"
          class="md:hidden text-gray-700 hover:text-green-600 transition-colors duration-300"
          aria-label="Toggle menu"
        >
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            class="h-6 w-6" 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor"
          >
            <path 
              v-if="!isMenuOpen"
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M4 6h16M4 12h16M4 18h16"
            />
            <path 
              v-else
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>
      
      <div class="flex-grow flex justify-end items-center">
        <!-- 主導航選單 -->
        <div class="hidden md:flex items-center space-x-8">
          <router-link 
            v-for="item in menuItems.filter(item => item.name !== '登入')" 
            :key="item.path"
            :to="item.path"
            class="text-gray-700 hover:text-green-600 transition-colors duration-300"
          >
            {{ item.name }}
          </router-link>
        </div>

        <!-- 右側工具列 -->
        <div class="hidden md:flex items-center ml-8">
          <!-- 搜索按鈕和搜索框 -->
          <div class="relative mr-8">
            <button 
              @click="toggleSearch"
              class="text-gray-700 hover:text-green-600 transition-colors duration-300"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </button>
            
            <!-- 搜索框 -->
            <div v-if="isSearchOpen" 
                class="absolute right-0 top-full mt-2 w-64 bg-white rounded-lg shadow-lg p-2 z-50">
              <input
                type="text"
                v-model="searchQuery"
                @keyup.enter="performSearch"
                @input="handleSearchInput"
                @blur="handleBlur"
                placeholder="搜索..."
                class="w-full px-4 py-2 border rounded-full focus:outline-none focus:ring-2 focus:ring-green-600"
                ref="searchInput"
              />
              <!-- 搜索建議 -->
              <div v-if="showSuggestions && filteredBrands.length > 0" 
                  class="mt-2">
                <ul>
                  <li v-for="brand in filteredBrands" 
                      :key="brand.name"
                      @mousedown="selectBrand(brand.name)"
                      class="px-4 py-2 hover:bg-gray-100 cursor-pointer text-gray-700">
                    {{ brand.name }}
                  </li>
                </ul>
              </div>
            </div>
          </div>

            <div class="relative">
              <template v-if="isLogin">
                <button 
                  @click="toggleUserMenu"
                  class="flex items-center space-x-2 text-gray-700 hover:text-green-600 transition-colors duration-300"
                >
                  <span>{{ user.name }}</span>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                  </svg>
                </button>
                
                <div v-if="showUserMenu" 
                    class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg py-2 z-50">
                  <router-link 
                    to="/app-management"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    @click="showUserMenu = false"
                  >
                    個人資料
                  </router-link>
                  <router-link 
                    to="/app-management"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    @click="showUserMenu = false"
                  >
                    管理面板
                  </router-link>
                  <button 
                    @click="handleLogout"
                    class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100"
                  >
                    登出
                  </button>
                </div>
              </template>
              
              <router-link 
                v-else
                to="/login"
                class="text-gray-700 hover:text-green-600 transition-colors duration-300"
              >
                登入
              </router-link>
            </div>
        </div>
      </div>
    </nav>

    <!-- mobile menu -->
    <transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform -translate-y-full opacity-0"
      enter-to-class="transform translate-y-0 opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="transform translate-y-0 opacity-100"
      leave-to-class="transform -translate-y-full opacity-0"
    >
      <div 
        v-show="isMenuOpen" 
        class="md:hidden border-t bg-white absolute w-full z-50 shadow-lg"
      >
        <div class="px-4 py-2">
          <router-link 
            v-for="item in menuItems" 
            :key="item.path"
            :to="item.path"
            class="block py-2 text-gray-700 hover:text-green-600"
            @click="isMenuOpen = false"
          >
            {{ item.name }}
          </router-link>
        </div>

        <!-- 行動裝置搜索框 -->
        <div class="px-4 pb-4">
          <div class="relative">
            <input
              type="text"
              v-model="searchQuery"
              @keyup.enter="performSearch"
              @input="handleSearchInput"
              placeholder="搜索..."
              class="w-full px-4 py-2 border rounded-full focus:outline-none focus:ring-2 focus:ring-green-600"
            />
            <!-- 搜索建議 -->
            <div v-if="showSuggestions && filteredBrands.length > 0" 
                  class="mt-2 bg-white rounded-lg shadow-lg">
              <ul>
                <li v-for="brand in filteredBrands" 
                    :key="brand.name"
                    @mousedown="selectBrand(brand.name)"
                    class="px-4 py-2 hover:bg-gray-100 cursor-pointer text-gray-700">
                    {{ brand.name }}
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '@/composables/useAuth';

const router = useRouter();
const { user, logout, isLogin, checkAuth } = useAuth();

const isMenuOpen = ref(false);
const isSearchOpen = ref(false);
const searchQuery = ref('');
const showSuggestions = ref(false);
const searchInput = ref(null);
const showUserMenu = ref(false);

const menuItems = [
  { name: '首頁', path: '/' },
  { name: '尋找展館', path: '/exhibition' },
  { name: '展商支援', path: '/brandSupport' },
  { name: '聯絡我們', path: '/about-us' },
];

const brands = [
  { name: '紅橘子' },
  { name: '紅吉子' },
  { name: '紅柿子' },
];

const filteredBrands = computed(() => {
  if (!searchQuery.value) return [];
  return brands.filter(brand => 
    brand.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

const toggleSearch = () => {
  isSearchOpen.value = !isSearchOpen.value;
  if (isSearchOpen.value) {
    nextTick(() => {
      searchInput.value?.focus();
    });
  }
};

const handleSearchInput = () => {
  showSuggestions.value = true;
};

const performSearch = () => {
  console.log('Searching for:', searchQuery.value);
  isSearchOpen.value = false;
  showSuggestions.value = false;
};

const handleBlur = () => {
  setTimeout(() => {
    isSearchOpen.value = false;
    showSuggestions.value = false;
  }, 200);
};

const selectBrand = (brandName: string) => {
  searchQuery.value = brandName;
  performSearch();
};

const toggleUserMenu = () => {
  // prevent closeUserMenu to close the user menu
  event.stopPropagation();
  showUserMenu.value = !showUserMenu.value;
};

const handleLogout = async () => {
  try {
    await logout();
    showUserMenu.value = false;
    router.push('/login');
  } catch (error) {
    console.error('Logout failed:', error);
  }
};

// close user menu when click outside
const closeUserMenu = (event: MouseEvent) => {
  const target = event.target as HTMLElement;
  if (!target.closest('.user-menu')) {
    showUserMenu.value = false;
  }
};

onMounted(() => {
  document.addEventListener('click', closeUserMenu);
});

onUnmounted(() => {
  document.removeEventListener('click', closeUserMenu);
});
</script>