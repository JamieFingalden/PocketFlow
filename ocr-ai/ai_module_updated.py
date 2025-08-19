# -*- coding: utf-8 -*-
"""
更新后的AI模型调用模块：支持用户自定义分类列表
"""

import requests
import json

# 默认分类标签
DEFAULT_CATEGORIES = ["餐饮", "交通", "购物", "医疗", "其他"]

# Ollama服务地址
OLLAMA_URL = "http://localhost:11434/api/generate"

def initialize_model():
    """
    初始化AI模型（Ollama版本）
    实际上不需要初始化，只需要确保Ollama服务正在运行
    """
    print("检查Ollama服务连接...")
    
    try:
        # 检查Ollama服务是否运行
        response = requests.get("http://localhost:11434/api/tags", timeout=5)
        if response.status_code == 200:
            print("Ollama服务连接成功")
            return True
        else:
            print("Ollama服务连接失败")
            return False
    except Exception as e:
        print(f"Ollama服务连接错误: {str(e)}")
        return False

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

def classify_expense(ocr_text, user_id=None):
    """
    使用Ollama本地部署的Qwen模型对OCR文本进行消费分类
    支持用户自定义分类列表
    
    Args:
        ocr_text (str): OCR识别出的文本
        user_id (int): 用户ID，用于获取用户自定义分类
        
    Returns:
        str: 分类结果
    """
    # 获取用户分类列表（包含默认分类）
    user_categories = get_user_categories(user_id)
    
    # 构造分类选项字符串
    categories_str = ",".join(user_categories)
    
    # 构造系统提示词（简化版以节省内存）
    prompt = f"""分类任务：从以下选项中选择一个最匹配的分类：
{categories_str}
输入文本：{ocr_text}
输出分类："""

    # 构造请求数据（优化参数以适应16GB内存）
    data = {
        "model": "qwen2:1.5b",  # 使用更小的1.5B版本模型以适应16GB内存
        "prompt": prompt,
        "stream": False,
        "options": {
            "temperature": 0.1,     # 降低随机性以获得更确定的结果
            "top_p": 0.9,
            "max_tokens": 15,       # 增加一点长度以适应更多分类
            "num_ctx": 2048         # 减少上下文长度以节省内存
        }
    }

    try:
        # 发送请求到Ollama API
        response = requests.post(OLLAMA_URL, json=data, timeout=30)
        
        if response.status_code == 200:
            result = response.json()
            response_text = result.get("response", "").strip()
            
            # 确保返回的分类在用户分类列表中
            # 使用更宽松的匹配方式
            response_text_lower = response_text.lower()
            for category in user_categories:
                if category in response_text or category.lower() in response_text_lower:
                    return category
            
            # 如果没有匹配的分类，返回用户分类列表中的第一个作为默认分类
            if user_categories:
                return user_categories[0]
            else:
                return "其他"
        else:
            raise RuntimeError(f"Ollama API调用失败: {response.status_code}")
            
    except Exception as e:
        raise RuntimeError(f"AI分类失败: {str(e)}")

# 测试代码（可选）
if __name__ == "__main__":
    # 示例分类
    test_text = "星巴克咖啡 价格: 35元"
    try:
        # 使用默认分类
        category = classify_expense(test_text)
        print(f"默认分类结果: {category}")
        
        # 使用用户自定义分类
        user_category = classify_expense(test_text, user_id=1)
        print(f"用户分类结果: {user_category}")
    except Exception as e:
        print(f"分类失败: {str(e)}")