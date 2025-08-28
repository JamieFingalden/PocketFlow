<template>
	<view class="container">
		<!-- 顶部导航 -->
		<view class="header">
			<text class="back-btn" @click="cancel">取消</text>
			<text class="title">记账</text>
			<text class="save-btn" @click="saveBill">保存</text>
		</view>
		
		<view class="content">
			<!-- 金额输入区 -->
			<view class="amount-section" @click="showNumberKeyboard">
				<text class="amount-label">金额</text>
				<text class="amount-display">{{ billForm.amount || '0.00' }}</text>
				<text class="edit-icon">✎</text>
			</view>
									
			<!-- 数字键盘弹窗 -->
			<view class="keyboard-popup" v-if="showKeyboard">
				<NumberKeyboard 
					:value="billForm.amount" 
					@confirm="onKeyboardConfirm"
					@input="onKeyboardInput"
				/>
			</view>
			
			<!-- 账单类型选择 -->
			<view class="type-selector">
				<view 
					class="type-item" 
					:class="{ active: billForm.type === 'expense' }"
					@click="selectType('expense')"
				>
					支出
				</view>
				<view 
					class="type-item" 
					:class="{ active: billForm.type === 'income' }"
					@click="selectType('income')"
				>
					收入
				</view>
			</view>
			
			<!-- 表单区域 -->
			<view class="form-section">
				<!-- 分类选择 -->
				<view class="form-item" @click="showCategoryPicker">
					<text class="item-label">分类</text>
					<view class="item-value">
						<text>{{ billForm.category || '请选择分类' }}</text>
						<text class="arrow">></text>
					</view>
				</view>
				
				<!-- 备注 -->
				<view class="form-item">
					<text class="item-label">备注</text>
					<input 
						class="item-input" 
						type="text" 
						placeholder="请输入备注" 
						v-model="billForm.remark"
					/>
				</view>
				
				<!-- 日期时间 -->
				<view class="form-item" @click="showDateTimePicker">
					<text class="item-label">日期时间</text>
					<view class="item-value">
						<text>{{ billForm.datetime || currentTime }}</text>
						<text class="arrow">></text>
					</view>
				</view>
			</view>
			
			<!-- OCR上传图片 -->
			<view class="ocr-section">
				<button class="ocr-btn" @click="uploadBillImage">
					📷 上传账单图片
				</button>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, reactive, onMounted } from 'vue'
	import NumberKeyboard from '@/components/NumberKeyboard.vue'
	
	// 响应式数据
	const billForm = reactive({
		type: 'expense', // expense: 支出, income: 收入
		amount: '',
		category: '',
		remark: '',
		datetime: ''
	})
	
	const showKeyboard = ref(false)
	
	// 获取当前时间
	const getCurrentTime = () => {
		const now = new Date()
		const year = now.getFullYear()
		const month = String(now.getMonth() + 1).padStart(2, '0')
		const day = String(now.getDate()).padStart(2, '0')
		const hours = String(now.getHours()).padStart(2, '0')
		const minutes = String(now.getMinutes()).padStart(2, '0')
		return `${year}-${month}-${day} ${hours}:${minutes}`
	}
	
	const currentTime = getCurrentTime()
	
	// 方法定义
	const selectType = (type) => {
		billForm.type = type
	}
	
	const showNumberKeyboard = () => {
		showKeyboard.value = true
	}
	
	const onKeyboardInput = (value) => {
		billForm.amount = value
	}
	
	const onKeyboardConfirm = (value) => {
		billForm.amount = value
		showKeyboard.value = false
	}
	
	const simulateOCRRecognition = () => {
		// 模拟OCR识别结果
		setTimeout(() => {
			billForm.amount = '35.00'
			billForm.category = '餐饮'
			billForm.remark = '星巴克咖啡'
			billForm.datetime = getCurrentTime()
			uni.showToast({
				title: 'OCR识别完成',
				icon: 'success'
			})
		}, 1000)
	}
	
	const showCategoryPicker = () => {
		// 显示分类选择器
		uni.showToast({
			title: '选择分类功能待实现',
			icon: 'none'
		})
	}
	
	const showDateTimePicker = () => {
		// 显示日期时间选择器
		uni.showToast({
			title: '选择日期时间功能待实现',
			icon: 'none'
		})
	}
	
	const uploadBillImage = () => {
		// 上传账单图片进行OCR识别
		uni.chooseImage({
			count: 1,
			success: (res) => {
				console.log('选择的图片:', res.tempFilePaths)
				uni.showToast({
					title: '图片上传成功，开始识别',
					icon: 'success'
				})
				// 模拟OCR识别结果
				billForm.amount = '35.00'
				billForm.category = '餐饮'
				billForm.remark = '星巴克咖啡'
				billForm.datetime = getCurrentTime()
			},
			fail: (err) => {
				console.error('选择图片失败:', err)
				uni.showToast({
					title: '选择图片失败',
					icon: 'none'
				})
			}
		})
	}
	
	const saveBill = () => {
		// 保存账单
		console.log('保存账单:', billForm)
		uni.showToast({
			title: '账单保存成功',
			icon: 'success'
		})
		// 返回账单列表页
		uni.navigateBack()
	}
	
	const cancel = () => {
		uni.navigateBack()
	}
	
	// 生命周期钩子
	onMounted(() => {
		// 接收从首页传递的图片路径
		const pages = getCurrentPages()
		const current = pages[pages.length - 1]
		const options = current.options || {}
		
		if (options.imagePath) {
			console.log('接收到图片路径:', options.imagePath)
			// 模拟OCR识别
			simulateOCRRecognition()
		}
	})
</script>

<style>
	.container {
		background-color: #f8f8f8;
		min-height: 100vh;
	}
	
	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 20rpx 30rpx;
		background-color: #fff;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		position: sticky;
		top: 0;
		z-index: 999;
	}
	
	.back-btn, .save-btn {
		font-size: 32rpx;
		color: #007AFF;
		padding: 10rpx 20rpx;
	}
	
	.title {
		font-size: 36rpx;
		font-weight: bold;
		color: #333;
	}
	
	.content {
		padding: 20rpx;
	}
	
	.amount-section {
		background-color: #fff;
		border-radius: 20rpx;
		padding: 40rpx 30rpx;
		margin-bottom: 20rpx;
		text-align: center;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		position: relative;
	}
	
	.amount-label {
		font-size: 28rpx;
		color: #999;
		margin-bottom: 20rpx;
		display: block;
	}
	
	.amount-display {
		font-size: 48rpx;
		font-weight: bold;
		color: #333;
	}
	
	.edit-icon {
		position: absolute;
		top: 30rpx;
		right: 30rpx;
		font-size: 28rpx;
		color: #999;
	}
	
	.keyboard-popup {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		z-index: 1000;
	}
	
	.type-selector {
		display: flex;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		overflow: hidden;
	}
	
	.type-item {
		flex: 1;
		text-align: center;
		padding: 30rpx;
		font-size: 32rpx;
		font-weight: bold;
	}
	
	.type-item.active {
		background-color: #007AFF;
		color: #fff;
	}
	
	.form-section {
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.form-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}
	
	.form-item:last-child {
		border-bottom: none;
	}
	
	.item-label {
		font-size: 32rpx;
		color: #333;
	}
	
	.item-value {
		display: flex;
		align-items: center;
		color: #999;
		font-size: 28rpx;
	}
	
	.item-input {
		flex: 1;
		text-align: right;
		border: none;
		outline: none;
		font-size: 28rpx;
		color: #666;
	}
	
	.arrow {
		margin-left: 10rpx;
	}
	
	.ocr-section {
		text-align: center;
		margin: 30rpx 0;
	}
	
	.ocr-btn {
		width: 100%;
		height: 80rpx;
		background-color: #007AFF;
		color: #fff;
		border-radius: 40rpx;
		font-size: 32rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		box-shadow: 0 4rpx 10rpx rgba(0, 122, 255, 0.3);
	}
</style>
