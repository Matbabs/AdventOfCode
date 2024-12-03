import re

l, m = open("input.txt", "r").read().replace("\n", ""), []
m.extend(int(x) * int(y) for x, y in re.findall(r"mul\((\d+),(\d+)\)", l))
print(sum(m))
