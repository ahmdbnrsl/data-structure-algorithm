def identity(o, c):
    if c <= 1: return o
    
    temp = 0
    left = 0
    for i in range(2, c + 1):
        if i == 2:
            temp = o * o - 2
            left = o
            continue
        temp_left = temp
        temp = temp * o - left
        left = temp_left
    
    return temp

print(identity(5, 2))