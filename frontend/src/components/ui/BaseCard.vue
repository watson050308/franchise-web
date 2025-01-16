<template>
  <div class="bg-white rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-all duration-300 group">
    <div class="relative">
      <!-- 圖片容器 -->
      <div class="h-48 overflow-hidden">
        <img 
          :src="getCategoryImage(currentCategory)" 
          :alt="currentCategory"
          class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-500"
        >
      </div>
      
      <!-- 標題區塊 -->
      <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-4">
        <slot name="header">
          <h3 class="text-white text-xl font-bold">{{ currentCategory }}</h3>
        </slot>
      </div>
    </div>

    <!-- 內容區塊 -->
    <div class="p-4">
      <slot name="content"></slot>
    </div>
  </div>
</template>

<script>
import { inject } from 'vue'

export default {
  name: 'BaseCard',
  props: {
    currentCategory: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const expoCategories = inject('expoCategories')

    const getCategoryImage = (categoryName) => {
      const category = expoCategories.value.find(
        item => item.category === categoryName
      )
      return category ? category.image : ''
    }

    return {
      getCategoryImage
    }
  }
}
</script>

<style scoped>
.group:hover {
  transform: translateY(-2px);
}
</style>