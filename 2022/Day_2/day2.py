result = 0
with open('Input/Inputs.txt') as f:
    for line in f:
        A = line.split()
        for i in range(2):
            if A[i] == 'X':
                A[i] = 'A'
            elif A[i] == 'Y':
                A[i] = 'B'
            elif A[i] == 'Z':
                A[i] = 'C'
        if A[1] == 'A':
            result += 1
        elif A[1] == 'B':
            result += 2
        elif A[1] == 'C':
            result += 3
        if A[0] == A[1]:
            result += 3
        elif A[1] == 'A' and A[0] == 'C':
            result += 6
        elif A[1] == 'B' and A[0] == 'A':
            result += 6
        elif A[1] == 'C' and A[0] == 'B':
            result += 6


print(result)