with open("Input/input.txt") as f:
    lines = f.readlines()

totalAdj = 0
exposed = 0

for i in range(len(lines)):
    lines[i] = lines[i].strip().replace(',', ' ').split()


def findAdj (coord, array):
    adj = 0
    for i in range(len(array)):
        if abs(int(coord[0]) - int(array[i][0])) == 1 and int(coord[1]) - int(array[i][1]) == 0 and int(coord[2]) - int(array[i][2]) == 0:
            adj += 1
        elif abs(int(coord[1]) - int(array[i][1])) == 1 and int(coord[0]) - int(array[i][0]) == 0 and int(coord[2]) - int(array[i][2]) == 0:
            adj += 1
        elif abs(int(coord[2]) - int(array[i][2])) == 1 and int(coord[0]) - int(array[i][0]) == 0 and int(coord[1]) - int(array[i][1]) == 0:
            adj += 1
    return adj

for i in range(len(lines)):
    totalAdj = findAdj(lines[i], lines)
    exposed += 6 - totalAdj

print(exposed)