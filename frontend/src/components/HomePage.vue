<!-- HomePage.vue -->
<template>
  <BaseSlide />
  <div class="flex-grow max-w-6xl w-full mx-auto px-4 sm:px-6 lg:px-8">
    <SearchSection />
    <CategoriesSection
      sectionTitle="熱門加盟展館" 
      :categories="franchiseCategories"
    />
    <CategoriesSection
      sectionTitle="熱門寵物展館" 
      :categories="petCategories"
    />
  </div>
</template>
  
  <script>
  
  import SearchSection from './SearchSection.vue'
  import CategoriesSection from './CategoriesSection.vue'
  import BaseSlide from './ui/BaseSlide.vue'

  export default {
  name: 'HomePage',
  components: {
    BaseSlide,
    SearchSection,
    CategoriesSection
  },
  data() {
    return {
      franchiseCategories: ['紅橘子', '紅吉子', '紅橙子', '紅葡子'],
      petCategories: ['寵物食品', '寵物用品', '寵物服務', '寵物醫療']
    }
  }
}
  
  </script>