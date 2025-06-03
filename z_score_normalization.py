from math import sqrt

def z_score(array):
    N = len(array)
    mean = sum(array) / N
    standard_deviation = sqrt(sum([(i - mean) ** 2 for i in array]) / N)
    
    return [round((X - mean) / standard_deviation, 4) for X in array]

print(z_score([10, 20, 30, 40, 50]))