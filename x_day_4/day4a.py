


def has_two_digits(num):
    num=str(num)
    for idx, char in enumerate(num):
        if idx+1 < len(num) and char == num[idx+1]:
            print("has_two_digits {} => {}".format(char, num))
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
