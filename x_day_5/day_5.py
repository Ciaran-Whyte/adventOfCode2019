

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

def run_computer(input_int):
    inputLines = readInputs("/Users/ciaran.whyte/dev/adventofcode2019/x_day_5/input.txt")
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
            inputLines = write(inputLines, pos, 1, input_int)
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

print("A: {0}".format(run_computer(1)))
print("B: {0}".format(run_computer(5)))
