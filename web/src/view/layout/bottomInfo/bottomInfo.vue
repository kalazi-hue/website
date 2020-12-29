<template>
  <div class="bottom-info">
    <div>
      <el-divider direction="vertical"></el-divider>
      <span>Copyright</span>
      <span>
      <a href="">牛牛技术团队</a>
      </span>
    </div>
  </div>
</template>
<script>
export default {
    name:"BottomInfo"
}
</script>
<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
