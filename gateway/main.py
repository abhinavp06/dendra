import asyncio
import uuid

from fastapi import FastAPi

app = FastAPI(lifespan=lifespan)
@app.post()
