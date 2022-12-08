with open("Input/input.txt") as f:
    lines = f.readlines()

## Part 1

visibleTree = 0
isVisible = 0

for i in range(1, len(lines)-1, 1):
    rows = lines[i].strip()
    for j in range(1, len(rows)-1, 1):
        # Check to Top
        for t in range(i - 1, -1, -1):
            if rows[j] <= lines[t][j]:
                isVisible = 0
                break
            elif t == 0:
                isVisible = 1
                visibleTree += 1
        # Check to bottom (if the tree has not been visible before)
        if not isVisible:
            for b in range(i + 1, len(lines), 1):
                if rows[j] <= lines[b][j]:
                    isVisible = 0
                    break
                elif b == len(lines)-1:
                    isVisible = 1
                    visibleTree += 1
        # Check to right
        if not isVisible:
            for r in range(j + 1, len(rows), 1):
                if rows[j] <= lines[i][r]:
                    isVisible = 0
                    break
                elif r == len(rows)-1:
                    isVisible = 1
                    visibleTree += 1
        #Check to left
        if not isVisible:
            for l in range(j - 1, -1, -1):
                if rows[j] <= lines[i][l]:
                    isVisible = 0
                    break
                elif l == 0:
                    isVisible = 1
                    visibleTree += 1

visibleTree += 2*len(rows) + 2*(len(lines)-2)

print('First part answer is: ' + str(visibleTree))

