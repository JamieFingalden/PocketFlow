<template>
	<button 
		class="custom-button" 
		:class="[type, sizeClass]"
		@click="onClick"
	>
		<slot></slot>
	</button>
</template>

<script setup>
	import { computed, defineProps, defineEmits } from 'vue'
	
	// 定义props
	const props = defineProps({
		type: {
			type: String,
			default: 'primary' // primary, secondary, danger
		},
		size: {
			type: String,
			default: 'normal' // small, normal, large
		}
	})
	
	// 定义emits
	const emits = defineEmits(['click'])
	
	// 计算属性
	const sizeClass = computed(() => {
		return `btn-${props.size}`
	})
	
	// 方法定义
	const onClick = () => {
		emits('click')
	}
</script>

<style>
	.custom-button {
		border-radius: 10rpx;
		font-size: 32rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		border: none;
		outline: none;
	}
	
	/* 按钮类型样式 */
	.primary {
		background-color: #007AFF;
		color: #fff;
	}
	
	.secondary {
		background-color: #f0f0f0;
		color: #666;
	}
	
	.danger {
		background-color: #ff6b6b;
		color: #fff;
	}
	
	/* 按钮尺寸样式 */
	.btn-small {
		height: 60rpx;
		padding: 0 20rpx;
		font-size: 28rpx;
	}
	
	.btn-normal {
		height: 80rpx;
		padding: 0 30rpx;
	}
	
	.btn-large {
		height: 100rpx;
		padding: 0 40rpx;
		font-size: 36rpx;
	}
</style>