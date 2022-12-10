import pandas as pd
# Part 1

count = 0

with open('Input/Input.txt') as f:
    for line in f:
        A = line.split(',')
        A1 = A[0].split('-')
        A2 = A[1].split('-')
        if (int(A1[0]) <= int(A2[0])) and (int(A1[1]) >= int(A2[1])):
            count += 1
        elif (int(A2[0]) <= int(A1[0])) and (int(A2[1]) >= int(A1[1])):
            count += 1

print('First part answer is: ' + str(count))

# Part 2

count = 0

with open('Input/Input.txt') as f:
    for line in f:
        A = line.split(',')
        A1 = A[0].split('-')
        A2 = A[1].split('-')
        i1 = pd.Interval(int(A1[0]), int(A1[1]), closed='both')
        i2 = pd.Interval(int(A2[0]), int(A2[1]), closed='both')
        if i1.overlaps(i2):
            count += 1

print('Second part answer is: ' + str(count))


