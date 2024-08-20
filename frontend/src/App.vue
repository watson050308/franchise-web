<template>
  <div class="bg-white min-h-screen flex flex-col">
    <BaseHeader />
      <router-view></router-view>
    <BaseFooter />
  </div>
</template>

<script>
import BaseHeader from './components/BaseHeader.vue'
import BaseFooter from './components/BaseFooter.vue'


export default {
  name: 'App',
  components: {
    BaseHeader,
    BaseFooter
  }
}
</script>