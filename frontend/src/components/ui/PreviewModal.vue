<!-- components/ui/PreviewModal.vue -->
<template>
    <div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
      <div class="bg-white p-4 rounded-lg w-3/4 h-3/4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold">網頁預覽</h3>
          <button @click="$emit('close')" class="text-gray-500 hover:text-gray-700">
            關閉
          </button>
        </div>
        <div class="bg-gray-100 h-full p-4 overflow-auto">
          <p>這裡將顯示網頁預覽內容</p>
        </div>
      </div>
    </div>
  </template>
  
<script>
export default {
  name: 'PreviewModal',
  emits: ['close']
};
</script>