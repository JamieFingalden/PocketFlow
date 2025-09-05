<template>
	<view class="container">
		<!-- 用户信息区 -->
		<view class="user-profile">
			<view class="avatar">
				<text>👤</text>
			</view>
			<view class="user-details">
				<text class="username">用户名</text>
				<text class="account-info">您已记账 {{ daysSinceStart }} 天</text>
				<text class="budget-info">本月预算：¥5,000.00  已用：¥3,150.00</text>
			</view>
		</view>

		<!-- 功能列表 -->
		<view class="functions">
			<view class="function-item" @click="goToPage('/pages/profile/user-info')">
				<text class="icon">👤</text>
				<text class="label">个人资料</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/categories')">
				<text class="icon">📂</text>
				<text class="label">分类管理</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/budget')">
				<text class="icon">💰</text>
				<text class="label">预算设置</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/templates')">
				<text class="icon">⚡</text>
				<text class="label">快捷记账模板</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/backup')">
				<text class="icon">☁️</text>
				<text class="label">数据备份与同步</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/settings')">
				<text class="icon">⚙️</text>
				<text class="label">设置</text>
				<text class="arrow">></text>
			</view>
			<view class="function-item" @click="goToPage('/pages/profile/help')">
				<text class="icon">❓</text>
				<text class="label">帮助与反馈</text>
				<text class="arrow">></text>
			</view>
		</view>

		<!-- 底部信息 -->
		<view class="footer">
			<text class="version">版本信息 v1.0.0</text>
			<text class="update">检查更新</text>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	
	// 响应式数据
	const daysSinceStart = ref(45) // 假设用户已记账45天
	
	// 方法定义
	const goToPage = (url) => {
		uni.navigateTo({
			url: url
		})
	}
</script>

<style>
	.container {
		min-height: 100vh;
		padding: 20rpx;
		background-color: #F0F3F5;
		box-sizing: border-box;
	}

	.user-profile {
		display: flex;
		align-items: center;
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.avatar {
		width: 100rpx;
		height: 100rpx;
		border-radius: 50rpx;
		background-color: #e0f7fa;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 50rpx;
		margin-right: 30rpx;
	}

	.user-details {
		display: flex;
		flex-direction: column;
	}

	.username {
		font-size: 36rpx;
		font-weight: bold;
		margin-bottom: 10rpx;
	}

	.account-info {
		font-size: 28rpx;
		color: #999;
		margin-bottom: 10rpx;
	}

	.budget-info {
		font-size: 28rpx;
		color: #999;
	}

	.functions {
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.function-item {
		display: flex;
		align-items: center;
		padding: 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}

	.function-item:last-child {
		border-bottom: none;
	}

	.icon {
		font-size: 40rpx;
		margin-right: 20rpx;
		width: 50rpx;
	}

	.label {
		flex: 1;
		font-size: 32rpx;
	}

	.arrow {
		color: #999;
		font-size: 36rpx;
	}

	.footer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}

	.version {
		font-size: 28rpx;
		color: #999;
	}

	.update {
		font-size: 28rpx;
		color: #007AFF;
	}
</style>