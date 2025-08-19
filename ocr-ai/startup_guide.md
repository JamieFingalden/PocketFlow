# 启动和测试指南

## 1. 安装和启动Ollama服务

### Windows系统安装步骤：
1. 访问 https://ollama.com/download/OllamaSetup.exe 下载Windows安装包
2. 运行安装程序，完成Ollama安装
3. 启动Ollama服务（通常会自动启动）

### 下载Qwen2 1.5B模型：
打开命令行终端，执行以下命令：
```bash
ollama pull qwen2:1.5b
```

### 验证Ollama服务：
```bash
ollama list
```

## 2. 准备测试图片

创建一个简单的测试图片，或者使用现有账单截图。图片应该包含一些消费信息，例如：
- "星巴克咖啡 35元"
- "出租车费 25元"
- "超市购物 120元"

## 3. 更新项目文件

将ai_module_ollama.py重命名为ai_module.py以替换原文件：
```bash
mv ai_module_ollama.py ai_module.py
```

## 4. 启动Flask应用

```bash
python app.py
```

## 5. 测试应用

使用curl命令测试：
```bash
curl -X POST -F "image=@your_test_image.jpg" http://localhost:5000/process_image
```

或者使用Python测试脚本：
```bash
python test_app.py
```

## 6. 验证结果

正常情况下应该返回类似以下的JSON响应：
```json
{
  "status": "success",
  "ocr_text": "星巴克咖啡 价格: 35元",
  "category": "餐饮"
}
```

## 故障排除

1. 如果Ollama服务无法启动，检查Windows防火墙设置
2. 如果模型下载失败，检查网络连接或使用代理
3. 如果Flask应用启动失败，检查端口是否被占用
4. 如果分类结果不准确，可以调整提示词或模型参数