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
  
<script setup>
import { provide, ref } from 'vue'
import CategoriesSection from './CategoriesSection.vue'
import BaseSlide from './ui/BaseSlide.vue'

// 直接導入圖片
import franchiseImg from '../assets/expo/franchise.png'
import petImg from '../assets/expo/pet.png'
import travelImg from '../assets/expo/travel.png'
import techImg from '../assets/expo/tech.png'
import artImg from '../assets/expo/art.png'
import foodImg from '../assets/expo/food.png'

const expoCategories = ref([
  {category: '加盟展', label: 'franchiseExpo', image: franchiseImg},
  {category: '寵物展', label: 'petExpo', image: petImg},
  {category: '旅遊展', label: 'travelExpo', image: travelImg},
  {category: '科技展', label: 'techExpo', image: techImg},
  {category: '藝術展', label: 'artExpo', image: artImg},
  {category: '美食展', label: 'foodExpo', image: foodImg}
])

provide('expoCategories', expoCategories)
</script>