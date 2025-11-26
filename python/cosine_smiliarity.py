def cosine_similiarity(string1, string2):
    set1 = set(string1)
    set2 = set(string2)
    basic = sorted(set1.union(set2))
    
    from collections import Counter
    from math import sqrt
    
    def freq_vec(s):
        freq = Counter(s)
        return [freq[char] for char in basic]
    
    A, B = freq_vec(string1), freq_vec(string2)
    dot_product = sum(a * b for a, b in zip(A, B))
    
    def magnitude(vec):
        return sqrt(sum(i**2 for i in vec))
    
    m_A, m_B = magnitude(A), magnitude(B)
    
    if m_A == 0 or m_B == 0: return 0.0
    
    return round(dot_product / (m_A * m_B), 4)

print(cosine_similiarity("machinelearning", "mathematics"))