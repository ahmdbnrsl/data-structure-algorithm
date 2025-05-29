def max_decode_alphabet_sum(string: str) -> int:
    upper = string.upper()
    chars = [chr(i) for i in range(65, 91)]
    
    concat = []
    for i, e in enumerate(upper):
        if i < len(upper) - 1:
            char = chars.index(e) + 1
            nextChar = chars.index(upper[i + 1]) + 1
            concatChar = str(char) + str(nextChar)
            
            if int(concatChar) <= 26:
                concat.insert(i, (i, i + 1, int(concatChar)))
    
    current_remove = ()
    for i, e in enumerate(concat):
        if e == 0: continue
        if i == len(concat) - 1: break
        if e[1] == concat[i + 1][0]:
            if e[2] == concat[i + 1][2]:
                concat[i + 1] = 0
                continue
            if e[2] < concat[i + 1][2]:
                temp = current_remove
                current_remove = e
                concat[i] = 0
                if i != 0 and len(temp) > 0 and temp[1] == e[0]:
                    concat.insert(len(concat), temp)
    
    while 0 in concat:
        concat.remove(0)
    
    indexs = []
    concat_sum = 0
    for i in concat:
        indexs.insert(0, i[0])
        indexs.insert(0, i[1])
        concat_sum += i[2]
    
    non_concat = []
    for i, e in enumerate(upper):
        if i not in indexs:
            non_concat.insert(0, chars.index(e) + 1)
    
    return sum(non_concat) + concat_sum

'''
[(0, 1, 11), (1, 2, 11), (2, 3, 12), (3, 4, 22), (4, 5, 23), (6, 7, 13), (8, 9, 11), (9, 10, 14)]

[(0, 1, 11), (2, 3, 12), (4, 5, 23), (6, 7, 13), (9, 10, 14)]
'''
print(max_decode_alphabet_sum("abcjkl"))