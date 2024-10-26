<!-- AboutUsPage.vue -->
<template>
    <div class="flex-grow max-w-6xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12">
      <WhyChooseUsSection />
      <section class="text-center mb-16">
        <h2 class="text-2xl font-semibold mb-4">準備好開始你的加盟之旅了嗎？</h2>
        <BaseButton class="bg-blue-600 hover:bg-blue-700 text-white">
          立即註冊
        </BaseButton>
      </section>
    </div>
</template>
    
<script>

import WhyChooseUsSection from './WhyChooseUsSection.vue'
import BaseButton from './ui/BaseButton.vue'

export default {
  name: 'AboutUsPage',
  components: {
    WhyChooseUsSection,
    BaseButton
}
}

</script>