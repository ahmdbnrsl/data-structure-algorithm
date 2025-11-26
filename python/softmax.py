def softmax(array):
    from math import exp
    
    exponentials = [exp(i) for i in array]
    n = sum(exponentials)
    
    return [round(i / n, 4) for i in exponentials]

print(softmax([2.0, 1.0, 0.1]))