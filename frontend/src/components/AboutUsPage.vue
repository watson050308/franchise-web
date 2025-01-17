<!-- AboutUsPage.vue -->
<template>
  <div class="flex-grow">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-indigo-600 to-blue-500 text-white py-20">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-5xl font-bold mb-6">關於 ExpoVerse</h1>
        <p class="text-xl md:text-2xl max-w-3xl mx-auto opacity-90">
          我們致力於為企業和消費者打造最優質的展覽體驗，連結產業、創造機會
        </p>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <Suspense>
        <template #default>
          <div class="space-y-20">
            <!-- Mission Section -->
            <section class="text-center">
              <h2 class="text-3xl font-bold mb-8 text-gray-800">我們的使命</h2>
              <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                <div class="p-6 bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow">
                  <div class="text-4xl mb-4">🌟</div>
                  <h3 class="text-xl font-semibold mb-3">創新展覽體驗</h3>
                  <p class="text-gray-600">打造數位化展覽平台，突破傳統展覽的限制</p>
                </div>
                <div class="p-6 bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow">
                  <div class="text-4xl mb-4">🤝</div>
                  <h3 class="text-xl font-semibold mb-3">連結產業價值</h3>
                  <p class="text-gray-600">促進產業交流，創造商機與合作機會</p>
                </div>
                <div class="p-6 bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow">
                  <div class="text-4xl mb-4">📈</div>
                  <h3 class="text-xl font-semibold mb-3">推動產業成長</h3>
                  <p class="text-gray-600">協助企業拓展市場，加速產業發展</p>
                </div>
              </div>
            </section>

            <!-- Why Choose Us Section -->
            <WhyChooseUsSection />

            <!-- CTA Section -->
            <section class="text-center bg-gradient-to-r from-teal-500 to-emerald-500 text-white py-16 rounded-2xl">
              <h2 class="text-3xl font-bold mb-6">準備好開始你的展覽之旅了嗎？</h2>
              <p class="text-xl mb-8 max-w-2xl mx-auto">
                加入 ExpoVerse，探索無限可能，開創事業新篇章
              </p>
              <BaseButton class="bg-white text-emerald-600 hover:bg-gray-100 font-bold px-8 py-3 rounded-full text-lg shadow-lg hover:shadow-xl transition-all duration-300">
                立即註冊
              </BaseButton>
            </section>

            <!-- Stats Section -->
            <section class="grid grid-cols-1 md:grid-cols-4 gap-8 text-center">
              <div class="p-6">
                <div class="text-4xl font-bold text-indigo-600 mb-2">1000+</div>
                <div class="text-gray-600">參展企業</div>
              </div>
              <div class="p-6">
                <div class="text-4xl font-bold text-indigo-600 mb-2">50+</div>
                <div class="text-gray-600">展覽類別</div>
              </div>
              <div class="p-6">
                <div class="text-4xl font-bold text-indigo-600 mb-2">100萬+</div>
                <div class="text-gray-600">月活用戶</div>
              </div>
              <div class="p-6">
                <div class="text-4xl font-bold text-indigo-600 mb-2">98%</div>
                <div class="text-gray-600">客戶滿意度</div>
              </div>
            </section>
          </div>
        </template>

        <template #fallback>
          <div class="flex items-center justify-center min-h-screen">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
          </div>
        </template>
      </Suspense>
    </div>
  </div>
</template>
    
<script setup>
import WhyChooseUsSection from './WhyChooseUsSection.vue'
import BaseButton from './ui/BaseButton.vue'
</script>

<style scoped>
.transition-shadow {
  transition: box-shadow 0.3s ease;
}
</style>