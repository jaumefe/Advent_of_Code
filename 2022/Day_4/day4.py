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


print(count)


