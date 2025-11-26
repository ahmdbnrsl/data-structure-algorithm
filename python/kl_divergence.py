def kl_divergence(P, Q):
    if len(P) != len(Q): return
    
    from math import log2
    
    epsilon = 1e-10
    P = [max(e, epsilon) for e in P]
    Q = [max(e, epsilon) for e in Q]
    
    return round(sum(e * log2(e / Q[i]) for i, e in enumerate(P)), 4)

P = [0.1, 0.4, 0.5]
Q = [0.2, 0.3, 0.5]

print(kl_divergence(P, Q))