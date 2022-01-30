<template>
  <div>
    <Header></Header>
    <div class="video-list-container">
      <Video v-for="video in videos"
             :video="video"
             :key="video.id"
      ></Video>
    </div>
  </div>
</template>

<script>
import Header from "../components/Header";
import Video from "../components/Video";
import API from "../api";

export default {
  name: "Home",
  components: {
    Header: Header,
    Video: Video
  },
  data() {
    return {
      videos: []
    }
  },
  async created() {
    this.videos = await API.getVideos();
  }
}
</script>

<style scoped>

.video-list-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
  width: 1600px;
  height: 1100px;
  padding-top: 80px;
}
</style>
