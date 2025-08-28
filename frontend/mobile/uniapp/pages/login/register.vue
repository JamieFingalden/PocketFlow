<template>
	<view class="container">
		<view class="header">
			<text class="title">注册</text>
			<text class="subtitle">创建您的PocketFlow账户</text>
		</view>
		
		<view class="form">
			<view class="input-group">
				<input 
					class="input" 
					type="text" 
					placeholder="请输入用户名" 
					v-model="registerForm.username"
				/>
			</view>
			
			<view class="input-group">
				<input 
					class="input" 
					type="text" 
					placeholder="请输入邮箱" 
					v-model="registerForm.email"
				/>
			</view>
			
			<view class="input-group">
				<input 
					class="input" 
					type="password" 
					placeholder="请输入密码" 
					v-model="registerForm.password"
				/>
			</view>
			
			<view class="input-group">
				<input 
					class="input" 
					type="password" 
					placeholder="请确认密码" 
					v-model="registerForm.confirmPassword"
				/>
			</view>
			
			<button class="register-btn" @click="handleRegister">注册</button>
			
			<view class="links">
				<text class="link" @click="goToLogin">已有账号？立即登录</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	
	// 响应式数据
	const registerForm = ref({
		username: '',
		email: '',
		password: '',
		confirmPassword: ''
	})
	
	// 方法定义
	const handleRegister = () => {
		// 简单验证
		if (!registerForm.value.username || !registerForm.value.email || 
			!registerForm.value.password || !registerForm.value.confirmPassword) {
			uni.showToast({
				title: '请填写所有字段',
				icon: 'none'
			})
			return
		}
		
		if (registerForm.value.password !== registerForm.value.confirmPassword) {
			uni.showToast({
				title: '密码不一致',
				icon: 'none'
			})
			return
		}
		
		// 模拟注册成功
		uni.showToast({
			title: '注册成功',
			icon: 'success'
		})
		
		// 延迟返回登录页面
		setTimeout(() => {
			uni.navigateBack()
		}, 1000)
	}
	
	const goToLogin = () => {
		uni.navigateBack()
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
	
	.register-btn {
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