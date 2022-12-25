rocks = [[['@','@','@','@']], [['.','@','.'], ['@','@','@'], ['.','@','.']], [['.','.','@'], ['.','.','@'], ['@','@','@']],
         [['@'], ['@'], ['@'], ['@']], [['@', '@'], ['@', '@']]]

with open("Input/test.txt") as f:
    lines = f.readlines()

pattern = [['+', '-', '-', '-', '-', '-', '-', '-', '+']]

def findLastRow (array, character):
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] == character:
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
            for m in range(len(rock[i])):
                array[rowLast + i][initialPos + m] = rock[i][m]
    return array

def moveRock (array, instr, rock):
    instr = instr[0].strip()
    for i in range(len(instr)):
        if instr[i] == '>':
            row = findLastRow(array, '@')
            for j in range(row, row + len(rock)):
                for m in range(len(array[j]) - 1, -1, -1):
                    if array[j][m] == '@':
                        if m + 1 >= len(array[j]) - 1:
                            break
                        else:
                            array[j][m+1] = '@'
                            array[j][m] = '.'
        elif instr[i] == '<':
            row = findLastRow(array, '@')
            for j in range(row, row + len(rock)):
                for m in range(0, len(array[j])):
                    if array[j][m] == '@':
                        if m - 1 <= 0:
                            break
                        else:
                            array[j][m - 1] = '@'
                            array[j][m] = '.'
    return array

for i in range(len(rocks)):
    if len(pattern) == 1:
        pattern = addRows(pattern, 4 + len(rocks[i]) - 1)
        lastRow = 0
    else:
        lastRow = findLastRow(pattern, '#')
        if lastRow  < 4 + len(pattern):
            pattern = addRows(pattern, lastRow + 4 + len(rocks[i]) - 1)
    pattern = initialRock(pattern, rocks[i], lastRow)
    pattern = moveRock(pattern, lines, rocks[i])

print("A")

