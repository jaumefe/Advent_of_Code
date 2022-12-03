import string

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

print(sum)