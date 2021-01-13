<template>
  <div class="container page-home">
    <video-list-home :showRecommendList="true" /> 
    <video-list-home :showLatestList="true" />
    <photo-list :isHome="isHome" />
  </div>
</template>

<script>
import { API } from '@/api'
import videoListHome from '@/components/videoListHome'
import photoList from '@/components/photoList'
export default {
  components: {
    videoListHome,
    photoList
  },
  data () {
    return {
      isHome: true
    }
  }
}
</script>

<style lang="scss" scoped>
  .page-home {
    min-height: 800px;
  }
</style>