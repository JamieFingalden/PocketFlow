<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
	
	/* 修复页面显示不全的问题 */
	uni-page-body {
		height: auto;
		min-height: 100vh;
	}
	
	/* 确保页面容器能正确显示所有内容 */
	#app {
		height: auto;
		min-height: 100vh;
	}
</style>
