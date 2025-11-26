def entropy_predictability(array):
    from collections import Counter
    from math import log2
    
    populations = Counter(array)
    n = len(array)
    
    return -sum([(p / n) * log2(p / n) for _, p in populations.items()])

labels = ["A", "B", "A", "A", "B", "A", "B", "B", "B", "B"]
print(entropy_predictability(labels))