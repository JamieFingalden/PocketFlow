<template>
	<view class="container">
		<!-- 顶部时间选择和筛选 -->
		<view class="header">
			<view class="date-display">
				<text class="date-text">{{ currentMonth }}</text>
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
				<view class="summary-item income-item">
					<view class="label-wrapper">
						<text class="label">收入</text>
					</view>
					<text class="amount income">¥5,230.00</text>
				</view>
				<view class="summary-item expense-item">
					<view class="label-wrapper">
						<text class="label">支出</text>
					</view>
					<text class="amount expense">¥3,150.00</text>
				</view>
				<view class="summary-item balance-item">
					<view class="label-wrapper">
						<text class="label">结余</text>
					</view>
					<text class="amount balance">¥2,080.00</text>
				</view>
			</view>
		</view>

		<!-- 图表区域 -->
		<view class="charts-section">
			<!-- 选项卡 -->
			<view class="tabs">
				<view 
					class="tab" 
					:class="{ active: activeTab === 'expense' }"
					@click="switchTab('expense')"
				>
					支出
				</view>
				<view 
					class="tab" 
					:class="{ active: activeTab === 'income' }"
					@click="switchTab('income')"
				>
					收入
				</view>
			</view>
			
			<!-- 收入趋势图 -->
			<view class="chart-card" v-if="activeTab === 'income'">
				<view class="chart-header">
					<text class="chart-title">收入趋势</text>
				</view>
				<view class="chart-container">
					<LineChart 
						:chart-data="incomeTrendData" 
						:x-labels="trendLabels"
						:width="300"
						:height="300"
						chartType="income"
					/>
				</view>
			</view>
			
			<!-- 支出趋势图 -->
			<view class="chart-card" v-if="activeTab === 'expense'">
				<view class="chart-header">
					<text class="chart-title">支出趋势</text>
				</view>
				<view class="chart-container">
					<LineChart 
						:chart-data="expenseTrendData" 
						:x-labels="trendLabels"
						:width="300"
						:height="300"
						chartType="expense"
					/>
				</view>
			</view>

			<!-- 分类支出占比 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">分类支出占比</text>
				</view>
				<view class="chart-container">
					<PieChart 
						:chart-data="categoryData" 
						:width="300"
						:height="300"
					/>
				</view>
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
	import { ref, computed, onMounted } from 'vue'
	import PieChart from '@/components/PieChart.vue'
	import LineChart from '@/components/LineChart.vue'
	
	// 响应式数据
	const timeRange = ref('month')
	const activeTab = ref('expense') // 默认选中支出

	const trendLabels = ref(['1月', '2月', '3月', '4月', '5月', '6月'])
	const categoryData = ref([
		{ name: '餐饮', value: 1200, percent: 35, color: '#4a90e2', amount: '1200.00', percentage: '35' },
		{ name: '交通', value: 500, percent: 15, color: '#5e72e4', amount: '500.00', percentage: '15' },
		{ name: '购物', value: 850, percent: 25, color: '#667eea', amount: '850.00', percentage: '25' },
		{ name: '医疗', value: 350, percent: 10, color: '#764ba2', amount: '350.00', percentage: '10' },
		{ name: '其他', value: 500, percent: 15, color: '#96ceb4', amount: '500.00', percentage: '15' }
	])
	
	// 原始数据
	const rawData = ref([
		{ income: 5000, expense: 3000 },
		{ income: 5200, expense: 3100 },
		{ income: 4800, expense: 3200 },
		{ income: 5500, expense: 2900 },
		{ income: 5100, expense: 3300 },
		{ income: 5300, expense: 3100 }
	])
	
	// 计算属性 - 收入趋势数据
	const incomeTrendData = computed(() => {
		return rawData.value.map(item => ({ value: item.income }))
	})
	
	// 计算属性 - 支出趋势数据
	const expenseTrendData = computed(() => {
		return rawData.value.map(item => ({ value: item.expense }))
	})
	
	// 计算属性 - 当前月份
	const currentMonth = computed(() => {
		const date = new Date()
		return `${date.getFullYear()}年${date.getMonth() + 1}月`
	})
	
	// 方法定义
	const selectTimeRange = (range) => {
		timeRange.value = range
	}
	
	// 切换选项卡
	const switchTab = (tab) => {
		activeTab.value = tab
	}
	
	// 初始化
	onMounted(() => {
		// 可以在这里添加初始化逻辑
	})
</script>

<style>
	.container {
		padding: 20rpx;
		background-color: #F0F3F5;
		min-height: 100vh;
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

	.date-display .date-text {
		font-size: 36rpx;
		font-weight: bold;
	}

	.filter-btn {
		background-color: rgba(255, 255, 255, 0.2);
		padding: 10rpx 20rpx;
		border-radius: 10rpx;
		font-size: 28rpx;
	}

	.summary-card {
		background-color: #FFFFFF;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 30rpx;
		color: #2E2E2E;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		border: 1rpx solid #f0f0f0;
	}

	.summary-header .card-title {
		font-size: 32rpx;
		font-weight: bold;
		margin-bottom: 20rpx;
	}

	.summary-content {
		display: flex;
		justify-content: space-between;
		text-align: center;
	}

	.summary-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		flex: 1;
		padding: 20rpx;
		border-radius: 15rpx;
		transition: all 0.3s ease;
	}

	.summary-item:hover {
		background-color: rgba(255, 255, 255, 0.1);
		transform: translateY(-3rpx);
	}

	.label-wrapper {
		margin-bottom: 10rpx;
	}

	.label {
		font-size: 24rpx;
		color: #6D7378;
	}

	.amount {
		font-size: 36rpx;
		font-weight: bold;
	}

	.income-item .amount {
		color: #8BC34A;
	}

	.expense-item .amount {
		color: #FF7F50;
	}

	.balance-item .amount {
		color: #5AA9E6;
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
		overflow: hidden;
		transition: all 0.3s ease;
	}

	.chart-card:hover {
		transform: translateY(-5rpx);
		box-shadow: 0 6rpx 25rpx rgba(0, 0, 0, 0.1);
	}

	.chart-header {
		padding: 20rpx 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
		background-color: #fafafa;
	}

	.chart-title {
		font-size: 30rpx;
		font-weight: bold;
		color: #333;
	}

	.chart-container {
		padding: 30rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.details-section {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.details-header .section-title {
		font-size: 30rpx;
		font-weight: bold;
		color: #333;
		margin-bottom: 20rpx;
		position: relative;
		padding-bottom: 15rpx;
	}

	.details-header .section-title::after {
		content: '';
		position: absolute;
		bottom: 0;
		left: 0;
		width: 60rpx;
		height: 4rpx;
		background: linear-gradient(90deg, #5AA9E6, #8BC34A);
		border-radius: 2rpx;
	}

	.category-list {
		padding-top: 10rpx;
	}

	.category-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 20rpx;
		padding: 20rpx;
		border-radius: 15rpx;
		background-color: #fafafa;
		transition: all 0.3s ease;
	}

	.category-item:hover {
		background-color: #f0f0f0;
		transform: translateY(-2rpx);
		box-shadow: 0 2rpx 8rpx rgba(0, 0, 0, 0.05);
	}

	.category-item:last-child {
		margin-bottom: 0;
	}

	.category-info {
		display: flex;
		align-items: center;
		flex: 1;
	}

	.category-color {
		width: 36rpx;
		height: 36rpx;
		border-radius: 18rpx;
		margin-right: 20rpx;
		box-shadow: 0 2rpx 4rpx rgba(0, 0, 0, 0.1);
	}

	.category-name {
		font-size: 28rpx;
		color: #2E2E2E;
		font-weight: 500;
	}

	.category-amount {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
	}

	.amount {
		font-size: 28rpx;
		font-weight: bold;
		margin-bottom: 5rpx;
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

	.percentage {
		font-size: 24rpx;
		color: #999;
		background-color: #f0f0f0;
		padding: 4rpx 12rpx;
		border-radius: 10rpx;
	}

	/* 选项卡样式 */
	.tabs {
		display: flex;
		background-color: #f0f0f0;
		border-radius: 20rpx;
		margin: 20rpx 30rpx;
		overflow: hidden;
	}

	.tab {
		flex: 1;
		text-align: center;
		padding: 20rpx;
		font-size: 28rpx;
		color: #666;
		background-color: #f0f0f0;
		transition: all 0.3s;
	}

	.tab.active {
		background: linear-gradient(135deg, #5AA9E6, #8BC34A);
		color: #fff;
		font-weight: bold;
	}
</style>