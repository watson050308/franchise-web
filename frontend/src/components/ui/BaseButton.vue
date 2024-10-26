<template>
  <button
    :class="[
      'px-4 py-2 rounded font-semibold focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500',
      variant === 'outline' ? 'border border-current' : 'text-white',
      block ? 'w-full' : '',
      $attrs.class
    ]"
  >
    <slot></slot>
  </button>
</template>
  
<script>
export default {
  name: 'BaseButton',
  props: {
    variant: {
      type: String,
      default: 'default'
    }
  }
}
</script>