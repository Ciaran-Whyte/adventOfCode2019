def readInputs(file_name):
    with open(file_name) as f:
        lineInput = f.readlines()
    return [x.strip() for x in lineInput]


in6 = readInputs("/Users/ciaran.whyte/dev/adventofcode2019/x_day_6/input.txt")
orbits = {}
for line in in6:
    source, orbiter = line.split(')')
    orbits[orbiter] = source
    # print("{} -> {}".format(orbiter, source))

num_orb = 0
for k in orbits.keys(): # for each orbiter
    s = orbits.get(k, None) 
    while s:  # while we haven't hit the root
        num_orb += 1 # add one
        s = orbits.get(s, None)  # get the orbiters source and check it for a source

print("A: {}".format(num_orb))


def get_path(orbit_dict, start):
    path = []
    obj = orbit_dict[start]
    while obj:
        path.append(obj)
        obj = orbit_dict.get(obj, None)
    return path

santas_path = get_path(orbits, "SAN")
our_path = get_path(orbits, "YOU")


min_path = 100000
for i, p in enumerate(our_path):
    try:
        san_way = santas_path.index(p)
    except ValueError:
        continue
    total_steps = i + san_way
    if total_steps < min_path:
        min_path = total_steps

print("B: {}".format(min_path))