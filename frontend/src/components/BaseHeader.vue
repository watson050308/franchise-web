<template>
    <header class="bg-blue-600 text-white p-4">
      <nav class="flex justify-between items-center max-w-6xl mx-auto">
        <div class="text-2xl font-bold">FranchiseConnect</div>
        <div class="space-x-4">
          <a href="#" class="hover:underline">尋找加盟機會</a>
          <a href="#" class="hover:underline">加盟商支援</a>
          <a href="#" class="hover:underline">關於我們</a>
          <BaseButton variant="outline" class="text-white border-white hover:bg-blue-700">登入</BaseButton>
        </div>
      </nav>
    </header>
  </template>
  
  <script>
  import BaseButton from './ui/BaseButton.vue'
  
  export default {
    name: 'BaseHeader',
    components: {
      BaseButton
    }
  }
  </script>