def max_digit_mod(digit):
    digit_str = str(digit)
    
    current_digit = {}
    for i, e in enumerate(digit_str):
        if e not in current_digit: current_digit[e] = 1
        else: current_digit[e] += 1
    
    maximal = max(current_digit.values())
    candidates = [int(d) for d, c in current_digit.items() if c == maximal]
    return (max(candidates), maximal)

print(max_digit_mod(1199199191999))