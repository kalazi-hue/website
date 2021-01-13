<template>
   <card-swipe
      :autoplay="true"
      :images="images"
      @change="onChange"
      class="swiper"
    />
</template>

<script>
export default {
  name: 'example',
  components: {
    'card-swipe': () => import('@/components/CardSwiper.vue')
  },
  data () {
    return {
      path: '',
      images: []
    }
  },
  methods: {
    onChange (index, oldIndex) {
      // console.log(`curIndex：${index}`, `oldIndex：${oldIndex}`)
    }
  },
  mounted () {
    const img = ['/www.1990tu.com/i/20201206105926tm7.png',
                  '/www.1990tu.com/i/20201206105904pyl.png',
                  '/www.1990tu.com/i/20201206105845mzv.png'
    ]
    this.images = img.map(item => `${this.path}/${item}`)
  }
}
</script>

<style scoped >
.swiper {
  margin: 0 auto; 
}
</style>