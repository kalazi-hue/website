<template>
  <special-banner :banner-list="bannerList" :swiper-config="swiperConfig"></special-banner>
</template>

<script>
  import specialBanner from '@/components/common/specialBanner.vue';
  export default {
    components: {
      specialBanner
    },
    data() {
        return {
            bannerList: [{
                picture: 'http://image.mishi.cn/r/yry_h5_test/detail/3_1535359279285.png',
                path: ''
            }, {
                picture: 'http://image.mishi.cn/r/yry_h5_test/detail/2_1535359240426.png',
                path: ''
            }, {
                picture: 'http://image.mishi.cn/r/yry_h5_test/detail/1_1535359204228.png',
                path: ''
            }, {
                picture: 'http://image.mishi.cn/r/yry_h5_test/detail/4_1535359327213.png',
                path: ''
            }],
            swiperConfig: {
                indicatorDots: true,
                indicatorColor: 'rgba(255, 255, 255, .4)',
                indicatorActiveColor: 'rgba(255, 255, 255, 1)',
                autoplay: false,
                interval: 3000,
                duration: 300,
                circular: true,
                previousMargin: '58rpx',
                nextMargin: '58rpx'
            }
        }
    }
  }
</script>

<style>
  .content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .logo {
    height: 200rpx;
    width: 200rpx;
    margin-top: 200rpx;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 50rpx;
  }
  .text-area {
    display: flex;
    justify-content: center;
  }
  .title {
    font-size: 36rpx;
    color: #8f8f94;
  }
</style>