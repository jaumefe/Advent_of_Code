with open("Input/input.txt") as f:
    lines = f.readlines()

count = 4

for i in range(len(lines[0].strip())-4):
    rep_item = 0
    marker = lines[0][i:(4+i)]
    for j in range(len(marker)):
        for k in range(len(marker)):
            if j < k:
                if marker[j] == marker[k]:
                    rep_item = 1
                    break
        if rep_item:
            count += 1
            break
    if not rep_item:
        break

print(count)
