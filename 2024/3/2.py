import re

l, m = open("input.txt", "r").read().replace("\n", ""), []
m.extend(
    int(x) * int(y)
    for x, y in re.findall(r"mul\((\d+),(\d+)\)", re.sub(r"don't\(\).*?do\(\)", "", l))
)
print(sum(m))
