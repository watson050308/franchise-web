<template>
  <div id="app">
    <header>
      <nav>
        <div class="logo">Tour.ink</div>
        <div class="nav-links">
          <a href="#" class="active">Home</a>
          <a href="#">Booking Form</a>
          <a href="#">Category</a>
          <a href="#">About Us</a>
        </div>
        <div class="auth-buttons">
          <button class="log-in">Log in</button>
          <button class="sign-up">Sign up</button>
        </div>
      </nav>
    </header>
    
    <main>
      <hero-section></hero-section>
      <search-section></search-section>
      <popular-search></popular-search>
    </main>
  </div>
</template>

<script>
import HeroSection from './components/HeroSection.vue'
import SearchSection from './components/SearchSection.vue'
import PopularSearch from './components/PopularSearch.vue'

export default {
  name: 'App',
  components: {
    HeroSection,
    SearchSection,
    PopularSearch
  }
}
</script>

<style>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #f8f7ff;
}

#app {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

button {
  cursor: pointer;
  border: none;
  background: none;
}

/* 導航欄樣式 */
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 0;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: #3d3b54;
}

.nav-links a {
  margin: 0 15px;
  text-decoration: none;
  color: #3d3b54;
}

.nav-links a.active {
  font-weight: bold;
}

.auth-buttons button {
  padding: 10px 20px;
  margin-left: 10px;
  border-radius: 20px;
}

.log-in {
  color: #3d3b54;
}

.sign-up {
  background-color: #3d3b54;
  color: white;
}
</style>