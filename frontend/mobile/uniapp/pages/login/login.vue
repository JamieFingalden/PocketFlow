<template>
	<view class="container">
		<view class="header">
			<text class="title">登录</text>
			<text class="subtitle">欢迎使用PocketFlow</text>
		</view>
		
		<view class="form">
			<view class="input-group">
				<input 
					class="input" 
					type="text" 
					placeholder="请输入用户名" 
					v-model="loginForm.username"
				/>
			</view>
			
			<view class="input-group">
				<input 
					class="input" 
					type="password" 
					placeholder="请输入密码" 
					v-model="loginForm.password"
				/>
			</view>
			
			<button class="login-btn" @click="handleLogin">登录</button>
			
			<view class="links">
				<text class="link" @click="goToRegister">还没有账号？立即注册</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	
	// 响应式数据
	const loginForm = ref({
		username: '',
		password: ''
	})
	
	// 方法定义
	const handleLogin = () => {
		// 简单验证
		if (!loginForm.value.username || !loginForm.value.password) {
			uni.showToast({
				title: '请输入用户名和密码',
				icon: 'none'
			})
			return
		}
		
		// 模拟登录成功
		uni.showToast({
			title: '登录成功',
			icon: 'success'
		})
		
		// 延迟跳转到首页
		setTimeout(() => {
			uni.switchTab({
				url: '/pages/index/index'
			})
		}, 1000)
	}
	
	const goToRegister = () => {
		uni.navigateTo({
			url: '/pages/login/register'
		})
	}
</script>

<style>
	.container {
		padding: 40rpx;
		background-color: #f8f8f8;
		min-height: 100vh;
	}
	
	.header {
		text-align: center;
		margin-bottom: 80rpx;
		margin-top: 100rpx;
	}
	
	.title {
		font-size: 48rpx;
		font-weight: bold;
		color: #333;
		display: block;
		margin-bottom: 20rpx;
	}
	
	.subtitle {
		font-size: 28rpx;
		color: #999;
	}
	
	.form {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 40rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.input-group {
		margin-bottom: 30rpx;
	}
	
	.input {
		width: 100%;
		height: 80rpx;
		border: 1rpx solid #e0e0e0;
		border-radius: 10rpx;
		padding: 0 20rpx;
		font-size: 32rpx;
		box-sizing: border-box;
	}
	
	.input:focus {
		border-color: #007AFF;
	}
	
	.login-btn {
		width: 100%;
		height: 80rpx;
		background-color: #007AFF;
		color: #fff;
		border-radius: 10rpx;
		font-size: 32rpx;
		margin-top: 20rpx;
	}
	
	.links {
		text-align: center;
		margin-top: 30rpx;
	}
	
	.link {
		color: #007AFF;
		font-size: 28rpx;
	}
</style>