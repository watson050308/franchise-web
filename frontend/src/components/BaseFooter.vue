<template>
  <footer class="bg-indigo-600 text-white text-center py-4 mt-8">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center text-white">
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