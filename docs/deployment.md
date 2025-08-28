# 部署指南

## 1. 环境要求

### 1.1 系统要求
- 操作系统：Windows/Linux/macOS
- 内存：至少8GB（推荐16GB或以上）
- 硬盘空间：至少10GB可用空间

### 1.2 软件依赖
- Go 1.25或更高版本
- Python 3.8或更高版本
- MySQL 5.7或更高版本
- Node.js 14或更高版本（前端）
- Docker（可选，用于容器化部署）

## 2. 后端服务部署 (Go/Gin)

### 2.1 数据库配置
1. 安装并启动MySQL服务
2. 创建数据库：
   ```sql
   CREATE DATABASE IF NOT EXISTS pocketflow CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```
3. 执行SQL脚本创建表结构：
   ```bash
   mysql -u root -p pocketflow < backend-go/SQL.sql
   ```

### 2.2 配置文件修改
编辑 `backend-go/config/config.yaml` 文件：
```yaml
server:
  port: "8080"
  mode: "release"

database:
  host: "localhost"
  port: "3306"
  user: "your_username"
  password: "your_password"
  name: "pocketflow"
  charset: "utf8mb4"
  parseTime: true
  loc: "Local"

jwt:
  secret: "your_jwt_secret_key"
  expire: 86400

ocr:
  service_url: "http://localhost:5000/ocr/process"
```

### 2.3 编译和运行
```bash
# 进入后端目录
cd backend-go

# 下载Go依赖
go mod tidy

# 编译项目
go build -o backend-go.exe ./cmd/api

# 运行项目
./backend-go.exe
```

## 3. OCR服务部署 (Python/Flask)

### 3.1 环境准备
1. 安装Python 3.8或更高版本
2. 安装Ollama服务：
   - 访问 https://ollama.com/download 下载对应系统的安装包
   - 安装并启动Ollama服务
   - 下载Qwen2 1.5B模型：
     ```bash
     ollama pull qwen2:1.5b
     ```

### 3.2 安装Python依赖
```bash
# 进入OCR目录
cd ocr-ai

# 安装依赖（使用国内镜像源）
pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
```

### 3.3 配置环境变量
1. 复制配置文件模板：
   ```bash
   cp .env.example .env
   ```
2. 编辑 `.env` 文件，填入阿里云百炼API密钥：
   ```env
   # 阿里云百炼API配置
   DASHSCOPE_API_KEY=your_dashscope_api_key_here
   ```

### 3.4 启动服务
```bash
# 启动Flask应用
python app.py
```

## 4. 前端部署

### 4.1 环境准备
1. 安装Node.js 14或更高版本
2. 安装前端依赖：
   ```bash
   npm install
   ```

### 4.2 构建和运行
```bash
# 开发模式运行
npm run dev

# 生产模式构建
npm run build

# 生产模式运行
npm run start
```

## 5. Docker部署（可选）

### 5.1 构建Docker镜像
```bash
# 构建后端服务镜像
docker build -t pocketflow-backend -f docker/backend.Dockerfile .

# 构建OCR服务镜像
docker build -t pocketflow-ocr -f docker/ocr.Dockerfile .
```

### 5.2 运行容器
```bash
# 运行后端服务
docker run -d -p 8080:8080 --name pocketflow-backend pocketflow-backend

# 运行OCR服务
docker run -d -p 5000:5000 --name pocketflow-ocr pocketflow-ocr
```

## 6. Nginx反向代理配置（生产环境）

```nginx
server {
    listen 80;
    server_name your_domain.com;

    location /api/ {
        proxy_pass http://localhost:8080/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location /ocr/ {
        proxy_pass http://localhost:5000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location / {
        proxy_pass http://localhost:3000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 7. 系统服务配置（Linux）

### 7.1 后端服务
创建 systemd 服务文件 `/etc/systemd/system/pocketflow-backend.service`：
```ini
[Unit]
Description=PocketFlow Backend Service
After=network.target

[Service]
Type=simple
User=your_user
WorkingDirectory=/path/to/pocketflow/backend-go
ExecStart=/path/to/pocketflow/backend-go/backend-go.exe
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

### 7.2 OCR服务
创建 systemd 服务文件 `/etc/systemd/system/pocketflow-ocr.service`：
```ini
[Unit]
Description=PocketFlow OCR Service
After=network.target

[Service]
Type=simple
User=your_user
WorkingDirectory=/path/to/pocketflow/ocr-ai
ExecStart=/usr/bin/python3 /path/to/pocketflow/ocr-ai/app.py
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

启动服务：
```bash
sudo systemctl daemon-reload
sudo systemctl enable pocketflow-backend
sudo systemctl enable pocketflow-ocr
sudo systemctl start pocketflow-backend
sudo systemctl start pocketflow-ocr
```