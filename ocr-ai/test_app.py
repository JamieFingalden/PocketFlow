# -*- coding: utf-8 -*-
"""
测试脚本：验证OCR和AI分类功能
"""

import requests
import os

def test_process_image():
    """
    测试图片处理接口
    """
    # 测试健康检查接口
    try:
        health_response = requests.get('http://localhost:5000/health')
        print("健康检查响应:", health_response.json())
    except Exception as e:
        print("健康检查失败:", str(e))
        print("请确保Flask应用正在运行")
        return
    
    # 准备测试图片（这里需要一个实际的测试图片文件）
    # 如果没有测试图片，可以创建一个简单的测试
    test_image_path = "test_bill.jpg"
    
    # 如果测试图片不存在，创建一个简单的测试
    if not os.path.exists(test_image_path):
        print(f"测试图片 {test_image_path} 不存在")
        print("请提供一个账单图片进行测试，或创建一个测试图片")
        return
    
    # 测试图片处理接口
    try:
        with open(test_image_path, 'rb') as f:
            files = {'image': f}
            response = requests.post('http://localhost:5000/process_image', files=files)
        
        print("图片处理响应:")
        print(response.json())
    except Exception as e:
        print("图片处理测试失败:", str(e))

if __name__ == "__main__":
    test_process_image()