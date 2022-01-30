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
    <div id="details">
      <!--   shows the video title, viewCount and publishDateInMonth   -->
      <h3 >{{ video.title }}</h3>
      <p><span>{{ video.viewCount }} views</span>&nbsp;•&nbsp;<span>{{ video.publishDateInMonth }} months ago</span></p>
      <!--   shows the owner name and owner image   -->
      <div id="owner-name-and-image">
        <img id="owner-image" :src="video.ownerImage" alt="product.name">
        <p id="owner-name">{{ video.ownerName }}</p>
      </div>
      <p>{{ video.description }}</p>
    </div>


  </div>
</template>

<script>
export default {
  name: "FavoriteVideo",
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
      this.$router.push({ path: '/watch', query: { id: this.video.id }});
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
  flex-direction: row;
  width: 1000px;
  height: 200px;
  margin-bottom: 50px;
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
the required height and width for the ownerImage
 */
#owner-image {
  width: 24px;
  height: 24px;
  margin-right: 10px;
  margin-left: 5px;
  border-radius: 20px;
}

/*
settings for the details div which is also using flexbox
 */
#details {
  display: flex;
  flex-direction: column;
  width: 750px;
  height: 160px;
}

#details > h3 {
  margin: 5px 0 5px 15px;
}

#details > p {
  margin: 5px 0 5px 15px;
}

#owner-name-and-image {
  display: flex;
  margin: 5px 0 5px 10px;
}

#owner-name-and-image > p {
  margin-top: 3px;
}
</style>
