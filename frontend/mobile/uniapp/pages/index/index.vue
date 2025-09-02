<template>
	<view class="container">
		<!-- 顶部用户信息区 -->
		<view class="header">
			<view class="user-info">
				<text class="welcome">欢迎回来</text>
				<text class="username">用户名</text>
			</view>
			<view class="date">{{ currentDate }}</view>
		</view>
		<!-- 确保没有残留元素 -->

				
		<!-- 预算进度条 -->
		<view class="budget-progress">
			<view class="budget-header">
				<text class="budget-label">本月预算</text>
				<text class="budget-amount">¥{{ budgetUsed }} / ¥{{ monthlyBudget }}</text>
			</view>
			<view class="progress-bar">
				<view 
					class="progress-fill" 
					:style="{ width: budgetProgress + '%' }"
				></view>
			</view>
			<view class="progress-text">
				<text :class="budgetStatusClass">{{ budgetStatusText }}</text>
			</view>
		</view>

		<!-- 快捷记账按钮 -->
		<view class="quick-add">
			<button class="add-btn" @click="goToAddBill">+</button>
		</view>
		
		<!-- OCR上传按钮 -->
		<view class="ocr-section">
			<button class="ocr-btn" @click="uploadBillImage">
				📷 上传账单图片
			</button>
		</view>

		<!-- 最近账单列表 -->
		<view class="recent-bills">
			<view class="section-title">
				<text>最近账单</text>
				<text class="more" @click="goToBills">更多 ></text>
			</view>
			<view class="bill-list">
				<view class="bill-item" v-for="(bill, index) in recentBills" :key="index">
					<view class="bill-icon">
						<text>{{ bill.categoryIcon }}</text>
					</view>
					<view class="bill-info">
						<text class="bill-category">{{ bill.category }}</text>
						<text class="bill-remark">{{ bill.remark }}</text>
					</view>
					<view class="bill-amount">
						<text class="amount" :class="bill.type">{{ bill.amount }}</text>
						<text class="bill-time">{{ bill.time }}</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed, onMounted, nextTick } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	
	// 响应式数据
	const monthlyBudget = ref(5000.00)
	const budgetUsed = ref(3150.00)
	const recentBills = ref([
		{ categoryIcon: '🍽️', category: '餐饮', remark: '星巴克咖啡', amount: '-35.00', time: '今天 12:30', type: 'expense' },
		{ categoryIcon: '🚗', category: '交通', remark: '地铁', amount: '-5.00', time: '今天 08:15', type: 'expense' },
		{ categoryIcon: '🛍️', category: '购物', remark: '京东购物', amount: '-128.50', time: '昨天 20:15', type: 'expense' },
		{ categoryIcon: '🏥', category: '医疗', remark: '药店', amount: '-45.00', time: '8月25日', type: 'expense' }
	])
	const overlayVisible = ref(false)
	
	// 计算属性
	const budgetProgress = computed(() => {
		if (monthlyBudget.value <= 0) return 0
		return Math.min((budgetUsed.value / monthlyBudget.value) * 100, 100)
	})
	
	const budgetStatusText = computed(() => {
		const remaining = monthlyBudget.value - budgetUsed.value
		if (remaining < 0) {
			return `超出预算 ¥${Math.abs(remaining).toFixed(2)}`
		} else {
			return `剩余 ¥${remaining.toFixed(2)}`
		}
	})
	
	const budgetStatusClass = computed(() => {
		const remaining = monthlyBudget.value - budgetUsed.value
		return remaining < 0 ? 'over-budget' : 'under-budget'
	})
	
	// 计算属性 - 当前日期
	const currentDate = computed(() => {
		const date = new Date()
		const year = date.getFullYear()
		const month = date.getMonth() + 1
		const day = date.getDate()
		const days = ['日', '一', '二', '三', '四', '五', '六']
		const weekDay = days[date.getDay()]
		return `${year}年${month}月${day}日 星期${weekDay}`
	})
	
	// 方法定义
	const goToAddBill = () => {
		uni.navigateTo({
			url: '/pages/bills/add-bill'
		})
	}
	
	const uploadBillImage = () => {
		// 上传账单图片
		uni.chooseImage({
			count: 1,
			success: (res) => {
				console.log('选择的图片:', res.tempFilePaths)
				uni.showToast({
					title: '图片上传成功',
					icon: 'success'
				})
				// 跳转到添加账单页面，并传递图片路径
				uni.navigateTo({
					url: `/pages/bills/add-bill?imagePath=${res.tempFilePaths[0]}`
				})
			},
			fail: (err) => {
				console.error('选择图片失败:', err)
				uni.showToast({
					title: '选择图片失败',
					icon: 'none'
				})
			}
		})
	}
	
	const goToBills = () => {
		uni.switchTab({
			url: '/pages/bills/bills'
		})
	}
	
	// 初始化
	onMounted(() => {
		// 强制更新组件以确保样式正确应用
		nextTick(() => {
			// 这里可以添加任何需要在DOM更新后执行的代码
			console.log('首页组件已挂载并更新')
		})
	})
	
	// 页面显示时的处理函数
	onShow(() => {
		// 当页面显示时强制更新样式
		console.log('首页页面显示')
		// 强制更新数据以触发重新渲染
		monthlyBudget.value = monthlyBudget.value
		budgetUsed.value = budgetUsed.value
		// 触发计算属性重新计算
		budgetProgress.value = budgetProgress.value
		
		// 添加调试信息
		setTimeout(() => {
			console.log('首页显示完成，检查是否有异常元素');
		}, 500);
	})
</script>

<style scoped>
	.container {
		height: 100%;
		padding: 20rpx;
		background-color: #F0F3F5;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background: linear-gradient(135deg, #5AA9E6, #8BC34A);
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 30rpx;
		color: #fff;
		box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.1);
	}

	.user-info {
		display: flex;
		flex-direction: column;
	}

	.welcome {
		font-size: 28rpx;
		color: #fff;
		opacity: 0.9;
	}

	.username {
		font-size: 36rpx;
		font-weight: bold;
		margin-top: 10rpx;
		color: #fff;
	}

	.date {
		font-size: 28rpx;
		color: #fff;
		opacity: 0.9;
	}

	.summary {
		display: flex;
		justify-content: space-between;
		background-color: #FFFFFF;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.summary-item {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.label {
		font-size: 28rpx;
		color: #6D7378;
		margin-bottom: 10rpx;
	}

	.amount {
		font-size: 32rpx;
		font-weight: bold;
	}

	.income {
		color: #8BC34A;
	}

	.expense {
		color: #FF7F50;
	}

	.balance {
		color: #5AA9E6;
	}

	.budget-progress {
		background-color: #FFFFFF;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.budget-header {
		display: flex;
		justify-content: space-between;
		margin-bottom: 20rpx;
	}

	.budget-label {
		font-size: 32rpx;
		font-weight: bold;
	}

	.budget-amount {
		font-size: 28rpx;
		color: #999;
	}

	.progress-bar {
		width: 100%;
		height: 20rpx;
		background-color: #f0f0f0;
		border-radius: 10rpx;
		margin-bottom: 20rpx;
		overflow: hidden;
	}

	.progress-fill {
		height: 100%;
		border-radius: 10rpx;
		background: linear-gradient(135deg, #5AA9E6, #8BC34A);
		transition: width 0.3s ease;
	}

	.progress-text {
		text-align: right;
		font-size: 24rpx;
	}

	.under-budget {
		color: #00b38a;
	}

	.over-budget {
		color: #ff6b6b;
	}

	.quick-add {
		display: flex;
		justify-content: center;
		margin: 50rpx 0;
	}

	.add-btn {
		width: 120rpx;
		height: 120rpx;
		border-radius: 50%;
		background: linear-gradient(135deg, #5AA9E6, #8BC34A);
		color: #fff;
		font-size: 60rpx;
		font-weight: bold;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 8rpx 20rpx rgba(90, 169, 230, 0.3);
	}

	.ocr-section {
		display: flex;
		justify-content: center;
		margin: 20rpx 0;
	}

	.ocr-btn {
		width: 80%;
		height: 80rpx;
		border-radius: 40rpx;
		background: linear-gradient(135deg, #5AA9E6, #8BC34A);
		color: #fff;
		font-size: 32rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 4rpx 10rpx rgba(90, 169, 230, 0.3);
	}

	.recent-bills {
		background-color: #FFFFFF;
		border-radius: 20rpx;
		padding: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}


	.more {
		font-size: 28rpx;
		color: #8BC34A;
	}

	.bill-list {
		display: flex;
		flex-direction: column;
		gap: 20rpx;
	}

	.bill-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 20rpx 0;
		border-bottom: 1rpx solid #f0f0f0;
	}

	.bill-item:last-child {
		border-bottom: none;
	}

	.bill-icon {
		width: 60rpx;
		height: 60rpx;
		border-radius: 50%;
		background-color: #e0f7fa;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-right: 20rpx;
	}

	.bill-info {
		flex: 1;
		display: flex;
		flex-direction: column;
	}

	.bill-category {
		font-size: 28rpx;
		font-weight: bold;
		margin-bottom: 5rpx;
		color: #2E2E2E;
	}

	.bill-remark {
		font-size: 24rpx;
		color: #999;
	}

	.bill-amount {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
	}

	.bill-time {
		font-size: 24rpx;
		color: #999;
		margin-top: 5rpx;
	}
</style>