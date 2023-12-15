from pydantic import BaseModel

class Story(BaseModel):
    word: str
    lines: int