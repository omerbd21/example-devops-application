import unittest
from app.internal import RabbitMQConnection
import os

class TestRabbitMQConnection(unittest.TestCase):

    def setUp(self) -> None:
        self.RabbitMQConnection = RabbitMQConnection(host='localhost', username=os.getenv('username'), password=os.getenv('password'))
    
    def test_publish_message(self):
        status  = self.RabbitMQConnection.publish_message("lines_test", "test")
        self.assertEqual(status, {"status": "Message published successfully"})


if __name__ == '__main__':
    unittest.main()