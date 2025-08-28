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
	import { ref, computed } from 'vue'
	
	// 响应式数据
	const currentDate = ref('2023年8月27日')
	const daysSinceStart = ref(45) // 假设用户已记账45天
	const income = ref('5,230.00')
	const expense = ref('3,150.00')
	const balance = ref('2,080.00')
	const monthlyBudget = ref(5000.00)
	const budgetUsed = ref(3150.00)
	const recentBills = ref([
		{ categoryIcon: '餐', category: '餐饮', remark: '星巴克咖啡', amount: '-35.00', time: '今天 12:30', type: 'expense' },
		{ categoryIcon: '交', category: '交通', remark: '地铁', amount: '-5.00', time: '今天 08:15', type: 'expense' },
		{ categoryIcon: '购', category: '购物', remark: '京东购物', amount: '-128.50', time: '昨天 20:15', type: 'expense' },
		{ categoryIcon: '医', category: '医疗', remark: '药店', amount: '-45.00', time: '8月25日', type: 'expense' }
	])
	
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
</script>

<style>
	.container {
		padding: 20rpx;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 30rpx;
	}

	.user-info {
		display: flex;
		flex-direction: column;
	}

	.welcome {
		font-size: 28rpx;
		color: #999;
	}

	.username {
		font-size: 36rpx;
		font-weight: bold;
		margin-top: 10rpx;
	}

	.date {
		font-size: 28rpx;
		color: #999;
	}

	.summary {
		display: flex;
		justify-content: space-between;
		background-color: #fff;
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
		color: #999;
		margin-bottom: 10rpx;
	}

	.amount {
		font-size: 32rpx;
		font-weight: bold;
	}

	.income {
		color: #00b38a;
	}

	.expense {
		color: #ff6b6b;
	}

	.balance {
		color: #4a90e2;
	}

	.budget-progress {
		background-color: #fff;
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
		background: linear-gradient(90deg, #4a90e2, #5cdbd3);
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
		background-color: #007AFF;
		color: #fff;
		font-size: 60rpx;
		font-weight: bold;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 8rpx 20rpx rgba(0, 122, 255, 0.3);
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
		background-color: #007AFF;
		color: #fff;
		font-size: 32rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 4rpx 10rpx rgba(0, 122, 255, 0.3);
	}

	.recent-bills {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.section-title {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 30rpx;
		font-size: 32rpx;
		font-weight: bold;
	}

	.more {
		font-size: 28rpx;
		color: #007AFF;
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