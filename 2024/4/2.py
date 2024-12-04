k = "MAS"
d, r = open("input.txt", "r").read().split("\n"), 0
s = lambda w1, w2: 1 if (w1 == k or w1 == k[::-1]) and (w2 == k or w2 == k[::-1]) else 0
for y in range(len(d)):
    for x in range(len(d[0])):
        if x + 2 < len(d[0]) and y + 2 < len(d):
            r += s(
                "".join(d[y + z][x + z] for z in range(3)),
                "".join(d[y + z][x + 2 - z] for z in range(3)),
            )
print(r)
