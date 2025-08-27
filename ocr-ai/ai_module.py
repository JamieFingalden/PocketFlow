# -*- coding: utf-8 -*-
"""
AI模型调用模块：使用本地Ollama服务进行消费分类，备用方案为阿里云百炼API
"""

import requests
import json
import os
from datetime import date
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()

# 默认分类标签
DEFAULT_CATEGORIES = ["餐饮", "交通", "购物", "医疗", "其他"]

# Ollama服务地址
OLLAMA_URL = "http://localhost:11434/api/generate"

# 阿里云百炼API配置
DASHSCOPE_API_KEY = os.getenv('DASHSCOPE_API_KEY', 'your_api_key_here')
DASHSCOPE_URL = "https://dashscope.aliyuncs.com/compatible-mode/v1"

def initialize_model():
    """
    初始化AI模型（Ollama版本）
    实际上不需要初始化，只需要确保Ollama服务正在运行
    """
    print("AI模块初始化完成（Ollama版本）")
    return True

def get_user_categories(user_id=None):
    """
    从数据库获取用户自定义分类列表
    这里是模拟实现，实际应用中需要连接真实数据库
    
    Args:
        user_id (int): 用户ID
        
    Returns:
        list: 用户分类列表
    """
    # 模拟数据库查询
    # 实际应用中这里应该查询数据库获取用户自定义分类
    if user_id is not None:
        # 模拟从数据库获取用户自定义分类
        # 这里应该使用真实的数据库查询
        try:
            # 示例：从SQLite数据库查询用户分类
            # import sqlite3
            # conn = sqlite3.connect('expenses.db')
            # cursor = conn.cursor()
            # cursor.execute("SELECT category_name FROM user_categories WHERE user_id = ?", (user_id,))
            # user_categories = [row[0] for row in cursor.fetchall()]
            # conn.close()
            # return user_categories
            
            # 临时返回模拟数据
            return ["餐饮", "交通", "购物", "医疗", "娱乐", "教育"]
        except Exception as e:
            print(f"获取用户分类失败: {str(e)}")
            return DEFAULT_CATEGORIES
    else:
        return DEFAULT_CATEGORIES

def classify_with_ollama(ocr_text, user_categories):
    """
    使用Ollama本地部署的Qwen模型对OCR文本进行消费分类
    
    Args:
        ocr_text (str): OCR识别出的文本
        user_categories (list): 用户分类列表
        
    Returns:
        dict: 分类结果
    """
    # 构造分类选项字符串
    categories_str = ",".join(user_categories)
    
    # 构造系统提示词（简化版以节省内存）
    prompt = f"""你是一个账单分类助手。
你将接收一段中文账单文字，任务是：
1. 判断账单属于哪一类（从以下类别中选择一个）：
[{categories_str}]
2. 提取消费金额（数字即可，可以有小数）。
3. 提取消费日期，输出完整的'YYYY-MM-DD'格式日期：
   - 如果日期是相对描述（如"今天"、"昨天"、"周一"），请根据当前日期推算具体日期。
   - 如果账单只提供"月-日"，则使用当前年份补全。
4. 严格按照指定的JSON格式返回，不要多余文字。

输出示例：'美团外卖 11:30 23.50元'
返回格式：{{"category": "类别", "amount": 金额数字, "date": "YYYY-MM-DD"}}

用户提供的文字是：{ocr_text}
请严格按照指定的JSON格式返回，不要多余文字。"""
    
    # 构造请求数据（优化参数以适应16GB内存）
    data = {
        "model": "qwen2:1.5b",  # 使用更小的1.5B版本模型以适应16GB内存
        "prompt": prompt,
        "stream": False,
        "options": {
            "temperature": 0.1,     # 降低随机性以获得更确定的结果
            "top_p": 0.9,
            "max_tokens": 50,       # 增加一点长度以适应更多分类
            "num_ctx": 2048         # 减少上下文长度以节省内存
        }
    }
    
    # 发送请求到Ollama API
    response = requests.post(OLLAMA_URL, json=data, timeout=30)
    
    if response.status_code == 200:
        result = response.json()
        response_text = result.get("response", "").strip()
        
        # 尝试解析JSON响应
        try:
            # 尝试从响应中提取JSON部分
            import re
            json_match = re.search(r'\{.*\}', response_text)
            if json_match:
                result_json = json_match.group()
                result = json.loads(result_json)
                return result
            else:
                raise ValueError("无法从响应中提取JSON")
        except Exception as e:
            raise RuntimeError(f"解析Ollama响应失败: {str(e)}")
    else:
        raise RuntimeError(f"Ollama API调用失败: {response.status_code}")

def classify_with_dashscope(ocr_text, user_categories):
    """
    使用阿里云百炼API对OCR文本进行消费分类
    
    Args:
        ocr_text (str): OCR识别出的文本
        user_categories (list): 用户分类列表
        
    Returns:
        dict: 分类结果
    """
    # 构造分类选项字符串
    categories_str = ",".join(user_categories)
    
    today_str = date.today().strftime("%Y-%m-%d")
    
    # 构造系统提示词
    system_prompt = f"""你是一个账单分类助手。
你将接收一段中文账单文字，任务是：
1. 判断账单属于哪一类（从以下类别中选择一个）：
[{categories_str}]
2. 提取消费金额（数字即可，可以有小数）。
3. 提取消费日期，输出完整的'YYYY-MM-DD'格式日期：
   - 如果日期是相对描述（如"今天"、"昨天"、"周一"），请根据当前日期推算具体日期。
   - 如果账单只提供"月-日"，则使用当前年份补全。
4. 严格按照指定的JSON格式返回，不要多余文字。

输出示例：'美团外卖 11:30 23.50元'
返回格式：{{"category": "类别", "amount": 金额数字, "date": "YYYY-MM-DD"}}"""
    
    # 构造请求数据
    data = {
        "model": "qwen-plus",
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "system", "content": f"这个是{today_str}当前时间"},
            {"role": "user", "content": ocr_text},
        ]
    }
    
    # 设置请求头
    headers = {
        "Authorization": f"Bearer {DASHSCOPE_API_KEY}",
        "Content-Type": "application/json"
    }
    
    # 发送请求到阿里云百炼API
    response = requests.post(DASHSCOPE_URL + "/chat/completions", 
                           json=data, headers=headers, timeout=30)
    
    if response.status_code == 200:
        result = response.json()
        response_text = result["choices"][0]["message"]["content"].strip()
        
        # 尝试解析JSON响应
        try:
            # 尝试从响应中提取JSON部分
            import re
            json_match = re.search(r'\{.*\}', response_text)
            if json_match:
                result_json = json_match.group()
                result = json.loads(result_json)
                return result
            else:
                raise ValueError("无法从响应中提取JSON")
        except Exception as e:
            raise RuntimeError(f"解析阿里云百炼API响应失败: {str(e)}")
    else:
        raise RuntimeError(f"阿里云百炼API调用失败: {response.status_code}")

def classify_expense(ocr_text, user_id=None):
    """
    使用AI模型对OCR文本进行消费分类
    优先使用本地Ollama服务，如果不可用则使用阿里云百炼API
    
    Args:
        ocr_text (str): OCR识别出的文本
        user_id (int): 用户ID，用于获取用户自定义分类
        
    Returns:
        dict: 分类结果
    """
    # 获取用户分类列表（包含默认分类）
    user_categories = get_user_categories(user_id)
    
    # 首先尝试使用本地Ollama服务
    try:
        print("正在尝试使用本地Ollama服务...")
        result = classify_with_ollama(ocr_text, user_categories)
        print("本地Ollama服务调用成功")
        return result
    except Exception as e:
        print(f"本地Ollama服务调用失败: {str(e)}")
        print("正在尝试使用阿里云百炼API...")
        
        # 如果Ollama服务不可用，尝试使用阿里云百炼API
        try:
            result = classify_with_dashscope(ocr_text, user_categories)
            print("阿里云百炼API调用成功")
            return result
        except Exception as e2:
            print(f"阿里云百炼API调用失败: {str(e2)}")
            # 如果两种方式都失败，返回默认值
            return {"category": "其他", "amount": 0.0, "date": "1900-01-01"}

# 测试代码（可选）
if __name__ == "__main__":
    # 初始化模型（在实际应用中，这应该在服务启动时调用一次）
    # initialize_model()
    
    # 示例分类
    test_text = "星巴克咖啡 价格: 35元"
    result = classify_expense(test_text)
    print(f"分类结果: {result}")