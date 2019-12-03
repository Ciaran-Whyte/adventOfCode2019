
from math import *

def parse_lines(cords):
    in_line = [(0, 0)]
    for next_p in cords:
        cur_p = in_line[-1]
        if next_p[0] == 'R':
            in_line.append((cur_p[0] + int(next_p[1:]), cur_p[1]))
        elif next_p[0] == 'L':
            in_line.append((cur_p[0] - int(next_p[1:]), cur_p[1]))
        elif next_p[0] == 'U':
            in_line.append((cur_p[0], cur_p[1] + int(next_p[1:])))
        elif next_p[0] == 'D':
            in_line.append((cur_p[0], cur_p[1] - int(next_p[1:])))

    expanded_lines = []
    for x in range(len(in_line)):
        if x+1 < len(in_line):
            expanded_lines.append((in_line[x], in_line[x+1]))
    return expanded_lines

def ccw(A,B,C):
    return (C[1]-A[1]) * (B[0]-A[0]) > (B[1]-A[1]) * (C[0]-A[0])

def intersect(A,B,C,D):
    return ccw(A,C,D) != ccw(B,C,D) and ccw(A,B,C) != ccw(A,B,D)

def findIntersection(x1,y1,x2,y2,x3,y3,x4,y4):
    px= ( (x1*y2-y1*x2)*(x3-x4)-(x1-x2)*(x3*y4-y3*x4) ) / ( (x1-x2)*(y3-y4)-(y1-y2)*(x3-x4) ) 
    py= ( (x1*y2-y1*x2)*(y3-y4)-(y1-y2)*(x3*y4-y3*x4) ) / ( (x1-x2)*(y3-y4)-(y1-y2)*(x3-x4) )
    return [px, py]

def readInputs(file_name):
    with open(file_name) as f:
        lineInput = f.readlines()
    lineInput = [x.strip() for x in lineInput]
    return parse_lines(lineInput[0].split(",")), parse_lines(lineInput[1].split(","))

cords_a, cords_b = readInputs(
    "/Users/ciaran.whyte/dev/adventofcode2019/x_day_3_a/src/input.txt")

match = []
for line_a in cords_a:
    for line_b in cords_b:
        A, B = line_a
        C, D = line_b
        if intersect(A,B,C,D):
            match.append(findIntersection(A[0],A[1],B[0],B[1],C[0],C[1],D[0],D[1]))

 
def manhattan_distance(point_a, point_b):
    return abs(point_a[0] - point_b[0]) + abs(point_a[1] - point_b[1])

answer = None
for m in match:
    d = manhattan_distance((0,0),m)

    if answer is None:
        answer = d
    if d < answer:
        answer = d

print("answer: {}".format(answer))
