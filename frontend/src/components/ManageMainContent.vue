<!-- components/ManageMainContent.vue -->
<template>
    <div class="flex-1 p-8">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-2xl font-bold">{{ activeTab }}</h2><br>
        <div>
          {{ activeTab }}
          
        </div>
        <button
          class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded flex items-center"
          @click="isPreviewOpen = true"
        >
          <Eye class="mr-2" :size="20" />
          預覽網頁
        </button>
      </div>
      <component :is="currentTabComponent" />
      <PreviewModal v-if="isPreviewOpen" @close="isPreviewOpen = false" />
    </div>
  </template>
  
<script>
import { ref, computed } from 'vue';
import { Eye } from 'lucide-vue-next';
import VendorInfoForm from './ui/VendorInfoForm.vue';
import PreviewModal from './ui/PreviewModal.vue';
import ManageFranchiseAbout from './ui/franchise/ManageFranchiseAbout.vue';

export default {
  name: 'ManageMainContent',
  components: {
    Eye,
    VendorInfoForm,
    PreviewModal
  },
  props: ['activeTab'],
  setup(props) {
    const isPreviewOpen = ref(false);

    const currentTabComponent = computed(() => {
      switch (props.activeTab) {
        case '關於品牌':
          return ManageFranchiseAbout;


        case '廠商資訊':
          return VendorInfoForm;
        default:
          return {
            template: `<p>這裡是 {{ activeTab }} 的內容編輯區域</p>`,
            props: ['activeTab']
          };
      }
    });

    return {
      isPreviewOpen,
      currentTabComponent
    };
  }
};
</script>