<template>
  <footer class="bg-gray-100 py-6 sm:py-8">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center text-gray-600">
      <p>&copy; {{ currentYear }} FranchiseConnect. 保留所有權利。</p>
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