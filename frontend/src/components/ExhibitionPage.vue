<template>
    <div class="bg-gray-50 font-poppins text-gray-800">
      <main class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 p-8">
        <ExhibitionCard v-for="exhibition in exhibitions" :key="exhibition.title" :exhibition="exhibition" />
      </main>
    </div>
  </template>
  
  <script>
  import ExhibitionCard from './ui/ExhibitionCard.vue'
  
  export default {
    name: 'ExhibitionHall',
    components: {
      ExhibitionCard
    },
    data() {
      return {
        exhibitions: [
          { title: '加盟展', description: '探索成功的商業模式，開啟你的創業之路', subcategories: ['早午餐', '咖啡廳', '手搖飲'], image: '/api/placeholder/300/300' },
          { title: '寵物展', description: '發現最新的寵物用品和服務，豐富毛孩生活', subcategories: ['狗狗', '貓咪', '小動物'], image: '/api/placeholder/300/300' },
          { title: '旅遊展', description: '探索夢想目的地，規劃你的下一次冒險', subcategories: ['國內旅遊', '國際旅遊', '主題之旅'], image: '/api/placeholder/300/300' },
          { title: '科技展', description: '體驗最新科技，洞察未來趨勢', subcategories: ['AI', 'IoT', '綠能科技'], image: '/api/placeholder/300/300' },
          { title: '藝術展', description: '感受藝術的魅力，探索創意的無限可能', subcategories: ['當代藝術', '攝影', '裝置藝術'], image: '/api/placeholder/300/300' },
          { title: '美食展', description: '品味各國美食，探索舌尖上的驚喜', subcategories: ['異國料理', '甜點', '健康飲食'], image: '/api/placeholder/300/300' }
        ]
      }
    }
  }
  </script>