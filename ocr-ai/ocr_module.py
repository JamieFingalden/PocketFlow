# -*- coding: utf-8 -*-
"""
OCR模块：使用PaddleOCR进行中文文本识别
"""

from paddleocr import PaddleOCR
import os

# 初始化PaddleOCR引擎（支持中英文）
# 使用轻量级模型以适应Mac mini M4的内存限制
ocr_engine = PaddleOCR(
    use_angle_cls=True,      # 启用文字方向分类
    lang='ch',               # 设置语言为中英文混合
    det_model_dir=None,      # 使用默认检测模型
    rec_model_dir=None,      # 使用默认识别模型
    cls_model_dir=None,      # 使用默认分类模型
    show_log=False           # 关闭日志输出
)

def ocr_image(image_path):
    """
    对图片进行OCR识别，返回拼接后的文本
    
    Args:
        image_path (str): 图片文件路径
        
    Returns:
        str: 识别出的文本内容
    """
    # 检查文件是否存在
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片文件不存在: {image_path}")
    
    # 执行OCR识别
    result = ocr_engine.ocr(image_path, cls=True)
    
    # 提取文本内容并拼接
    texts = []
    for line in result:
        if line is not None:  # 检查是否有识别结果
            for item in line:
                if item is not None and len(item) > 1:
                    text_info = item[1]
                    if text_info is not None and len(text_info) > 0:
                        texts.append(text_info[0])
    
    # 返回拼接后的文本
    return '\n'.join(texts)

# 测试代码（可选）
if __name__ == "__main__":
    # 示例用法
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        ocr_text = ocr_image(test_image)
        print("OCR识别结果:")
        print(ocr_text)
    else:
        print(f"测试图片 {test_image} 不存在")