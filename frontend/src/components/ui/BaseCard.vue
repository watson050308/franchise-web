<template>
  <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-all duration-300">
    <slot name="content"></slot>
  </div>
</template>

<script setup>
// BaseCard 不需要任何 props 或其他邏輯，
// 只需要提供一個具名插槽 "content" 供內容插入
</script>

<style scoped>
.group:hover {
  transform: translateY(-2px);
}
</style>