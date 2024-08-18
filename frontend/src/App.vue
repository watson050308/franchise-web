<template>
  <div class="bg-white min-h-screen flex flex-col">
    <BaseHeader />
    <main class="flex-grow max-w-6xl w-full mx-auto px-4 py-12">
      <SearchSection />
      <CategoriesSection />
      <WhyChooseUsSection />
      <section class="text-center mb-16">
        <h2 class="text-2xl font-semibold mb-4">準備好開始你的加盟之旅了嗎？</h2>
        <BaseButton class="bg-blue-600 hover:bg-blue-700 text-white">
          立即註冊
        </BaseButton>
      </section>
    </main>
    <BaseFooter />
  </div>
</template>

<script>
import BaseHeader from './components/BaseHeader.vue'
import BaseFooter from './components/BaseFooter.vue'
import SearchSection from './components/SearchSection.vue'
import CategoriesSection from './components/CategoriesSection.vue'
import WhyChooseUsSection from './components/WhyChooseUsSection.vue'
import BaseButton from './components/ui/BaseButton.vue'

export default {
  name: 'App',
  components: {
    BaseHeader,
    BaseFooter,
    SearchSection,
    CategoriesSection,
    WhyChooseUsSection,
    BaseButton
  }
}
</script>