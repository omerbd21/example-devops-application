import pika
from pydantic import BaseModel

class RabbitMQConnection():
    def __init__(self, host: str, username: str, password: str):
        credentials = pika.PlainCredentials(username=username, password=password)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(
                host=host,
                credentials=credentials
            )
        )
        self.channel = self.connection.channel()
    
    def publish_message(self, routing_key: str, message: str):
        try:
            self.channel.basic_publish(exchange='', routing_key=routing_key, body=message)
            return {"status": "Message published successfully"}

        except Exception as e:
            raise Exception(f"Error publishing message: {str(e)}")