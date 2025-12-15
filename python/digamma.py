import mpmath as mp
mp.dbs = 100

def digamma_positive(x):
    if x < 1: raise ValueError("x must >= 1")
    if x == 1: return -mp.euler
    else: return digamma_positive(x - 1) + 1/(x - 1)

result = digamma_positive(4)
print(result)