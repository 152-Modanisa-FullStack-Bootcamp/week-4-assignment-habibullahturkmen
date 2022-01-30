<template>
  <div class="video-container">
    <!--  The coverImage of the video, when mouseover, shows the hoverImage, else coverImage  -->
    <div>
      <span @mouseover="hover = true" @mouseleave="hover = false" @click="navigateToVideo">
        <img v-if="hover" class="cover-image" :src="video.hoverImage" alt="product.name">
        <img v-else class="cover-image" :src="video.coverImage" alt="product.name">
      </span>
    </div>

    <!--  This div contains the ownerImage as well as video details  -->
    <div id="details-and-owner-image">
      <!--   shows the ownerImage   -->
      <img id="ownerImage" :src="video.ownerImage" alt="product.name">
      <!--   shows the video title, ownerName, viewCount and publishDateInMonth   -->
      <div id="details">
        <h4 id="title">{{ video.title }}</h4>
        <p>{{ video.ownerName }}</p>
        <p><span>{{ video.viewCount }} views</span>&nbsp;•&nbsp;<span>{{ video.publishDateInMonth }} months ago</span></p>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "Video",
  props: {
    video: Object
  },
  data() {
    return {
      hover: false
    }
  },
  methods: {
    navigateToVideo() {
      this.$router.push({ path: 'watch', query: { id: this.video.id }});
    }
  }
}
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}

/*
this is the main container for every single video and it's using flexbox
 */
.video-container {
  display: flex;
  flex-direction: column;
  width: 360px;
  height: 300px;
}

/*
the required height and width for the cover image
 */
.cover-image {
  width: 360px;
  height: 200px;
}

.cover-image:hover {
  cursor: pointer;
}

/*
this is the container for ownerImage and video details for every single video and it's also using flexbox
 */
#details-and-owner-image {
  display: flex;
  flex-direction: row;
  width: 360px;
  height: 80px;
  margin-top: 10px;
}

/*
the required height and width for the ownerImage
 */
#ownerImage {
  width: 36px;
  height: 36px;
  margin-right: 10px;
  border-radius: 20px;
}

/*
settings for the details div which is also using flexbox
 */
#details {
  display: flex;
  flex-direction: column;
  width: 324px;
  height: 70px;
}

#details > h4 {
  margin: 5px 0;
}

</style>
