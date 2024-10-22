<!-- App.vue -->
<template>
  <div class="flex min-h-screen">
    <SideBar :activeTab="activeTab" @setActiveTab="setActiveTab" />
    <ManageMainContent :activeTab="activeTab" />
  </div>
</template>

<script>
import { ref } from 'vue';
import SideBar from './components/ui/Sidebar.vue';
import ManageMainContent from './components/ManageMainContent.vue';

export default {
  name: 'AppManagement',
  components: {
    SideBar,
    ManageMainContent
  },
  setup() {
    const activeTab = ref('主頁');
    const setActiveTab = (tab) => {
      activeTab.value = tab;
    };

    return {
      activeTab,
      setActiveTab
    };
  }
};
</script>