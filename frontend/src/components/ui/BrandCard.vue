<template>
  <div class="relative bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-300 hover:shadow-lg hover:-translate-y-1">
    <!-- 認證標章 -->
    <div 
      v-if="brand.isVerified" 
      class="absolute left-2 top-2 z-10"
      title="已認證品牌"
    >
      <img 
        src="../../assets/medal.png" 
        alt="認證標章" 
        class="h-10 w-10"
      />
    </div>

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

<script setup>
defineProps({
  brand: {
    type: Object,
    required: true
  }
})
</script>