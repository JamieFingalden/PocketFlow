<template>
	<view class="container">
		<!-- 顶部时间选择和筛选 -->
		<view class="header">
			<view class="date-display">
				<text class="date-text">2023年8月</text>
			</view>
			<view class="filter-actions">
				<text class="filter-btn">筛选</text>
			</view>
		</view>

		<!-- 收支概览卡片 -->
		<view class="summary-card">
			<view class="summary-header">
				<text class="card-title">本月收支</text>
			</view>
			<view class="summary-content">
				<view class="summary-item">
					<text class="label">收入</text>
					<text class="amount income">¥5,230.00</text>
				</view>
				<view class="summary-item">
					<text class="label">支出</text>
					<text class="amount expense">¥3,150.00</text>
				</view>
				<view class="summary-item">
					<text class="label">结余</text>
					<text class="amount balance">¥2,080.00</text>
				</view>
			</view>
		</view>

		<!-- 图表区域 -->
		<view class="charts-section">
			<!-- 收支趋势图 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">收支趋势</text>
				</view>
				<BarChart 
					:chart-data="trendData" 
					:x-labels="trendLabels"
					:width="300"
					:height="300"
				/>
			</view>

			<!-- 分类支出占比 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">分类支出占比</text>
				</view>
				<PieChart 
					:chart-data="categoryData" 
					:width="300"
					:height="300"
				/>
			</view>
		</view>

		<!-- 详细分类列表 -->
		<view class="details-section">
			<view class="details-header">
				<text class="section-title">支出详情</text>
			</view>
			<view class="category-list">
				<view class="category-item" v-for="(item, index) in categoryData" :key="index">
					<view class="category-info">
						<view class="category-color" :style="{ backgroundColor: item.color }"></view>
						<text class="category-name">{{ item.name }}</text>
					</view>
					<view class="category-amount">
						<text class="amount expense">¥{{ item.amount }}</text>
						<text class="percentage">{{ item.percentage }}%</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import PieChart from '@/components/PieChart.vue'
	import BarChart from '@/components/BarChart.vue'
	
	// 响应式数据
	const timeRange = ref('month')
	const trendData = ref([
		{ income: 5000, expense: 3000 },
		{ income: 5200, expense: 3100 },
		{ income: 4800, expense: 3200 },
		{ income: 5500, expense: 2900 },
		{ income: 5100, expense: 3300 },
		{ income: 5300, expense: 3100 }
	])
	const trendLabels = ref(['1月', '2月', '3月', '4月', '5月', '6月'])
	const categoryData = ref([
		{ name: '餐饮', value: 1200, percent: 35, color: '#FF6B6B', amount: '1200.00', percentage: '35' },
		{ name: '交通', value: 500, percent: 15, color: '#4ECDC4', amount: '500.00', percentage: '15' },
		{ name: '购物', value: 850, percent: 25, color: '#45B7D1', amount: '850.00', percentage: '25' },
		{ name: '医疗', value: 350, percent: 10, color: '#96CEB4', amount: '350.00', percentage: '10' },
		{ name: '其他', value: 500, percent: 15, color: '#FFEAA7', amount: '500.00', percentage: '15' }
	])
	
	// 方法定义
	const selectTimeRange = (range) => {
		timeRange.value = range
	}
</script>

<style>
	.container {
		padding: 20rpx;
	}

	.time-selector {
		display: flex;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.time-option {
		flex: 1;
		text-align: center;
		padding: 20rpx;
		font-size: 32rpx;
	}

	.time-option.active {
		background-color: #007AFF;
		color: #fff;
		border-radius: 20rpx;
	}

	.charts-section {
		display: flex;
		flex-direction: column;
		gap: 30rpx;
		margin-bottom: 30rpx;
	}
	
	.chart-card {
		background-color: #fff;
		border-radius: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.chart-header {
		padding: 20rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}

	.title {
		font-size: 32rpx;
		font-weight: bold;
	}

	.chart-container {
		padding: 30rpx;
	}

	.summary-item {
		display: flex;
		justify-content: space-between;
		margin-bottom: 20rpx;
	}

	.summary-item:last-child {
		margin-bottom: 0;
	}

	.label {
		font-size: 28rpx;
		color: #999;
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

	.category-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 20rpx;
	}

	.category-item:last-child {
		margin-bottom: 0;
	}

	.category-info {
		display: flex;
		align-items: center;
	}

	.category-color {
		width: 30rpx;
		height: 30rpx;
		border-radius: 15rpx;
		margin-right: 20rpx;
	}

	.category-name {
		font-size: 28rpx;
	}

	.category-percent {
		font-size: 28rpx;
		font-weight: bold;
		color: #666;
	}
</style>