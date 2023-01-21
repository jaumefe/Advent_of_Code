from collections import deque

with open("Input/input.txt") as f:
    lines = f.readlines()

encryptedFile = deque([])

for i in range(len(lines)):
    encryptedFile.append((i, int(lines[i].strip())))

def mix(enumerated: deque):
    for original_index in range(len(enumerated)):
        while enumerated[0][0] != original_index:
            enumerated.rotate(-1)

        current_pair = enumerated.popleft()
        shift = current_pair[1] % len(enumerated)
        enumerated.rotate(-shift)
        enumerated.append(current_pair)
    return enumerated

def value_at_n(values: list, n: int):
    digit_posn = (values.index(0)+n) % len(values)
    return values[digit_posn]

enumerateFile = mix(encryptedFile)

coord_sum = 0
for n in (1000, 2000, 3000):
    # Turn our enumerated list into a list
    coord_sum += value_at_n([val[1] for val in enumerateFile], n)
print(f"Part 1: {coord_sum}")