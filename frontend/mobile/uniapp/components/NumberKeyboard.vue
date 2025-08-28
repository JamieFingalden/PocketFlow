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
	const emits = defineEmits(['input', 'confirm'])
	
	// 响应式数据
	const currentValue = ref(props.value.toString())
	const expression = ref('') // 存储表达式
	const lastInputType = ref('') // 记录上次输入的类型('number'或'operator')
	
	// 生命周期钩子
	onMounted(() => {
		console.log('NumberKeyboard组件已挂载, 接收到的value:', props.value)
	})
	
	// 计算属性
	const displayValue = computed(() => {
		// 如果有表达式，显示表达式，否则显示当前值
		if (expression.value) {
			return expression.value
		}
		return currentValue.value || '0.00'
	})
	
	// 方法定义
	const inputNumber = (num) => {
		// 如果上次输入的是等号，清空表达式重新开始
		if (lastInputType.value === 'equals') {
			expression.value = ''
			currentValue.value = '0.00'
			lastInputType.value = ''
		}
		
		if (currentValue.value === '0.00' || lastInputType.value === 'equals') {
			currentValue.value = num
		} else {
			currentValue.value += num
		}
		lastInputType.value = 'number'
		
		// 限制小数点后两位
		if (currentValue.value.includes('.')) {
			const parts = currentValue.value.split('.')
			if (parts[1].length > 2) {
				currentValue.value = parts[0] + '.' + parts[1].substring(0, 2)
			}
		}
	}
	
	const inputOperator = (operator) => {
		// 如果是等号，计算结果
		if (operator === '=') {
			calculateResult()
			lastInputType.value = 'equals'
			return
		}
		
		// 如果表达式为空且当前值不为0，则将当前值加入表达式
		if (!expression.value && currentValue.value !== '0.00') {
			expression.value = currentValue.value
		}
		
		// 如果上次输入的是数字，则将数字和运算符加入表达式
		if (lastInputType.value === 'number') {
			expression.value += operator
		} 
		// 如果上次输入的是运算符，则替换运算符
		else if (lastInputType.value === 'operator' && expression.value.length > 0) {
			expression.value = expression.value.slice(0, -1) + operator
		}
		
		currentValue.value = '0.00'
		lastInputType.value = 'operator'
	}
	
	const inputDecimal = () => {
		// 如果上次输入的是等号，清空表达式重新开始
		if (lastInputType.value === 'equals') {
			expression.value = ''
			currentValue.value = '0.00'
			lastInputType.value = ''
		}
		
		if (!currentValue.value.includes('.')) {
			currentValue.value += '.'
		}
		lastInputType.value = 'number'
	}
	
	const backspace = () => {
		// 如果有表达式，删除表达式的最后一个字符
		if (expression.value) {
			expression.value = expression.value.slice(0, -1)
			// 如果表达式为空了，重置当前值
			if (!expression.value) {
				currentValue.value = '0.00'
			}
		} 
		// 如果没有表达式但有当前值，删除当前值的最后一个字符
		else if (currentValue.value.length > 1) {
			currentValue.value = currentValue.value.slice(0, -1)
		} else {
			currentValue.value = '0'
		}
	}
	
	const clear = () => {
		// 清空所有内容
		currentValue.value = '0.00'
		expression.value = ''
		lastInputType.value = ''
	}
	
	const calculateResult = () => {
		// 如果有表达式，计算结果
		if (expression.value) {
			try {
				// 处理表达式中的乘除号显示
				let evalExpression = expression.value
					.replace(/×/g, '*')
					.replace(/÷/g, '/')
				
				// 如果最后一个字符是运算符，去掉它
				if (['+', '-', '*', '/'].includes(evalExpression.slice(-1))) {
					evalExpression = evalExpression.slice(0, -1)
				}
				
				// 添加当前值（如果有）
				if (currentValue.value !== '0.00') {
					evalExpression += currentValue.value
				}
				
				// 计算结果
				const result = eval(evalExpression)
				currentValue.value = parseFloat(result.toFixed(2)).toString()
				expression.value = ''
			} catch (error) {
				console.error('计算错误:', error)
				currentValue.value = '0.00'
				expression.value = ''
			}
		}
	}
	
	const confirm = () => {
		// 如果有表达式，先计算结果
		if (expression.value) {
			calculateResult()
		}
		
		// 格式化金额
		let formattedValue = parseFloat(currentValue.value || '0').toFixed(2)
		emits('input', formattedValue)
		emits('confirm', formattedValue)
		
		// 确认后清空表达式
		expression.value = ''
		lastInputType.value = 'equals'
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