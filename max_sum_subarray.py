def max_subarray_sum(arr):
    max_sum = current_sum = arr[0]
    for num in arr[1:]:
        current_sum = max(num, current_sum + num)
        max_sum = max(max_sum, current_sum)
        print(current_sum, max_sum)
    return max_sum
    
max_subarray_sum([1, 2, 3, 4, -5, 9, -7])