import json
import os
import urllib.request

from transformers import AutoTokenizer


TOKENIZER_REPO = "deepseek-ai/DeepSeek-V3.1"
MESSAGES = [{"role": "user", "content": "Hello, world!"}]


def count_local_tokens():
    tokenizer = AutoTokenizer.from_pretrained(TOKENIZER_REPO)
    prompt = tokenizer.apply_chat_template(
        MESSAGES,
        tokenize=False,
        add_generation_prompt=True,
    )
    tokens = tokenizer.encode(prompt)
    return len(tokens), tokenizer.convert_ids_to_tokens(tokens)


def main() -> None:
    local_prompt_tokens, local_tokens = count_local_tokens()

    api_key = os.getenv("DEEPSEEK_API_KEY")
    payload = {
        "model": "deepseek-chat",
        "messages": MESSAGES,
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
            "local_prompt_tokens": local_prompt_tokens,
            "local_tokens": local_tokens,
            "api_prompt_tokens": usage["prompt_tokens"],
            "match": local_prompt_tokens == usage["prompt_tokens"],
            "input_tokens": usage["prompt_tokens"],
            "output_tokens": usage["completion_tokens"],
            "total_tokens": usage["total_tokens"],
            "usage": usage,
        },
        ensure_ascii=False,
    ))


if __name__ == "__main__":
    main()
