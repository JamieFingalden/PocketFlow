# -*- coding: utf-8 -*-
"""
AI模型调用模块：使用本地Qwen模型进行消费分类
"""

import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
import requests
import json
import os

# 全局变量用于存储模型和分词器
model = None
tokenizer = None

# 定义支持的分类标签
CATEGORIES = ["餐饮", "交通", "购物", "医疗", "其他"]

def initialize_model(model_path="Qwen/Qwen-7B-Chat"):
    """
    初始化AI模型和分词器，在服务启动时调用一次
    
    Args:
        model_path (str): 模型路径，可以是HuggingFace模型名或本地路径
    """
    global model, tokenizer
    
    print("正在加载AI模型，请稍候...")
    
    # 加载分词器
    tokenizer = AutoTokenizer.from_pretrained(
        model_path,
        trust_remote_code=True
    )
    
    # 加载模型（使用float16以节省内存）
    model = AutoModelForCausalLM.from_pretrained(
        model_path,
        torch_dtype=torch.float16,
        device_map="auto",
        trust_remote_code=True
    )
    
    # 设置为评估模式
    model.eval()
    
    print("AI模型加载完成")

def search_web(query):
    """
    简单的联网搜索功能（可选实现）
    这里使用一个简单的API示例，实际应用中可能需要替换为更稳定的搜索服务
    
    Args:
        query (str): 搜索查询
        
    Returns:
        str: 搜索结果摘要
    """
    try:
        # 使用一个简单的搜索API（示例，可能需要替换）
        search_url = f"https://api.peacefulapi.com/search?q={query}"
        response = requests.get(search_url, timeout=5)
        if response.status_code == 200:
            data = response.json()
            # 提取前几个结果
            results = data.get("results", [])[:3]
            return "\n".join([f"- {r.get('title', '')}: {r.get('snippet', '')}" for r in results])
        else:
            return "搜索失败"
    except Exception as e:
        return f"搜索出错: {str(e)}"

def classify_expense(ocr_text, enable_search=False):
    """
    使用AI模型对OCR文本进行消费分类
    
    Args:
        ocr_text (str): OCR识别出的文本
        enable_search (bool): 是否启用联网搜索功能
        
    Returns:
        str: 分类结果
    """
    global model, tokenizer
    
    if model is None or tokenizer is None:
        raise RuntimeError("模型未初始化，请先调用initialize_model()")
    
    # 构造系统提示词
    system_prompt = f"""你是一个账单分类助手。
你将接收一段中文账单文字，任务是：
1. 判断账单属于哪一类（从以下类别中选择一个）：
["餐饮", "交通", "购物", "医疗", "其他"]
2. 提取消费金额（数字即可，可以有小数）。
3. 提取消费日期，输出完整的'YYYY-MM-DD'格式日期：
   - 如果日期是相对描述（如"今天"、"昨天"、"周一"），请根据当前日期推算具体日期。
   - 如果账单只提供"月-日"，则使用当前年份补全。
4. 严格按照指定的JSON格式返回，不要多余文字。

用户提供的文字是：{ocr_text}
请严格按照指定的JSON格式返回，不要多余文字。
输出示例：'美团外卖 11:30 23.50元'
返回格式：{{"category": "类别", "amount": 金额数字, "date": "YYYY-MM-DD"}}"""
    
    # 如果启用搜索功能，可以先进行搜索补充信息
    if enable_search:
        # 提取关键词进行搜索（简单实现）
        search_results = search_web(ocr_text[:50])  # 取前50个字符作为搜索词
        system_prompt += f"\n\n参考信息：\n{search_results}"
    
    # 使用分词器编码输入
    inputs = tokenizer(system_prompt, return_tensors="pt").to(model.device)
    
    # 生成响应（设置最大长度和停止条件以提高响应速度）
    with torch.no_grad():
        outputs = model.generate(
            **inputs,
            max_new_tokens=10,     # 限制生成长度以提高速度
            temperature=0.1,       # 降低随机性以获得更确定的结果
            top_p=0.9,
            do_sample=False,       # 禁用采样以获得确定性结果
            pad_token_id=tokenizer.pad_token_id,
            eos_token_id=tokenizer.eos_token_id
        )
    
    # 解码输出
    response = tokenizer.decode(outputs[0], skip_special_tokens=True)
    
    # 提取JSON格式的结果
    # 假设模型会返回JSON格式的字符串，我们尝试解析它
    try:
        # 尝试从响应中提取JSON部分
        import re
        json_match = re.search(r'\{.*\}', response)
        if json_match:
            result_json = json_match.group()
            result = json.loads(result_json)
            # 验证返回的分类是否在预定义的分类列表中
            if result.get("category") in CATEGORIES:
                return result
            else:
                # 如果分类不在预定义列表中，返回"其他"
                result["category"] = "其他"
                return result
        else:
            # 如果没有找到JSON格式，返回默认值
            return {"category": "其他", "amount": 0.0, "date": "1900-01-01"}
    except Exception as e:
        # 如果解析失败，返回默认值
        return {"category": "其他", "amount": 0.0, "date": "1900-01-01"}

# 测试代码（可选）
if __name__ == "__main__":
    # 初始化模型（在实际应用中，这应该在服务启动时调用一次）
    # initialize_model()
    
    # 示例分类
    test_text = "星巴克咖啡 价格: 35元"
    result = classify_expense(test_text)
    print(f"分类结果: {result}")