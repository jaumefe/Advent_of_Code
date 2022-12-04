import string

# Part 1

item_type = []
rep_item = ''
sum = 0

with open('Input/Input.txt') as f:
    for line in f:
        A = line.strip()
        length = len(A)
        item1 = A[0:int(length/2)]
        item2 = A[int(length/2):]
        for i in range(int(length/2)):
            for j in range (int(length/2)):
                if item1[i] == item2[j]:
                    rep_item = item1[i]
        item_type.append(rep_item)

for i in range(len(item_type)):
    if (item_type[i]).isupper():
        sum += string.ascii_uppercase.index(item_type[i]) + 1 + 26
    else:
        sum += string.ascii_lowercase.index(item_type[i]) + 1

print('First part answer is: ' + str(sum))

# Part 2
item_type = []
rep_item = ''
sum = 0
n = 0

with open('Input/Input.txt') as f:
    lines = f.readlines()

while n < len(lines):
    A = lines[n:(3+n)]
    for i in range(len(A[0].strip())):
        for j in range(len(A[1].strip())):
            for k in range(len(A[2].strip())):
                if A[0][i] == A[1][j] and A[0][i] == A[2][k]:
                    rep_item = A[0][i]
    item_type.append(rep_item)
    n += 3

for i in range(len(item_type)):
    if (item_type[i]).isupper():
        sum += string.ascii_uppercase.index(item_type[i]) + 1 + 26
    else:
        sum += string.ascii_lowercase.index(item_type[i]) + 1

print('Second part answer is: ' + str(sum))

