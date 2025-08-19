# -*- coding: utf-8 -*-
"""
更新后的AI模型调用模块：使用Ollama本地部署的Qwen模型进行消费分类
针对16GB内存设备优化
"""

import requests
import json

# 定义支持的分类标签
CATEGORIES = ["餐饮", "交通", "购物", "医疗", "其他"]

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

def classify_expense(ocr_text):
    """
    使用Ollama本地部署的Qwen模型对OCR文本进行消费分类
    针对16GB内存设备优化
    
    Args:
        ocr_text (str): OCR识别出的文本
        
    Returns:
        str: 分类结果
    """
    # 构造系统提示词（详细版，与a.py保持一致）
    from datetime import date
    today_str = date.today().strftime("%Y-%m-%d")
    
    prompt = f"""账单分类助手任务：
                 1. 从["餐饮", "交通", "购物", "医疗", "其他"]中选择分类
                 2. 提取金额数字（可含小数，通常为￥xx.xx或$xx.xx格式）
                 3. 提取日期转为'YYYY-MM-DD'格式：
                    - "昨天"、"星期一"等相对日期需推算
                    - "8月4日"格式补全年份
                    - 提取"18:30"等具体时间
                    - 如果没有时间信息，日期字段保持为空字符串""
                 5. 特殊规则：
                    - 京东/淘宝/美团等平台按时间判断（11-13点、16-21点：餐饮；0-6点：购物）
                    - 金额极小（<5元）：可能是服务费，倾向于"其他"或"医疗"
                    - 特定商户+金额特征：如"眼镜店 ￥1"可能是清洗服务，归为"其他"
                 5. 返回JSON格式：{{"category": "类别", "amount": 金额, "date": "YYYY-MM-DD"}}
                    - 如果无法提取时间，date字段应为""
                 
                 当前时间：{today_str}
                 账单：{ocr_text}
                 务必放回JSON格式结果："""

    # 构造请求数据（优化参数以提高速度和准确性）
    data = {
        "model": "qwen2.5:3B",  # 使用更大的3B版本模型以提高准确性
        "prompt": prompt,
        "stream": False,
        "options": {
            "temperature": 0.1,     # 降低随机性以获得更确定的结果
            "top_p": 0.9,
            "max_tokens": 50,       # 限制生成长度以提高速度
            "num_ctx": 1024,        # 减少上下文长度以提高速度
            "repeat_penalty": 1.2,  # 防止重复
            "top_k": 40             # 限制词汇选择以提高速度
        }
    }

    try:
        # 发送请求到Ollama API
        response = requests.post(OLLAMA_URL, json=data, timeout=30)
        
        if response.status_code == 200:
            result = response.json()
            response_text = result.get("response", "").strip()
            
            # 尝试解析JSON格式的响应
            try:
                import re
                # 使用正则表达式提取JSON部分
                json_match = re.search(r'\{[^}]+\}', response_text)
                if json_match:
                    json_str = json_match.group()
                    # 修复转义问题
                    json_str = json_str.replace('"category": "类别"', '"category": "其他"')
                    json_str = json_str.replace('"amount": 金额数字', '"amount": 0.0')
                    json_str = json_str.replace('"date": "YYYY-MM-DD"', '"date": "1900-01-01"')
                    
                    parsed_result = json.loads(json_str)
                    
                    # 验证分类是否有效
                    if parsed_result.get("category") in CATEGORIES:
                        return parsed_result
                    else:
                        # 如果分类无效，返回默认值
                        parsed_result["category"] = "其他"
                        return parsed_result
                else:
                    # 如果没有找到JSON格式，返回默认值
                    return {"category": "其他", "amount": 0.0, "date": "1900-01-01"}
            except json.JSONDecodeError:
                # 如果JSON解析失败，尝试原来的分类方法
                # 确保返回的分类在预定义的分类列表中
                # 使用更宽松的匹配方式
                response_text_lower = response_text.lower()
                for category in CATEGORIES:
                    if category in response_text or category in response_text_lower:
                        return {"category": category, "amount": 0.0, "date": "1900-01-01"}
                
                # 如果没有匹配的分类，返回"其他"
                return {"category": "其他", "amount": 0.0, "date": "1900-01-01"}
        else:
            raise RuntimeError(f"Ollama API调用失败: {response.status_code}")
            
    except Exception as e:
        raise RuntimeError(f"AI分类失败: {str(e)}")

def classify_multiple_expenses(ocr_texts):
    """
    批量处理多个OCR文本进行消费分类
    
    Args:
        ocr_texts (list): OCR识别出的文本列表
        
    Returns:
        list: 分类结果列表，每个元素都是JSON格式的字典
    """
    results = []
    for ocr_text in ocr_texts:
        try:
            result = classify_expense(ocr_text)
            results.append(result)
        except Exception as e:
            # 如果某个账单处理失败，添加默认值
            results.append({"category": "其他", "amount": 0.0, "date": "1900-01-01", "error": str(e)})
    return results

# 测试代码（可选）
if __name__ == "__main__":
    # 示例分类
    test_text = "星巴克咖啡 价格: 35元"
    try:
        result = classify_expense(test_text)
        print(f"分类结果: {result}")
    except Exception as e:
        print(f"分类失败: {str(e)}")
    
    # 批量处理示例（包含不同格式的时间信息）
    test_texts = [
        "12:30 星巴克咖啡 ￥35.00",
        "昨天 18:45 地铁 ￥5.00",
        "20:15 京东购物 ￥128.50",
        "星期一 23:41 光明眼镜店 ￥1.00",
        "美团外卖 ￥28.80",  # 没有时间信息
        "$15.50 咖啡店"  # 没有时间信息，金额格式不同
    ]
    try:
        results = classify_multiple_expenses(test_texts)
        print(f"批量分类结果: {results}")
    except Exception as e:
        print(f"批量分类失败: {str(e)}")