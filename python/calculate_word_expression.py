def find_index(array, target):
    for i, e in enumerate(array):
        if e == target: return i
    return -1

def calculate_word_expression(word: str) -> int:
    if " " in word: return
    
    list_number = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
    expressions = ["plus", "minus"]
    
    
    number_of_word = " ".join(" ".join(word.split("plus")).split("minus")).split(" ")
    
    result = 0
    current_word = word
    for i, e in enumerate(number_of_word):
        if e not in list_number: return 0
        if i == len(number_of_word) - 1: break
        if i == 0: result += list_number.index(e)
        
        plus = current_word.split("plus", 1)
        minus = current_word.split("minus", 1)
        if find_index(list_number, plus[0]) != -1:
            result += list_number.index(number_of_word[i + 1])
            current_word = plus[1]
        if find_index(list_number, minus[0]) != -1:
            result -= list_number.index(number_of_word[i + 1])
            current_word = minus[1]
    
    
    return result

def calculate_word_expression2(word: str) -> int:
    if " " in word: return 0
    
    list_number = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
    
    import re
    tokens = re.split(r"(plus|minus)", word)
    
    result = list_number.index(tokens[0])
    
    i = 1
    while i < len(tokens):
        operator = tokens[i]
        number_word = tokens[i + 1]

        if number_word not in list_number:
            return 0

        number = list_number.index(number_word)
        if operator == "plus":
            result += number
        elif operator == "minus":
            result -= number
        i += 2

    return result

print(calculate_word_expression("oneplusfourminustwoplustwominusnineminusninepluseight"))
print(calculate_word_expression2("oneplusfourminustwoplustwominusnineminusninepluseight"))