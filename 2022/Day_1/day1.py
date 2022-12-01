max_sum = 0
sum = 0
elf_index = 0
elf_index_max = 0

with open('Input/Input.txt') as f:
    for line in f:
            if line != '\n':
                sum += int(line.strip())
            else:
                elf_index += 1
                if sum > max_sum:
                    max_sum = sum
                    elf_index_max = elf_index
                    sum = 0


print(max_sum)
print('\n')
print(elf_index_max)
print('\n')
