def decode_ways(number: str, start=0):
    alphabet = [chr(i) for i in range(65, 91)]
    
    if list(number)[0] == '0': raise ValueError("String must not started with 0")
    
    int(number)
    
    if '00' in number: return 0
    for i in range(3, 10):
        if str(i) + '0' in number: return 0
    
    list_number = list(number)
    count_of_ways = 0
    index_of_zero = -1
    
    for i, el in enumerate(list_number):
        if el == '0': index_of_zero = i
    
    if index_of_zero == -1: count_of_ways += 1
    
    two_digit_index = []
    
    for i, el in enumerate(list_number):
        if i < len(list_number) - 1:
            after = list_number[i + 1]
            if el != '0':
                cond = int(el + after)
                if cond <= 26:
                    two_digit_index.insert(0, i)
    
    two_digit_index.reverse()
    exist = []
    
    for i, e in enumerate(two_digit_index):
        epoch_1 = []
        for z in range(len(two_digit_index)):
            e_count = 1
            index = []
            epoch_2 = []
            for j, el in enumerate(list_number):
                if j == e:
                    epoch_2.insert(j, el + list_number[j + 1])
                    continue
                if j < len(list_number) - 1:
                    if j == e + 1:
                        continue
                    after = list_number[j + 1]
                    if el != '0' and j + 1 != e:
                        cond = int(el + after)
                        if cond <= 26:
                            if e_count > z: break
                            if len(index) != 0 and index[0] == j: continue
                            epoch_2.insert(j, el + list_number[j + 1])
                            index.insert(0, j + 1)
                            e_count += 1
            if len(epoch_2) != 0: epoch_1.insert(z, epoch_2)
        if len(epoch_1) != 0: exist.insert(i, epoch_1)
    
    for i, el in enumerate(exist):
        if len(el) == 0: continue
        length = 0
        for j, e in enumerate(el):
            if j == 0:
                length += 1
                continue
            if e == el[j - 1]: continue
            length += 1
        count_of_ways += length
        
    return count_of_ways
    
print(decode_ways('22612'))

def num_decodings(s: str) -> int:
    if not s or s[0] == '0':
        return 0

    dp = [0] * (len(s) + 1)
    dp[0], dp[1] = 1, 1

    for i in range(2, len(s) + 1):
        one = int(s[i-1:i])
        two = int(s[i-2:i])
        
        if 1 <= one <= 9:
            dp[i] += dp[i - 1]
        if 10 <= two <= 26:
            dp[i] += dp[i - 2]

    return dp[-1]
    
print(num_decodings('226123112'))






'''
22612
22 6 1 2
22 6 12
2 26 12
2 26 1 2
2 2 6 12

'''