<template>
  <section>
    <h2 class="text-2xl font-bold mb-8">{{ sectionTitle }}</h2>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <BaseCard 
        v-for="(expoCategorie, index) in categories" 
        :key="index" 
        :current-category="expoCategorie.category"
      >
        <template #content>
          <router-link 
            :to="{path: '/showbrand', query: {category: expoCategorie.label}}" 
            class="inline-flex items-center text-blue-600 hover:text-blue-800 font-medium"
          >
            查看機會 
            <ChevronRight class="ml-1 w-4 h-4" />
          </router-link>
        </template>
      </BaseCard>
    </div>
  </section>
</template>
  
<script>
import { ChevronRight } from 'lucide-vue-next'
import BaseCard from './ui/BaseCard.vue'
// import CardTitle from './ui/BaseCard.vue'

export default {
  name: 'CategoriesSection',
  components: {
    BaseCard,
    // CardTitle,
    ChevronRight
  },
  props: {
    sectionTitle: {
      type: String,
      default: '熱門品牌展館'
    },
    categories: {
      type: Array,
      default: () => ['餐飲', '零售', '教育', '美容健康']
    }
  }
}
</script>