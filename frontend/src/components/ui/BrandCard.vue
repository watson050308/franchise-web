<template>
    <div class="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-300 hover:shadow-lg hover:-translate-y-1">
      <img :src="brand.image" :alt="brand.name" class="w-full h-48 object-cover" />
      <div class="p-4">
        <h3 class="text-lg font-semibold mb-1 truncate">{{ brand.name }}</h3>
        <p class="text-sm text-gray-600 mb-2 truncate">{{ brand.category }}</p>
        <div class="text-xs text-gray-500">
          <p class="truncate">初始投資: {{ brand.initialInvestment }}</p>
          <p class="truncate">年營業額: {{ brand.annualRevenue }}</p>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'BrandCard',
    props: {
      brand: {
        type: Object,
        required: true
      }
    }
  };
  </script>