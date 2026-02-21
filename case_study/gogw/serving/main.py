import asyncio
import time
import uuid
from typing import List, Optional
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

# --- SCHEMAS ---
class Message(BaseModel):
    role: str
    content: str

class ChatRequest(BaseModel):
    model: str = "fake-gpt-4"
    messages: List[Message]
    temperature: Optional[float] = 0.7

# --- ENDPOINT ---
@app.post("/v1/chat/completions")
async def create_chat_completion(request: ChatRequest):
    # 1. Simulate "Inference" time (e.g., 3 seconds)
    # Using 'await asyncio.sleep' is critical so you don't block the server
    await asyncio.sleep(3)

    # 2. Get the last user message to make the "fake" reply relevant
    user_input = request.messages[-1].content
    fake_reply = f"This is a simulated response to: '{user_input}'"

    # 3. Return OpenAI-compatible format
    return {
        "id": f"chatcmpl-{uuid.uuid4()}",
        "object": "chat.completion",
        "created": int(time.time()),
        "model": request.model,
        "choices": [
            {
                "index": 0,
                "message": {
                    "role": "assistant",
                    "content": fake_reply,
                },
                "finish_reason": "stop",
            }
        ],
        "usage": {
            "prompt_tokens": len(user_input),
            "completion_tokens": len(fake_reply),
            "total_tokens": len(user_input) + len(fake_reply),
        }
    }