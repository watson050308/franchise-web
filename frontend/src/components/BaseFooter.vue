<template>
  <footer class="bg-indigo-600 text-white text-center py-4">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center text-white">
      <p>&copy; {{ currentYear }} ExpoVerse. 保留所有權利。</p>
    </div>
  </footer>
</template>
  
<script setup>
import { computed } from 'vue'

const currentYear = computed(() => {
  return new Date().getFullYear()
})
</script>