<template>
	<view class="number-keyboard">
		<view class="keyboard-display">
			<text class="display-text">{{ displayValue }}</text>
		</view>
				
		<view class="keyboard-grid">
			<view class="keyboard-row">
				<view class="key" @click="inputNumber('7')">7</view>
				<view class="key" @click="inputNumber('8')">8</view>
				<view class="key" @click="inputNumber('9')">9</view>
				<view class="key operator" @click="inputOperator('/')">÷</view>
			</view>
			
			<view class="keyboard-row">
				<view class="key" @click="inputNumber('4')">4</view>
				<view class="key" @click="inputNumber('5')">5</view>
				<view class="key" @click="inputNumber('6')">6</view>
				<view class="key operator" @click="inputOperator('*')">×</view>
			</view>
			
			<view class="keyboard-row">
				<view class="key" @click="inputNumber('1')">1</view>
				<view class="key" @click="inputNumber('2')">2</view>
				<view class="key" @click="inputNumber('3')">3</view>
				<view class="key operator" @click="inputOperator('-')">-</view>
			</view>
			
			<view class="keyboard-row">
				<view class="key" @click="inputNumber('0')">0</view>
				<view class="key special" @click="inputDecimal">.</view>
				<view class="key operator" @click="inputOperator('+')">+</view>
				<view class="key operator" @click="inputOperator('=')">=</view>
			</view>
			
			<view class="keyboard-row">
				<view class="key special" @click="clear">C</view>
				<view class="key special" @click="backspace">⌫</view>
				<view class="key confirm" @click="confirm">确定</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed, onMounted, defineProps, defineEmits } from 'vue'
	
	// 定义props
	const props = defineProps({
		value: {
			type: [String, Number],
			default: '0.00'
		}
	})
	
	// 定义emits
	const emits = defineEmits(['input', 'confirm', 'close'])
	
	// 响应式数据
	const currentValue = ref('0') // 当前显示的值
	const previousValue = ref(null) // 上一个值
	const operation = ref(null) // 当前操作符
	const waitingForOperand = ref(false) // 是否在等待操作数
	
	// 生命周期钩子
	onMounted(() => {
		console.log('NumberKeyboard组件已挂载, 接收到的value:', props.value)
		if (props.value && props.value !== '0.00') {
			currentValue.value = props.value.toString()
		}
	})
	
	// 计算属性
	const displayValue = computed(() => {
		return currentValue.value
	})
	
	// 方法定义
	const inputNumber = (num) => {
		console.log('输入数字:', num, 'waitingForOperand:', waitingForOperand.value)
		if (waitingForOperand.value) {
			currentValue.value = num
			waitingForOperand.value = false
		} else {
			// 如果当前值是"0"或者以运算符结尾，则替换为新数字，否则追加
			if (currentValue.value === '0' || 
				currentValue.value.endsWith(' + ') || 
				currentValue.value.endsWith(' - ') || 
				currentValue.value.endsWith(' × ') || 
				currentValue.value.endsWith(' ÷ ')) {
				currentValue.value = num
			} else {
				currentValue.value = currentValue.value === '0' ? num : currentValue.value + num
			}
		}
		// 发出input事件，让父组件实时显示当前值
		emits('input', currentValue.value)
	}
	
	const inputDecimal = () => {
		if (waitingForOperand.value) {
			currentValue.value = '0.'
			waitingForOperand.value = false
			return
		}
		
		if (!currentValue.value.includes('.')) {
			currentValue.value += '.'
		}
	}
	
	const clear = () => {
		currentValue.value = '0'
		previousValue.value = null
		operation.value = null
		waitingForOperand.value = false
	}
	
	const backspace = () => {
		if (waitingForOperand.value) {
			return
		}
		
		if (currentValue.value.length === 1 || 
			(currentValue.value.length === 2 && currentValue.value.startsWith('-'))) {
			currentValue.value = '0'
		} else {
			currentValue.value = currentValue.value.slice(0, -1)
		}
	}
	
	const performOperation = () => {
		const current = parseFloat(currentValue.value)
		
		if (previousValue.value !== null && operation.value) {
			const previous = previousValue.value
			let result
			
			switch (operation.value) {
				case '+':
					result = previous + current
					break
				case '-':
					result = previous - current
					break
				case '*':
					result = previous * current
					break
				case '/':
					result = previous / current
					break
				default:
					return
			}
			
			// 格式化结果，保留两位小数
			currentValue.value = parseFloat(result.toFixed(2)).toString()
			previousValue.value = result
		}
	}
	
	const getOperatorSymbol = (operator) => {
		switch (operator) {
			case '+': return '+'
			case '-': return '-'
			case '*': return '×'
			case '/': return '÷'
			case '=': return '='
			default: return operator
		}
	}
	
	const inputOperator = (operator) => {
		console.log('输入运算符:', operator, '当前值:', currentValue.value, '前一个值:', previousValue.value, '当前运算符:', operation.value)
		if (operator === '=') {
			if (operation.value && previousValue.value !== null) {
				performOperation()
				// 计算完成后，将结果直接应用到金额输入框
				let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
				emits('input', formattedValue)
				emits('confirm', formattedValue)
				operation.value = null
				previousValue.value = null
				waitingForOperand.value = true
			} else {
				// 如果没有操作符，直接将当前值应用到金额输入框
				let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
				emits('input', formattedValue)
				emits('confirm', formattedValue)
			}
			return
		}
		
		// 如果之前有操作符和前一个值，说明现在要计算前两个数的结果
		if (operation.value && previousValue.value !== null) {
			performOperation()
			// 自动计算完成后，将结果直接应用到金额输入框，但不关闭键盘
			let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
			emits('input', formattedValue)
			// 发送一个特殊的事件来更新金额但不关闭键盘
			emits('autoCalculate', formattedValue)
		}
		
		// 设置新的操作符
		if (previousValue.value === null) {
			previousValue.value = parseFloat(currentValue.value)
		}
		operation.value = operator
		waitingForOperand.value = true
		// 显示运算符表达式
		const displayText = previousValue.value !== null ? `${previousValue.value} ${getOperatorSymbol(operator)} ` : `${currentValue.value} ${getOperatorSymbol(operator)} `
		emits('input', displayText)
	}
	
	const confirm = () => {
		// 如果有待处理的操作和前一个值，先计算结果
		if (operation.value && previousValue.value !== null) {
			performOperation()
		}
		
		// 格式化金额
		let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
		emits('input', formattedValue)
		emits('confirm', formattedValue)
		
		// 重置状态
		currentValue.value = '0'
		previousValue.value = null
		operation.value = null
		waitingForOperand.value = false
	}
	
	const closeKeyboard = () => {
		console.log('数字键盘 closeKeyboard 函数被调用')
		// 如果有待处理的操作和前一个值，先计算结果
		if (operation.value && previousValue.value !== null) {
			console.log('有待处理的计算操作，先执行计算')
			performOperation()
		}
		
		// 格式化金额
		let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
		console.log('格式化金额:', formattedValue)
		emits('input', formattedValue)
		emits('confirm', formattedValue)
		emits('close') // 发送关闭事件
		console.log('已发送 close 事件')
		
		// 重置状态
		currentValue.value = '0'
		previousValue.value = null
		operation.value = null
		waitingForOperand.value = false
	}
</script>

<style>
	.number-keyboard {
		background-color: #f8f8f8;
		border-top-left-radius: 20rpx;
		border-top-right-radius: 20rpx;
		overflow: hidden;
	}
	
	.keyboard-display {
		padding: 30rpx;
		text-align: center;
		background-color: #fff;
		border-bottom: 1rpx solid #f0f0f0;
	}
	
	.display-text {
		font-size: 48rpx;
		font-weight: bold;
		color: #333;
	}
	
	.keyboard-grid {
		padding: 20rpx;
	}
	
	.keyboard-row {
		display: flex;
		justify-content: space-between;
		margin-bottom: 20rpx;
	}
	
	.keyboard-row:last-child {
		margin-bottom: 0;
	}
	
	.key {
		flex: 1;
		height: 80rpx;
		background-color: #fff;
		border-radius: 10rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 36rpx;
		font-weight: bold;
		margin: 0 10rpx;
		box-shadow: 0 4rpx 10rpx rgba(0, 0, 0, 0.1);
	}
	
	.key:first-child {
		margin-left: 0;
	}
	
	.key:last-child {
		margin-right: 0;
	}
	
	.key.special {
		background-color: #f0f0f0;
		color: #666;
	}
	
	.key.operator {
		background-color: #f0f0f0;
		color: #666;
	}
	
	.key.equals {
		background-color: #4a90e2;
		color: #fff;
	}

	.key.confirm {
		background-color: #007AFF;
		color: #fff;
	}
</style>