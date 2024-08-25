<!-- components/ui/VendorInfoForm.vue -->
<template>
    <form @submit.prevent="submitForm" class="space-y-4">
      <div v-for="field in fields" :key="field.name">
        <label class="block mb-1">{{ field.label }}</label>
        <input
          :type="field.type"
          v-model="formData[field.name]"
          class="w-full border rounded p-2"
        />
      </div>
      <button type="submit" class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">
        保存
      </button>
    </form>
  </template>
  
  <script>
  import { reactive } from 'vue';
  
  export default {
    name: 'VendorInfoForm',
    setup() {
      const fields = [
        { name: 'account', label: '帳號', type: 'text' },
        { name: 'password', label: '密碼', type: 'password' },
        { name: 'brandName', label: '品牌名稱', type: 'text' },
        { name: 'companyName', label: '公司名稱', type: 'text' },
        { name: 'contactPerson', label: '聯絡人', type: 'text' },
        { name: 'phoneNumber', label: '聯絡電話', type: 'tel' },
        { name: 'address', label: '聯絡地址', type: 'text' },
        { name: 'businessNumber', label: '統一編號', type: 'text' },
      ];
  
      const formData = reactive({});
  
      const submitForm = () => {
        console.log('Form submitted:', formData);
        // Here you would typically send the data to your backend
      };
  
      return {
        fields,
        formData,
        submitForm
      };
    }
  };
  </script>