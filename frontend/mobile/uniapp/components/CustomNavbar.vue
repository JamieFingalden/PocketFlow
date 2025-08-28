<template>
	<view class="navbar">
		<view class="nav-item" v-for="(item, index) in items" :key="index" @click="onSelect(index)">
			<text class="nav-text" :class="{ active: activeIndex === index }">{{ item }}</text>
		</view>
	</view>
</template>

<script setup>
	import { ref, defineProps, defineEmits } from 'vue'
	
	// 定义props
	const props = defineProps({
		items: {
			type: Array,
			default: () => []
		},
		defaultActive: {
			type: Number,
			default: 0
		}
	})
	
	// 定义emits
	const emits = defineEmits(['select'])
	
	// 响应式数据
	const activeIndex = ref(props.defaultActive)
	
	// 方法定义
	const onSelect = (index) => {
		activeIndex.value = index
		emits('select', index)
	}
</script>

<style>
	.navbar {
		display: flex;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.nav-item {
		flex: 1;
		text-align: center;
		padding: 20rpx;
		font-size: 32rpx;
	}
	
	.nav-text {
		color: #999;
	}
	
	.nav-text.active {
		color: #007AFF;
		font-weight: bold;
		border-bottom: 4rpx solid #007AFF;
		padding-bottom: 10rpx;
	}
</style>