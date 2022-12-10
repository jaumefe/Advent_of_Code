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
                return reg
            reg += int(row[1])
        elif row[0] == 'noop':
            n += 1
            if n >= cycle:
                return reg

## Part 1

result = 0

for i in range(20, 220 + 40, 40):
    result += i*signalStrength(i, lines)

print(result)

## Part 2
spriteLen = 3
CRT = [40, 6]
screen = []

result = 0

for i in range(CRT[1]):
    screen.append(['']*CRT[0])

for j in range(0, 6):
    for i in range(0, 40):
        result = signalStrength(i+1+40*j, lines)
        if i == result - 1:
            screen[j][i] = '#'
        elif i == result:
            screen[j][i] = '#'
        elif i == result + 1:
            screen[j][i] = '#'
        else:
            screen[j][i] = '.'

for i in range(len(screen)):
    print(screen[i])
