<template>
  <div class="header-container">
    <button id="home-button" @click="goToHome"><strong>Home</strong></button>
    <button id="favorites-button" @click="goToFavorites"><strong>Favorites</strong></button>
  </div>
</template>

<script>
export default {
  name: "Header",
  methods: {
    goToFavorites() {
      this.$router.push("/favorites/bootcamp");
    },
    goToHome() {
      this.$router.push("/");
    }
  }
}
</script>

<style scoped>
.header-container {
  display: flex;
  width: 100%;
  height: 50px;
  background-color: rebeccapurple;
  justify-content: space-between;
}

#favorites-button {
  margin: 10px 20px 10px 0;
  height: 30px;
  width: 100px;
}

#favorites-button {
  cursor: pointer;
}

#home-button {
  margin: 10px 0 10px 20px;
  height: 30px;
  width: 100px;
}

#home-button {
  cursor: pointer;
}
</style>
