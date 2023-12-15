from fastapi import APIRouter, Request
import pika
import os
from internal import Story, EdenAI


router = APIRouter()

@router.post("/story/queue/", tags=["story"])
async def queue_line(story: Story, request: Request):
    eden = EdenAI(story.word, story.lines)
    generated_story = eden.run()
    request.app.rabbitmq_connection.channel.basic_publish(exchange='', routing_key=os.getenv("QUEUE"), body=story, properties=pika.BasicProperties(
                          headers={'Content-Type': 'text/plain;charset=UTF-8'}),)
    return generated_story


