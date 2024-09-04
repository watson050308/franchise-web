<!-- HomePage.vue -->
<template>
  <div class="flex-grow max-w-6xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12">
    <SearchSection />
    <CategoriesSection />
    <CategoriesSection />
  </div>
  </template>
  
  <script>
  
  import SearchSection from './SearchSection.vue'
  import CategoriesSection from './CategoriesSection.vue'

  export default {
  name: 'HomePage',
  components: {
    SearchSection,
    CategoriesSection
  }
}
  
  </script>