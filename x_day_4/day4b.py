
from collections import Counter

def has_two_digits(num):
    num=str(num)
    c = Counter(num) 
    if 2 in c.values():
        for double_digits in [x[0] for x in c.items() if x[1] is 2]:
            idx = num.index(double_digits)
            if num[idx] and num[idx+1]:
                return True
    return False


def is_increasing(num):
    num=str(num)
    for idx, char in enumerate(num):
        if idx+1 < len(num) and not (int(char) <= int(num[idx+1])):
            return False
    return True


start=347312
end=805915

count=0
while start <= end:
    start+=1  
    if is_increasing(start) and has_two_digits(start):
        count+=1

print(count)
