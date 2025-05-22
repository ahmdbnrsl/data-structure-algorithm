def findInCapitalOrSmall(capital: list, small: list, target: int) -> (bool, int):
    for i, e in enumerate(capital):
        if ord(target) == e: return True, i
    for i, e in enumerate(small):
        if ord(target) == e: return False, i

def LongestNiceSubstring(string: str) -> str:
    capital = [i for i in range(65, 91)]
    small = [i for i in range(97, 123)]
    
    temp = [findInCapitalOrSmall(capital, small, el) for _, el in enumerate(list(string))]
    current_list = [i for i, el in enumerate(temp) if not (not el[0], el[1]) in temp]
    
    new_string = string
    for i, el in enumerate(current_list):
        new_string = new_string.replace(string[el], " ", 1)
    
    splitted_string = new_string.split(" ")
    list_posibility = ""
    
    for e in splitted_string:
        temp_result = ""
        is_temp = [findInCapitalOrSmall(capital, small, el) for i, el in enumerate(list(e))]
        for i, el in enumerate(is_temp):
            if not (not el[0], el[1]) in is_temp:
                temp_result += ","
                continue
            temp_result += e[i]
        temp_result += ","
        list_posibility += temp_result
    
    splited_posibility = list_posibility.split(",")
    len_list_result = [len(i) for i in splited_posibility if i != ""]
    list_result = [i for i in splited_posibility if i != ""]
    longest_nice = [i for i in list_result if len(i) == max(len_list_result)]
    return longest_nice[0]

print(LongestNiceSubstring("aAbBcCdDxyXYzZaABmMnNoOpPqQrRsStTuUvVwWxXyYzZ"))
# aAa