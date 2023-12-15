from fastapi import FastAPI
from contextlib import asynccontextmanager
import uvicorn
import os
from internal import RabbitMQConnection
from routers import router

@asynccontextmanager
async def lifespan(app: FastAPI):
    app.rabbitmq_connection = RabbitMQConnection(host=os.getenv("RABBIT_HOST"), username=os.getenv('username'), password=os.getenv('password'))
    
    yield
    
    app.rabbitmq_connection.connection.close()

app = FastAPI(lifespan=lifespan)
app.include_router(router)