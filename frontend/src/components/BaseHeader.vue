<template>
  <header class="bg-gradient-to-r from-indigo-500 to-teal-400 text-white py-4 px-8 sm:px-6 lg:px-8">
    <nav class="flex justify-between items-center max-w-6xl mx-auto">
      <div>
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
          <text x="75" y="32" font-family="Arial, sans-serif" font-size="20" fill="#333333">ExpoVerse</text>
        </svg>
      </router-link>
  </div>
      <div class="hidden sm:flex space-x-4 items-center">
        <router-link to="/" class="hover:opacity-80 transition-opacity duration-300">首頁</router-link>
        <router-link to="/exhibition" class="hover:opacity-80 transition-opacity duration-300">尋找展館</router-link>
        <!-- <router-link to="/login" class="hover:underline">登入</router-link> -->
        <!-- <a href="#" class="hover:underline">尋找加盟機會</a> -->
        <a href="#" class="hover:opacity-80 transition-opacity duration-300">展商支援</a>
        <a href="#" class="hover:opacity-80 transition-opacity duration-300">關於我們</a>
        <BaseButton variant="outline" class="text-white border-white hover:opacity-80 transition-opacity duration-300"><router-link to="/login" class="hover:underline">登入</router-link></BaseButton>
      </div>
      <button @click="toggleMenu" class="sm:hidden">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
        </svg>
      </button>
    </nav>
    <transition
      enter-active-class="transition ease-out duration-100 transform"
      enter-from-class="opacity-0 scale-95"
      enter-to-class="opacity-100 scale-100"
      leave-active-class="transition ease-in duration-75 transform"
      leave-from-class="opacity-100 scale-100"
      leave-to-class="opacity-0 scale-95"
    >
      <div v-show="isMenuOpen" class="sm:hidden mt-4">
        <router-link to="/exhibition" class="block py-2 px-4 hover:opacity-80 transition-opacity duration-300">尋找加盟機會</router-link>
        <a href="#" class="block py-2 px-4 hover:opacity-80 transition-opacity duration-300">加盟商支援</a>
        <a href="#" class="block py-2 px-4 hover:opacity-80 transition-opacity duration-300">關於我們</a>
        <BaseButton variant="outline" class="mt-2 w-full text-white border-white hover:opacity-80 transition-opacity duration-300">登入</BaseButton>
      </div>
    </transition>
  </header>
</template>

<script>
import { ref } from 'vue';
import BaseButton from './ui/BaseButton.vue';

export default {
  name: 'BaseHeader',
  components: {
    BaseButton
  },
  setup() {
    const isMenuOpen = ref(false);
    const toggleMenu = () => {
      isMenuOpen.value = !isMenuOpen.value;
    };
    return { isMenuOpen, toggleMenu };
  }
}
</script>