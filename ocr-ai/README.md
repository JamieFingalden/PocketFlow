# 记账应用OCR+AI分类模块

## 项目概述
这是一个完整的记账应用模块，包含OCR文字识别和AI智能分类功能。用户上传账单图片后，系统自动识别文字内容并进行消费分类。

## 文件结构
```
test/
├── app.py              # Flask应用入口
├── ocr_module.py       # OCR识别模块
├── ai_module_updated.py # AI分类模块（支持用户自定义分类）
├── requirements.txt    # 依赖管理文件
├── startup_guide.md    # 启动和测试指南
└── development_summary.md # 开发总结文档
```

## 使用方法

### 1. 复制必要文件
将以下文件复制到您的项目目录：
- `app.py`
- `ocr_module.py`
- `ai_module_updated.py` (重命名为 `ai_module.py`)
- `requirements.txt`

### 2. 安装依赖
```bash
pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
```

### 3. 配置环境变量
复制 `.env.example` 文件并重命名为 `.env`，然后在其中设置您的阿里云百炼API密钥：
```bash
cp .env.example .env
```

编辑 `.env` 文件，将 `your_dashscope_api_key_here` 替换为您的实际API密钥。

### 4. 安装Ollama服务
1. 访问 https://ollama.com/download 下载对应系统的安装包
2. 安装并启动Ollama服务
3. 下载Qwen2 1.5B模型：
   ```bash
   ollama pull qwen2:1.5b
   ```

### 5. 启动应用
```bash
python app.py
```

### 6. 调用接口
```bash
curl -X POST -F "image=@bill.jpg" http://localhost:5000/process_image
```

## 功能特点

### OCR识别
- 使用PaddleOCR进行中英文混合文字识别
- 自动处理图片预处理和文本后处理

### AI分类
- 使用本地部署的Qwen2 1.5B模型
- 支持默认分类和用户自定义分类
- 针对16GB内存设备优化

### API接口
- `/process_image`: 处理上传的账单图片
- `/health`: 健康检查接口

## 自定义分类配置

### 数据库表结构示例
```sql
CREATE TABLE user_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    category_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 修改数据库连接
在`ai_module.py`中修改`get_user_categories`函数以适配您的数据库系统。

## 性能优化

### 内存优化
- 使用Qwen2 1.5B小模型适配16GB内存设备
- 模型在服务启动时一次性加载
- OCR和AI处理结果缓存

### 响应速度优化
- 限制模型生成长度
- 降低模型温度参数提高确定性
- 简化提示词减少处理时间

## 部署建议

### 生产环境部署
1. 使用WSGI服务器如Gunicorn
2. 配置Nginx反向代理
3. 设置系统服务自启动

### Docker部署
可以创建Docker镜像简化部署：
```dockerfile
FROM python:3.11-slim
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
CMD ["python", "app.py"]
```

## 注意事项

1. 确保Ollama服务在应用启动前运行
2. 模型首次运行时需要下载，可能需要较长时间
3. 根据实际硬件配置调整模型参数
4. 定期更新依赖包以获取安全修复

## 故障排除

### 常见问题
1. **依赖安装失败**: 使用国内镜像源
2. **模型下载失败**: 检查网络连接或使用代理
3. **分类结果不准确**: 调整提示词或增加训练数据
4. **内存不足**: 使用更小的模型版本

### 日志查看
查看应用输出日志以诊断问题：
```bash
python app.py > app.log 2>&1
```