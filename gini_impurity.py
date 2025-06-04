def gini_impurity(array):
    from collections import Counter
    
    populations = Counter(array)
    n = len(array)
    
    return 1 - sum([(p / n) ** 2 for _, p in populations.items()])

labels = ["A", "B", "A", "A", "B", "B"]
print(gini_impurity(labels))