<template>
    <div class="min-h-screen bg-gray-50">
      <SearchBar @search="handleSearch" />
      <BrandGrid :brands="brands" />
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import SearchBar from './ui/SearchBar.vue';
  import BrandGrid from './ui/BrandGrid.vue';
  
  export default {
    name: 'FranchiseBrandShowcase',
    components: {
      SearchBar,
      BrandGrid
    },
    setup() {
      const brands = ref([
        {
          id: 1,
          name: "超級漢堡",
          image: "/api/placeholder/400/300",
          category: "餐飲",
          initialInvestment: "NT$500萬-800萬",
          annualRevenue: "NT$1000萬+"
        },
        {
          id: 2,
          name: "時尚服飾",
          image: "/api/placeholder/400/300",
          category: "零售",
          initialInvestment: "NT$300萬-500萬",
          annualRevenue: "NT$800萬+"
        },
        {
          id: 3,
          name: "健康生活",
          image: "/api/placeholder/400/300",
          category: "服務",
          initialInvestment: "NT$200萬-400萬",
          annualRevenue: "NT$600萬+"
        },
        // ... 更多品牌數據
      ]);

      const handleSearch = (searchTerm) => {
        // 實現搜索邏輯
        console.log('Searching for:', searchTerm);
      };

      const handleFilter = (filterType, value) => {
        // 實現過濾邏輯
        console.log('Filtering:', filterType, value);
      };

      return {
        brands,
        handleSearch,
        handleFilter
      };
    }
  };
  </script>