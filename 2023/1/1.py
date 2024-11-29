lns = open("input.txt", "r").read().split("\n")
r = 0
for l in lns:
    n = "".join(c for c in l if c.isdigit())
    r += int(n[0] + n[-1])
print(r)
