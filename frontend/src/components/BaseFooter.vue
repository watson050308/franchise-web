<template>
    <footer class="bg-gray-100 py-8">
      <div class="max-w-6xl mx-auto px-4 text-center text-gray-600">
        <p>© {{ currentYear }} FranchiseConnect. 保留所有權利。</p>
      </div>
    </footer>
  </template>
  
  <script>
  export default {
    name: 'BaseFooter',
    computed: {
      currentYear() {
        return new Date().getFullYear()
      }
    }
  }
  </script>