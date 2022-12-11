with open('Input/input.txt') as f:
    lines = f.readlines()

actionsMonkey = 7 # Number of lines that describe every monkey action
numberMonkeys = int((len(lines) + 1) / actionsMonkey)
itemsMonkeys = []
operationMonkeys = []
testMonkeys = [0]*numberMonkeys
inspectItems = [0]*numberMonkeys
toMonkeys = []
rounds = 0

for i in range(numberMonkeys):
    itemsMonkeys.append([])
    operationMonkeys.append([])
    toMonkeys.append([])

# Extracting data from the file
for i in range(numberMonkeys):
    actions = lines[0+i*actionsMonkey:actionsMonkey+i*actionsMonkey]
    monkey = int(actions[0].strip().split()[1].replace(':',''))
    initialItems = actions[1].replace(',','').strip().split()[2:]
    itemsMonkeys[monkey].extend(initialItems)
    operation = actions[2].strip().split()[4:]
    operationMonkeys[i].extend(operation)
    test = int(actions[3].strip().split()[3])
    testMonkeys[i] = test
    nextMonkey = [int(actions[4].strip().split()[5]), int(actions[5].strip().split()[5])]
    toMonkeys[i].extend(nextMonkey)

while (rounds < 20):
    rounds += 1
    for i in range(numberMonkeys):
        toRemove = []
        for j in range(len(itemsMonkeys[i])):
            initialWorry = itemsMonkeys[i]
            worry = 0
            inspectItems[i] += 1
            if operationMonkeys[i][0] == '+':
                if operationMonkeys[i][1] == 'old':
                    worry = int(initialWorry[j]) + int(initialWorry[j])
                else:
                    worry = int(initialWorry[j]) + int(operationMonkeys[i][1])
            elif operationMonkeys[i][0] == '*':
                if operationMonkeys[i][1] == 'old':
                    worry = int(initialWorry[j]) * int(initialWorry[j])
                else:
                    worry = int(initialWorry[j]) * int(operationMonkeys[i][1])
            worry /= 3
            worry = int(worry)
            if worry % testMonkeys[i] == 0:
                itemsMonkeys[toMonkeys[i][0]].append(str(worry))
            else:
                itemsMonkeys[toMonkeys[i][1]].append(str(worry))
            toRemove.append(itemsMonkeys[i][j])
        for j in range(len(toRemove)):
            itemsMonkeys[i].remove(toRemove[j])

print(inspectItems)

max1 = max(inspectItems)
inspectItems.remove(max1)
max2 = max(inspectItems)

print(max1*max2)

