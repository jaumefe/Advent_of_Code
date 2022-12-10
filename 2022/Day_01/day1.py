sum = 0
elf_index = 0
elf_index_max = 0
max_sums = [0, 0, 0]

with open('Input/Input.txt') as f:
    for line in f:
            if line != '\n':
                sum += int(line)
            else:
                elf_index += 1
                if sum > max(max_sums):
                    max_sums[2] = max_sums[1]
                    max_sums[1] = max_sums[0]
                    max_sums[0] = sum
                    elf_index_max = elf_index
                elif sum > min(max_sums):
                    if sum > max_sums[1]:
                        max_sums[2] = max_sums[1]
                        max_sums[1] = sum
                    elif sum > max_sums[2] and sum < max_sums[1]:
                        max_sums[2] = 2
                sum = 0
print(max(max_sums))
print(max_sums[0] + max_sums[1] + max_sums[2])
print(elf_index_max)
