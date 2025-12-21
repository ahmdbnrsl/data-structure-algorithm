import numpy as np

A = np.array([
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12],
    [13, 14, 15, 16]
])

v_filter = np.array([
    [2, 2],
    [2, 2]
])

oh = A.shape[0] - v_filter.shape[0] + 1
ow = A.shape[1]- v_filter.shape[1] + 1

result = np.zeros((oh, ow))

for i in range(oh):
    for j in range(ow):
        vertical_start = i
        vertical_end = i + v_filter.shape[0]
        horizontal_start = j
        horizontal_end = j + v_filter.shape[1]
        inputs = A[
            vertical_start:vertical_end,
            horizontal_start: horizontal_end
        ]
        result[i, j] = np.sum(inputs * v_filter)
        print(inputs)
        print(result)

print(result)