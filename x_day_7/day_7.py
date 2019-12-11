from itertools import permutations

def readInputs(file_name):
    with open(file_name) as f:
        lineInput = f.readlines()
    return [int(x) for x in lineInput[0].split(",")]

def read(data, pos, step, mode):
    p = pos + step
    if mode != '0':
        return data[p]
    else:
        return data[data[p]]

def write(data, pos, step, value):
    data[data[pos+step]] = value
    return data

def run_computer(inputLines, input_int):
    pos = 0
    output = None
    while True:
        A, B, C, D, E = str(inputLines[pos]).zfill(5)
        op = int(D+E)

        if op == 99: # halt
            break

        if op == 1: # add
            arg_one = read(inputLines, pos, 1, C)
            arg_two = read(inputLines, pos, 2, B)
            inputLines = write(inputLines, pos, 3, arg_one + arg_two)
            pos += 4
        elif op == 2: # mul
            arg_one = read(inputLines, pos, 1, C)
            arg_two = read(inputLines, pos, 2, B)
            inputLines = write(inputLines, pos, 3, arg_one * arg_two)
            pos += 4
        elif op == 3: # write
            if len(input_int) == 0:
                raise Exception("BAD PLACE TO BE")
            arg = input_int.pop(0)
            inputLines = write(inputLines, pos, 1, arg)
            pos += 2
        elif op == 4: # print
            output = read(inputLines, pos, 1, C)
            # print("{}".format(output))
            pos += 2
        elif op == 5: # jump-if-true
            val1 = read(inputLines, pos, 1, C)
            val2 = read(inputLines, pos, 2, B)
            if val1 != 0:
                pos = val2
            else:
                pos +=3   
        elif op == 6: # jump-if-false
            val1 = read(inputLines, pos, 1, C)
            val2 = read(inputLines, pos, 2, B)
            if val1 == 0:
                pos = val2
            else:
                pos +=3
        elif op == 7: # less than
            val1 = read(inputLines, pos, 1, C)
            val2 = read(inputLines, pos, 2, B)
            inputLines = write(inputLines, pos, 3, 1 if val1 < val2 else 0)
            pos += 4
        elif op == 8: # equals
            val1 = read(inputLines, pos, 1, C)
            val2 = read(inputLines, pos, 2, B)
            inputLines = write(inputLines, pos, 3, 1 if val1 == val2 else 0)
            pos += 4
        else:
            print("ERROR: op should not be ==> {}".format(op))
            break
    return output


# PART 1
perm = permutations([0, 1, 2, 3, 4])
inputLines = readInputs("/Users/ciaran.whyte/dev/adventofcode2019/x_day_7/input.txt")
m = 0
code = []
for p in perm:
    o = 0
    for phase in p:
        o = run_computer(inputLines, [phase, o])
    if o > m:
        code = p[:]
        m = o
print(m)


