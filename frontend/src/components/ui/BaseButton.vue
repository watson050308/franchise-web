<template>
    <button
      :class="[
        'px-4 py-2 rounded',
        variant === 'outline' ? 'border border-current' : 'text-white',
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