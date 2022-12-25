with open("Input/test.txt") as f:
    lines = f.readlines()

for i in range(len(lines)):
    lines[i] = lines[i].strip().replace('.', '. ').replace('#', '# ').split()

priority = ['N', 'S', 'W', 'E']

rounds = 0

while rounds < 10:
    new_i = []
    old_i = []
    new_j = []
    old_j = []
    iD = []
    rounds += 1
    toMove = priority[0]
    for i in range(len(lines)):
        for j in range(len(lines[i])):
            if lines[i][j] == '#':
                if toMove == 'N':
                    if i - 1 <= 0:
                        A = [['.']*len(lines[0])]
                        A.extend(lines)
                    new_i.append(i)
                    new_j.append(j)
                    old_i.append(i+1)
                    old_j.append(j)
    for i in range(len(new_i)):
        for j in range(len(new_j)):
            A[new_i][new_j] = '#'
            A[old_i][old_j] = '.'


    priority.remove(toMove)
    priority.append(toMove)