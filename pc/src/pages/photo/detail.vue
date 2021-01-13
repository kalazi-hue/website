<template>
  <div class="container">
  	<h2 class="title">图集浏览<span class="total">（共{{ photoList.length || 0}}张）</span></h2>
    <div class="photo-detail">
  		<el-image 
  			v-for="(item, index) in photoList" 
  			:key="index" 
  			:src="item" 
		    :preview-src-list="photoList"
		    title="点击可放大浏览"
  			lazy>
  			</el-image>
  			<p class="nomore">已到底了~</p>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
    	photoList: []
    }
  },
  created () {
    this.getDetail()
  },
  methods: {
    getDetail () {
      if (this.$route.query.content) {
        this.photoList = []
        this.photoList = this.$route.query.content.split(',')
      }
    }
  },
  watch: {
    '$route':'getDetail'
  }
}
</script>

<style scoped lang="scss">
.container {
		margin: 0 0 100px;
	.total {
		font-size: 24px;
		color: #e9d919;
	}
	.photo-detail {
		.el-image {
			display: block;
			width: 100%;
			margin: 0 auto;
		}
		.nomore {
	    text-align: center;
	    margin: 30px auto;
	    color: #ccc;
	    font-size: 20px;			
		}
	}
}
</style>

<style>	
.el-image__preview {
  width: auto!important;
  display: block;
  margin: 0 auto 20px;
}
</style>