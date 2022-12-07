import numpy as np

# Part 1

with open('Input/input.txt') as f:
    lines = f.readlines()

n = 0
crates = []

# Obtaining crate structure and formatting
for i in range(len(lines)):
    if i > 0 and (len(lines[i].strip()) == len(lines[i-1].strip())):
        crates.append(lines[i].strip())
    elif i == 0:
        crates.append(lines[i].strip())
    elif len(lines[i].strip()) != len(lines[i-1].strip()):
        num = lines[i].split()
        n = i
        break

rows = i
crates_np = np.array([[""]*rows]*len(num))

# Movements obtained
lines = lines[i+2:len(lines)]

for i in range(len(crates)):
    crates[i] = crates[i].replace("    ", " [-]")
    crates[i] = crates[i].split()

# Crate movements

for i in range(len(lines)):
    count = 0
    toMove = []
    A = lines[i].split()
    crateToMove = int(A[1])
    initialColumn = int(A[3])
    finalColumn = int(A[5])

    for j in range(len(crates)):
        if count >= crateToMove:
            break
        if crates[j][initialColumn-1] != '[-]':
            toMove.append(crates[j][initialColumn-1])
            crates[j][initialColumn - 1] = '[-]'
            count += 1

    toMoveRemove = []
    if crates[0][finalColumn - 1] == '[-]':
        lenToMove = len(toMove)

        for j in range(lenToMove):
            for k in range(rows):
                if crates[rows - 1 - k][finalColumn - 1] == '[-]':
                    crates[rows - 1 - k][finalColumn - 1] = toMove[j]
                    toMoveRemove.append(toMove[j])
                    count -= 1
                    break

    for j in range(len(toMoveRemove)):
        toMove.remove(toMoveRemove[j])
    if count != 0:
        if crates[0][finalColumn - 1] != '[-]':
            for j in range(len(toMove)):
                C = [['[-]'] * 9]
                C[0][finalColumn - 1] = toMove[j]
                C.extend(crates)
                crates = C
                count -= 1
                rows += 1
                if count == 0:
                    break

print(crates)
print(crates[0])

# Part 2

with open('Input/input.txt') as f:
    lines = f.readlines()

n = 0
crates = []

# Obtaining crate structure and formatting
for i in range(len(lines)):
    if i > 0 and (len(lines[i].strip()) == len(lines[i-1].strip())):
        crates.append(lines[i].strip())
    elif i == 0:
        crates.append(lines[i].strip())
    elif len(lines[i].strip()) != len(lines[i-1].strip()):
        num = lines[i].split()
        n = i
        break

rows = i
crates_np = np.array([[""]*rows]*len(num))

# Movements obtained
lines = lines[i+2:len(lines)]

for i in range(len(crates)):
    crates[i] = crates[i].replace("    ", " [-]")
    crates[i] = crates[i].split()

# Crate movements

for i in range(len(lines)):
    count = 0
    toMove = []
    A = lines[i].split()
    crateToMove = int(A[1])
    initialColumn = int(A[3])
    finalColumn = int(A[5])

    for j in range(len(crates)):
        if count >= crateToMove:
            break
        if crates[j][initialColumn-1] != '[-]':
            toMove.append(crates[j][initialColumn-1])
            crates[j][initialColumn - 1] = '[-]'
            count += 1

    toMoveRemove = []
    if crates[0][finalColumn - 1] == '[-]':
        lenToMove = len(toMove)

        for j in range(lenToMove - 1, -1, -1):
            for k in range(rows):
                if crates[rows - 1 - k][finalColumn - 1] == '[-]':
                    crates[rows - 1 - k][finalColumn - 1] = toMove[j]
                    toMoveRemove.append(toMove[j])
                    count -= 1
                    break

    for j in range(len(toMoveRemove)):
        toMove.remove(toMoveRemove[j])
    if count != 0:
        if crates[0][finalColumn - 1] != '[-]':
            for j in range(len(toMove)- 1, -1, -1):
                C = [['[-]'] * 9]
                C[0][finalColumn - 1] = toMove[j]
                C.extend(crates)
                crates = C
                count -= 1
                rows += 1
                if count == 0:
                    break

print(crates)
print(crates[0])