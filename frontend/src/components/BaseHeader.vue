<template>
  <header class="bg-gradient-to-r from-indigo-500 to-teal-400 text-white py-4 px-8 sm:px-6 lg:px-8">
    <nav class="flex justify-between items-center max-w-6xl mx-auto">
      <router-link to="/" class="text-2xl font-bold">FranchiseConnect</router-link>
      <div class="hidden sm:flex space-x-4 items-center">
        <router-link to="/" class="hover:opacity-80 transition-opacity duration-300">首頁</router-link>
        <router-link to="/exhibition" class="hover:opacity-80 transition-opacity duration-300">尋找加盟機會</router-link>
        <!-- <router-link to="/login" class="hover:underline">登入</router-link> -->
        <!-- <a href="#" class="hover:underline">尋找加盟機會</a> -->
        <a href="#" class="hover:opacity-80 transition-opacity duration-300">加盟商支援</a>
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
        <a href="#" class="block py-2 px-4 hover:opacity-80 transition-opacity duration-300">尋找加盟機會</a>
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