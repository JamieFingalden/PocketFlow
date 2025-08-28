<template>
	<view class="pie-chart">
		<canvas 
			class="chart-canvas" 
			:canvas-id="canvasId" 
			:id="canvasId"
			@touchstart="handleTouchStart"
		></canvas>
		<view class="legend">
			<view 
				class="legend-item" 
				v-for="(item, index) in chartData" 
				:key="index"
			>
				<view class="legend-color" :style="{ backgroundColor: item.color }"></view>
				<text class="legend-label">{{ item.name }}</text>
				<text class="legend-value">¥{{ item.value }}</text>
				<text class="legend-percent">{{ item.percent }}%</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, onMounted, watch, defineProps, defineEmits } from 'vue'
	
	// 定义props
	const props = defineProps({
		chartData: {
			type: Array,
			default: () => []
		},
		width: {
			type: [String, Number],
			default: 300
		},
		height: {
			type: [String, Number],
			default: 300
		}
	})
	
	// 定义emits
	const emits = defineEmits(['chart-click'])
	
	// 响应式数据
	const canvasId = ref('pieChartCanvas')
	
	// 方法定义
	const drawChart = () => {
		const ctx = uni.createCanvasContext(canvasId.value, this)
		
		const width = props.width
		const height = props.height
		const centerX = width / 2
		const centerY = height / 2
		const radius = Math.min(width, height) / 2 - 30
		
		let startAngle = -Math.PI / 2
		const total = props.chartData.reduce((sum, item) => sum + item.value, 0)
		
		// 绘制背景
		ctx.setFillStyle('#ffffff')
		ctx.fillRect(0, 0, width, height)
		
		// 绘制饼图
		props.chartData.forEach((item, index) => {
			const sliceAngle = (item.value / total) * 2 * Math.PI
			const endAngle = startAngle + sliceAngle
			
			// 绘制扇形
			ctx.beginPath()
			ctx.moveTo(centerX, centerY)
			ctx.arc(centerX, centerY, radius, startAngle, endAngle)
			ctx.closePath()
			ctx.setFillStyle(item.color)
			ctx.fill()
			
			// 绘制边框
			ctx.beginPath()
			ctx.arc(centerX, centerY, radius, startAngle, endAngle)
			ctx.setStrokeStyle('#fff')
			ctx.setLineWidth(2)
			ctx.stroke()
			
			startAngle = endAngle
		})
		
		// 绘制中心圆（可选）
		ctx.beginPath()
		ctx.arc(centerX, centerY, radius * 0.2, 0, 2 * Math.PI)
		ctx.setFillStyle('#fff')
		ctx.fill()
		ctx.setStrokeStyle('#eee')
		ctx.setLineWidth(1)
		ctx.stroke()
		
		// 绘制百分比标签
		startAngle = -Math.PI / 2
		props.chartData.forEach((item, index) => {
			const sliceAngle = (item.value / total) * 2 * Math.PI
			const midAngle = startAngle + sliceAngle / 2
			const labelRadius = radius * 0.7
			const labelX = centerX + Math.cos(midAngle) * labelRadius
			const labelY = centerY + Math.sin(midAngle) * labelRadius
			
			const percent = Math.round((item.value / total) * 100)
			if (percent > 5) { // 只显示占比大于5%的标签
				ctx.setFontSize(12)
				ctx.setFillStyle('#fff')
				ctx.setTextAlign('center')
				ctx.setTextBaseline('middle')
				ctx.fillText(`${percent}%`, labelX, labelY)
			}
			
			startAngle = startAngle + sliceAngle
		})
		
		ctx.draw()
	}
	
	const darkenColor = (color, percent) => {
		// 简单的颜色变暗函数
		const num = parseInt(color.replace("#", ""), 16)
		const amt = Math.round(2.55 * percent)
		const R = (num >> 16) - amt
		const G = (num >> 8 & 0x00FF) - amt
		const B = (num & 0x0000FF) - amt
		return "#" + (0x1000000 + (R < 255 ? R < 1 ? 0 : R : 255) * 0x10000 +
			(G < 255 ? G < 1 ? 0 : G : 255) * 0x100 +
			(B < 255 ? B < 1 ? 0 : B : 255)).toString(16).slice(1)
	}
	
	const handleTouchStart = (e) => {
		// 处理点击事件（可选）
		emits('chart-click', e)
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
	.pie-chart {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	
	.chart-canvas {
		width: 300px;
		height: 300px;
	}
	
	.legend {
		width: 100%;
		margin-top: 20rpx;
	}
	
	.legend-item {
		display: flex;
		align-items: center;
		padding: 10rpx 0;
	}
	
	.legend-color {
		width: 20rpx;
		height: 20rpx;
		border-radius: 10rpx;
		margin-right: 10rpx;
	}
	
	.legend-label {
		flex: 1;
		font-size: 28rpx;
		color: #333;
	}
	
	.legend-value {
		font-size: 28rpx;
		color: #666;
		margin-right: 10rpx;
	}
	
	.legend-percent {
		font-size: 28rpx;
		color: #999;
	}
</style>