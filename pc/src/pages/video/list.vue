<template>
  <div class="container">
    <div class="video-page">

      <tagList />

      <video-list />

    </div>
  </div>
</template>

<script>
import videoList from '@/components/videoList'
import tagList from '@/components/tags'
export default {
  components: {
    videoList,
    tagList
  },
  data () {
    return {
    }
  }
}
</script>

<style>
</style>
