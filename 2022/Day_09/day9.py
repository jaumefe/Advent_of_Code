with open('Input/test.txt') as f:
    lines = f.readlines()

motion = [["H"]]

def findPos (name, array):
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] == name:
                return i, j

def moveT (array):
    rowH, colH = findPos('H', array)
    try:
        rowT, colT = findPos('T', array)
    except TypeError:
        rowT, colT = 0,0
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
        if step > len(array[rowH]):
            stepsToAdd = step - len(array[rowH])
            for m in range(stepsToAdd + 1):
                array[rowH].append("-")
        for m in range(colH + 1, colH + step, 1):
            if array[rowH][m - 1] == 'H':
                array[rowH][m - 1] = '-'
            array[rowH][m] = 'H'
            moveT(array)
    elif direction == 'L':
        if step > len(array[rowH]):
            stepsToAdd = step - len(array[rowH])
            C = []
            for m in range(stepsToAdd + 1):
                C.append('-')
            C.extend(array[rowH])
            array[rowH] = C
            rowH, colH = findPos('H', array)
        for m in range(colH - 1, colH - 1 - step, -1):
            if array[rowH][m + 1] == 'H':
                array[rowH][m + 1] = '-'
            array[rowH][m] = 'H'
            moveT(array)
    elif direction == 'U':
        if step > len(array):
            stepsToAdd = step - len(array)
            C = []
            for m in range(stepsToAdd + 1):
                C.append(['-']*len(array[0]))
            C.extend(array)
            array = C
            rowH, colH = findPos('H', array)
        for m in range(rowH + 1, rowH + 1 - step, -1):
            if array[m-1][colH] == 'H':
                array[m - 1][colH] = '-'
            array[m - 2][colH] = 'H'
            moveT(array)
    elif direction == 'D':
        if step > len(array):
            stepsToAdd = step - len(array)
            C = []
            for m in range(stepsToAdd + 1):
                C.append(['-']*len(array[0]))
            array.extend(C)
            rowH, colH = findPos('H', array)
        for m in range(rowH + 1, rowH - 1 + step, -1):
            if array[m-1][colH] == 'H':
                array[m - 1][colH] = '-'
            array[m][colH] = 'H'
            moveT(array)
    return array


for i in range(len(lines)):
    order = lines[i].strip().split()
    steps = int(order[1])
    motion = moveH(order[0], motion, steps)
