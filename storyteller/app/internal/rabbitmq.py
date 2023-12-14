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

        except pika.exceptions.AMQPConnectionError as e:
            return {"status": "Can't connect to AMQP server", "error": e}
        except (pika.exceptions.ChannelClosedByServer, pika.exceptions.ChannelWrongStateError) as e:
            return {"status": "Can't publish to channel", "error": e}
        except pika.exceptions.UnroutableError as e:
            return {"status": "Can't direct message to queue", "error": e}
        except pika.exceptions.NackError as e:
            return {"status": "Server can't acknowledge message", "error": e}