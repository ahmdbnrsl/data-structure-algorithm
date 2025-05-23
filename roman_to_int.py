def roman_to_int(roman: str) -> int:
    roman = roman.upper()
    list_roman = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }
    
    if len(roman) == 1: return list_roman[roman]
    
    index = 0
    for key, val in list_roman.items():
        index += 1
        
        if index % 2 == 0 and key * 2 in roman: return 0
        elif key * 4 in roman: return 0
    
    result = 0
    for i, el in enumerate(roman):
        if el not in list_roman: return 0
        
        if str(list_roman[el]).startswith("5"):
            if i < len(roman) - 1 and list_roman[roman[i + 1]] >= list_roman[el]: return 0
                
            if i == 0:
                result += list_roman[el]
                continue
            
            if list_roman[roman[i - 1]] < list_roman[el]:
                if list_roman[el] / list_roman[roman[i - 1]] != 5:
                    return 0
                else:
                    result += list_roman[el] - list_roman[roman[i - 1]]
                    continue
            if list_roman[roman[i - 1]] > list_roman[el]: result += list_roman[el]
            
        elif str(list_roman[el]).startswith("1"):
            
            if i < len(roman) - 1 and list_roman[el] < list_roman[roman[i + 1]]:
                if i != 0 and list_roman[roman[i - 1]] <= list_roman[el]:
                    return 0
                if list_roman[roman[i - 1]] == list_roman[roman[i + 1]] and str(list_roman[roman[i + 1]]).startswith("5"):
                    return 0
                continue
        
            if i == 0:
                result += list_roman[el]
                continue
            
            if list_roman[roman[i - 1]] < list_roman[el]:
                if list_roman[el] / list_roman[roman[i - 1]] != 10:
                    return 0
                else:
                    if i < len(roman) - 1 and list_roman[el] == list_roman[roman[i + 1]]: return 0
                    result += list_roman[el] - list_roman[roman[i - 1]]
                    continue
            if list_roman[roman[i - 1]] >= list_roman[el]: result += list_roman[el]
        
    return result
    
    
print(roman_to_int("MCDLXXXIX"))