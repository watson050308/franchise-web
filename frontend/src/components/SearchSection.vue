<template>
    <section class="search">
      <div class="search-box">
        <div class="location">
          <label>Location</label>
          <input type="text" placeholder="Where do you want to go?" />
        </div>
        <div class="date">
          <label>Date</label>
          <select>
            <option>Choose date</option>
          </select>
        </div>
        <div class="guests">
          <label>Guests</label>
          <div class="guest-input">
            <button>-</button>
            <span>Add guest</span>
            <button>+</button>
          </div>
        </div>
        <button class="search-button">Search</button>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    name: 'SearchSection'
  }
  </script>
  
  <style scoped>
.search-box {
  display: flex;
  justify-content: space-between;
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
}

.search-box > div {
  flex: 1;
  margin: 0 10px;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #666;
}

input, select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.guest-input {
  display: flex;
  align-items: center;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.guest-input button {
  padding: 10px;
  font-size: 18px;
  color: #3d3b54;
}

.guest-input span {
  flex: 1;
  text-align: center;
}

.search-button {
  background-color: #3d3b54;
  color: white;
  padding: 10px 20px;
  border-radius: 5px;
}
</style>