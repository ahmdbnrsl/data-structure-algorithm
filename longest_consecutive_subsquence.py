def longest_consecutive_subsequence(array: list) -> int:
    temp = []
    for i in array:
        if i + 1 in array:
            temp += [i, i+1]
        elif i - 1 in array:
            temp += [i, i-1]
    
    unique = list(set(temp))

    result = [1]
    for i, e in enumerate(unique):
        if i != len(unique) - 1:
            if e + 1 != unique[i + 1]:
                result.append(1)
                continue
            result[len(result) - 1] += 1
    
    return max(result)

array_1 = [100, 4, 200, 1, 3, 2]          # Expected output: 4
array_2 = [10, 5, 12, 3, 55, 30, 4, 11, 13]  # Expected output: 4
array_3 = [1, 9, 3, 10, 4, 20, 2]         # Expected output: 4
array_4 = [8, 20, 7, 30, 6, 4, 5, 3]      # Expected output: 6
array_5 = [5]                             # Expected output: 1
array_6 = [45, 50, 46, 48, 90, 47, 49, 100, 10, 11, 12, 13]  # Expected output: 6
print(longest_consecutive_subsequence(array_1))
print(longest_consecutive_subsequence(array_2))
print(longest_consecutive_subsequence(array_3))
print(longest_consecutive_subsequence(array_4))
print(longest_consecutive_subsequence(array_5))
print(longest_consecutive_subsequence(array_6))