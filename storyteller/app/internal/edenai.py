import json
import requests
import os

HEADERS = {"Authorization": "Bearer " + os.getenv("EDENAI_TOKEN")} 
URL = "https://api.edenai.run/v2/text/generation"
PROVIDER = "openai"

class EdenAI:
    def __init__(self, text, lines):
        self.text = text
        self.lines = lines
    def run(self):
        payload = {
            "providers": PROVIDER,
            "text": "Can you please write a " + str(self.lines) +" line story regarding " + self.text + "?",
            "temperature": 0.2
        }
        response = requests.post(URL, json=payload, headers=HEADERS)
        return json.loads(response.text)[PROVIDER]['generated_text']