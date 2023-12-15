import json
import requests
import os

HEADERS = {"Authorization": "Bearer " + os.getenv("EDENAI_TOKEN")} 

class EdenAI():
    def __init__(self, text: str, lines: int):
        self.text = text
        self.lines = lines
    def run(self):
        payload = {
            "providers": os.getenv("EDENAI_PROVIDER"),
            "text": f"Can you please write a {self.lines} line story regarding {self.text}?",
            "temperature": 0.2
        }
        response = requests.post(os.getenv("EDENAI_URL"), json=payload, headers=HEADERS)
        return json.loads(response.text)[os.getenv("EDENAI_PROVIDER")]['generated_text']