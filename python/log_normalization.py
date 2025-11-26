from math import log

def log_normalization(array):
    maximum = max(array)
    log_max = log(maximum)
    
    return [round(log(X) / log_max, 4) if X != 0 else float(0) for X in array]
    
print(log_normalization([0, 20, 30, 40, 50]))