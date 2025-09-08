# PocketFlow API 文档 (含测试标注)

本文档详细描述了 PocketFlow 后端 API 的所有接口，并包含基于代码静态分析的测试标注。

**基础 URL**: `/`

**通用响应格式**:

所有 API 响应都遵循以下标准格式：

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

- `code`: HTTP 状态码或自定义业务码。
- `message`: 对响应结果的描述。
- `data`: 实际的响应数据，可以是对象、数组或 null。

---

## 1. 公共接口

这些接口无需任何认证即可访问。

### 1.1. 健康检查

检查服务是否正常运行。

- **URL**: `/health`
- **Method**: `GET`
- **认证**: 无
- **成功响应**:
  - **Code**: `200 OK`

### 1.2. OCR 票据识别

- **URL**: `/ocr/process`
- **Method**: `POST`
- **认证**: 无
- 📝 **注意**: 此接口当前为占位符，未实现具体功能。代码中明确指出它只是一个占位响应。

### 1.3. API 文档 (Swagger)

提供 Swagger UI 界面来浏览和测试 API。

- **URL**: `/swagger/*any`
- **Method**: `GET`
- **认证**: 无

---

## 2. 认证接口

### 2.1. 用户注册

- **URL**: `/api/v1/register`
- **Method**: `POST`
- **认证**: 无
- **请求体**: `{"username": "string", "password": "string", "email": "string"}`
- **成功响应**: `200 OK` (返回创建的用户信息)
- **错误响应**: `400 Bad Request`, `500 Internal Server Error` (例如用户名已存在)

### 2.2. 用户登录

- **URL**: `/api/v1/login`
- **Method**: `POST`
- **认证**: 无
- **请求体**: `{"username": "string", "password": "string"}`
- **成功响应**: `200 OK` (返回 JWT Token)
- **错误响应**: `400 Bad Request`, `401 Unauthorized`

### 2.3. 管理员登录

- **URL**: `/api/v1/admin/login`
- **Method**: `POST`
- **认证**: 无
- **请求体**: `{"username": "string", "password": "string"}`
- **成功响应**: `200 OK` (返回 JWT Token)
- **错误响应**: `400 Bad Request`, `401 Unauthorized`

### 2.4. 管理员登出

- **URL**: `/api/v1/admin/logout`
- **Method**: `POST`
- **认证**: 无
- 📝 **注意**: 这是一个无操作接口。代码注释表明，登出逻辑依赖客户端删除 Token。这意味着即使用户调用了登出，其 Token 在过期前依然有效。对于高安全要求的应用，服务端应实现 Token 黑名单机制。

---

## 3. 用户接口 (需 JWT 认证)

以下所有接口都需要在请求头中携带 `Authorization: Bearer <token>`。

⚠️ **潜在安全问题 (IDOR)**: 对于所有通过 `:id` 操作资源的接口 (如 `GET /bills/:id`)，后端服务层必须严格校验该资源是否属于当前通过 JWT 认证的用户。如果仅通过 `id` 查询而不关联 `user_id`，会导致用户可以越权访问或操作其他用户的数据。

### 3.1. 账单管理 (Bills)

#### 3.1.1. 创建账单

- **URL**: `/api/v1/bills`
- **Method**: `POST`
- **请求体**: `{"type": "string", "category_id": "int", "amount": "float", "remark": "string", "bill_time": "time"}`
- **成功响应**: `200 OK` (返回创建的账单对象)

#### 3.1.2. 获取当前用户的所有账单

- **URL**: `/api/v1/bills`
- **Method**: `GET`
- **成功响应**: `200 OK` (返回账单对象数组)
- ⚠️ **潜在性能问题**: 此接口没有实现分页。如果单个用户的账单数量非常多，一次性返回所有数据可能导致服务器和客户端性能问题。

#### 3.1.3. 获取单个账单详情

- **URL**: `/api/v1/bills/:id`
- **Method**: `GET`
- **成功响应**: `200 OK` (返回单个账单对象)

#### 3.1.4. 更新账单

- **URL**: `/api/v1/bills/:id`
- **Method**: `PUT`
- **请求体**: (与创建账单类似)
- **成功响应**: `200 OK` (返回更新后的账单对象)
- 📝 **注意**: 请求体中的 `user_id` 字段应被忽略。更新操作应仅依赖于 JWT 中的 `user_id` 来确定归属，防止用户尝试将账单修改到其他用户下。

#### 3.1.5. 删除账单

- **URL**: `/api/v1/bills/:id`
- **Method**: `DELETE`
- **成功响应**: `200 OK`

### 3.2. 统计 (Statistics)

(所有统计接口均需要 `year` 和 `month` 作为 Query 参数)

#### 3.2.1. 获取月度收支概览

- **URL**: `/api/v1/statistics/monthly-summary`
- **Method**: `GET`

#### 3.2.2. 获取分类统计

- **URL**: `/api/v1/statistics/category`
- **Method**: `GET`
- **Query**: `type=income|expense`

#### 3.2.3. 获取每日收支统计

- **URL**: `/api/v1/statistics/daily`
- **Method**: `GET`

#### 3.2.4. 获取年度月度图表统计

- **URL**: `/api/v1/statistics/monthly-chart`
- **Method**: `GET`
- **Query**: `year`

---

## 4. 管理员接口 (需 Admin 认证)

以下接口需要管理员权限的 JWT Token。

### 4.1. 用户管理 (Users)

#### 4.1.1. 获取用户列表 (分页)

- **URL**: `/api/v1/admin/users`
- **Method**: `GET`
- **Query**: `page`, `page_size`

#### 4.1.2. 获取单个用户信息

- **URL**: `/api/v1/admin/users/:id`
- **Method**: `GET`

#### 4.1.3. 更新用户信息

- **URL**: `/api/v1/admin/users/:id`
- **Method**: `PUT`
- ⚠️ **潜在问题**: 此接口直接将请求体绑定到 `models.User` 结构。这意味着管理员可以更新包括 `role`, `enabled` 在内的所有字段。这种宽泛的绑定可能带来意想不到的副作用，建议使用专门的 `UpdateUserDTO` 对象，只暴露允许管理员修改的字段。

#### 4.1.4. 禁用用户

- **URL**: `/api/v1/admin/users/:id/disable`
- **Method**: `PUT`

#### 4.1.5. 启用用户

- **URL**: `/api/v1/admin/users/:id/enable`
- **Method**: `PUT`

### 4.2. 账单管理 (Bills)

#### 4.2.1. 获取所有账单 (分页)

- **URL**: `/api/v1/admin/bills`
- **Method**: `GET`
- **Query**: `page`, `page_size`, `user_id` (可选)

#### 4.2.2. 获取任意账单详情

- **URL**: `/api/v1/admin/bills/:id`
- **Method**: `GET`

#### 4.2.3. 更新任意账单

- **URL**: `/api/v1/admin/bills/:id`
- **Method**: `PUT`

#### 4.2.4. 删除任意账单

- **URL**: `/api/v1/admin/bills/:id`
- **Method**: `DELETE`

### 4.3. 统计管理 (Statistics)

#### 4.3.1. 获取系统整体统计

- **URL**: `/api/v1/admin/statistics/system`
- **Method**: `GET`

#### 4.3.2. 获取用户统计 (分页)

- **URL**: `/api/v1/admin/statistics/users`
- **Method**: `GET`
- **Query**: `page`, `page_size`

#### 4.3.3. 获取账单统计

- **URL**: `/api/v1/admin/statistics/bills`
- **Method**: `GET`
- **Query**: `year` (可选), `month` (可选)