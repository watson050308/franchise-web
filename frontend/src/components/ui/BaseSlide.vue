<template>
  <div class="relative h-[500px] overflow-hidden" 
       @mouseenter="isHovering = true" 
       @mouseleave="isHovering = false">
    <transition-group name="fade">
      <div v-for="(slide, index) in slides" :key="index" v-show="currentSlide === index" class="absolute top-0 left-0 w-full h-full">
        <img :src="slide.image" :alt="slide.alt" class="w-full h-full object-cover">
      </div>
    </transition-group>

    <!-- Left Control -->
    <transition name="fade">
      <button v-show="isHovering" 
              @click="prevSlide" 
              class="absolute left-4 top-1/2 transform -translate-y-1/2 bg-black bg-opacity-50 text-white p-2 rounded-full focus:outline-none hover:bg-opacity-75 transition duration-300">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
        </svg>
      </button>
    </transition>

    <!-- Right Control -->
    <transition name="fade">
      <button v-show="isHovering" 
              @click="nextSlide" 
              class="absolute right-4 top-1/2 transform -translate-y-1/2 bg-black bg-opacity-50 text-white p-2 rounded-full focus:outline-none hover:bg-opacity-75 transition duration-300">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </button>
    </transition>

    <!-- Slide Indicators -->
    <div class="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex space-x-2">
      <button v-for="(slide, index) in slides" 
              :key="index" 
              @click="goToSlide(index)"
              :class="['w-3 h-3 rounded-full focus:outline-none transition duration-300', 
                       currentSlide === index ? 'bg-white' : 'bg-white bg-opacity-50']">
      </button>
    </div>

  </div>
</template>

<script>
import orange from '@/assets/category/redorange.png'
import combuy from '@/assets/category/combuy.png'

export default {
  name: 'BaseSlide',
  data() {
    return {
      currentSlide: 0,
      slides: [
        { image: orange, alt: 'orange' },
        { image: combuy, alt: 'combuy' },
      ],
      intervalId: null,
      isHovering: false
    }
  },
  mounted() {
    this.startSlideShow()
  },
  beforeUnmount() {
    this.stopSlideShow()
  },
  methods: {
    startSlideShow() {
      this.intervalId = setInterval(this.nextSlide, 5000)
    },
    stopSlideShow() {
      clearInterval(this.intervalId)
    },
    nextSlide() {
      this.currentSlide = (this.currentSlide + 1) % this.slides.length
    },
    prevSlide() {
      this.currentSlide = (this.currentSlide - 1 + this.slides.length) % this.slides.length
    },
    goToSlide(index) {
      this.currentSlide = index
    },
    resetSlideShowTimer() {
      this.stopSlideShow()
      this.startSlideShow()
    }
  },
  watch: {
    currentSlide() {
      this.resetSlideShowTimer()
    }
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease-in-out;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>