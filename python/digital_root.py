def digital_root(number: int) -> int:
    temp = 0
    iterable = [int(i) for i in str(number)]
    for i in iterable: temp += i
    
    length = len(list(str(temp)))
    
    if length > 1:
        return digital_root(temp)
    else: return temp
    
print(digital_root(9875))
print(digital_root(999999999999))