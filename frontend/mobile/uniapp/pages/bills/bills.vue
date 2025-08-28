<template>
	<view class="container">
		<!-- 顶部筛选区域 -->
		<view class="filter-bar">
			<view class="time-filter">
				<picker mode="date" :value="selectedDate" fields="month" @change="onDateChange">
					<view class="picker">
						{{ selectedDate }}
					</view>
				</picker>
			</view>
			<view class="action-icons">
				<text class="icon" @click="toggleSearch">🔍</text>
				<text class="icon">⚙️</text>
			</view>
		</view>
		
		<!-- 搜索框 -->
		<view class="search-bar" v-if="showSearch">
			<input 
				class="search-input" 
				type="text" 
				placeholder="搜索账单备注或分类" 
				v-model="searchKeyword"
				@input="onSearchInput"
			/>
			<text class="clear-btn" @click="clearSearch">✕</text>
		</view>

		<!-- 收支概览 -->
		<view class="summary">
			<view class="summary-item">
				<text class="label">收入</text>
				<text class="amount income">¥{{ income }}</text>
			</view>
			<view class="summary-item">
				<text class="label">支出</text>
				<text class="amount expense">¥{{ expense }}</text>
			</view>
			<view class="summary-item">
				<text class="label">结余</text>
				<text class="amount balance">¥{{ balance }}</text>
			</view>
		</view>

		<!-- 账单列表 -->
		<view class="bills-list">
			<view class="bill-date-group" v-for="(group, index) in filteredBillGroups" :key="index">
				<view class="date-header">{{ group.date }}</view>
				<view class="bill-item" v-for="(bill, billIndex) in group.bills" :key="billIndex">
					<view class="bill-icon">
						<text>{{ bill.categoryIcon }}</text>
					</view>
					<view class="bill-info">
						<text class="bill-category">{{ bill.category }}</text>
						<text class="bill-remark">{{ bill.remark }}</text>
					</view>
					<view class="bill-amount">
						<text class="amount" :class="bill.type">{{ bill.amount }}</text>
					</view>
				</view>
			</view>
		</view>

		<!-- 悬浮添加按钮 -->
		<view class="floating-btn" @click="goToAddBill">
			<text class="add-icon">+</text>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed } from 'vue'
	
	// 响应式数据
	const selectedDate = ref('2023-08')
	const showSearch = ref(false)
	const searchKeyword = ref('')
	const income = ref('5,230.00')
	const expense = ref('3,150.00')
	const balance = ref('2,080.00')
	const billGroups = ref([
		{
			date: '今天',
			bills: [
				{ categoryIcon: '餐', category: '餐饮', remark: '星巴克咖啡', amount: '-35.00', type: 'expense' },
				{ categoryIcon: '交', category: '交通', remark: '地铁', amount: '-5.00', type: 'expense' }
			]
		},
		{
			date: '昨天',
			bills: [
				{ categoryIcon: '购', category: '购物', remark: '京东购物', amount: '-128.50', type: 'expense' },
				{ categoryIcon: '医', category: '医疗', remark: '药店', amount: '-45.00', type: 'expense' }
			]
		},
		{
			date: '8月24日',
			bills: [
				{ categoryIcon: '餐', category: '餐饮', remark: '麦当劳', amount: '-28.00', type: 'expense' }
			]
		}
	])
	
	// 计算属性
	const filteredBillGroups = computed(() => {
		if (!searchKeyword.value) {
			return billGroups.value
		}
		
		const keyword = searchKeyword.value.toLowerCase()
		return billGroups.value.map(group => {
			const filteredBills = group.bills.filter(bill => 
				bill.remark.toLowerCase().includes(keyword) || 
				bill.category.toLowerCase().includes(keyword)
			)
			return {
				...group,
				bills: filteredBills
			}
		}).filter(group => group.bills.length > 0)
	})
	
	// 方法定义
	const onDateChange = (e) => {
		selectedDate.value = e.target.value
	}
	
	const toggleSearch = () => {
		showSearch.value = !showSearch.value
		if (!showSearch.value) {
			searchKeyword.value = ''
		}
	}
	
	const onSearchInput = (e) => {
		searchKeyword.value = e.detail.value
	}
	
	const clearSearch = () => {
		searchKeyword.value = ''
	}
	
	const goToAddBill = () => {
		uni.navigateTo({
			url: '/pages/bills/add-bill'
		})
	}
</script>

<style>
	.container {
		padding: 20rpx;
		padding-bottom: 120rpx;
	}

	.filter-bar {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background-color: #fff;
		border-radius: 20rpx;
		padding: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.time-filter {
		font-size: 32rpx;
		font-weight: bold;
	}

	.picker {
		padding: 10rpx 20rpx;
	}

	.action-icons {
		display: flex;
		gap: 30rpx;
	}

	.icon {
		font-size: 36rpx;
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
		flex: 1;
	}

	.summary-item:not(:last-child) {
		border-right: 1rpx solid #f0f0f0;
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

	.bills-list {
		display: flex;
		flex-direction: column;
		gap: 30rpx;
	}

	.bill-date-group {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.date-header {
		font-size: 28rpx;
		font-weight: bold;
		color: #999;
		margin-bottom: 20rpx;
		padding-bottom: 10rpx;
		border-bottom: 1rpx solid #f0f0f0;
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

	.amount {
		font-size: 32rpx;
		font-weight: bold;
	}

	.expense {
		color: #ff6b6b;
	}

	.income {
		color: #00b38a;
	}

	.floating-btn {
		position: fixed;
		bottom: 120rpx; /* 高于tabBar */
		right: 50rpx;
		width: 90rpx;
		height: 90rpx;
		border-radius: 50%;
		background-color: #007AFF;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 8rpx 20rpx rgba(0, 122, 255, 0.3);
		z-index: 999;
	}

	.add-icon {
		color: #fff;
		font-size: 60rpx;
		font-weight: bold;
	}
</style>