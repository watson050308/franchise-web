<template>
  <section class="mb-16">
    <h2 class="text-2xl font-semibold mb-6">為什麼選擇FranchiseConnect</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 md:gap-8">
      <BaseCard v-for="(item, index) in reasons" :key="index" class="text-center">
        <template #content>
          <component :is="item.icon" class="mx-auto mb-4 text-blue-600" :size="48" />
          <h3 class="text-xl font-semibold mb-2">{{ item.title }}</h3>
          <p class="text-gray-600">{{ item.desc }}</p>
        </template>
      </BaseCard>
    </div>
  </section>
</template>
  
  <script>
  import { Users, BarChart, Award } from 'lucide-vue-next'
  import BaseCard from './ui/BaseCard.vue'
  
  export default {
    name: 'WhyChooseUsSection',
    components: {
      BaseCard,
      Users,
      BarChart,
      Award
    },
    data() {
      return {
        reasons: [
          { icon: Users, title: "龐大的加盟網路", desc: "連接數千個品牌和潛在加盟商" },
          { icon: BarChart, title: "數據驅動決策", desc: "提供詳細的市場分析和績效報告" },
          { icon: Award, title: "專業支持", desc: "全程指導，確保你的加盟成功" }
        ]
      }
    }
  }
  </script>