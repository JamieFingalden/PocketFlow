from openai import OpenAI
from datetime import date
import json

today_str = date.today().strftime("%Y-%m-%d")

client = OpenAI(
    # 若没有配置环境变量，请用百炼API Key将下行替换为：api_key="sk-xxx",
    api_key="sk-4033889abffa4e0d88b1c98d783820ea",
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
)

completion = client.chat.completions.create(
    # 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
    model="qwen-plus",
    messages=[
        {"role": "system", "content": "你是一个账单分类助手。"
                                      "你将接收一段中文账单文字，任务是："
                                      "1. 判断账单属于哪一类（从以下类别中选择一个）：" +
                                      '["餐饮", "交通", "购物", "医疗", "其他"]' +
                                      "2. 提取消费金额（数字即可，可以有小数）。"
                                      "3. 提取消费日期，输出完整的'YYYY-MM-DD'格式日期："
                                      "   - 如果日期是相对描述（如“今天”、“昨天”、“周一”），请根据当前日期推算具体日期。"
                                      "   - 如果账单只提供“月-日”，则使用当前年份补全。"
                                      "4. 严格按照指定的JSON格式返回，不要多余文字。"
                                      ""
                                      "输出示例：'美团外卖 11:30 23.50元'"
                                      "返回格式：{'category': '类别', 'amount': 金额数字, 'date': 'YYYY-MM-DD'}"},
        {"role": "system", "content": f"这个是{today_str}当前时间"},
        {"role": "user", "content": "京东"
                                    "18元"
                                    "周一"},
    ],
    # Qwen3模型通过enable_thinking参数控制思考过程（开源版默认True，商业版默认False）
    # 使用Qwen3开源版模型时，若未启用流式输出，请将下行取消注释，否则会报错
    # extra_body={"enable_thinking": False},
)
data = json.loads(completion.model_dump_json())
print(data["choices"][0]["message"]["content"])