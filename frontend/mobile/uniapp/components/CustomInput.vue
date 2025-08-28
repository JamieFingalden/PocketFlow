<template>
	<view class="input-container">
		<input 
			class="custom-input" 
			:type="type" 
			:placeholder="placeholder" 
			:value="value"
			@input="onInput"
		/>
	</view>
</template>

<script setup>
	import { defineProps, defineEmits } from 'vue'
	
	// 定义props
	const props = defineProps({
		type: {
			type: String,
			default: 'text'
		},
		placeholder: {
			type: String,
			default: ''
		},
		value: {
			type: [String, Number],
			default: ''
		}
	})
	
	// 定义emits
	const emits = defineEmits(['update:value'])
	
	// 方法定义
	const onInput = (e) => {
		emits('update:value', e.detail.value)
	}
</script>

<style>
	.input-container {
		margin-bottom: 30rpx;
	}
	
	.custom-input {
		width: 100%;
		height: 80rpx;
		border: 1rpx solid #e0e0e0;
		border-radius: 10rpx;
		padding: 0 20rpx;
		font-size: 32rpx;
		box-sizing: border-box;
	}
	
	.custom-input:focus {
		border-color: #007AFF;
	}
</style>