with open("Input/test.txt") as f:
    lines = f.readlines()

for i in range(len(lines)):
    lines[i] = lines[i].strip().replace('.', '. ').replace('#', '# ').replace('v', 'v ').replace('^', '^ ').replace('<', '< ').replace('>', '> ').split()

def get_unique_numbers(numbers):

    list_of_unique_numbers = []

    unique_numbers = set(numbers)

    for number in unique_numbers:
        list_of_unique_numbers.append(number)

    return list_of_unique_numbers

# Find initial position
for i in range(len(lines[0])):
    if lines[0][i] != '#':
        lines[0][i] = 'E'

new_i = []
old_i = []
new_j = []
old_j = []
iD = []

# Motion of blizzards
for i in range(1, len(lines)):
    for j in range(1, len(lines[i])):
        if lines[i][j] == '>':
            if j + 1 == len(lines[i]) - 1:
                new_j.append(1)
            else:
                new_j.append(j + 1)
            new_i.append(i)
            old_i.append(i)
            old_j.append(j)
            iD.append('>')
        elif lines[i][j] == '<':
            if j - 1 == 0:
                new_j.append(len(lines[i]) - 2)
            else:
                new_j.append(j - 1)
            new_i.append(i)
            old_i.append(i)
            old_j.append(j)
            iD.append('<')
        elif lines[i][j] == '^':
            if i - 1 == 0:
                new_i.append(len(lines) - 2)
            else:
                new_i.append(i - 1)
            new_j.append(j)
            old_j.append(j)
            old_i.append(i)
            iD.append('^')
        elif lines[i][j] == 'v':
            if i + 1 == len(lines) - 1:
                new_i.append(1)
            else:
                new_i.append(i + 1)
            new_j.append(j)
            old_j.append(j)
            old_i.append(i)
            iD.append('v')

for i in range(len(new_i)):
    lines[new_i[i]][new_j[i]] = iD[i]

toRemove = []

for i in range(len(old_i)):
    for j in range(len(new_i)):
        if (old_i[i] == new_i[j]) and (old_j[i] == new_j[j]):
            toRemove.append(i)

toRemove = get_unique_numbers(toRemove)

for i in range(len(toRemove) - 1, -1, -1):
    old_i.pop(toRemove[i])
    old_j.pop(toRemove[i])

for i in range(len(old_i)):
    lines[old_i[i]][old_j[i]] = '.'

# Determine if a spot has more than one element

indexRep = []
freqRep = [1]*len(new_i)

for i in range(len(new_i)):
    for j in range(len(new_i)):
        if i < j:
            if new_i[i] == new_i[j] and new_j[i] == new_j[j]:
                indexRep.append(i)
                freqRep[i] += 1

print("B")

print("A")