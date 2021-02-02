import unittest
import calculatePrice
import numpy as np

class TestIT(unittest.TestCase):
    
    def test_numberDaysPerClass(self):
        self.assertEqual(calculatePrice.numberDaysPerClass(np.array([1,1,5, 10, 19, 20, 60, 101, 200])), {"standard":140, "ia":81, "glacier":60, "deep_archive":219})

if __name__ == "__main__":
    unittest.main()
