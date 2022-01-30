<template>
  <div>
    <Header></Header>
    <div class="video-list-container">
      <FavoriteVideo v-for="video in getFavorites"
                     :video="video"
                     :key="video.id"
      ></FavoriteVideo>
    </div>
  </div>
</template>

<script>
import Header from "../components/Header";
import FavoriteVideo from "../components/FavoriteVideo";
import axios from "axios";

export default {
  name: "Favorite",
  components: {
    Header: Header,
    FavoriteVideo: FavoriteVideo
  },
  data() {
    return {
      videos: []
    }
  },
  computed: {
    getFavorites() {
      return this.videos.filter(n => n.favorite === true);
    }
  },
  async mounted() {
    const response = await axios.get("https://my-json-server.typicode.com/modanisa/bootcamp-video-db/videos");
    this.videos = response.data;
    console.log(this.videos);
  }
}
</script>

<style scoped>
.video-list-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  flex-direction: column;
  width: 1600px;
  height: 1100px;
  padding-top: 80px;
  padding-left: 250px;
}
</style>
