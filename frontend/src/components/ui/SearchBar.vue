<template>
    <div class="container mx-auto px-4 py-8">
      <div class="bg-white rounded-lg shadow-lg p-4 md:p-6">
        <div class="flex flex-col md:flex-row md:items-center space-y-4 md:space-y-0 md:space-x-4">
          <div class="flex-1 flex items-center border-r border-gray-300 px-4">
            <Search :size="20" class="text-gray-400 mr-2" />
            <input
              type="text"
              placeholder="搜索品牌名稱"
              v-model="brandName"
              class="w-full bg-transparent border-none focus:outline-none"
            />
          </div>
          <div class="flex-1 flex items-center border-r border-gray-300 px-4">
            <MapPin :size="20" class="text-gray-400 mr-2" />
            <select 
              v-model="region"
              class="w-full bg-transparent border-none focus:outline-none"
            >
              <option value="">選擇區域</option>
              <option value="north">北部</option>
              <option value="south">南部</option>
              <option value="east">東部</option>
              <option value="west">西部</option>
            </select>
            <!-- <ChevronDown :size="20" class="text-gray-400 ml-2" /> -->
          </div>
          <div class="flex-1 flex items-center border-r border-gray-300 px-4">
            <Tag :size="20" class="text-gray-400 mr-2" />
            <select
              v-model="category"
              class="w-full bg-transparent border-none focus:outline-none"
            >
              <option value="">選擇類別</option>
              <option v-for="option in categoryOptions" :key="option.value" :value="option.value">
                {{ option.label }}
              </option>
              <!-- <option value="food">寵物展</option>
              <option value="retail">零售</option>
              <option value="service">服務</option> -->
            </select>
            <!-- <ChevronDown :size="20" class="text-gray-400 ml-2" /> -->
          </div>
          <div class="flex-1 flex items-center px-4">
            <DollarSign :size="20" class="text-gray-400 mr-2" />
            <select
              v-model="priceRange"
              class="w-full bg-transparent border-none focus:outline-none"
            >
              <option value="">選擇價位</option>
              <option value="low">低價位 (&lt; 100萬)</option>
              <option value="medium">中價位 (100-500萬)</option>
              <option value="high">高價位 (&gt; 500萬)</option>
            </select>
            <!-- <ChevronDown :size="20" class="text-gray-400 ml-2" /> -->
          </div>
          <button
            @click="handleSearch"
            class="bg-black text-white rounded-full px-6 py-2 ml-4 hover:bg-gray-800 transition-colors"
          >
            搜索品牌
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import { Search, MapPin, DollarSign, Tag } from 'lucide-vue-next';
import { watch } from 'vue';
  
  export default {
    name: 'CustomSearchBar',
    components: {
      Search,
      MapPin,
      DollarSign,
      Tag
    },
    props: {
      defaultCategory: {
        type: String,
        default: ''
      },
      categoryOptions: {
        type: Array,
        default: () => []
      }
    },
    emits: ['search'],
    setup(props, { emit }) {
      const brandName = ref('');
      const region = ref('');
      const category = ref(props.defaultCategory);
      const priceRange = ref('');

      watch(() => props.defaultCategory, () => {
        console.log("new value");
        category.value = 'food';
      });
  
      const handleSearch = () => {
        emit('search', { brandName: brandName.value, region: region.value, category: category.value, priceRange: priceRange.value });
      };
  
      return {
        brandName,
        region,
        category,
        priceRange,
        handleSearch
      };
    }
  };
  </script>