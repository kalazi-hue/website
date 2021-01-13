<template>
</template>

<script>
export default {
  data () {
    return {
    }
  },
	computed: {
	  contractFile() {
	    return this.$store.state.adList
	  }
	},
	methods: {
		fetchAdList (adList) {
	  	let subLunboList = [],subTopList = [],subBottomList = [],subLeftList = [],subRightList = []
	    for (let i = 0; i < adList.length; i++) {
	      if (adList[i].position.indexOf('轮播') > -1) {
	        subLunboList.push(adList[i])
	      } else if (adList[i].position.indexOf('顶部') > -1) {
	        subTopList.push(adList[i])
	      } else if (adList[i].position.indexOf('底部') > -1) {
	        subBottomList.push(adList[i])
	      } else if (adList[i].position.indexOf('左飘') > -1) {
	        subLeftList.push(adList[i])
	      } else if (adList[i].position.indexOf('右飘') > -1) {
	        subRightList.push(adList[i])
	      }
	    }
	    this.$emit('lunboData', subLunboList)
	    this.$emit('adTopData', subTopList)
	    this.$emit('adBottomData', subBottomList)
	    this.$emit('adLeftData', subLeftList)
	    this.$emit('adRightData', subRightList)
		}
	},
	watch: {
    contractFile: {
			immediate: true,
			handler: function (newval) {
  			this.fetchAdList(newval)
			}
    }
  }
}
</script>

