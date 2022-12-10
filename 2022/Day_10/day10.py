with open('Input/input.txt') as f:
    lines = f.readlines()

def signalStrength (cycle, inst):
    n = 0
    reg = 1
    for i in range(len(inst)):
        row = inst[i].strip().split()
        if row[0] == 'addx':
            n += 2
            if n >= cycle:
                return cycle*reg
            reg += int(row[1])
        elif row[0] == 'noop':
            n += 1
            if n >= cycle:
                return cycle*reg

result = 0

for i in range(20, 220 + 40, 40):
    result += signalStrength(i, lines)

print(result)
