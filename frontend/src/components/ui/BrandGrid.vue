<template>
    <div class="container mx-auto grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4 py-8">
      <BrandCard v-for="brand in brands" :key="brand.id" :brand="brand" />
    </div>
  </template>
  
  <script>
  import BrandCard from './BrandCard.vue';
  
  export default {
    name: 'BrandGrid',
    components: {
      BrandCard
    },
    props: {
      brands: {
        type: Array,
        required: true
      }
    }
  };
  </script>