def unique_paths(m: int, n: int):
    right_step = (m-1)
    down_step = (n-1)
    total_step = right_step + down_step
    
    total_step_factorial = 1
    for i in range(1, total_step + 1):
        total_step_factorial *= i
    
    right_step_factorial = 1
    for i in range(1, right_step + 1):
        right_step_factorial *= i
        
    down_step_factorial = 1
    for i in range(1, down_step + 1):
        down_step_factorial *= i
    
    return total_step_factorial / (right_step_factorial * down_step_factorial)

print(unique_paths(5, 3))