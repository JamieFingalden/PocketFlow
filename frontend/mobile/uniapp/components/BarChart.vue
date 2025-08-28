<template>
	<view class="bar-chart">
		<canvas 
			class="chart-canvas" 
			:canvas-id="canvasId" 
			:id="canvasId"
		></canvas>
		<view class="x-axis">
			<text 
				class="x-label" 
				v-for="(label, index) in xLabels" 
				:key="index"
			>
				{{ label }}
			</text>
		</view>
	</view>
</template>

<script setup>
	import { ref, onMounted, watch, defineProps } from 'vue'
	
	// 定义props
	const props = defineProps({
		chartData: {
			type: Array,
			default: () => []
		},
		xLabels: {
			type: Array,
			default: () => []
		},
		width: {
			type: [String, Number],
			default: 300
		},
		height: {
			type: [String, Number],
			default: 200
		}
	})
	
	// 响应式数据
	const canvasId = ref('barChartCanvas')
	
	// 方法定义
	const drawChart = () => {
		const ctx = uni.createCanvasContext(canvasId.value, this)
		
		const width = props.width
		const height = props.height
		const padding = 20
		const barWidth = (width - padding * 2) / (props.chartData.length * 2 + props.chartData.length - 1) // 考虑收入和支出两个柱子以及它们之间的间距
		const barSpacing = barWidth * 0.5 // 柱子之间的间距
		const maxValue = Math.max(...props.chartData.map(item => Math.max(item.income, item.expense)))
		
		// 绘制网格线和背景
		ctx.setFillStyle('#ffffff')
		ctx.fillRect(0, 0, width, height)
		
		// 绘制水平网格线
		ctx.setStrokeStyle('#f0f0f0')
		ctx.setLineWidth(1)
		for (let i = 0; i <= 5; i++) {
			const y = padding + (height - padding * 2) * (1 - i / 5)
			ctx.beginPath()
			ctx.moveTo(padding, y)
			ctx.lineTo(width - padding, y)
			ctx.stroke()
			
			// 绘制Y轴标签
			const value = Math.round((maxValue * i / 5) * 100) / 100
			ctx.setFontSize(10)
			ctx.setFillStyle('#999999')
			ctx.setTextAlign('right')
			ctx.fillText(value, padding - 5, y + 4)
		}
		
		// 绘制柱状图
		props.chartData.forEach((item, index) => {
			const barHeightIncome = (item.income / maxValue) * (height - padding * 2)
			const barHeightExpense = (item.expense / maxValue) * (height - padding * 2)
			
			// 收入柱子
			const incomeX = padding + index * (barWidth * 2 + barSpacing)
			const incomeY = height - padding - barHeightIncome
			ctx.setFillStyle('#4a90e2')
			ctx.fillRect(incomeX, incomeY, barWidth, barHeightIncome)
			
			// 支出柱子
			const expenseX = incomeX + barWidth + barSpacing/2
			const expenseY = height - padding - barHeightExpense
			ctx.setFillStyle('#ff6b6b')
			ctx.fillRect(expenseX, expenseY, barWidth, barHeightExpense)
		})
		
		// 绘制X轴
		ctx.setStrokeStyle('#333')
		ctx.setLineWidth(1)
		ctx.beginPath()
		ctx.moveTo(padding, height - padding)
		ctx.lineTo(width - padding, height - padding)
		ctx.stroke()
		
		// 绘制Y轴
		ctx.beginPath()
		ctx.moveTo(padding, padding)
		ctx.lineTo(padding, height - padding)
		ctx.stroke()
		
		// 绘制图例
		const legendY = 15
		ctx.setFontSize(12)
		
		// 收入图例
		ctx.setFillStyle('#4a90e2')
		ctx.fillRect(width - 100, legendY, 10, 10)
		ctx.setFillStyle('#333')
		ctx.fillText('收入', width - 85, legendY + 10)
		
		// 支出图例
		ctx.setFillStyle('#ff6b6b')
		ctx.fillRect(width - 40, legendY, 10, 10)
		ctx.setFillStyle('#333')
		ctx.fillText('支出', width - 25, legendY + 10)
		
		ctx.draw()
	}
	
	// 生命周期钩子
	onMounted(() => {
		drawChart()
	})
	
	// 监听器
	watch(() => props.chartData, () => {
		drawChart()
	}, { deep: true })
</script>

<style>
	.bar-chart {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	
	.chart-canvas {
		width: 300px;
		height: 200px;
	}
	
	.x-axis {
		display: flex;
		justify-content: space-around;
		width: 100%;
		padding: 0 20rpx;
		box-sizing: border-box;
	}
	
	.x-label {
		font-size: 24rpx;
		color: #666;
		text-align: center;
		flex: 1;
	}
</style>