rocks = [[['#','#','#','#']], [['.','#','.'], ['#','#','#'], ['.','#','.']], [['.','.','#'], ['.','.','#'], ['#','#','#']],
         [['#'], ['#'], ['#'], ['#']], [['#', '#'], ['#', '#']]]

with open("Input/test.txt") as f:
    f.readlines()

pattern = [['+', '-', '-', '-', '-', '-', '-', '-', '+']]

def findLastRow (array):
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] == '#':
                return i

def addRows (array, rowsToAdd):
    n = 0
    while n < rowsToAdd:
        toAdd_temp = [['|', '.', '.', '.', '.', '.', '.', '.', '|']]
        toAdd_temp.extend(array)
        n += 1
        array = toAdd_temp
    return array

def initialRock (array, rock, rowLast):
    initialPos = 2
    for i in range(len(rock)):
        for j in range(len(rock)):
            for m in range(len(rock[j])):
                array[rowLast + i][initialPos + m] = rock[j][m]
    return array

for i in range(len(rocks)):
    if len(pattern) == 1:
        pattern = addRows(pattern, 4 + len(rocks[i]) - 1)
        lastRow = 0
    else:
        lastRow = findLastRow(pattern)
        if lastRow + 4 > len(pattern):
            pattern = addRows(pattern, lastRow + 4)
        elif lastRow + 4 + len(rocks[i]) - 1 > len(pattern):
            pattern = addRows(pattern, lastRow + 4 + len(rocks[i]) - 1)
    pattern = initialRock(pattern, rocks[i], lastRow)

print("A")

