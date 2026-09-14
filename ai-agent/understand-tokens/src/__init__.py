import json
import os
import urllib.request


def main() -> None:
    api_key = os.getenv("DEEPSEEK_API_KEY")
    payload = {
        "model": "deepseek-chat",
        "messages": [{"role": "user", "content": "Hello, world!"}],
    }
    request = urllib.request.Request(
        "https://api.deepseek.com/v1/chat/completions",
        data=json.dumps(payload).encode(),
        headers={
            "Authorization": f"Bearer {api_key}",
            "Content-Type": "application/json",
        },
    )

    with urllib.request.urlopen(request) as response:
        result = json.load(response)

    text = result["choices"][0]["message"]["content"]
    usage = result["usage"]
    print(json.dumps(
        {
            "model": "deepseek-chat",
            "text": text,
            "input_tokens": usage["prompt_tokens"],
            "output_tokens": usage["completion_tokens"],
            "total_tokens": usage["total_tokens"],
            "usage": usage,
        },
        ensure_ascii=False,
    ))


if __name__ == "__main__":
    main()
