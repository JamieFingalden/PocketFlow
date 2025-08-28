<template>
	<view class="container">
		<view class="header">
			<text class="title">分类管理</text>
		</view>
		
		<view class="tabs">
			<view 
				class="tab" 
				:class="{ active: activeTab === 'expense' }"
				@click="switchTab('expense')"
			>
				支出分类
			</view>
			<view 
				class="tab" 
				:class="{ active: activeTab === 'income' }"
				@click="switchTab('income')"
			>
				收入分类
			</view>
		</view>
		
		<view class="category-list">
			<view 
				class="category-item" 
				v-for="(category, index) in filteredCategories" 
				:key="index"
			>
				<view class="category-info">
					<text class="category-icon">{{ category.icon }}</text>
					<text class="category-name">{{ category.name }}</text>
				</view>
				<view class="category-actions">
					<text class="edit-btn" @click="editCategory(category)">编辑</text>
					<text class="delete-btn" @click="deleteCategory(category)">删除</text>
				</view>
			</view>
		</view>
		
		<view class="add-section">
			<button class="add-btn" @click="showAddModal">+ 添加分类</button>
		</view>
		
		<!-- 添加/编辑分类弹窗 -->
		<view class="modal" v-if="showModal">
			<view class="modal-content">
				<view class="modal-header">
					<text class="modal-title">{{ isEdit ? '编辑分类' : '添加分类' }}</text>
				</view>
				
				<view class="modal-form">
					<view class="form-group">
						<text class="label">分类名称</text>
						<input 
							class="input" 
							type="text" 
							placeholder="请输入分类名称" 
							v-model="categoryForm.name"
						/>
					</view>
					
					<view class="form-group">
						<text class="label">图标</text>
						<input 
							class="input" 
							type="text" 
							placeholder="请输入图标字符" 
							v-model="categoryForm.icon"
						/>
					</view>
				</view>
				
				<view class="modal-actions">
					<button class="cancel-btn" @click="closeModal">取消</button>
					<button class="confirm-btn" @click="saveCategory">{{ isEdit ? '更新' : '添加' }}</button>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref, computed } from 'vue'
	
	// 响应式数据
	const activeTab = ref('expense')
	const showModal = ref(false)
	const isEdit = ref(false)
	const editCategoryData = ref(null)
	const categoryForm = ref({
		name: '',
		icon: ''
	})
	const categories = ref([
		{ id: 1, name: '餐饮', icon: '餐', type: 'expense' },
		{ id: 2, name: '交通', icon: '交', type: 'expense' },
		{ id: 3, name: '购物', icon: '购', type: 'expense' },
		{ id: 4, name: '医疗', icon: '医', type: 'expense' },
		{ id: 5, name: '工资', icon: '薪', type: 'income' },
		{ id: 6, name: '奖金', icon: '奖', type: 'income' }
	])
	
	// 计算属性 - 根据当前标签页过滤分类
	const filteredCategories = computed(() => {
		return categories.value.filter(category => category.type === activeTab.value)
	})
	
	// 方法定义
	const switchTab = (tab) => {
		activeTab.value = tab
	}
	
	const showAddModal = () => {
		isEdit.value = false
		categoryForm.value = { name: '', icon: '' }
		showModal.value = true
	}
	
	const editCategory = (category) => {
		isEdit.value = true
		editCategoryData.value = category
		categoryForm.value = { ...category }
		showModal.value = true
	}
	
	const deleteCategory = (category) => {
		uni.showModal({
			title: '确认删除',
			content: `确定要删除分类 "${category.name}" 吗？`,
			success: (res) => {
				if (res.confirm) {
					// 删除分类逻辑
					categories.value = categories.value.filter(item => item.id !== category.id)
					uni.showToast({
						title: '删除成功',
						icon: 'success'
					})
				}
			}
		})
	}
	
	const closeModal = () => {
		showModal.value = false
	}
	
	const saveCategory = () => {
		if (!categoryForm.value.name) {
			uni.showToast({
				title: '请输入分类名称',
				icon: 'none'
			})
			return
		}
		
		if (isEdit.value) {
			// 更新分类
			const index = categories.value.findIndex(item => item.id === editCategoryData.value.id)
			if (index !== -1) {
				categories.value[index] = { ...editCategoryData.value, ...categoryForm.value }
			}
			uni.showToast({
				title: '更新成功',
				icon: 'success'
			})
		} else {
			// 添加分类
			const newCategory = {
				id: Date.now(),
				name: categoryForm.value.name,
				icon: categoryForm.value.icon,
				type: activeTab.value
			}
			categories.value.push(newCategory)
			uni.showToast({
				title: '添加成功',
				icon: 'success'
			})
		}
		
		closeModal()
	}
</script>

<style>
	.container {
		padding: 20rpx;
		background-color: #f8f8f8;
		min-height: 100vh;
	}
	
	.header {
		text-align: center;
		padding: 20rpx;
		background-color: #fff;
		margin-bottom: 20rpx;
		border-radius: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.title {
		font-size: 36rpx;
		font-weight: bold;
		color: #333;
	}
	
	.tabs {
		display: flex;
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.tab {
		flex: 1;
		text-align: center;
		padding: 20rpx;
		font-size: 32rpx;
	}
	
	.tab.active {
		background-color: #007AFF;
		color: #fff;
		border-radius: 20rpx;
	}
	
	.category-list {
		background-color: #fff;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
	}
	
	.category-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}
	
	.category-item:last-child {
		border-bottom: none;
	}
	
	.category-info {
		display: flex;
		align-items: center;
	}
	
	.category-icon {
		width: 60rpx;
		height: 60rpx;
		border-radius: 30rpx;
		background-color: #e0f7fa;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-right: 20rpx;
		font-size: 28rpx;
	}
	
	.category-name {
		font-size: 32rpx;
	}
	
	.category-actions {
		display: flex;
		gap: 20rpx;
	}
	
	.edit-btn, .delete-btn {
		font-size: 28rpx;
	}
	
	.edit-btn {
		color: #007AFF;
	}
	
	.delete-btn {
		color: #ff6b6b;
	}
	
	.add-section {
		text-align: center;
	}
	
	.add-btn {
		width: 100%;
		height: 80rpx;
		background-color: #007AFF;
		color: #fff;
		border-radius: 10rpx;
		font-size: 32rpx;
	}
	
	/* 弹窗样式 */
	.modal {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0, 0, 0, 0.5);
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 999;
	}
	
	.modal-content {
		width: 80%;
		background-color: #fff;
		border-radius: 20rpx;
		padding: 30rpx;
	}
	
	.modal-header {
		text-align: center;
		margin-bottom: 30rpx;
	}
	
	.modal-title {
		font-size: 36rpx;
		font-weight: bold;
	}
	
	.modal-form {
		margin-bottom: 30rpx;
	}
	
	.form-group {
		margin-bottom: 20rpx;
	}
	
	.label {
		display: block;
		font-size: 28rpx;
		color: #666;
		margin-bottom: 10rpx;
	}
	
	.input {
		width: 100%;
		height: 80rpx;
		border: 1rpx solid #e0e0e0;
		border-radius: 10rpx;
		padding: 0 20rpx;
		font-size: 32rpx;
		box-sizing: border-box;
	}
	
	.modal-actions {
		display: flex;
		gap: 20rpx;
	}
	
	.cancel-btn, .confirm-btn {
		flex: 1;
		height: 80rpx;
		border-radius: 10rpx;
		font-size: 32rpx;
	}
	
	.cancel-btn {
		background-color: #f0f0f0;
		color: #666;
	}
	
	.confirm-btn {
		background-color: #007AFF;
		color: #fff;
	}
</style>