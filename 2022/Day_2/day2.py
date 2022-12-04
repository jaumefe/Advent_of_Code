result = 0
# Part 1
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
print("First part result: " + str(result))

# Part 2
result = 0
with open('Input/Inputs.txt') as f:
    for line in f:
        A = line.split()
        # Need to lose
        if A[1] == 'X':
            result += 0
            if A[0] == 'A':
                result += 3
            elif A[0] == 'B':
                result += 1
            elif A[0] == 'C':
                result += 2
        # Need to draw
        elif A[1] == 'Y':
            result += 3
            if A[0] == 'A':
                result += 1
            elif A[0] == 'B':
                result += 2
            elif A[0] == 'C':
                result += 3
        # Need to win
        elif A[1] == 'Z':
            result += 6
            if A[0] == 'A':
                result += 2
            elif A[0] == 'B':
                result += 3
            elif A[0] == 'C':
                result += 1


print("Second part result: " + str(result))