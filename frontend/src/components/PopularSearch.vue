<template>
    <section class="popular-search">
      <h3>Popular Search</h3>
      <div class="search-tags">
        <button v-for="tag in popularTags" :key="tag">{{ tag }}</button>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    name: 'PopularSearch',
    data() {
      return {
        popularTags: [
          'Lotus Temple', 'Eiffel Tower', 'Batu Gong', 'Prambanan Temple',
          'Borobudur Temple', 'Taman Sari', 'Labuan Bajo', 'Seminyak Beach',
          'The Grand Palace', 'Wat Phra Kaew'
        ]
      }
    }
  }
  </script>
  
  <style scoped>
.popular-search {
  margin-top: 40px;
}

h3 {
  color: #3d3b54;
  margin-bottom: 20px;
}

.search-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.search-tags button {
  background-color: white;
  border: 1px solid #ddd;
  padding: 10px 20px;
  border-radius: 20px;
  color: #3d3b54;
  transition: all 0.3s ease;
}

.search-tags button:hover {
  background-color: #3d3b54;
  color: white;
}
</style>