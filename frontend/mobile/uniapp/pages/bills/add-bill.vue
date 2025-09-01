<template>
	<view class="container">
		<!-- 顶部导航 -->
		<view class="header">
			<text class="back-btn" @click="cancel">取消</text>
			<text class="title">记账</text>
			<text class="save-btn" @click="saveBill">保存</text>
		</view>
		
		<view class="content">
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
			
			<!-- 分类选择 -->
			<view class="category-selector" v-if="billForm.type">
				<view 
					class="category-item" 
					v-for="(category, index) in categories" 
					:key="index"
					:class="{ active: billForm.category === category.name }"
					@click="selectCategory(category.name)"
				>
					<text class="category-icon">{{ category.icon }}</text>
					<text class="category-name">{{ category.name }}</text>
				</view>
			</view>
			
			<!-- 金额输入 -->
			<view class="form-item" @click="showNumberKeyboard" v-if="billForm.category">
				<text class="item-label">金额</text>
				<view class="item-value">
					<text>{{ temporaryAmount || billForm.amount || '0.00' }}</text>
					<text class="arrow">></text>
				</view>
			</view>
			
			<!-- 备注 -->
			<view class="form-item" v-if="billForm.amount">
				<text class="item-label">备注</text>
				<input 
					class="item-input" 
					type="text" 
					placeholder="请输入备注" 
					v-model="billForm.remark"
				/>
			</view>
			
			<!-- 日期 -->
			<view class="form-item" v-if="billForm.amount">
				<text class="item-label">日期</text>
				<view class="item-value">
					<text>{{ billForm.datetime || currentDate }}</text>
				</view>
			</view>
		</view>
		
		<!-- 数字键盘弹窗 -->
		<view class="keyboard-popup" v-if="showKeyboard" @click="closeKeyboard">
			<view class="keyboard-container">
				<NumberKeyboard 
					ref="numberKeyboardRef"
					:value="billForm.amount" 
					@confirm="onKeyboardConfirm"
					@input="onKeyboardInput"
					@close="onKeyboardClose"
				/>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, reactive, computed, onMounted } from 'vue'
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
	const temporaryAmount = ref('')
	const numberKeyboardRef = ref(null)
	
	// 分类数据(带图标占位符)
	const expenseCategories = [
		{ name: '餐饮', icon: '🍽️' },
		{ name: '交通', icon: '🚗' },
		{ name: '购物', icon: '🛍️' },
		{ name: '娱乐', icon: '🎮' },
		{ name: '医疗', icon: '🏥' },
		{ name: '住房', icon: '🏠' },
		{ name: '其他', icon: '📦' }
	]
	const incomeCategories = [
		{ name: '工资', icon: '💰' },
		{ name: '奖金', icon: '🏆' },
		{ name: '投资', icon: '📈' },
		{ name: '兼职', icon: '💼' },
		{ name: '其他', icon: '📦' }
	]
	
	// 计算属性
	const categories = computed(() => {
		return billForm.type === 'expense' ? expenseCategories : incomeCategories
	})
	
	// 获取当前日期
	const getCurrentDate = () => {
		const now = new Date()
		const year = now.getFullYear()
		const month = String(now.getMonth() + 1).padStart(2, '0')
		const day = String(now.getDate()).padStart(2, '0')
		return `${year}-${month}-${day}`
	}
	
	const currentDate = getCurrentDate()
	billForm.datetime = currentDate
	
	// 方法定义
	const selectType = (type) => {
		billForm.type = type
		billForm.category = ''
		billForm.amount = ''
		billForm.remark = ''
	}
	
	const selectCategory = (category) => {
		billForm.category = category
		showNumberKeyboard()
	}
	
	const showNumberKeyboard = () => {
		if (!billForm.category) {
			uni.showToast({
				title: '请先选择分类',
				icon: 'none'
			})
			return
		}
		showKeyboard.value = true
	}
	
	const onKeyboardInput = (value) => {
		// 实时更新显示，但不更新表单数据
		// 这样用户可以看到正在输入的数值
		console.log('键盘输入值:', value)
		temporaryAmount.value = value
	}
	
	const onKeyboardConfirm = (value) => {
		console.log('键盘确认值:', value)
		billForm.amount = value
		showKeyboard.value = false
		temporaryAmount.value = ''
	}
	
	const saveBill = () => {
		// 验证表单
		if (!billForm.type) {
			uni.showToast({
				title: '请选择收入或支出',
				icon: 'none'
			})
			return
		}
		
		if (!billForm.category) {
			uni.showToast({
				title: '请选择分类',
				icon: 'none'
			})
			return
		}
		
		if (!billForm.amount) {
			uni.showToast({
				title: '请输入金额',
				icon: 'none'
			})
			return
		}
		
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
	
	// 空函数，保留以避免错误
	
	const closeKeyboard = () => {
		console.log('关闭键盘函数被调用')
		// 调用数字键盘组件的 closeKeyboard 方法来完成计算
		if (numberKeyboardRef.value && typeof numberKeyboardRef.value.closeKeyboard === 'function') {
			console.log('调用数字键盘组件的 closeKeyboard 方法')
			numberKeyboardRef.value.closeKeyboard()
		} else {
			console.log('直接关闭键盘')
			showKeyboard.value = false
			temporaryAmount.value = ''
		}
	}
	
	const onKeyboardClose = (value) => {
		console.log('数字键盘关闭事件:', value)
		billForm.amount = value
		showKeyboard.value = false
		temporaryAmount.value = ''
	}
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
	
	.form-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
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
	
	.keyboard-popup {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		z-index: 1000;
		padding: 20rpx;
	}
	
	
	.category-selector {
		display: flex;
		flex-wrap: wrap;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
		padding: 20rpx;
	}
	
	.category-selector .category-item {
		flex: 0 0 33.33%;
		text-align: center;
		padding: 20rpx;
		font-size: 28rpx;
		box-sizing: border-box;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
	
	.category-selector .category-item.active {
		background-color: #007AFF;
		color: #fff;
		border-radius: 10rpx;
	}
	
	.category-icon {
		font-size: 40rpx;
		margin-bottom: 10rpx;
	}
	
	
	.category-name {
		font-size: 28rpx;
	}
</style>