def largest_unique_number(numbers: list) -> int :
    for j in range(len(numbers) - 1):
        for i in range(len(numbers) -1):
            if numbers[i] < numbers[i + 1]:
                temp = numbers[i]
                numbers[i] = numbers[i + 1]
                numbers[i + 1] = temp
                
    for index, el in enumerate(numbers):
        if index == 0 and numbers[index + 1] != el: return el
        elif index != len(numbers) - 1 and numbers[index - 1] > el > numbers[index + 1]:
            return el

print(largest_unique_number([10, 6, 5, 10, 7, 5, 6, 3, 2, 8, 8, 9, 3, 4]))