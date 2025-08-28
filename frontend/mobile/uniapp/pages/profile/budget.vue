<template>
	<view class="container">
		<view class="header">
			<text class="title">预算设置</text>
		</view>
		
		<view class="budget-card">
			<view class="budget-header">
				<text class="month">{{ currentMonth }}月预算</text>
			</view>
			
			<view class="budget-amount">
				<text class="label">预算金额</text>
				<input 
					class="amount-input" 
					type="digit" 
					placeholder="0.00" 
					v-model="budgetAmount"
				/>
			</view>
			
			<view class="budget-progress" v-if="budgetAmount > 0">
				<view class="progress-info">
					<text class="progress-label">当前支出</text>
					<text class="progress-value">¥{{ currentExpense }} / ¥{{ budgetAmount }}</text>
				</view>
				<view class="progress-bar">
					<view 
						class="progress-fill" 
						:style="{ width: progressWidth }"
					></view>
				</view>
				<view class="progress-text">
					<text :class="progressClass">{{ progressText }}</text>
				</view>
			</view>
			
			<button class="save-btn" @click="saveBudget">保存预算</button>
		</view>
		
		<view class="history-section">
			<view class="section-title">
				<text>历史预算</text>
			</view>
			
			<view class="history-list">
				<view 
					class="history-item" 
					v-for="(item, index) in budgetHistory" 
					:key="index"
				>
					<text class="history-month">{{ item.month }}</text>
					<text class="history-amount">¥{{ item.amount }}</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed } from 'vue'
	
	// 响应式数据
	const currentMonth = ref(new Date().getMonth() + 1)
	const budgetAmount = ref('')
	const currentExpense = ref('1250.00')
	const budgetHistory = ref([
		{ month: '2023-07', amount: '5000.00' },
		{ month: '2023-06', amount: '4500.00' },
		{ month: '2023-05', amount: '5200.00' }
	])
	
	// 计算属性
	const progressWidth = computed(() => {
		if (!budgetAmount.value) return '0%'
		const progress = (parseFloat(currentExpense.value) / parseFloat(budgetAmount.value)) * 100
		return Math.min(progress, 100) + '%'
	})
	
	const progressText = computed(() => {
		if (!budgetAmount.value) return ''
		const progress = (parseFloat(currentExpense.value) / parseFloat(budgetAmount.value)) * 100
		if (progress <= 100) {
			return `已使用 ${progress.toFixed(1)}%`
		} else {
			return `超出预算 ¥${(parseFloat(currentExpense.value) - parseFloat(budgetAmount.value)).toFixed(2)}`
		}
	})
	
	const progressClass = computed(() => {
		if (!budgetAmount.value) return ''
		const progress = (parseFloat(currentExpense.value) / parseFloat(budgetAmount.value)) * 100
		return progress > 100 ? 'over-budget' : 'normal'
	})
	
	// 方法定义
	const saveBudget = () => {
		if (!budgetAmount.value) {
			uni.showToast({
				title: '请输入预算金额',
				icon: 'none'
			})
			return
		}
		
		// 保存预算逻辑
		uni.showToast({
			title: '预算保存成功',
			icon: 'success'
		})
		
		// 添加到历史记录
		const currentMonthStr = new Date().getFullYear() + '-' + String(currentMonth.value).padStart(2, '0')
		budgetHistory.value.unshift({
			month: currentMonthStr,
			amount: budgetAmount.value
		})
	}
</script>

<style>
	.container {
		padding: 20rpx;
		background-color: #f8f8f8;
		min-height: 100vh;
	}
	
	.header {
		text-align: center;
		padding: 20rpx;
		background-color: #fff;
		margin-bottom: 20rpx;
		border-radius: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.title {
		font-size: 36rpx;
		font-weight: bold;
		color: #333;
	}
	
	.budget-card {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.budget-header {
		text-align: center;
		margin-bottom: 30rpx;
	}
	
	.month {
		font-size: 32rpx;
		font-weight: bold;
		color: #333;
	}
	
	.budget-amount {
		margin-bottom: 30rpx;
	}
	
	.label {
		display: block;
		font-size: 28rpx;
		color: #666;
		margin-bottom: 15rpx;
	}
	
	.amount-input {
		width: 100%;
		height: 80rpx;
		border: 1rpx solid #e0e0e0;
		border-radius: 10rpx;
		padding: 0 20rpx;
		font-size: 36rpx;
		box-sizing: border-box;
		text-align: center;
	}
	
	.budget-progress {
		margin-bottom: 30rpx;
	}
	
	.progress-info {
		display: flex;
		justify-content: space-between;
		margin-bottom: 15rpx;
	}
	
	.progress-label, .progress-value {
		font-size: 28rpx;
		color: #666;
	}
	
	.progress-bar {
		width: 100%;
		height: 20rpx;
		background-color: #f0f0f0;
		border-radius: 10rpx;
		overflow: hidden;
		margin-bottom: 15rpx;
	}
	
	.progress-fill {
		height: 100%;
		background-color: #007AFF;
		transition: width 0.3s ease;
	}
	
	.progress-text {
		text-align: center;
	}
	
	.normal {
		color: #007AFF;
		font-size: 28rpx;
	}
	
	.over-budget {
		color: #ff6b6b;
		font-size: 28rpx;
		font-weight: bold;
	}
	
	.save-btn {
		width: 100%;
		height: 80rpx;
		background-color: #007AFF;
		color: #fff;
		border-radius: 10rpx;
		font-size: 32rpx;
	}
	
	.history-section {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.section-title {
		font-size: 32rpx;
		font-weight: bold;
		margin-bottom: 20rpx;
	}
	
	.history-list {
		display: flex;
		flex-direction: column;
		gap: 20rpx;
	}
	
	.history-item {
		display: flex;
		justify-content: space-between;
		padding: 20rpx;
		background-color: #f8f8f8;
		border-radius: 10rpx;
	}
	
	.history-month {
		font-size: 28rpx;
		color: #666;
	}
	
	.history-amount {
		font-size: 28rpx;
		font-weight: bold;
		color: #333;
	}
</style>