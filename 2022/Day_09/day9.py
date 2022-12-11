with open('Input/test.txt') as f:
    lines = f.readlines()

motion = [["H"]]

def findPos (name, array):
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] == name:
                return i, j

def moveT (array, direction):
    rowH, colH = findPos('H', array)
    try:
        rowT, colT = findPos('T', array)
    except TypeError:
        if direction == 'U':
            rowT, colT = rowH, colH - 1
        elif direction == 'D':
            rowT, colT = rowH, colH + 1
        elif direction == 'L':
            rowT, colT = rowH, colH + 1
        elif direction == 'R':
            rowT, colT = rowH, colH - 1
        array[rowT][colT] = 'T'
        rowT, colT = rowH, colH
    if (abs(colH - colT) > 1) and (abs(rowH - rowT) <= 1):
        if colH > colT:
            array[rowH][colH - 1] = 'T'
            array[rowT][colT] = '#'
        else:
            array[rowH][colH + 1] = 'T'
            array[rowT][colT] = '#'
    elif (abs(colH - colT) <= 1) and (abs(rowH - rowT) > 1):
        if rowT > rowH:
            array[rowH+1][colH] = 'T'
            array[rowT][colT] = '#'
        else:
            array[rowH - 1][colH] = 'T'
            array[rowT][colT] = '#'


def moveH (direction, array, step):
    rowH, colH = findPos('H', array)
    if direction == 'R':
        if step + colH >= len(array[rowH]):
            stepsToAdd = step - len(array[rowH]) + 1
            for m in range(stepsToAdd + 1):
                array[rowH].append("-")
        for m in range(colH, colH + step, 1):
            if array[rowH][m] == 'H':
                array[rowH][m] = '-'
            array[rowH][m + 1] = 'H'
            moveT(array, direction)
    elif direction == 'L':
        if step > colH:
            stepsToAdd = step - len(array[rowH]) + 1
            C = []
            for m in range(stepsToAdd + 1):
                C.append('-')
            C.extend(array[rowH])
            array[rowH] = C
            rowH, colH = findPos('H', array)
        for m in range(colH, colH - step, -1):
            if array[rowH][m] == 'H':
                array[rowH][m] = '-'
            array[rowH][m - 1] = 'H'
            moveT(array, direction)
    elif direction == 'U':
        if step > rowH:
            stepsToAdd = step - len(array) + 1
            C = []
            for m in range(stepsToAdd + 1):
                C.append(['-']*len(array[0]))
            C.extend(array)
            array = C
            rowH, colH = findPos('H', array)
        for m in range(rowH, rowH - step, -1):
            if array[m][colH] == 'H':
                array[m][colH] = '-'
            array[m - 1][colH] = 'H'
            moveT(array, direction)
    elif direction == 'D':
        if step + rowH >= len(array):
            stepsToAdd = step - len(array) + 1
            C = []
            for m in range(stepsToAdd + 1):
                C.append(['-']*len(array[0]))
            array.extend(C)
            rowH, colH = findPos('H', array)
        for m in range(rowH, rowH - step, -1):
            if array[m][colH] == 'H':
                array[m][colH] = '-'
            array[m+1][colH] = 'H'
            moveT(array, direction)
    return array


for i in range(len(lines)):
    order = lines[i].strip().split()
    steps = int(order[1])
    motion = moveH(order[0], motion, steps)
#    for j in range(len(motion)):
#        print(motion[j])

for i in range(len(motion)):
    print(motion[i])
