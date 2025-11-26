def min_max_scaling(array):
    minimum = min(array)
    maximum = max(array)

    if minimum == maximum:
        return [0.0] * len(array)
    
    return [(x - minimum) / (maximum - minimum) for x in array]

print(min_max_scaling([10, 20, 30, 40, 50]))
print(min_max_scaling([10, 10, 10, 10, 10, 20]))