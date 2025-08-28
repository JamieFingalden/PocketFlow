<template>
	<view class="card">
		<view class="card-header" v-if="title">
			<text class="card-title">{{ title }}</text>
		</view>
		<view class="card-body">
			<slot></slot>
		</view>
	</view>
</template>

<script setup>
	import { defineProps } from 'vue'
	
	// 定义props
	defineProps({
		title: {
			type: String,
			default: ''
		}
	})
</script>

<style>
	.card {
		background-color: #fff;
		border-radius: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		margin-bottom: 20rpx;
	}
	
	.card-header {
		padding: 20rpx 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}
	
	.card-title {
		font-size: 32rpx;
		font-weight: bold;
		color: #333;
	}
	
	.card-body {
		padding: 30rpx;
	}
</style>