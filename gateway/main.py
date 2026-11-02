import asyncio
import uuid
import uvicorn

from fastapi import FastAPI

app = FastAPI()


@app.post("/register/claim")
async def register_claim():
    print("Hi!")

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000, log_level="info", timeout_graceful_shutdown=5)
