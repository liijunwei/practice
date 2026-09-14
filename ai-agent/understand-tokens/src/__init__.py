from openai import OpenAI
import os
import json
from dotenv import load_dotenv

load_dotenv()

# 初始化DeepSeek客户端
client = OpenAI(
    api_key=os.getenv("DEEPSEEK_API_KEY"),
    base_url="https://api.deepseek.com/v1"
)

async def main():
    # 对应截图的 await generateText({model, prompt})
    response = client.chat.completions.create(
        model="deepseek-chat", # 可换成 deepseek-reasoner
        messages=[
            {"role": "user", "content": "Hello, world!"}
        ]
    )

    # 等价于 anthropicResult.text
    text = response.choices[0].message.content
    # 等价于 anthropicResult.usage
    usage = response.usage

    print(json.dumps({
        "model": "deepseek-chat",
        "text": text,
        "usage": usage.model_dump() if hasattr(usage, "model_dump") else dict(usage),
    }, ensure_ascii=False))

if __name__ == "__main__":
    import asyncio
    asyncio.run(main())
