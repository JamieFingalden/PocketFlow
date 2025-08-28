# 开发指南

## 1. 项目结构

```
PocketFlow/
├── backend-go/          # Go后端服务
│   ├── api/            # 控制器层
│   ├── cmd/            # 程序入口
│   ├── config/         # 配置文件
│   ├── database/       # 数据库初始化
│   ├── middleware/     # 中间件
│   ├── models/         # 数据模型
│   ├── repository/     # 数据访问层
│   ├── router/         # 路由注册
│   ├── services/       # 业务逻辑层
│   ├── utils/          # 工具函数
│   ├── docs/           # Swagger文档
│   ├── go.mod          # Go模块文件
│   ├── SQL.sql         # 数据库表结构
│   └── README.md       # 后端说明文档
├── ocr-ai/             # Python OCR+AI服务
│   ├── app.py          # Flask应用入口
│   ├── ocr_module.py   # OCR识别模块
│   ├── ai_module.py    # AI分类模块
│   ├── requirements.txt # Python依赖
│   └── README.md       # OCR服务说明文档
├── frontend/           # 前端项目
│   ├── mobile/         # 移动端应用
│   │   ├── ios/        # iOS原生应用
│   │   └── uniapp/     # Uniapp跨平台移动应用(Android)
│   └── web-admin/      # Web管理界面
├── docs/               # 项目文档
├── docker/             # Docker配置文件
```

## 2. 后端开发 (Go/Gin)

### 2.1 技术栈
- Go 1.25
- Gin Web框架
- GORM ORM库
- MySQL数据库
- JWT认证
- Swagger API文档

### 2.2 代码规范
1. 遵循Go官方代码风格
2. 使用清晰的命名规范
3. 添加必要的注释和文档
4. 编写单元测试

### 2.3 添加新API接口
1. 在`api/`目录创建对应控制器文件
2. 在`services/`目录实现业务逻辑
3. 在`repository/`目录添加数据访问方法
4. 在`router/router.go`注册路由
5. 更新Swagger注解

### 2.4 数据库操作
1. 在`models/`目录定义数据模型
2. 在`repository/`目录实现CRUD操作
3. 使用事务处理复杂业务逻辑
4. 注意SQL注入防护

## 3. OCR服务开发 (Python/Flask)

### 3.1 技术栈
- Python 3.8+
- Flask Web框架
- PaddleOCR文字识别
- Ollama本地AI服务
- 阿里云百炼API（备用）

### 3.2 代码规范
1. 遵循PEP 8代码规范
2. 使用类型提示
3. 添加docstring文档
4. 编写测试用例

### 3.3 OCR模块扩展
1. 在`ocr_module.py`中添加新的OCR处理逻辑
2. 支持更多语言和字体
3. 优化图像预处理算法
4. 添加错误处理和日志记录

### 3.4 AI分类优化
1. 在`ai_module.py`中调整提示词
2. 支持更多分类标签
3. 优化模型调用参数
4. 添加用户自定义分类功能

## 4. 前端开发

### 4.1 技术栈（假设）
- React/Vue框架
- TypeScript
- Ant Design/Ant Design Mobile
- Axios HTTP客户端
- React Router路由

### 4.2 开发环境
1. Node.js 14+
2. npm/yarn包管理器
3. VS Code编辑器
4. Chrome DevTools调试

### 4.3 组件开发
1. 遵循组件化开发思想
2. 使用Redux/Vuex状态管理
3. 实现响应式设计
4. 添加国际化支持

## 5. 测试

### 5.1 后端测试
1. 单元测试：使用Go testing包
2. 集成测试：测试API接口
3. 数据库测试：使用测试数据库
4. 性能测试：压力测试和负载测试

### 5.2 OCR服务测试
1. 单元测试：测试OCR和AI模块
2. 集成测试：测试Flask接口
3. 图像测试：使用多种账单图片测试
4. 性能测试：处理速度和准确率测试

### 5.3 前端测试
1. 单元测试：使用Jest/React Testing Library
2. E2E测试：使用Cypress/TestCafe
3. UI测试：视觉回归测试
4. 兼容性测试：多浏览器测试

## 6. 部署和运维

### 6.1 CI/CD流程
1. 使用GitHub Actions/GitLab CI
2. 自动化测试和构建
3. 自动化部署到测试环境
4. 手动审批部署到生产环境

### 6.2 监控和日志
1. 应用性能监控（APM）
2. 错误日志收集和分析
3. 数据库性能监控
4. 用户行为分析

### 6.3 安全措施
1. 定期安全扫描
2. 依赖包安全检查
3. 数据备份和恢复
4. 访问控制和权限管理

## 7. 贡献指南

### 7.1 提交代码
1. Fork项目到个人仓库
2. 创建功能分支
3. 编写代码和测试
4. 提交Pull Request

### 7.2 代码审查
1. 遵循代码规范
2. 添加必要的测试
3. 编写清晰的提交信息
4. 响应审查意见

### 7.3 问题报告
1. 使用问题模板
2. 提供详细的复现步骤
3. 包含环境信息
4. 添加相关日志和截图