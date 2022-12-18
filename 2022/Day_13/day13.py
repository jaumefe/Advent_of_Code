with open("Input/test.txt") as f:
    lines = f.readlines()



def checkOrder(messageA, messageB):
    if len(messageA) > len(messageB):
        long = len(messageB)
    else:
        long = len(messageA)
    for i in range(long):
        if int(messageA[i]) < int(messageB[i]):
            return True
        elif int(messageA[i]) > int(messageB[i]):
            return False
        elif i == long - 1:
            if len(messageA) > len(messageB):
                return False
            else:
                return True

def listToTest (message, index):
    mess = []
    for i in range(index + 1, len(message)):
        if message[i] == '[':
            return listToTest(message, i)
        elif message[i] != ']':
            mess.append(int(message[i]))
        elif len(mess) == 0:
            return [0]
        else:
            return mess

n = 0
index = 1
res = 0

while (n < len(lines) - 2):
    message1 = lines[n].strip().replace(",", "")
    message2 = lines[n+1].strip().replace(",", "")
    if len(message1) > len(message2):
        length = message2
    else:
        length = message1
    for i in range(len(length)):
        if message1[i] == '[' and message2[i] == '[':
            if checkOrder(listToTest(message1, i), listToTest(message2, i)):
                res = res + index
                print(index)
                break
        elif message1[i] == '[' or message2[i] == '[':
                if message1[i] != '[':
                    if checkOrder(listToTest(message1, i-1), listToTest(message2, i)):
                        res = res + index
                        print(index)
                        break
                elif message2[i] != '[':
                    if checkOrder(listToTest(message1, i), listToTest(message2, i - 1)):
                        res = res + index
                        print(index)
                        break
    index += 1
    n += 3

print(res)