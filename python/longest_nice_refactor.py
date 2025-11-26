def is_nice(sub: str) -> bool:
    char_set = set(sub)
    for ch in sub:
        if ch.lower() not in char_set or ch.upper() not in char_set:
            return False
    return True

def longest_nice_substring(s: str) -> str:
    max_nice = ""
    n = len(s)

    for i in range(n):
        for j in range(i + 1, n + 1):
            sub = s[i:j]
            if is_nice(sub) and len(sub) > len(max_nice):
                max_nice = sub

    return max_nice

# Contoh tes:
print(longest_nice_substring("aAbBcCdDxyXYzZaABmMnNoOpPqQrRsStTuUvVwWxXyYzZ"))