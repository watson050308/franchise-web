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
import { provide, ref } from 'vue'
import CategoriesSection from './CategoriesSection.vue'
import BaseSlide from './ui/BaseSlide.vue'

export default {
  name: 'HomePage',
  components: {
    BaseSlide,
    CategoriesSection
  },
  setup() {
    const expoCategories = ref([
      {category: '加盟展', label: 'franchiseExpo', image: './src/assets/expo/franchise.png'},
      {category: '寵物展', label: 'petExpo', image: './src/assets/expo/pet.png'},
      {category: '旅遊展', label: 'travelExpo', image: './src/assets/expo/travel.png'},
      {category: '科技展', label: 'techExpo', image: './src/assets/expo/tech.png'},
      {category: '藝術展', label: 'artExpo', image: './src/assets/expo/art.png'},
      {category: '美食展', label: 'foodExpo', image: './src/assets/expo/food.png'}
    ])
    
    provide('expoCategories', expoCategories)
    
    return {
      expoCategories
    }
  }
}
</script>