from math import pow

with open("Input/input.txt") as f:
    lines = f.readlines()

table = {'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}

def snafutoDec (snafu):
    result = 0
    for i in range(len(snafu)):
        if snafu[i] == '=':
            result -= 2*int(pow(5, len(snafu)-1-i))
        elif snafu[i] == '-':
            result -= int(pow(5, len(snafu)-1-i))
        else:
            result += int(snafu[i])*int(pow(5, len(snafu)-1-i))
    return result

sum = 0

def decToSnafu (dec):
    result = []
    result_temp = []
    result_str = ''
    while int(dec / 5) >= 5:
        result_temp.append(str(int(dec % 5)))
        dec = int(dec / 5)
    result_temp.append(str(int(dec % 5)))
    result_temp.append(str(int(dec/5)))
    for i in range(len(result_temp) - 1, -1, -1):
        result.append(result_temp[i])
    for i in range(len(result) - 1, -1, -1):
        if int(result[i]) == 4:
            result[i] = '-'
            if i - 1 < 0:
                result_temp = ['0']
                result_temp.extend(result)
                i += 1
                result = result_temp
            result[i-1] =str(int(result[i - 1]) + 1)
        elif int(result[i]) == 3:
            result[i] = '='
            if i - 1 < 0:
                result_temp = ['0']
                result_temp.extend(result)
                i += 1
                result = result_temp
            result[i - 1] = str(int(result[i - 1]) + 1)
        elif int(result[i]) == 5:
            result[i] = '0'
            if i - 1 < 0:
                result_temp = ['0']
                result_temp.extend(result)
                i += 1
                result = result_temp
            result[i - 1] = str(int(result[i - 1]) + 1)
    for i in range(len(result)):
        result_str += result[i]
    return result_str

for i in range(len(lines)):
    line = lines[i].strip()
    sum += snafutoDec(line)

print(decToSnafu(sum))