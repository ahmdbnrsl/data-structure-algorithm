def tf_idf(documents):
    from collections import Counter
    from math import log10
    
    results = {}
    len_documents = len(documents)
    
    for i, document in enumerate(documents):
        temp = {}
        
        split_document = document.split(" ")
        count = Counter(split_document)
        len_doc = len(split_document)
        
        for index, word in enumerate(split_document):
            freq = 1
            for idx, el in enumerate(documents):
                split_doc = el.split(" ")
                if idx != i and word in split_doc: freq += 1
            
            temp[word] = round(count[word] / len_doc * log10(len_documents / freq), 4)
        
        results[i] = temp
    
    return results

documents = [
  "machine learning is amazing",
  "learning about mathematics is fun",
  "machine learning involves mathematics and statistics"
]
print(tf_idf(documents))