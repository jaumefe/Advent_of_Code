from anytree import Node, RenderTree

dirName = ""
subDir = []
file = []
subLevel = 0
subLevelMax = 0
parentTemp = 0
size = 0

# Tree structure for data

with open("Input/input.txt") as f:
    for line in f:
        line = line.strip().split()
        if line[0] == "$":
            if line[1] == "cd":
                if line[2] == '/':
                    root = Node('/', size=0)
                elif line[2] == '..':
                    subLevel -= 1
                else:
                    subLevel += 1
                    dirName = line[2]
                    if subLevel > subLevelMax:
                        subLevelMax = subLevel
        elif line[0] == "dir":
            if subLevel == 0:
                parentTemp = root
            else:
                for i in range(len(subDir)):
                    if subDir[i].name == dirName:
                        parentTemp = subDir[i]
            subdir_temp = Node(line[1], parent=parentTemp, size=0, sublevel=subLevel)
            subDir.append(subdir_temp)
        else:
            if subLevel == 0:
                parentTemp = root
            else:
                for i in range(len(subDir)):
                    if subDir[i].name == dirName:
                        parentTemp = subDir[i]
            fileTemp = Node(line[1], parent=parentTemp, size=int(line[0]), sublevel=subLevel)
            file.append(fileTemp)

## Compute the size of each directory

# Adding the file sizes to the directories
for i in range(len(subDir)):
    for j in range(len(file)):
        if file[j].parent == subDir[i]:
            subDir[i].size += file[j].size

# Computing the size of each directory

for i in range(len(subDir)):
    for j in range(len(subDir)):
        if subDir[j].parent == subDir[i]:
            subDir[i].size += subDir[j].size

# Check directories with a size lower than 100000
for i in range(len(subDir)):
    if subDir[i].size < 100000:
        size += subDir[i].size

print(size)
print(RenderTree(root))