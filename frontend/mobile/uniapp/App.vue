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
uni-page-wrapper::after {
	display: none !important;
	content: none !important;
}

uni-page-body {
	height: 100%;
}
</style>
