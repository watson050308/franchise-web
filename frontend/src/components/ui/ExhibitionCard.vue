<template>
  <router-link :to="{path: '/showbrand', query: {category: exhibition.label}}" class="block">
    <div class="relative h-72 rounded-lg overflow-hidden shadow-md cursor-pointer transition duration-300 transform hover:-translate-y-2 hover:shadow-xl group">
      <img :src="imageSrc" :alt="exhibition.title" class="w-full h-full object-cover transition duration-300">
      <div class="absolute inset-0 bg-black opacity-0 group-hover:opacity-70 transition-opacity duration-300"></div>
      <div class="absolute inset-0 flex flex-col justify-center items-center p-4 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <h2 class="text-2xl font-semibold text-white text-center mb-2">{{ exhibition.title }}</h2>
        <p class="text-white text-center mb-4">{{ exhibition.description }}</p>
        <div class="text-sm text-pink-300">
          {{ exhibition.subcategories.join(' • ') }}
        </div>
      </div>
    </div>
  </router-link>
</template>

<script>
export default {
  name: 'ExhibitionCard',
  props: {
    exhibition: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      imageSrc: ''
    }
  },
  created() {
    this.loadImage()
  },
  methods: {
    async loadImage() {
      try {
        if (typeof this.exhibition.image === 'function') {
          const { default: src } = await this.exhibition.image()
          this.imageSrc = src
        } else {
          this.imageSrc = this.exhibition.image
        }
      } catch (error) {
        console.error('Error loading image:', error)
      }
    }
  }
}
</script>