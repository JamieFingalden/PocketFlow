# API文档

## 1. 后端API (Go/Gin)

### 1.1 用户相关接口

#### 注册
- **URL**: `POST /api/v1/register`
- **描述**: 用户注册接口
- **请求参数**:
  ```json
  {
    "username": "string",
    "password": "string",
    "email": "string"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "message": "User registered successfully",
    "data": {
      "id": 1,
      "username": "string",
      "email": "string",
      "role": "USER",
      "enabled": true,
      "create_time": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### 登录
- **URL**: `POST /api/v1/login`
- **描述**: 用户登录接口
- **请求参数**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Login successful",
    "data": {
      "token": "jwt_token_string"
    }
  }
  ```

### 1.2 账单相关接口

#### 创建账单
- **URL**: `POST /api/v1/bills`
- **描述**: 创建新的账单记录
- **请求头**: `Authorization: Bearer {token}`
- **请求参数**:
  ```json
  {
    "type": "EXPENSE|INCOME",
    "category_id": 1,
    "amount": 100.00,
    "remark": "string",
    "bill_time": "2023-01-01T12:00:00Z"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Bill created successfully",
    "data": {
      "id": 1,
      "user_id": 1,
      "type": "EXPENSE",
      "category_id": 1,
      "amount": 100.00,
      "remark": "string",
      "bill_time": "2023-01-01T12:00:00Z",
      "create_time": "2023-01-01T12:00:00Z"
    }
  }
  ```

#### 获取账单列表
- **URL**: `GET /api/v1/bills`
- **描述**: 获取当前用户的所有账单记录
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Bills retrieved successfully",
    "data": [
      {
        "id": 1,
        "user_id": 1,
        "type": "EXPENSE",
        "category_id": 1,
        "amount": 100.00,
        "remark": "string",
        "bill_time": "2023-01-01T12:00:00Z",
        "create_time": "2023-01-01T12:00:00Z"
      }
    ]
  }
  ```

#### 获取单个账单
- **URL**: `GET /api/v1/bills/{id}`
- **描述**: 获取指定ID的账单记录
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Bill retrieved successfully",
    "data": {
      "id": 1,
      "user_id": 1,
      "type": "EXPENSE",
      "category_id": 1,
      "amount": 100.00,
      "remark": "string",
      "bill_time": "2023-01-01T12:00:00Z",
      "create_time": "2023-01-01T12:00:00Z"
    }
  }
  ```

#### 更新账单
- **URL**: `PUT /api/v1/bills/{id}`
- **描述**: 更新指定ID的账单记录
- **请求头**: `Authorization: Bearer {token}`
- **请求参数**:
  ```json
  {
    "type": "EXPENSE|INCOME",
    "category_id": 1,
    "amount": 100.00,
    "remark": "string",
    "bill_time": "2023-01-01T12:00:00Z"
  }
  ```
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Bill updated successfully",
    "data": {
      "id": 1,
      "user_id": 1,
      "type": "EXPENSE",
      "category_id": 1,
      "amount": 100.00,
      "remark": "string",
      "bill_time": "2023-01-01T12:00:00Z",
      "create_time": "2023-01-01T12:00:00Z"
    }
  }
  ```

#### 删除账单
- **URL**: `DELETE /api/v1/bills/{id}`
- **描述**: 删除指定ID的账单记录
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "Bill deleted successfully",
    "data": null
  }
  ```

### 1.3 统计相关接口

#### 月度汇总
- **URL**: `GET /api/v1/statistics/monthly-summary`
- **描述**: 获取月度收支汇总统计
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "total_income": 5000.00,
      "total_expense": 3000.00,
      "net_income": 2000.00
    }
  }
  ```

#### 分类统计
- **URL**: `GET /api/v1/statistics/category`
- **描述**: 获取按分类的收支统计
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "success",
    "data": [
      {
        "category_id": 1,
        "category_name": "餐饮",
        "type": "EXPENSE",
        "total_amount": 1000.00,
        "count": 10
      }
    ]
  }
  ```

#### 日常统计
- **URL**: `GET /api/v1/statistics/daily`
- **描述**: 获取日常收支统计
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "success",
    "data": [
      {
        "date": "2023-01-01",
        "total_income": 1000.00,
        "total_expense": 500.00
      }
    ]
  }
  ```

#### 月度图表统计
- **URL**: `GET /api/v1/statistics/monthly-chart`
- **描述**: 获取月度图表统计数据
- **请求头**: `Authorization: Bearer {token}`
- **响应**:
  ```json
  {
    "code": 200,
    "message": "success",
    "data": [
      {
        "month": "2023-01",
        "total_income": 5000.00,
        "total_expense": 3000.00
      }
    ]
  }
  ```

## 2. OCR服务API (Python/Flask)

### 2.1 图片处理
- **URL**: `POST /process_image`
- **描述**: 处理上传的账单图片，进行OCR识别和AI分类
- **请求格式**: `multipart/form-data`
- **请求参数**: `image` (图片文件)
- **响应**:
  ```json
  {
    "status": "success",
    "ocr_text": "识别出的文本内容",
    "category": {
      "category": "餐饮",
      "amount": 35.0,
      "date": "2023-01-01"
    }
  }
  ```

### 2.2 健康检查
- **URL**: `GET /health`
- **描述**: 检查OCR服务运行状态
- **响应**:
  ```json
  {
    "status": "healthy",
    "message": "服务运行正常"
  }
  ```