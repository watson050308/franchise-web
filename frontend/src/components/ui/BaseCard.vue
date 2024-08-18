<template>
  <div class="bg-white shadow rounded-lg overflow-hidden">
    <div v-if="$slots.header" class="px-4 py-5 sm:px-6">
      <slot name="header"></slot>
    </div>
    <div class="px-4 py-5 sm:p-6">
      <slot name="content"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BaseCard'
}
</script>