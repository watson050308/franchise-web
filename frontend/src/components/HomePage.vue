<!-- HomePage.vue -->
<template>
  <BaseSlide />
  <div class="flex-grow max-w-6xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-16">
    <CategoriesSection
      sectionTitle="熱門展館" 
      :categories="expoCategories"
      class="mb-16"
    />
    <!--<CategoriesSection
      sectionTitle="熱門寵物展館" 
      :categories="petCategories"
    />-->
  </div>
</template>
  
  <script>
  
  import CategoriesSection from './CategoriesSection.vue'
  import BaseSlide from './ui/BaseSlide.vue'

  export default {
  name: 'HomePage',
  components: {
    BaseSlide,
    CategoriesSection
  },
  data() {
    return {
      expoCategories: [
        {category: '加盟展', label: 'franchiseExpo'},
        {category: '寵物展', label: 'petExpo'},
        {category: '旅遊展', label: 'travelExpo'},
        {category: '科技展', label: 'techExpo'},
        {category: '藝術展', label: 'artExpo'},
        {category: '美食展', label: 'foodExpo'}
      ],
      petCategories: ['寵物食品', '寵物用品', '寵物服務', '寵物醫療']
    }
  }
}
  
  </script>