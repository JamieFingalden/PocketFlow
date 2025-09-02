<template>
	<view class="line-chart">
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
		},
		chartType: {
			type: String,
			default: 'income' // 'income' 或 'expense'
		}
	})
	
	// 响应式数据
	const canvasId = ref('lineChartCanvas')
	
	// 方法定义
	const drawChart = () => {
		const ctx = uni.createCanvasContext(canvasId.value, this)
		
		const width = props.width
		const height = props.height
		const padding = 20
		
		// 获取最大值
		let maxValue = 0
		if (props.chartData.length > 0) {
			if (props.chartData[0].value !== undefined) {
				// 新的数据格式 { value: number }
				maxValue = Math.max(...props.chartData.map(item => item.value))
			} else {
				// 旧的数据格式 { income: number, expense: number }
				maxValue = Math.max(...props.chartData.map(item => Math.max(item.income, item.expense)))
			}
		}
		
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
		
		// 计算点的位置
		const points = []
		const spacing = (width - padding * 2) / (props.chartData.length - 1)
		
		props.chartData.forEach((item, index) => {
			const x = padding + index * spacing
			let y = 0
			
			if (item.value !== undefined) {
				// 新的数据格式 { value: number }
				y = height - padding - (item.value / maxValue) * (height - padding * 2)
			} else {
				// 旧的数据格式 { income: number, expense: number }
				const value = props.chartType === 'income' ? item.income : item.expense
				y = height - padding - (value / maxValue) * (height - padding * 2)
			}
			
			points.push({ x, y })
		})
		
		// 确定颜色和标签
		const color = props.chartType === 'income' ? '#4a90e2' : '#ff6b6b'
		const label = props.chartType === 'income' ? '收入' : '支出'
		
		// 绘制折线
		ctx.setStrokeStyle(color)
		ctx.setLineWidth(2)
		ctx.beginPath()
		points.forEach((point, index) => {
			if (index === 0) {
				ctx.moveTo(point.x, point.y)
			} else {
				ctx.lineTo(point.x, point.y)
			}
		})
		ctx.stroke()
		
		// 绘制数据点
		ctx.setFillStyle(color)
		points.forEach(point => {
			ctx.beginPath()
			ctx.arc(point.x, point.y, 4, 0, 2 * Math.PI)
			ctx.fill()
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
		
		ctx.setFillStyle(color)
		ctx.fillRect(width - 50, legendY, 10, 10)
		ctx.setFillStyle('#333')
		ctx.fillText(label, width - 35, legendY + 10)
		
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
	.line-chart {
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