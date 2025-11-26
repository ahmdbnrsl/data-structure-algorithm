def get_smallest_value(arr):
    if not isinstance(arr, list):
        raise ValueError("Parameter must be a list")
    result = 0
    for i, n in enumerate(arr):
        if not isinstance(n, (int, float)):
            raise ValueError("List must containing a integer and float")
        if i == 0:
            result = n
        elif i > 0:
            if result > n:
                result = n
    print(result)
    return result
get_smallest_value([15,3,4,8,3, 2, 1, 12, 1, 60])

class MinMax:
    def __init__(self, arr):
        if not isinstance(arr, list):
            raise ValueError("Parameter must be a list")
        for i, n in enumerate(arr):
            if not isinstance(n, (int, float)):
                raise ValueError("List must containing a integer and float")
        self.arr = arr
        self.minResult = 0
        self.maxResult = 0
    def min(self):
        for i, n in enumerate(self.arr):
            if i == 0:
                self.minResult = n
            elif i > 0:
                if self.minResult > n:
                    self.minResult = n
        return self.minResult
    def max(self):
        for i, n in enumerate(self.arr):
            if i == 0:
                self.maxResult = n
            elif i > 0:
                if self.maxResult < n:
                    self.maxResult = n
        return self.maxResult

arr = [15,3,4,8,3, 2, 1, 12, 1, 60, 8, 4]
minmax = MinMax(arr)
print(minmax.min())
print(minmax.max())