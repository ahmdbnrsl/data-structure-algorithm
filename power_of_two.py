import math as m

def min_remove_el_power_of_two(number: str) -> int:
    for i in range(60, 0, -1):
        if len(str(2**i)) > len(number): continue
        
        target = str(2**i)
        nothing = False
        target_in_number = ""
        
        for i, e in enumerate(target):
            if e not in number:
                nothing = True
                break
        
        if nothing: continue
    
        for i, e in enumerate(number):
            if e in target: target_in_number += e
        
        if target not in target_in_number: continue
        
        return len(number) - len(str(target)), (2, m.log(int(target)) / m.log(2))

minimum_digit, power_of_two = min_remove_el_power_of_two("123456")
print(minimum_digit, power_of_two)

