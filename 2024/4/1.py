k = "XMAS"
d, r = open("input.txt", "r").read().split("\n"), 0
s = lambda w: 1 if w == k or w == k[::-1] else 0
for y in range(len(d)):
    for x in range(len(d[0])):
        if x + 3 < len(d[0]):
            r += s(d[y][x : x + 4])
        if y + 3 < len(d):
            r += s("".join([d[y + z][x] for z in range(4)]))
        if x + 3 < len(d[0]) and y + 3 < len(d):
            r += s("".join([d[y + z][x + z] for z in range(4)]))
        if x - 3 >= 0 and y + 3 < len(d):
            r += s("".join([d[y + z][x - z] for z in range(4)]))
print(r)
