<template>
  <section class="mb-16">
    <h2 class="text-2xl font-semibold mb-6">熱門加盟類別</h2>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <BaseCard v-for="(category, index) in categories" :key="index" class="hover:shadow-lg transition-shadow">
        <template #header>
          <CardTitle class="text-lg font-semibold">{{ category }}</CardTitle>
        </template>
        <template #content>
          <a href="#" class="text-blue-600 flex items-center hover:underline">
            查看機會 <ChevronRight class="ml-1" :size="16" />
          </a>
        </template>
      </BaseCard>
    </div>
  </section>
</template>
  
  <script>
  import { ChevronRight } from 'lucide-vue-next'
  import BaseCard from './ui/BaseCard.vue'
  import CardTitle from './ui/BaseCard.vue'
  
  export default {
    name: 'CategoriesSection',
    components: {
      BaseCard,
      CardTitle,
      ChevronRight
    },
    data() {
      return {
        categories: ['餐飲', '零售', '教育', '美容健康']
      }
    }
  }
  </script>