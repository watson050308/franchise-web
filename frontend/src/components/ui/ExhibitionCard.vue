<template>
  <div class="relative h-72 rounded-lg overflow-hidden shadow-md cursor-pointer transition duration-300 transform hover:-translate-y-2 hover:shadow-xl group">
    <img :src="exhibition.image" :alt="exhibition.title" class="w-full h-full object-cover transition duration-300">
    <div class="absolute inset-0 bg-black opacity-0 group-hover:opacity-70 transition-opacity duration-300"></div>
    <div class="absolute inset-0 flex flex-col justify-center items-center p-4 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
      <h2 class="text-2xl font-semibold text-white text-center mb-2">{{ exhibition.title }}</h2>
      <p class="text-white text-center mb-4">{{ exhibition.description }}</p>
      <div class="text-sm text-pink-300">
        {{ exhibition.subcategories.join(' • ') }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ExhibitionCard',
  props: {
    exhibition: {
      type: Object,
      required: true
    }
  }
}
</script>