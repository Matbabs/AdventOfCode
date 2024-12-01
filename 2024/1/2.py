d = open("input.txt", "r").read().split("\n")
l, r, s = [], [], 0
for e in d:
    v = list(map(int, e.split()))
    l.append(v[0])
    r.append(v[-1])
for el in l:
    c = 0
    for er in r:
        if el == er:
            c += 1
    print(c, r.count(el))
    s += el * c
print(s)
