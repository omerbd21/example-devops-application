from fastapi import APIRouter
from internal import Story, RabbitMQConnection, EdenAI
import pika
import os

router = APIRouter()
rabbitmq_connection = RabbitMQConnection(host='localhost', username=os.getenv('username'), password=os.getenv('password'))

@router.post("/story/queue/", tags=["story"])
async def queue_line(story: Story):
    eden = EdenAI(story.word, story.lines)
    story = eden.run()
    rabbitmq_connection.channel.basic_publish(exchange='', routing_key='lines', body=story, properties=pika.BasicProperties(
                          headers={'Content-Type': 'text/plain;charset=UTF-8'}),)
    return story


