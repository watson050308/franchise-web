<template>
    <div class="container mx-auto grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4 py-8">
      <router-link v-for="brand in brands" 
                 :key="brand.id" 
                 to="/vendor" 
                 class="block">
      <BrandCard :brand="brand" />
    </router-link>
    </div>
  </template>
  
<script setup>
import BrandCard from './BrandCard.vue'

defineProps({
  brands: {
    type: Array,
    required: true
  }
})
</script>