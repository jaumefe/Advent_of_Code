with open('Input/input.txt') as f:
    lines = f.readlines()

for i in range(len(lines)):
    lines[i] = lines[i].strip().replace(":", '').split()

monkeyDict = {}

def monkey_eval (monkeyname, monkeydict):
    monkeyyell = monkeydict[monkeyname]
    if isinstance(monkeyyell, int):
        return monkeyyell
    else:
        if monkeyyell[1] == '+':
            return monkey_eval(monkeyyell[0], monkeydict) + monkey_eval(monkeyyell[2], monkeydict)
        elif monkeyyell[1] == '-':
            return monkey_eval(monkeyyell[0], monkeydict) - monkey_eval(monkeyyell[2], monkeydict)
        elif monkeyyell[1] == '*':
            return monkey_eval(monkeyyell[0], monkeydict) * monkey_eval(monkeyyell[2], monkeydict)
        elif monkeyyell[1] == '/':
            return monkey_eval(monkeyyell[0], monkeydict) / monkey_eval(monkeyyell[2], monkeydict)

for i in lines:
    monkeyName, monkeyYell = i[0], i[1:]
    if len(monkeyYell) == 1:
        monkeyYell = int(monkeyYell[0])
    monkeyDict[monkeyName] = monkeyYell

print(int(monkey_eval('root', monkeyDict)))

