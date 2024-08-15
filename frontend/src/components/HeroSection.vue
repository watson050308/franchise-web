<template>
    <section class="hero">
      <h2>All you need is Tourink</h2>
      <h1>Explore Beautiful Places</h1>
      <div class="carousel">
        <button class="prev">&lt;</button>
        <span>Lotus Temple</span>
        <button class="next">&gt;</button>
      </div>
      <img src="path_to_lotus_temple_image.jpg" alt="Lotus Temple" />
    </section>
  </template>
  
  <script>
  export default {
    name: 'HeroSection'
  }
  </script>
  
  <style scoped>
.hero {
  text-align: center;
  padding: 50px 0;
  background: linear-gradient(to bottom, #ffeeff, #f8f7ff);
  border-radius: 20px;
}

h2 {
  color: #666;
  font-size: 18px;
}

h1 {
  color: #3d3b54;
  font-size: 48px;
  margin: 20px 0;
}

.carousel {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px 0;
}

.carousel button {
  font-size: 24px;
  color: #3d3b54;
}

.carousel span {
  margin: 0 20px;
  font-size: 20px;
  color: #3d3b54;
}

img {
  max-width: 100%;
  height: auto;
  border-radius: 10px;
}
</style>