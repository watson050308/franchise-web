<template>
  這是測試
  <div class="container mx-auto p-4">
    <EditableArea
      v-model="content"
      :min-height="300"
      @change="handleContentChange"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import EditableArea from './EditableArea.vue'

const content = ref('')

const handleContentChange = (newContent) => {
  console.log('Content updated:', newContent)
}
</script>