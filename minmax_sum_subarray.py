def max_subarray_sum(arr):
    max_sum = current_sum = arr[0]
    for num in arr[1:]:
        current_sum = max(num, current_sum + num)
        max_sum = max(max_sum, current_sum)
        print(current_sum, max_sum)
    return max_sum
def min_subarray_sum(arr):
    min_sum = current_sum = arr[0]
    for num in arr[1:]:
        current_sum = min(num, current_sum + num)
        min_sum = min(min_sum, current_sum)
        print(current_sum, min_sum)
    return min_sum

# max_subarray_sum([1, 2, 3, 4, -5, 9, -7])
min_subarray_sum([1, 2, 3, 4, -5, 9, -7, 1, -10])