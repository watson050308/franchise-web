<template>
    <div class="min-h-screen bg-gray-50">
      <SearchBar @search="handleSearch" :defaultCategory="defaultCategory" :categoryOptions="categoryOptions" />
      <BrandGrid :brands="brands" />
    </div>
  </template>
  
  <script>
  import { ref, computed } from 'vue';

  import SearchBar from './ui/SearchBar.vue';
  import BrandGrid from './ui/BrandGrid.vue';
import { useRoute } from 'vue-router';
  
  export default {
    name: 'FranchiseBrandShowcase',
    components: {
      SearchBar,
      BrandGrid
    },
    setup() {
      const route = useRoute();
      const defaultCategory = computed(() => route.params.category || '')

      const categoryOptions = computed(() => {
        const options = {
          franchiseExpo: [
            { value: 'fastfood', label: '快餐' },
            { value: 'restaurant', label: '餐廳' },
            { value: 'cafe', label: '咖啡廳' }
          ],
          petExpo: [
            { value: 'clothing', label: '服飾' },
            { value: 'electronics', label: '電子產品' },
            { value: 'furniture', label: '家具' }
          ],
          travelExpo: [
            { value: 'food', label: '餐飲' },
            { value: 'retail', label: '零售' },
            { value: 'service', label: '服務' }
          ]
        };

        console.log("category", route.query)
        return options[route.query.category];
      });

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
        defaultCategory,
        categoryOptions,
        brands,
        handleSearch,
        handleFilter
      };
    }
  };
  </script>