<template>
  <div class="contact-page">
    <h5>联系我们</h5>
    <div class="contact-list">
    	<h6>商务合作</h6>
			<p>Telegram：
				<el-tooltip class="item" effect="dark" content="点击打开" placement="right">
					<a href="https://t.me/niuav">@niuav</a>
    		</el-tooltip>
			</p>
			<p>Skype
				<el-tooltip class="item" effect="dark" content="点击打开" placement="right">
					<a href="https://t.me/niuav">@niuav</a>
    		</el-tooltip>
			</p>
			<p>Potato
				<el-tooltip class="item" effect="dark" content="点击打开" placement="right">
					<a href="https://t.me/niuav">@niuav</a>
    		</el-tooltip>
			</p>
			<p>邮箱
				<el-tooltip class="item" effect="dark" content="点击打开" placement="right">
					<a href="mailto:111@gmail.com" title="点击发送邮件">mailto:111@gmail.com</a>
    		</el-tooltip>
			</p>
    </div>
  </div>
</template>

<style lang="scss" scoped>
	.contact-page {
		min-height: 550px;
		font-size: 20px;
		padding: 20px 0 100px;
		h5 {
			font-size: 36px;
			padding: 20px 0;
			border-bottom: 1px solid #615a4e;
		}
		.contact-list {
			h6 {
				font-size: 20px;
				margin-top: 50px;
			}
			p {
    		color: #ccc;
				font-size: 18px;
				margin-top: 20px;
				a {
					color: #ccc;
					margin-left: 10px;
				}
			}
		}
	}
</style>