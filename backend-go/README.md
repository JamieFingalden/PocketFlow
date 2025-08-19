# PocketFlow 后端项目

这是一个使用 Go Gin 框架开发的记账应用后端项目。

## 项目结构

```
backend-go/
├── api/                 # 控制器层，处理HTTP请求和响应
├── cmd/api/            # 程序入口点
├── config/             # 配置文件
├── database/           # 数据库初始化
├── middleware/         # 中间件
├── models/             # 数据库模型
├── repository/         # 数据访问层
├── router/             # 路由注册
├── services/           # 业务逻辑层
├── utils/              # 工具函数
├── go.mod             # Go模块文件
├── go.sum             # Go模块校验和
└── SQL.sql            # 数据库表结构
```

## 功能特性

1. 用户注册与登录（JWT鉴权）
2. 账单的增删改查
3. 健康检查接口
4. OCR代理接口（转发到Python服务）

## 数据库模型

- User: 用户账户信息
- UserProfile: 用户业务信息（预算、余额等）
- Category: 账单分类
- Bill: 记账账单
- QuickTemplate: 快捷记账模板

## API接口

### 用户相关
- `POST /api/v1/register` - 用户注册
- `POST /api/v1/login` - 用户登录

### 账单相关
- `POST /api/v1/bills` - 创建账单
- `GET /api/v1/bills` - 获取账单列表
- `GET /api/v1/bills/:id` - 获取单个账单
- `PUT /api/v1/bills/:id` - 更新账单
- `DELETE /api/v1/bills/:id` - 删除账单

### 其他
- `GET /health` - 健康检查
- `POST /ocr/process` - OCR处理

## 数据库设置

在运行程序之前，需要确保 MySQL 数据库已安装并创建了 pocketflow 数据库：

```sql
CREATE DATABASE IF NOT EXISTS pocketflow CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

配置文件位于 `config/config.yaml`，包括：
- 服务器配置
- 数据库连接信息（默认用户: root, 密码: 123456）
- JWT密钥
- OCR服务URL

## 编译和运行

```bash
# 编译项目
go build -o backend-go.exe ./cmd/api

# 运行项目
./backend-go.exe
```

## 依赖

- Gin: HTTP框架
- GORM: ORM库
- MySQL驱动: 数据库连接
- JWT: 用户认证
- Bcrypt: 密码加密