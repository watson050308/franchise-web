<!-- components/ui/SideBar.vue -->
<template>
    <div class="w-64 min-h-screen bg-gray-800 text-white p-4 flex flex-col">
      <h1 class="text-2xl font-bold mb-8">後台管理</h1>
      <nav class="flex-grow">
        <button
          v-for="tab in tabs"
          :key="tab.name"
          class="flex items-center w-full p-2 mb-2 rounded"
          :class="{ 'bg-blue-600': activeTab === tab.name, 'hover:bg-gray-700': activeTab !== tab.name }"
          @click="$emit('setActiveTab', tab.name)"
        >
          <component :is="tab.icon" class="mr-2" :size="20" />
          {{ tab.name }}
        </button>
      </nav>
    </div>
  </template>
  
  <script>
  import { Home, Info, Package, Award, GitBranch, Gift, MessageSquare, User } from 'lucide-vue-next';
  
  export default {
    name: 'SideBar',
    components: {
      Home, Info, Package, Award, GitBranch, Gift, MessageSquare, User
    },
    props: ['activeTab'],
    emits: ['setActiveTab'],
    setup() {
      const tabs = [
        { name: '主頁', icon: Home },
        { name: '關於品牌', icon: Info },
        { name: '產品服務', icon: Package },
        { name: '加盟優勢', icon: Award },
        { name: '加盟流程', icon: GitBranch },
        { name: '優惠活動', icon: Gift },
        { name: '聯繫我們', icon: MessageSquare },
        { name: '廠商資訊', icon: User },
      ];
  
      return {
        tabs
      };
    }
  };
  </script>