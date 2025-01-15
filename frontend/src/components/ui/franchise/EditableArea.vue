<!-- EditableArea.vue -->
<template>
    <div class="editable-container">
      <!-- 工具欄 -->
      <div class="toolbar bg-gray-100 p-2 rounded-t-lg border-b flex flex-wrap gap-2">
        <!-- 文字樣式 -->
        <div class="flex gap-2">
          <button 
            v-for="(style, index) in textStyles"
            :key="index"
            @click="applyStyle(style.command)"
            class="p-2 hover:bg-gray-200 rounded"
            :class="{ 'bg-gray-300': isStyleActive(style.command) }">
            <span :class="style.icon">{{ style.label }}</span>
          </button>
        </div>
  
        <!-- 分隔線 -->
        <div class="h-6 w-px bg-gray-300"></div>
  
        <!-- 對齊方式 -->
        <div class="flex gap-2">
          <button 
            v-for="align in alignments"
            :key="align.command"
            @click="applyStyle(align.command)"
            class="p-2 hover:bg-gray-200 rounded">
            {{ align.label }}
          </button>
        </div>
  
        <!-- 顏色選擇器 -->
        <div class="flex items-center gap-2">
          <input 
            type="color" 
            v-model="selectedColor"
            @change="applyColor"
            class="w-8 h-8 rounded cursor-pointer">
        </div>

        <div class="flex items-center gap-2">
        <!-- 在原有工具欄的分隔線後面添加 -->
        <select 
            @change="applyFontSize"
            class="p-2 border rounded hover:bg-gray-200"
            v-model="selectedFontSize">
            <option value="1">最小</option>
            <option value="2">小</option>
            <option value="3">正常</option>
            <option value="4">大</option>
            <option value="5">較大</option>
            <option value="6">最大</option>
        </select>
        </div>
      </div>

      
  
      <!-- 編輯區域 -->
      <div
        ref="editableDiv"
        class="edit-area p-4 min-h-[200px] bg-white border rounded-b-lg"
        contenteditable="true"
        @input="handleInput"
        @paste="handlePaste"
        v-html="content">
      </div>
  
      <!-- 調整大小控制器 -->
      <div 
        class="resize-handle"
        @mousedown="startResize">
      </div>
    </div>

    <!-- 在工具欄中添加圖片上傳按鈕 -->
    <div class="flex items-center gap-2">
    <input 
        type="file" 
        ref="fileInput" 
        accept="image/*" 
        class="hidden"
        @change="handleImageUpload">
    <button 
        @click="triggerImageUpload"
        class="p-2 hover:bg-gray-200 rounded">
        插入圖片
    </button>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onUnmounted } from 'vue'
  
  // 定義屬性
  const props = defineProps({
    modelValue: {
      type: String,
      default: ''
    },
    minHeight: {
      type: Number,
      default: 200
    }
  })
  
  // 定義事件
  const emit = defineEmits(['update:modelValue', 'change'])
  
  // 參考和狀態
  const editableDiv = ref(null)
  const selectedColor = ref('#000000')
  const content = ref(props.modelValue)
  const isResizing = ref(false)
  const startY = ref(0)
  const startHeight = ref(0)
  
  // 文字樣式選項
  const textStyles = [
    { label: 'B', command: 'bold', icon: 'font-bold' },
    { label: 'I', command: 'italic', icon: 'italic' },
    { label: 'U', command: 'underline', icon: 'underline' }
  ]
  
  // 對齊方式
  const alignments = [
    { label: '靠左', command: 'justifyLeft' },
    { label: '置中', command: 'justifyCenter' },
    { label: '靠右', command: 'justifyRight' }
  ]

  // 設定字體大小
  const selectedFontSize = ref('3')

  const applyFontSize = () => {
  document.execCommand('fontSize', false, selectedFontSize.value)
  editableDiv.value.focus()
}
  
  // 應用樣式
  const applyStyle = (command) => {
    document.execCommand(command, false, null)
    editableDiv.value.focus()
  }
  
  // 應用顏色
  const applyColor = () => {
    document.execCommand('foreColor', false, selectedColor.value)
    editableDiv.value.focus()
  }
  
  // 檢查樣式是否激活
  const isStyleActive = (command) => {
    return document.queryCommandState(command)
  }
  
  // 處理輸入
  const handleInput = () => {
    const newContent = editableDiv.value.innerHTML
    emit('update:modelValue', newContent)
    emit('change', newContent)
  }
  
  // 處理貼上，去除格式
  const handlePaste = (e) => {
    e.preventDefault()
    const text = e.clipboardData.getData('text/plain')
    document.execCommand('insertText', false, text)
  }
  
  // 調整大小相關函數
  const startResize = (e) => {
    isResizing.value = true
    startY.value = e.clientY
    startHeight.value = editableDiv.value.offsetHeight
    document.addEventListener('mousemove', handleResize)
    document.addEventListener('mouseup', stopResize)
  }
  
  const handleResize = (e) => {
    if (!isResizing.value) return
    const newHeight = startHeight.value + (e.clientY - startY.value)
    if (newHeight >= props.minHeight) {
      editableDiv.value.style.height = `${newHeight}px`
    }
  }
  
  const stopResize = () => {
    isResizing.value = false
    document.removeEventListener('mousemove', handleResize)
    document.removeEventListener('mouseup', stopResize)
  }

  const fileInput = ref(null)

const triggerImageUpload = () => {
  fileInput.value.click()
}
// 處理圖像上傳
const handleImageUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      document.execCommand('insertImage', false, e.target.result)
    }
    reader.readAsDataURL(file)
  }
}
  
  // 生命週期
  onMounted(() => {
    if (props.modelValue) {
      content.value = props.modelValue
    }
  })
  
  onUnmounted(() => {
    document.removeEventListener('mousemove', handleResize)
    document.removeEventListener('mouseup', stopResize)
  })
  </script>
  
  <style scoped>
  .editable-container {
    position: relative;
    width: 100%;
    border-radius: 8px;
    overflow: hidden;
  }
  
  .edit-area {
    outline: none;
    overflow-y: auto;
    word-break: break-word;
  }
  
  .resize-handle {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 5px;
    cursor: ns-resize;
    background: transparent;
  }
  
  .resize-handle:hover {
    background: rgba(0, 0, 0, 0.1);
  }
  
  /* 自定義滾動條樣式 */
  .edit-area::-webkit-scrollbar {
    width: 8px;
  }
  
  .edit-area::-webkit-scrollbar-track {
    background: #f1f1f1;
  }
  
  .edit-area::-webkit-scrollbar-thumb {
    background: #888;
    border-radius: 4px;
  }
  
  .edit-area::-webkit-scrollbar-thumb:hover {
    background: #555;
  }
  </style>