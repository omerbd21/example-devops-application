import unittest
from app.internal import EdenAI

class TestEdenAI(unittest.TestCase):

    def setUp(self) -> None:
        self.edenAI = EdenAI("idan", 3)
    
    def test_run(self):
        story = self.edenAI.run()
        self.assertEqual(self.edenAI.lines, len(str.splitlines(story)))
        self.assertIn(self.edenAI.text.lower(), story.lower())


if __name__ == '__main__':
    unittest.main()